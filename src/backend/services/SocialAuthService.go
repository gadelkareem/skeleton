package services

import (
    "fmt"
    "net/http"
    "strings"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "github.com/astaxie/beego/logs"
    "github.com/danilopolani/gocialite/structs"
    h "github.com/gadelkareem/go-helpers"
    "golang.org/x/oauth2"
    "gopkg.in/danilopolani/gocialite.v1"
)

type (
    SocialAuthService struct {
        us  *UserService
        jwt *JWTService
        s   SocialHandler
    }
    SocialHandler interface {
        New() *gocialite.Gocial
        Handle(state, code string) (*structs.User, *oauth2.Token, error)
    }
)

var providers map[string]map[string]string

func NewSocialAuthService(us *UserService, jwt *JWTService, s SocialHandler) *SocialAuthService {
    initProviderSecrets()
    return &SocialAuthService{us: us, jwt: jwt, s: s}
}

func (s *SocialAuthService) Redirect(provider string) (a *models.SocialAuth, err error) {
    var (
        p map[string]string
        k bool
    )
    if p, k = providers[provider]; !k {
        return nil, internal.ErrInvalidSocialProvider
    }

    a = new(models.SocialAuth)
    d := s.s.New().Driver(provider)
    if p["scope"] != "" {
        d = d.Scopes([]string{p["scope"]})
    }
    a.RedirectUri, err = d.Redirect(
        p["clientID"],
        p["clientSecret"],
        fmt.Sprintf("%s/auth/social-callback/?p=%s", kernel.App.FrontEndURL, provider),
    )

    return
}

func (s *SocialAuthService) Authenticate(r *models.SocialAuth) (*models.AuthToken, error) {
    u, _, err := s.s.Handle(r.State, r.Code)
    if err != nil {
        logs.Debug("Social login error: %s",err)
        if strings.Contains(err.Error(), "invalid CSRF token") {
            return nil, internal.ErrInvalidCSRFToken
        }
        return nil, err
    }
    if u.Email == "" {
        return nil, internal.ErrEmailRequired
    }
    m := models.NewUser()
    m.Email = u.Email
    existingUser, err := s.us.FindUser(m, false)
    if err != nil && err != internal.ErrNotFound {
        return nil, err
    }
    if existingUser != nil {
        if existingUser.AuthenticatorEnabled {
            return nil, internal.ErrAuthenticatorCodeMissing
        }
        if existingUser.MobileVerified {
            return nil, internal.ErrMobileCodeMissing
        }
        if existingUser.SocialLogin {
            return s.jwt.SocialToken(existingUser)
        }
        return nil, internal.Errorf(http.StatusBadRequest,
            fmt.Sprintf("The email %s already has an account with us."+
                "|Please login using your login credintials or reset your password to continue.", m.Email))
    }
    m.AvatarURL = u.Avatar
    setFullName(m, u)
    setAddress(m, u)
    err = s.setUsername(m, u)
    if err != nil {
        return nil, err
    }

    err = s.us.SignUpSocial(m)
    if err != nil {
        return nil, err
    }

    go s.us.UpdateLoginAt(m)

    return s.jwt.SocialToken(m)
}

func setFullName(m *models.User, u *structs.User) {
    m.FirstName = u.FirstName
    m.LastName = u.LastName
    if m.FirstName == "" && m.LastName == "" {
        ls := strings.Split(u.FullName, " ")
        if len(ls) > 0 {
            m.FirstName = ls[0]
        }
        if len(ls) > 1 {
            m.LastName = ls[1]
        }
    }
    return
}

func setAddress(m *models.User, u *structs.User) {
    // location: Alexandra, Egypt
    if l, k := u.Raw["location"]; k {
        ls := strings.Split(l.(string), ",")
        if len(ls) > 0 {
            m.Address.City = ls[0]
        }
        if len(ls) > 1 {
            m.Country = ls[1]
        }
    }
}

func (s *SocialAuthService) setUsername(m *models.User, u *structs.User) error {
    if l, k := u.Raw["login"]; k {
        m.Username = l.(string)
    }
    if m.Username == "" {
        ls := strings.Split(m.Email, "@")
        if len(ls) > 0 {
            m.Username = ls[0]
        }
    }

    for {
        existingUser, err := s.us.FindUser(&models.User{Username: m.Username}, false)
        if err != nil && err != internal.ErrNotFound {
            return err
        }
        if existingUser == nil {
            return nil
        }
        m.Username = fmt.Sprintf("%s_%d", m.Username, h.RandomNumber(10000, 100000))
    }
}

func initProviderSecrets() {
    providers = map[string]map[string]string{
        // https://github.com/settings/developers
        "github": {
            "clientID":     kernel.App.Config.String("social::githubClientID"),
            "clientSecret": kernel.App.Config.String("social::githubClientSecret"),
            "scope":        "",
        },
        // https://developers.google.com/identity/sign-in/web/sign-in
        "google": {
            "clientID":     kernel.App.Config.String("social::googleClientID"),
            "clientSecret": kernel.App.Config.String("social::googleClientSecret"),
            "scope":        "",
        },
        // https://developers.facebook.com/apps/
        "facebook": {
            "clientID":     kernel.App.Config.String("social::facebookClientID"),
            "clientSecret": kernel.App.Config.String("social::facebookClientSecret"),
            "scope":        "",
        },
        // https://www.linkedin.com/developers/apps/
        "linkedin": {
            "clientID":     kernel.App.Config.String("social::linkedinClientID"),
            "clientSecret": kernel.App.Config.String("social::linkedinClientSecret"),
            "scope":        "",
        },
    }
}
