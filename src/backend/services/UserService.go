package services

import (
    "strings"
    "time"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "backend/rbac"
    "backend/utils/paginator"
    "github.com/astaxie/beego/logs"
    h "github.com/gadelkareem/go-helpers"
)

type (
    UserService struct {
        r   *models.UserRepository
        es  *EmailService
        sms *SMSService
        p   *rbac.RBAC
        c   *CacheService
        py  *PaymentService
    }
)

func NewUserService(r *models.UserRepository, es *EmailService, sms *SMSService, p *rbac.RBAC, c *CacheService, py *PaymentService) *UserService {
    return &UserService{r: r, es: es, sms: sms, p: p, c: c, py: py}
}

func (s *UserService) Save(m *models.User, columns ...string) error {
    err := s.r.Save(m, columns...)
    if err != nil {
        return err
    }
    go s.c.InvalidateModel(m)
    return nil
}

func (s *UserService) getCacheUser(params ...interface{}) (m *models.User, cID string) {
    m = models.NewEmptyUser()
    cID, _ = s.c.Get(m, nil, params...)
    if m == nil || m.ID == 0 {
        m = nil
    }
    return m, cID
}

func (s *UserService) cacheUser(m *models.User, cID string, tags ...string) {
    go s.c.Put(m, cID, nil, tags...)
}

func (s *UserService) UserById(id int64) (m *models.User, err error) {
    m, cID := s.getCacheUser(id)
    if m != nil {
        return
    }

    m, err = s.r.FindById(id, true)
    if err != nil {
        return
    }

    go s.cacheUser(m, cID)

    return
}

func (s *UserService) FindUser(u *models.User, checkActive bool) (m *models.User, err error) {
    m, cID := s.getCacheUser(u)
    if m != nil && (!checkActive || m.Active) {
        return
    }

    if u.ID != 0 {
        m, err = s.r.FindById(u.ID, checkActive)
    } else if u.Email != "" {
        m, err = s.r.FindByEmail(u.Email, checkActive)
    } else if u.Username != "" {
        m, err = s.r.FindByUsername(u.Username, checkActive)
    }

    if err != nil {
        return nil, err
    }
    if m == nil {
        return nil, internal.ErrNotFound
    }

    go s.cacheUser(m, cID)

    return m, nil
}

func (s *UserService) DeleteUser(id int64) error {
    u, err := s.r.FindById(id, false)
    if err != nil {
        return err
    }
    u.Delete()
    return s.Save(u)
}

func (s *UserService) MakeAdmin(username string) error {
    if !kernel.App.IsCLI {
        return internal.ErrForbidden
    }
    u, err := s.FindUser(&models.User{Username: username}, true)
    if err != nil {
        return err
    }
    u.MakeAdmin()
    return s.Save(u)
}

func (s *UserService) Authenticate(username, password string) (m *models.User, err error) {
    m, err = s.r.CheckPassword(username, password)
    return m, err
}

func (s *UserService) UpdateLoginAt(m *models.User) error {
    m.LastLoginAt = time.Now()
    return s.Save(m)
}

func (s *UserService) SignUp(m *models.User) (err error) {
    err = s.r.ValidateUser(m)
    if err != nil {
        return
    }
    if m.Password == "" {
        return internal.ValidationError("Password.Required", "Password cannot be empty.")
    }
    if h.InArray(m.Username, models.ForbiddenUsernames) {
        return internal.ValidationError("Username", "Username is not allowed.")
    }
    m.HashPass()
    m.GenerateEmailVerificationHash()
    m.SetDefaultAvatar()
    m.AddNotification("Welcome to Skeleton!", "/dashboard/home/")
    err = s.Save(m)
    if err != nil {
        switch err.Error() {
        case internal.ErrEmailExists.Error():
            return internal.ValidationError("Email.Unique", "Email already exists in our system.")
        case internal.ErrUsernameExists.Error():
            return internal.ValidationError("Username.Unique", "Username already exists in our system.")
        }
        return
    }
    // send verify email link
    err = s.es.VerifyUserEmail(m.GetFullName(), m.Email, m.GetEmailVerificationURL())
    if err != nil {
        return
    }
    // @todo add a worker
    go func() {
        cus := &models.Customer{}
        if m.CustomerID != "" {
            cus.ID = m.CustomerID
            cus, err = s.py.UpdateCustomer(cus, m)
        } else {
            cus, err = s.py.CreateCustomer(m)
        }
        if err != nil {
            logs.Error("Error getting payment customer ID %v", err)
            return
        }
        m.CustomerID = cus.ID
        err = s.Save(m)
        if err != nil {
            logs.Error("Error saving user %v", err)
        }
    }()

    return
}

