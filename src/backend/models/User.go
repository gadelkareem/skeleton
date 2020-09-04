package models

import (
    "fmt"
    "hash/fnv"
    "net/url"
    "strconv"
    "strings"
    "time"

    "backend/internal"
    "backend/kernel"
    "github.com/astaxie/beego/logs"
    h "github.com/gadelkareem/go-helpers"
    "golang.org/x/crypto/bcrypt"
)

const (
    RoleGuest = "guest"
    RoleUser  = "user"
    RoleAdmin = "admin"
)

var ForbiddenUserames = []string{"admin", "mod", "root"}

type (
    User struct {
        Base      `pg:" inherit"`
        tableName struct{} `pg:"s_users,alias:u,discard_unknown_columns"`

        ID                          int64             `pg:"id,pk" jsonapi:"primary,users" fake:"skip"`
        Username                    string            `valid:"Required;Match(/^[A-Za-z0-9]+(?:[-_.][A-Za-z0-9]+)*$/)" pg:"username,notnull" jsonapi:"attr,username" fake:"{person.last}-###"`
        Password                    string            `valid:"MinSize(6);MaxSize(100)" pg:"-" jsonapi:"attr,password" fake:"?????##??"`
        PasswordHash                string            `pg:"password_hash,use_zero" json:"-" fake:"skip"`
        ForgotPasswordHash          string            `pg:"forgot_password_hash" json:"-" fake:"skip"`
        ForgotPasswordHashCreatedAt time.Time         `pg:"forgot_password_hash_created_at" json:"-" fake:"skip"`
        Email                       string            `valid:"Required;Email;MaxSize(100)" pg:"email,notnull" jsonapi:"attr,email" fake:"{person.first}####@{person.last}.{internet.domain_suffix}"`
        EmailVerifyHash             string            `pg:"email_verify_hash,notnull" json:"-" fake:"skip"`
        EmailVerifyCreatedAt        time.Time         `pg:"email_verify_created_at" json:"-" fake:"skip"`
        FirstName                   string            `valid:"MaxSize(100)" pg:"first_name" jsonapi:"attr,first_name" fake:"{person.first}"`
        LastName                    string            `valid:"MaxSize(100)" pg:"last_name" jsonapi:"attr,last_name" fake:"{person.last}"`
        AvatarURL                   string            `pg:"avatar_url" jsonapi:"attr,avatar_url" fake:"https://{person.last}.{internet.domain_suffix}/?????????????????????????.jpg"`
        Mobile                      string            `pg:"mobile" jsonapi:"attr,mobile" fake:"+1-202-555-####"`
        MobileVerifyCode            string            `pg:"mobile_verify_code" json:"-" fake:"skip"`
        MobileVerifyCreatedAt       time.Time         `pg:"mobile_verify_created_at" json:"-" fake:"skip"`
        MobileVerified              bool              `pg:"mobile_verified"  jsonapi:"attr,mobile_verified" fake:"skip"`
        RecoveryQuestions           map[string]string `pg:"recovery_questions"  json:"-" fake:"skip"`
        RecoveryQuestionsSet        bool              `pg:"recovery_questions_set"  jsonapi:"attr,recovery_questions_set" fake:"skip"`
        Roles                       []string          `pg:"roles,array" jsonapi:"attr,roles" json:"roles" fake:"skip"`
        Active                      bool              `pg:"active" jsonapi:"attr,active" fake:"skip"`
        SocialLogin                 bool              `pg:"social_login" jsonapi:"attr,social_login" fake:"skip"`
        Language                    string            `valid:"MaxSize(100);Alpha" pg:"language" jsonapi:"attr,language" fake:"{languages.short}"`
        Address                     Address           `pg:"address" jsonapi:"attr,address"`
        AuthenticatorEnabled        bool              `pg:"authenticator_enabled" jsonapi:"attr,authenticator_enabled" json:"enabled" fake:"skip"`
        AuthenticatorSecret         string            `pg:"authenticator_secret" json:"-" fake:"skip"`
        Country                     string            `valid:"MaxSize(100)" pg:"country" jsonapi:"attr,country" fake:"{address.country}"`
        LastLoginAt                 time.Time         `pg:"last_login_at" json:"-" fake:"skip"`
        DeletedAt                   time.Time         `pg:"deleted_at,type:TIMESTAMPTZ" json:"-"`
    }
    Address struct {
        Street  string `jsonapi:"attr,street" json:"street"  fake:"{address.street_name} ###"`
        City    string `jsonapi:"attr,city" json:"city"  fake:"{address.city}"`
        ZipCode string `jsonapi:"attr,zip_code" json:"zip_code"  fake:"{address.zip}"`
    }
)

func NewUser() *User {
    return &User{
        Base:        NewBaseModel(),
        Active:      false,
    }
}

func NewEmptyUser() *User {
    return new(User)
}

func (m *User) GetID() string {
    return fmt.Sprintf("%d", m.ID)
}

func (m *User) GetResetPasswordURL() string {
    q := make(url.Values)
    q.Set("t", m.ForgotPasswordHash)
    q.Set("email", m.Email)
    return fmt.Sprintf("%s/auth/reset-password/?%s", kernel.App.FrontEndURL, q.Encode())
}

func (m *User) GetEmailVerificationURL() string {
    q := make(url.Values)
    q.Set("t", m.EmailVerifyHash)
    q.Set("email", m.Email)
    return fmt.Sprintf("%s/auth/verify-email/?%s", kernel.App.FrontEndURL, q.Encode())
}

func (m *User) GetFullName() string {
    if m.FirstName != "" || m.LastName != "" {
        return fmt.Sprintf("%s %s", m.FirstName, m.LastName)
    }
    return m.Username
}

