package controllers_test

import (
    "fmt"
    "net/http"
    "os"
    "reflect"
    "strings"
    "testing"
    "time"

    "backend/internal"
    "github.com/pquerna/otp/totp"

    "backend/kernel"
    "backend/models"
    "backend/tests"
    "github.com/brianvoe/gofakeit/v4"
    "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    tests.Bootstrap()
    retCode := m.Run()
    os.Exit(retCode)
}

func TestUserController_SignUp(t *testing.T) {
    u := new(models.User)
    gofakeit.Struct(&u)

    sub := fmt.Sprintf("Verify your %s email address", kernel.SiteName)
    tests.ExpectEmail(t,
        "",
        []string{strings.ToLower(u.Email)},
        sub,
        "")

    tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/users", Model: u,
        Status: http.StatusCreated})

    w, _ := tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/users", Model: u,
        Status: http.StatusBadRequest})
    tests.AssertError(t, w.Body, internal.ErrValidationError.Error())
    tests.AssertValidationError(t, w.Body, "Email.Unique", "Email already exists in our system.")

    tests.CheckEmailRetry(t, sub)

}

func TestUserController_VerifyEmail(t *testing.T) {
    // t.Parallel()
    u := tests.User(false)

    sub := fmt.Sprintf("Welcome to %s", kernel.SiteName)
    tests.ExpectEmail(t,
        "",
        []string{u.Email},
        sub,
        "")

    r := &models.VerifyEmail{Email: u.Email, Token: gofakeit.Word()}
    req := &tests.Request{T: t,
        Method: http.MethodPost,
        URI:    "/users/verify-email",
        Model:  r,
        Status: http.StatusBadRequest}
    tests.ApiRequest(req)

    r.Token = u.EmailVerifyHash
    req.Status = http.StatusOK
    tests.ApiRequest(req)
    // wait for go routine sending the welcome mail
    time.Sleep(100 * time.Millisecond)
    tests.CheckEmailRetry(t, sub)

}

func TestUserController_ForgotPassword(t *testing.T) {
    // t.Parallel()
    u := tests.User(true)

    sub := fmt.Sprintf("Reset your %s password", kernel.SiteName)
    tests.ExpectEmail(t,
        "",
        []string{u.Email},
        sub,
        "")

    r := &models.ResetPassword{Email: u.Email}
    tests.ApiRequest(&tests.Request{T: t,
        Method: http.MethodPost, URI:
        "/users/forgot-password",
        Model:  r,
        Status: http.StatusOK})
    tests.CheckEmailRetry(t, sub)
}

func TestUserController_ResetPassword(t *testing.T) {
    t.Parallel()
    tests.ResetEmailService()
    u := tests.User(true)
    err := tests.C.UserService.ForgotPassword(u.Email, "")
    tests.FailOnErr(t, err)

    u, err = tests.C.UserRepository.FindByEmail(u.Email, false)
    tests.FailOnErr(t, err)

    pass := gofakeit.Password(false, false, false, false, false, 6)
    r := &models.ResetPassword{Email: u.Email, Token: u.ForgotPasswordHash, Password: pass}
    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method: http.MethodPost,
        URI:    "/users/reset-password",
        Model:  r,
        Status: http.StatusOK})

    u, err = tests.C.UserRepository.FindByEmail(u.Email, false)
    tests.FailOnErr(t, err)

    w = tests.Login(t, u.Username, pass)
    assert.Equal(t, http.StatusOK, w.Code)
    tests.ReadToken(t, w.Body.Bytes())
}

func TestUserController_Update(t *testing.T) {
    // t.Parallel()
    u, tk := tests.UserWithToken(true)

    // test changing first and last name
    u.LastName = strings.ToLower(gofakeit.LastName())
    u.FirstName = strings.ToLower(gofakeit.FirstName())
    req := &tests.Request{T: t,
        Method:    http.MethodPatch,
        URI:       fmt.Sprintf("/users/%d", u.ID),
        Model:     u,
        AuthToken: tk,
        Status:    http.StatusOK}
    tests.ApiRequest(req)

    newUser := tests.RefreshUser(t, u, true)
    assert.NotEmpty(t, newUser)
    assert.Equal(t, newUser.LastName, u.LastName)
    assert.Equal(t, newUser.FirstName, u.FirstName)
    assert.True(t, newUser.Active)

    // test change email
    u.Email = strings.ToLower(gofakeit.Email())
    sub := fmt.Sprintf("Verify your %s email address", kernel.SiteName)
    tests.ExpectEmail(t,
        "",
        []string{strings.ToLower(u.Email)},
        sub,
        "")
    tests.ApiRequest(req)

    newUser = tests.RefreshUser(t, u, false)

    assert.NotEmpty(t, newUser)
    assert.False(t, newUser.Active)
    assert.Equal(t, newUser.Email, u.Email)
    tests.CheckEmailRetry(t, sub)
}

