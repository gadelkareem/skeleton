package services

import (
    "bytes"
    "encoding/base64"
    "image/png"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "github.com/pquerna/otp/totp"
)

type (
    AuthenticatorService struct {
        us *UserService
    }
)

func NewAuthenticatorService(us *UserService) *AuthenticatorService {
    return &AuthenticatorService{us: us}
}

func (s *AuthenticatorService) Process(m *models.User, r *models.Authenticator) error {
    if m.AuthenticatorEnabled {
        return s.DisableAuthenticator(m, r)
    }
    return s.EnableAuthenticator(m, r)
}

func (s *AuthenticatorService) Generate(m *models.User) (*models.Authenticator, error) {
    if !m.RecoveryQuestionsSet {
        return nil, internal.ErrRecoveryQuestionNotSet
    }
    if m.AuthenticatorEnabled {
        return nil, internal.ErrAuthenticatorAlreadyEnabled
    }
    key, err := totp.Generate(totp.GenerateOpts{
        Issuer:      kernel.App.Host,
        AccountName: m.Email,
    })
    if err != nil {
        return nil, err
    }
    var buf bytes.Buffer
    img, err := key.Image(200, 200)
    if err != nil {
        return nil, err
    }
    err = png.Encode(&buf, img)
    if err != nil {
        return nil, err
    }
    a := new(models.Authenticator)
    a.Seed = key.Secret()
    a.URL = key.URL()
    a.Image = base64.StdEncoding.EncodeToString(buf.Bytes())
    m.AuthenticatorSecret = key.Secret()
    err = s.us.Save(m)
    if err != nil {
        return nil, err
    }
    return a, nil
}

func (s *AuthenticatorService) DisableMFA(r *models.DisableMFA) error {
    u, err := s.us.Authenticate(r.Username, r.Password)
    if err != nil {
        return err
    }
    // b, err := s.us.ValidateRecoveryQuestions(u, r.RecoveryQuestions)
    // if err != nil {
    //     return err
    // }
    // if !b {
    //     return internal.ErrBadRecoveryAnswers
    // }
    u.DisableAuthenticator()
    u.UnVerifyMobile()

    return s.us.Save(u)
}

func (s *AuthenticatorService) DisableAuthenticator(m *models.User, r *models.Authenticator) error {
    if !totp.Validate(r.Code, m.AuthenticatorSecret) {
        return internal.ErrInvalidAuthenticatorCode
    }
    m.DisableAuthenticator()
    return s.us.Save(m)
}

func (s *AuthenticatorService) EnableAuthenticator(m *models.User, r *models.Authenticator) error {
    if !m.RecoveryQuestionsSet {
        return internal.ErrRecoveryQuestionNotSet
    }
    if !totp.Validate(r.Code, m.AuthenticatorSecret) {
        return internal.ErrInvalidAuthenticatorCode
    }
    m.AuthenticatorEnabled = true
    err := s.us.Save(m)
    if err != nil {
        return err
    }
    return nil
}

func (s *AuthenticatorService) ValidateAuthenticator(m *models.User, code string) error {
    if !totp.Validate(code, m.AuthenticatorSecret) {
        return internal.ErrInvalidAuthenticatorCode
    }
    return nil
}