func (m *User) CleanStrings() {
    m.Username = strings.ToLower(h.TrimLine(m.Username))
    m.Email = strings.ToLower(h.TrimLine(m.Email))
    m.FirstName = h.TrimLine(m.FirstName)
    m.LastName = h.TrimLine(m.LastName)
    m.Mobile = h.TrimLine(m.Mobile)
    m.Language = strings.ToUpper(h.SubString(h.TrimLine(m.Language), 2))
    m.Country = h.TrimLine(m.Country)
}

func (m *User) HashPass() {
    m.PasswordHash = hashAndSalt(m.Password)
    m.Password = ""
}

func (m *User) IsValidPass(s string) bool {
    s = strings.TrimSpace(s)
    if s == "" {
        return false
    }
    return comparePasswords(m.PasswordHash, s)
}

func (m *User) DisableAuthenticator() {
    m.AuthenticatorEnabled = false
    m.AuthenticatorSecret = ""
}

func (m *User) UnVerifyMobile() {
    m.MobileVerified = false
}

func (m *User) GenerateEmailVerificationHash() {
    m.EmailVerifyHash = h.Md5(strconv.Itoa(h.RandomNumber(1, 1000)))
    m.EmailVerifyCreatedAt = time.Now()
}

func (m *User) SetDefaultAvatar() {
    if m.AvatarURL == "" {
        m.AvatarURL = fmt.Sprintf("https://www.gravatar.com/avatar/%s.jpg", h.Md5(m.Email))
    }
}

func (m *User) GenerateForgotPasswordHash() {
    m.ForgotPasswordHash = h.Md5(strconv.Itoa(h.RandomNumber(1, 1000)))
    m.ForgotPasswordHashCreatedAt = time.Now()
}

func (m *User) Activate() {
    m.EmailVerifyHash = ""
    m.Active = true
}

func (m *User) VerifyMobile() {
    m.MobileVerifyCode = ""
    m.MobileVerified = true
    m.MobileVerifyCreatedAt = time.Time{}
}

func (m *User) GenerateVerifyMobileCode() {
    m.MobileVerifyCode = fmt.Sprintf("%d", h.RandomNumber(100000, 999999))
    m.MobileVerifyCreatedAt = time.Now()
}

func (m *User) IsValidMobileCode(s string) bool {
    s = strings.TrimSpace(s)
    if s == "" {
        return false
    }
    return m.MobileVerifyCode == s && m.MobileVerifyCreatedAt.After(time.Now().Add(-1*time.Hour))
}

func (m *User) IsValidEmailVerificationHash(s string) bool {
    s = strings.TrimSpace(s)
    if s == "" {
        return false
    }
    return strings.ToLower(s) == m.EmailVerifyHash &&
        m.EmailVerifyCreatedAt.After(time.Now().Add(-24*time.Hour))
}

func (m *User) IsValidForgotPasswordHash(s string) bool {
    s = strings.TrimSpace(s)
    if s == "" {
        return false
    }
    return m.ForgotPasswordHash == s &&
        m.ForgotPasswordHashCreatedAt.After(time.Now().Add(-1*time.Hour))
}

func (m *User) AddRole(r string) {
    if !h.InArray(r, m.Roles) {
        m.Roles = append(m.Roles, r)
    }
}

func (m *User) MakeAdmin() {
    m.AddRole("admin")
}

func (m *User) AddRecoveryQuestions(q []*RecoveryQuestion) error {
    m.RecoveryQuestions = make(map[string]string)
    for _, rc := range q {
        if h.TrimWhitespace(rc.Answer) != "" {
            m.RecoveryQuestions[rc.Question] = hashAndSalt(rc.Answer)
        }
    }
    if len(m.RecoveryQuestions) < 3 {
        return internal.ErrRecoveryQuestionNum
    }
    m.RecoveryQuestionsSet = true
    return nil
}

func (m *User) Sanitize() {
    m.Password = ""
    m.ForgotPasswordHash = ""
    m.RecoveryQuestions = nil
    m.MobileVerifyCode = ""
    m.AuthenticatorSecret = ""
    m.EmailVerifyHash = ""
}

func (m *User) Delete() {
    m.DeletedAt = time.Now()
    hs := int(hash(m.Username)) + h.RandomNumber(0, 1000)
    m.Username = fmt.Sprintf("deleted_%d", hs)
    m.Email = fmt.Sprintf("deleted_%d@example.com", hs)
    m.FirstName = ""
    m.LastName = ""
    m.AvatarURL = ""
    m.Mobile = ""
    m.RecoveryQuestions = nil
    m.Roles = nil
    m.Language = ""
    m.Address = Address{}
    m.AuthenticatorEnabled = false
    m.Active = false
    m.Password = ""
    m.RecoveryQuestions = nil
    m.MobileVerifyCode = ""
    m.AuthenticatorSecret = ""
}

func (m *User) IsValidRecoveryQuestions(q []*RecoveryQuestion) bool {
    if len(m.RecoveryQuestions) < len(q) {
        return false
    }
    for _, rc := range q {
        s, k := m.RecoveryQuestions[rc.Question]
        if !k {
            return false
        }
        if !comparePasswords(s, rc.Answer) {
            fmt.Printf("\n\nQuestions: %+q\n\n", rc)
            return false
        }
    }
    return true
}

func hashAndSalt(pass string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
    if err != nil {
        logs.Error("Error hashing password %v", err)
    }
    hs := string(hash)
    if hs == "" {
        hs = h.RandomString(60)
    }
    return hs
}

func comparePasswords(hashedPass, plainPass string) bool {
    if len(plainPass) == 0 {
        return false
    }
    err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
    if err != nil {
        if err != bcrypt.ErrMismatchedHashAndPassword {
            logs.Error("Error comparing passwords %v", err)
        }
        return false
    }

    return true
}

func hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}