func TestUserController_UpdateByAdmin(t *testing.T) {
    t.Parallel()
    u, tk := tests.UserWithToken(true)
    admin, adminTk := tests.UserWithToken(true)
    admin.MakeAdmin()
    tests.SaveUser(t, admin)

    // test if user can access other users
    f := func(user *models.User, token string, status int) {
        tests.ApiRequest(&tests.Request{T: t,
            Method:    http.MethodPatch,
            URI:       fmt.Sprintf("/users/%d", user.ID),
            Model:     user,
            AuthToken: token,
            Status:    status})
    }
    f(admin, tk, http.StatusForbidden)

    // test admin changing user
    u.LastName = strings.ToLower(gofakeit.LastName())
    u.FirstName = strings.ToLower(gofakeit.FirstName())
    u.Active = true
    u.AuthenticatorEnabled = true
    f(u, adminTk, http.StatusOK)

    newUser := tests.RefreshUser(t, u, false)
    assert.NotEmpty(t, newUser)
    assert.Equal(t, newUser.LastName, u.LastName)
    assert.Equal(t, newUser.FirstName, u.FirstName)
    assert.True(t, newUser.Active)
    assert.True(t, newUser.AuthenticatorEnabled)
}

func TestUserController_UpdatePassword(t *testing.T) {
    t.Parallel()
    u, tk := tests.UserWithToken(true)
    r := &models.UpdatePassword{
        OldPassword: u.Password,
        Password:    gofakeit.Password(false, false, false, false, false, 6),
    }

    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodPatch,
        URI:       fmt.Sprintf("/users/%d/password", u.ID),
        Model:     r,
        AuthToken: tk,
        Status:    http.StatusOK})

    w = tests.Login(t, u.Username, r.Password)
    assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_GetUser(t *testing.T) {
    t.Parallel()
    u, tk := tests.UserWithToken(true)
    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       fmt.Sprintf("/users/%d", u.ID),
        AuthToken: tk,
        Status:    http.StatusOK})
    nu := new(models.User)
    tests.ParseModel(t, nu, w.Body)
    assert.Equal(t, nu.Email, u.Email)
    assert.Equal(t, nu.FirstName, u.FirstName)
    assert.Equal(t, nu.LastName, u.LastName)
    assert.Equal(t, nu.Username, u.Username)

    // test get another user account
    u2 := tests.User(true)
    w, _ = tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       fmt.Sprintf("/users/%d", u2.ID),
        AuthToken: tk,
        Status:    http.StatusForbidden})
}

func TestUserController_GenerateAuthenticator(t *testing.T) {
    t.Parallel()
    u, tk := tests.UserWithToken(true)
    u.RecoveryQuestionsSet = true
    tests.SaveUser(t, u)
    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodPatch,
        URI:       fmt.Sprintf("/users/%d/generate-auth-code", u.ID),
        AuthToken: tk,
        Status:    http.StatusOK,
    })

    nu := new(models.Authenticator)
    tests.ParseModel(t, nu, w.Body)
    assert.NotEmpty(t, nu.Image)
    assert.NotEmpty(t, nu.URL)
    u = tests.RefreshUser(t, u, false)
    assert.Equal(t, nu.Seed, u.AuthenticatorSecret)
}

func TestUserController_Authenticator(t *testing.T) {
    t.Parallel()

    // test RecoveryQuestions not Set
    u, tk := tests.UserWithToken(true)
    code := "986944"
    f := func(enable bool, status int) {
        tests.ApiRequest(&tests.Request{T: t,
            Method:    http.MethodPatch,
            URI:       fmt.Sprintf("/users/%d/authenticator", u.ID),
            Model:     &models.Authenticator{Enable: enable, Code: code},
            AuthToken: tk,
            Status:    status})
    }
    f(true, http.StatusBadRequest)

    // test enable
    u.RecoveryQuestionsSet = true
    r, err := tests.C.AuthenticatorService.Generate(u)
    tests.FailOnErr(t, err)
    code, err = totp.GenerateCode(r.Seed, time.Now())
    tests.FailOnErr(t, err)
    tests.SaveUser(t, u)
    f(true, http.StatusOK)

    u = tests.RefreshUser(t, u, false)
    assert.True(t, u.AuthenticatorEnabled)

    // test disable
    f(false, http.StatusOK)
    u = tests.RefreshUser(t, u, false)
    assert.False(t, u.AuthenticatorEnabled)
}

func TestUserController_GetUsers(t *testing.T) {
    _, tk := tests.AdminWithToken(t)

    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       "/users",
        AuthToken: tk,
        Status:    http.StatusOK})
    users, payload := tests.ParseModels(t, reflect.TypeOf(new(models.User)), w.Body)
    assert.NotEmpty(t, users)
    totalUsers, err := tests.C.UserRepository.Count()
    tests.FailOnErr(t, err)
    count := totalUsers
    if totalUsers > kernel.ListLimit {
        count = kernel.ListLimit
    }
    assert.Len(t, users, count)

    m := *payload.Meta
    assert.Equal(t, totalUsers, int(m["page"].(map[string]interface{})["total"].(float64)))
    assert.Equal(t, kernel.ListLimit, int(m["page"].(map[string]interface{})["size"].(float64)))
    assert.NotEmpty(t, users[0].(*models.User).ID)
    assert.NotEmpty(t, users[1].(*models.User).Username)

    // test admin access only
    _, tk = tests.UserWithToken(true)
    w, _ = tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       "/users",
        AuthToken: tk,
        Status:    http.StatusForbidden})

}

