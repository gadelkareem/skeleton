package controllers_test

import (
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"
    "time"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "backend/services"
    "backend/tests"
    h "github.com/gadelkareem/go-helpers"
    "github.com/gbrlsnchs/jwt/v3"
    "github.com/pquerna/otp/totp"
    "github.com/stretchr/testify/assert"
)

func TestAuthController_Token(t *testing.T) {
    t.Parallel()
    u := tests.User(false)

    // test inactive user
    w := tests.Login(t, u.Username, u.Password)
    assert.Equal(t, w.Code, http.StatusUnauthorized)
    tests.AssertError(t, w.Body, internal.ErrInvalidPass.Error())

    // test wrong password
    u = tests.User(true)
    w = tests.Login(t, u.Username, u.Password+"x")
    assert.Equal(t, w.Code, http.StatusUnauthorized)
    tests.AssertError(t, w.Body, internal.ErrInvalidPass.Error())

    // test normal login
    w = tests.Login(t, u.Username, u.Password)
    assert.Equal(t, http.StatusOK, w.Code)
    tk := tests.ReadToken(t, w.Body.Bytes())
    pl := tests.C.JWTService.ParseToken(tk.RefreshToken)
    assert.False(t, pl.RememberMe)
    assert.Equal(t, time.Now().Add(services.RefreshExpiryExtend).Unix(), pl.ExpirationTime.Unix())

    // test remember me expiry
    r := &models.Login{Username: u.Username, Password: u.Password, RememberMe: true}
    w, _ = tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/token", Model: r, Status: http.StatusOK})
    tk = tests.ReadToken(t, w.Body.Bytes())
    pl = tests.C.JWTService.ParseToken(tk.RefreshToken)
    assert.True(t, pl.RememberMe)
    assert.Equal(t, time.Now().Add(services.RememberMeExpiryExtend).Unix(), pl.ExpirationTime.Unix())

    // test Authenticator
    u.RecoveryQuestionsSet = true
    tests.SaveUser(t, u)
    authenticator, err := tests.C.AuthenticatorService.Generate(u)
    tests.FailOnErr(t, err)
    u.AuthenticatorEnabled = true
    tests.SaveUser(t, u)
    w = tests.Login(t, u.Username, u.Password)
    assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
    tests.AssertErrorObj(t, w.Body, internal.ErrAuthenticatorCodeMissing)

    code, err := totp.GenerateCode(authenticator.Seed, time.Now())
    tests.FailOnErr(t, err)
    r = &models.Login{Username: u.Username, Password: u.Password, RememberMe: true, Code: code}
    w, _ = tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/token", Model: r,
        Status: http.StatusOK})

    // test Verified mobile
    u = tests.User(true)
    u.RecoveryQuestionsSet = true
    u.MobileVerified = true
    tests.SaveUser(t, u)
    w = tests.Login(t, u.Username, u.Password)
    assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
    tests.AssertErrorObj(t, w.Body, internal.ErrMobileCodeMissing)

    nu := tests.RefreshUser(t, u, true)
    r = &models.Login{Username: u.Username, Password: u.Password, RememberMe: true, Code: nu.MobileVerifyCode}
    w, _ = tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/token", Model: r, Status: http.StatusOK})
}

func TestAuthController_RefreshToken(t *testing.T) {
    t.Parallel()
    u := tests.User(true)
    rt, err := tests.C.JWTService.RefreshToken(u, false)
    tests.FailOnErr(t, err)

    w, _ := tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/refresh-token",
        Model: &models.AuthToken{RefreshToken: rt}, Status: http.StatusOK})
    assert.NotEqual(t, rt, tests.ReadToken(t, w.Body.Bytes()).Token)

    now := time.Now()
    pl := &services.Payload{
        Payload: jwt.Payload{
            ExpirationTime: jwt.NumericDate(now),
            NotBefore:      jwt.NumericDate(now.Add(10 * time.Hour)),
            IssuedAt:       jwt.NumericDate(now),
        },
        UserId: u.ID,
    }

    tk1, err := jwt.Sign(pl, jwt.NewHS256([]byte("137954aa8b5fa851c20d7f1d6c52636565f22947c3faaa7a84c14616ff97d33f"))) // invalid key
    tests.FailOnErr(t, err)
    w, _ = tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/refresh-token",
        Model: &models.AuthToken{RefreshToken: string(tk1)}, Status: http.StatusUnauthorized})
}