func (s *UserService) SignUpSocial(m *models.User) (err error) {
    err = s.r.ValidateUser(m)
    if err != nil {
        return
    }
    m.Password = h.RandomString(10)
    m.HashPass()
    m.Activate()
    m.SetDefaultAvatar()
    m.SocialLogin = true
    m.AddNotification("Welcome to Skeleton!", "/dashboard/home/")
    err = s.Save(m)
    if err != nil {
        switch err.Error() {
        case internal.ErrEmailExists.Error():
            return internal.ValidationError("Email.Unique", "Email already exists in our system.")
        case internal.ErrUsernameExists.Error():
            return internal.ValidationError("Username.Unique", "Username already exists in our system.")
        }
        return
    }
    // send welcome email
    err = s.es.WelcomeEmail(m.GetFullName(), m.Email)

    return
}

func (s *UserService) VerifyEmail(email, verificationHash string) error {
    m, err := s.FindUser(&models.User{Email: email}, false)
    if err != nil {
        if err == internal.ErrNotFound {
            return internal.ErrInvalidActivationCode
        }
        return err
    }
    if m.Active {
        return internal.ErrEmailAlreadyVerified
    }

    if !m.IsValidEmailVerificationHash(verificationHash) {
        return internal.ErrInvalidActivationCode
    }

    m.Activate()
    err = s.Save(m)
    if err != nil {
        return err
    }

    if m.LastLoginAt.IsZero() {
        // send welcome email
        err = s.es.WelcomeEmail(m.GetFullName(), m.Email)
    }

    return err
}

func (s *UserService) SendVerifySMS(m *models.User) (err error) {
    if !m.RecoveryQuestionsSet {
        return internal.ErrRecoveryQuestionNotSet
    }
    if m.Mobile == "" {
        return internal.ErrMobileRequired
    }
    m.GenerateVerifyMobileCode()
    m.AddNotification("Verify your mobile", "/dashboard/account/verify-mobile/")
    err = s.Save(m)
    if err != nil {
        return
    }

    err = s.sms.Enqueue(m.Mobile, m.MobileVerifyCode)

    return
}

func (s *UserService) VerifyMobile(code string, m *models.User) (err error) {
    if !m.RecoveryQuestionsSet {
        return internal.ErrRecoveryQuestionNotSet
    }
    if m.MobileVerified {
        return internal.ErrMobileAlreadyVerified
    }
    if !m.IsValidMobileCode(code) {
        return internal.ErrInvalidSMSCode
    }
    m.VerifyMobile()
    err = s.Save(m)

    return
}

func (s *UserService) ForgotPassword(email, username string) (err error) {
    m := models.NewEmptyUser()
    if email != "" {
        m.Email = email
    } else if username == "" {
        return internal.ErrEmailRequired // username and email are empty
    } else {
        m.Username = username
    }
    m, err = s.FindUser(m, true)
    if err != nil {
        return err
    }
    if m == nil {
        return internal.ErrEmailNotExist
    }
    // reset pass hash should not be regenerated within 2 hours
    if m.ForgotPasswordHash != "" && m.UpdatedAt.After(time.Now().Add(-2*time.Hour)) {
        return internal.ErrResetPasswordAlreadyGenerated
    }

    m.GenerateForgotPasswordHash()
    err = s.Save(m)
    if err != nil {
        return err
    }

    // send reset password link
    err = s.es.ForgotPasswordEmail(m)

    return err
}

func (s *UserService) ResetPassword(email, forgotPasswordHash, pass string) error {
    m, err := s.FindUser(&models.User{Email: email}, false)
    if err != nil {
        return err
    }
    if m == nil {
        return internal.ErrEmailNotExist
    }

    if !m.IsValidForgotPasswordHash(forgotPasswordHash) {
        return internal.ErrInvalidResetPassHash
    }
    m.ForgotPasswordHash = ""
    m.Password = pass
    err = s.r.ValidateUser(m)
    if err != nil {
        return err
    }
    m.HashPass()

    err = s.Save(m)
    if err != nil {
        return err
    }

    return err
}

func (s *UserService) UpdatePassword(u *models.User, oldPass, pass string) error {
    if oldPass == pass {
        return internal.ValidationError("Password.Password", "Both old and new passwords are the same.")
    }
    m, err := s.r.CheckPassword(u.Username, oldPass)
    if err != nil {
        return err
    }

    m.Password = pass
    err = s.r.ValidateUser(m)
    if err != nil {
        return err
    }
    m.HashPass()

    err = s.Save(m)
    if err != nil {
        return err
    }

    return err
}