func TestUserController_SendVerifySMS(t *testing.T) {
    t.Parallel()

    // test endpoint
    u, tk := tests.UserWithToken(true)
    u.RecoveryQuestionsSet = true
    tests.SaveUser(t, u)
    tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodPatch,
        URI:       fmt.Sprintf("/users/%d/send-verify-sms", u.ID),
        AuthToken: tk,
        Status:    http.StatusOK})
}

func TestUserController_VerifyMobile(t *testing.T) {
    t.Parallel()
    // test recovery questions not set
    u, tk := tests.UserWithToken(true)
    u.GenerateVerifyMobileCode()
    tests.SaveUser(t, u)

    f := func(code string, status int) {
        tests.ApiRequest(&tests.Request{T: t,
            Method:    http.MethodPatch,
            URI:       fmt.Sprintf("/users/%d/verify-mobile", u.ID),
            Model:     &models.VerifyMobile{Code: code},
            AuthToken: tk,
            Status:    status})
    }
    f(u.MobileVerifyCode, http.StatusBadRequest)

    // test verify mobile code
    u.RecoveryQuestionsSet = true
    tests.SaveUser(t, u)
    f(u.MobileVerifyCode, http.StatusOK)
    u = tests.RefreshUser(t, u, true)
    assert.True(t, u.MobileVerified)
    assert.Empty(t, u.MobileVerifyCode)

    // test wrong verify mobile code
    u, tk = tests.UserWithToken(true)
    f(gofakeit.Word(), http.StatusBadRequest)
    u = tests.RefreshUser(t, u, true)
    assert.False(t, u.MobileVerified)
}

func recoveryQuestions() models.RecoveryQuestions {
    return models.RecoveryQuestions{Questions: []*models.RecoveryQuestion{{Question: "q1", Answer: "a1"},
        {Question: "q2", Answer: "a2"}, {Question: "q3", Answer: "a3"}}}
}

func TestUserController_RecoveryQuestions(t *testing.T) {
    t.Parallel()

    // test add recovery questions
    u, tk := tests.UserWithToken(true)
    r := recoveryQuestions()
    tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodPatch,
        URI:       fmt.Sprintf("/users/%d/recovery-questions", u.ID),
        Model:     &r,
        AuthToken: tk,
        Status:    http.StatusOK})
    u = tests.RefreshUser(t, u, true)
    assert.True(t, u.RecoveryQuestionsSet)
    b := u.IsValidRecoveryQuestions(r.Questions)
    assert.True(t, b)

}

func TestUserController_GetRecoveryQuestions(t *testing.T) {
    t.Parallel()

    // test add recovery questions
    u := tests.User(true)
    r := recoveryQuestions()
    l := &models.Login{Username: u.Username, Password: u.Password}
    err := u.AddRecoveryQuestions(r.Questions)
    tests.FailOnErr(t, err)
    tests.SaveUser(t, u)
    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method: http.MethodPost,
        URI:    "/users/recovery-questions",
        Model:  l,
        Status: http.StatusOK,
    })
    r2 := new(models.RecoveryQuestions)
    tests.ParseModel(t, r2, w.Body)
    var qs []string
    for _, q := range r.Questions {
        qs = append(qs, q.Question)
    }

    for _, q := range r2.Questions {
        assert.Empty(t, q.Answer)
        assert.Contains(t, qs, q.Question)
    }
}

func TestUserController_DisableMFA(t *testing.T) {
    t.Parallel()

    // test disable MFA
    r := recoveryQuestions()
    u := tests.User(true)
    err := u.AddRecoveryQuestions(r.Questions)
    tests.FailOnErr(t, err)
    // r, err := tests.C.AuthenticatorService.Generate(u)
    // tests.FailOnErr(t, err)
    u.AuthenticatorEnabled = true
    u.VerifyMobile()
    tests.SaveUser(t, u)
    f := func(status int) {
        tests.ApiRequest(&tests.Request{T: t,
            Method: http.MethodPost,
            URI:    "/users/disable-mfa",
            Model:  &models.DisableMFA{Username: u.Username, Password: u.Password, RecoveryQuestions: r.Questions},
            Status: status,
        })
    }
    f(http.StatusOK)
    u = tests.RefreshUser(t, u, false)
    assert.False(t, u.AuthenticatorEnabled)

    r.Questions[0].Answer = "wrong answer"
    f(http.StatusUnauthorized)
}

func TestUserController_DeleteUser(t *testing.T) {
    t.Parallel()

    // test delete
    u, tk := tests.UserWithToken(true)
    tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodDelete,
        URI:       fmt.Sprintf("/users/%d", u.ID),
        AuthToken: tk,
        Status:    http.StatusNoContent,
    })

    // test deleted user cannot login
    w := tests.Login(t, u.Username, u.Password)
    assert.Equal(t, w.Code, http.StatusUnauthorized)
    tests.AssertError(t, w.Body, internal.ErrInvalidPass.Error())
}