var xhr = map[string]string{"X-Requested-With": "XMLHttpRequest"}

func TestAuthController_RefreshCookie(t *testing.T) {
    t.Parallel()
    u, cks := ajaxLogin(t)

    w, _ := tests.ApiRequest(&tests.Request{T: t, Method: http.MethodPost, URI: "/auth/refresh-cookie",
        Headers: xhr, Cookies: cks, Status: http.StatusOK})

    rt, err := tests.C.JWTService.Token(u)
    tests.FailOnErr(t, err)
    assert.Equal(t, rt, tests.ReadToken(t, w.Body.Bytes()).Token)

}

func TestAuthController_Logout(t *testing.T) {
    t.Parallel()
    ajaxLogin(t)

    w, _ := tests.ApiRequest(&tests.Request{T: t, Method: http.MethodGet, URI: "/auth/logout", Headers: xhr,
        Status: http.StatusOK})
    cks := w.Result().Cookies()
    assert.Equal(t, time.Time{}, cks[0].Expires)

}

func ajaxLogin(t *testing.T) (*models.User, []*http.Cookie) {
    u := tests.User(true)

    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:  http.MethodPost,
        URI:     "/auth/token",
        Model:   &models.Login{Username: u.Username, Password: u.Password},
        Headers: xhr,
        Status:  http.StatusOK})
    cks := w.Result().Cookies()
    assert.Greater(t, len(cks), 0)
    return u, cks
}

func TestAuthController_SocialRedirect(t *testing.T) {
    t.Parallel()
    f := func(t *testing.T, provider string) (w *httptest.ResponseRecorder, rq *http.Request) {
        return tests.ApiRequest(&tests.Request{T: t,
            Method: http.MethodPost,
            URI:    "/auth/social/redirect",
            Model:  &models.SocialAuth{Provider: provider},
        })
    }
    w, _ := f(t, "google")
    assert.Equal(t, w.Code, http.StatusOK)

    m := new(models.SocialAuth)
    tests.ParseModel(t, m, w.Body)
    u, err := url.Parse(m.RedirectUri)
    tests.FailOnErr(t, err)
    q, _ := url.ParseQuery(u.RawQuery)
    q.Set("state", "")
    u.RawQuery = q.Encode()
    assert.Equal(t, u, h.ParseUrl("https://accounts.google.com/o/oauth2/auth?client_id="+
        kernel.App.Config.String("social::googleClientID")+
        "&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fsocial-callback%2F%3Fp%3Dgoogle&response_type=code&scope=profile+email&state="))

    w, _ = f(t, "someprovider")
    assert.Equal(t, w.Code, http.StatusBadRequest)
    tests.AssertError(t, w.Body, internal.ErrInvalidSocialProvider.Error())
}

func TestAuthController_SocialCallback(t *testing.T) {
    // t.Parallel()
    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method: http.MethodPost,
        URI:    "/auth/social/callback",
        Model:  &models.SocialAuth{State: "sometoken", Code: "somecode"},
        Status: http.StatusOK,
    })
    tk := tests.ReadToken(t, w.Body.Bytes())
    assert.NotEmpty(t, tk)
    u, err := tests.C.UserService.FindUser(&models.User{Email: tests.GocialiteMockUser.Email}, true)
    tests.FailOnErr(t, err)
    assert.NotEmpty(t, u)
    assert.Equal(t, tests.GocialiteMockUser.LastName, u.LastName)
    assert.Equal(t, tests.GocialiteMockUser.FirstName, u.FirstName)
    assert.True(t, u.Active)
    assert.True(t, u.SocialLogin)
    assert.Empty(t, u.EmailVerifyHash)
}