func (s *UserService) UpdateProfile(m *models.User) (*models.User, error) {
    err := s.r.ValidateUser(m)
    if err != nil {
        return nil, err
    }
    u, err := s.FindUser(m, true)
    if err != nil {
        return nil, err
    }
    cls := []string{"email", "first_name", "last_name", "language", "address", "country"}
    if m.Email != u.Email {
        m.GenerateEmailVerificationHash()
        if strings.Contains(m.AvatarURL, "www.gravatar.com") {
            m.AvatarURL = ""
            m.SetDefaultAvatar()
        }
        m.Active = false
        cls = append(cls, "email_verify_hash", "email_verify_created_at", "active", "avatar_url")
    }
    if m.Mobile != u.Mobile {
        m.UnVerifyMobile()
        m.GenerateVerifyMobileCode()
        cls = append(cls, "mobile", "mobile_verified", "mobile_verify_created_at", "mobile_verify_code")
    }
    err = s.Save(m, cls...)
    if err != nil {
        if err.Error() == internal.ErrEmailExists.Error() {
            return nil, internal.ValidationError("Email.Unique", "Email already exists in our system.")
        }
        return nil, err
    }
    if m.Email != u.Email {
        // send verify email link
        err = s.es.VerifyUserEmail(m.GetFullName(), m.Email, m.GetEmailVerificationURL())
    }
    return m, err
}

func (s *UserService) UpdateUser(m, admin *models.User) (*models.User, error) {
    if !s.p.HasPermission(admin, "users_edit") {
        return nil, internal.ErrForbidden
    }
    _, err := s.validateAndFind(m, false)
    if err != nil {
        return nil, err
    }
    cls := []string{"username", "password_hash", "email", "first_name", "last_name",
        "mobile", "language", "address", "country", "avatar_url", "active", "authenticator_enabled"}

    if m.Password != "" {
        m.HashPass()
    }
    err = s.Save(m, cls...)
    if err != nil {
        return nil, err
    }

    return m, err
}

func (s *UserService) RecoveryQuestions(r *models.Login) (*models.RecoveryQuestions, error) {
    u, err := s.Authenticate(r.Username, r.Password)
    if err != nil {
        return nil, err
    }

    if !u.RecoveryQuestionsSet {
        return nil, internal.ErrNotFound
    }

    rc := new(models.RecoveryQuestions)
    for q := range u.RecoveryQuestions {
        rc.Questions = append(rc.Questions, &models.RecoveryQuestion{Question: q})
    }

    return rc, nil
}

func (s *UserService) SaveRecoveryQuestions(m *models.User, r *models.RecoveryQuestions) error {
    if m.RecoveryQuestionsSet {
        return internal.ErrRecoveryChangeDisallowed
    }
    err := m.AddRecoveryQuestions(r.Questions)
    if err != nil {
        return err
    }

    return s.Save(m)
}

func (s *UserService) ReadNotification(m *models.User, r *models.Notification) error {
    m.ReadNotification(r)
    m.CleanNotifications()

    return s.Save(m)
}

func (s *UserService) DisableMFA(r *models.DisableMFA) error {
    u, err := s.Authenticate(r.Username, r.Password)
    if err != nil {
        return err
    }
    b := u.IsValidRecoveryQuestions(r.RecoveryQuestions)
    if !b {
        return internal.ErrBadRecoveryAnswers
    }
    u.DisableAuthenticator()
    u.UnVerifyMobile()

    return s.Save(u)
}

func (s *UserService) validateAndFind(m *models.User, checkActive bool) (u *models.User, err error) {
    err = s.r.ValidateUser(m)
    if err != nil {
        return
    }
    u, err = s.FindUser(m, checkActive)
    if err != nil {
        return nil, err
    }
    return
}

func (s *UserService) PaginateUsers(p *paginator.Paginator) (*paginator.Paginator, error) {
    var err error
    var ms []*models.User

    p.Sort = s.sanitizeSort(p.Sort)
    // caching
    id, _ := s.c.Get(&ms, &p)
    if len(ms) > 0 {
        var msi []interface{}
        for _, m := range ms {
            msi = append(msi, m)
        }
        p.Models = msi
        return p, nil
    }

    ms, p.Size, err = s.r.Paginate(p.Filter, p.Sort, p.Offset, p.Limit, []string{
        "id", "first_name", "last_name", "email", "username", "country", "address",
    })
    if err != nil {
        return p, err
    }
    for _, m := range ms {
        p.Models = append(p.Models, m)
    }

    go s.c.Put(p.Models, id, &*p)

    return p, nil
}

func (s *UserService) sanitizeSort(sort map[string]string) map[string]string {
    sort2 := make(map[string]string)
    for c, d := range sort {
        c = strings.ToLower(c)
        if c != "id" && c != "first_name" && c != "last_name" && c != "email" && c != "username" {
            c = "id"
        }
        sort2[c] = d
    }

    return sort2
}
