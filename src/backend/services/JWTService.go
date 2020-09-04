package services

import (
    "time"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "github.com/astaxie/beego/logs"
    "github.com/gbrlsnchs/jwt/v3"
)

type (
    JWTService struct {
        hmacKey *jwt.HMACSHA
        us      *UserService
        a       *AuthenticatorService
    }

    Payload struct {
        UserId     int64 `json:"i"`
        Type       int   `json:"t"`
        RememberMe bool  `json:"r"`
        jwt.Payload
    }
)

const (
    mainType = iota
    refreshType

    expiryExtend              = 15 * time.Minute
    notBeforeExtend           = 10 * time.Minute
    RefreshCookieMaxAge       = 365 * 24 * 60 * 60
    RefreshExpiryExtend       = 24 * time.Hour
    refreshNotBeforeExtend    = 20 * time.Hour
    RememberMeExpiryExtend    = 365 * 24 * time.Hour
    rememberMeNotBeforeExtend = 360 * 24 * time.Hour
    bearerSchema              = "Bearer "
    bearerSchemaLn            = len(bearerSchema)
    RefreshTokenCookieName    = "REFRESH_TOKEN"
)

func NewJWTService(k string, us *UserService) *JWTService {
    return &JWTService{hmacKey: jwt.NewHS256([]byte(k)), us: us}
}

func (s *JWTService) Authenticate(r *models.Login) (a *models.AuthToken, err error) {
    var u *models.User
    u, err = s.us.Authenticate(r.Username, r.Password)
    if err != nil {
        return
    }
    if u.AuthenticatorEnabled {
        if r.Code == "" {
            return nil, internal.ErrAuthenticatorCodeMissing
        }
        err = s.a.ValidateAuthenticator(u, r.Code)
        if err != nil {
            return
        }
    } else if u.MobileVerified && u.Mobile != "" && kernel.App.Config.DefaultBool("enableLoginSMS", true) {
        if r.Code == "" {
            err = s.us.SendVerifySMS(u)
            if err != nil {
                return nil, err
            }
            return nil, internal.ErrMobileCodeMissing
        }
        if !u.IsValidMobileCode(r.Code) {
            return nil, internal.ErrInvalidSMSCode
        }
    }
    go s.us.UpdateLoginAt(u)
    a = &models.AuthToken{UserId: u.ID}
    a.Token, err = s.Token(u)
    if err != nil {
        return
    }
    a.RefreshToken, err = s.RefreshToken(u, r.RememberMe)
    return
}

func (s *JWTService) SocialToken(u *models.User) (a *models.AuthToken, err error) {
    a = &models.AuthToken{UserId: u.ID}
    a.Token, err = s.Token(u)
    if err != nil {
        return
    }
    a.RefreshToken, err = s.RefreshToken(u, false)
    return
}

func (s *JWTService) Token(u *models.User) (string, error) {
    return s.token(u, mainType, expiryExtend, notBeforeExtend, false)
}

func (s *JWTService) RefreshToken(u *models.User, rememberMe bool) (string, error) {
    expiry := RefreshExpiryExtend
    expiryExtend := refreshNotBeforeExtend
    if rememberMe {
        expiry = RememberMeExpiryExtend
        expiryExtend = rememberMeNotBeforeExtend
    }
    return s.token(u, refreshType, expiry, expiryExtend, rememberMe)
}

func (s *JWTService) token(u *models.User, t int, expiryExtend, notBeforeExtend time.Duration, rememberMe bool) (string, error) {

    now := time.Now()
    pl := &Payload{
        Payload: jwt.Payload{
            ExpirationTime: jwt.NumericDate(now.Add(expiryExtend)),
            NotBefore:      jwt.NumericDate(now.Add(notBeforeExtend)),
            IssuedAt:       jwt.NumericDate(now),
            // Issuer:         uuid.New().String(),
            // Subject:        uuid.New().String(),
            // Audience: jwt.Audience{aud},
            // JWTID:          uuid.New().String(),
        },
        UserId:     u.ID,
        Type:       t,
        RememberMe: rememberMe,
    }

    return s.createToken(pl)
}

func (s *JWTService) Header(t string) string {
    return bearerSchema + t
}

func (s *JWTService) ParseHeader(t string) *Payload {
    if len(t) < bearerSchemaLn {
        return nil
    }
    pl := s.ParseToken(t[bearerSchemaLn:])
    if pl == nil || pl.Type != mainType {
        return nil
    }
    return pl
}

func (s *JWTService) ParseToken(t string) *Payload {
    if t == "" {
        return nil
    }
    var pl *Payload
    _, err := jwt.Verify([]byte(t), s.hmacKey, &pl, jwt.ValidateHeader)
    if err != nil || pl == nil {
        return nil
    }
    if pl.ExpirationTime.Before(time.Now()) {
        return nil
    }

    return pl
}

func (s *JWTService) AuthenticateRefreshToken(t string) (a *models.AuthToken, err error) {
    pl := s.ParseToken(t)
    if pl == nil || pl.Type != refreshType {
        return nil, internal.ErrInvalidJWTToken
    }
    var u *models.User
    u, err = s.us.UserById(pl.UserId)
    if err != nil {
        if err == internal.ErrNotFound {
            return nil, internal.ErrInvalidJWTToken
        }
        return
    }

    a = &models.AuthToken{UserId: u.ID}
    a.Token, err = s.Token(u)
    return
}

func (s *JWTService) createToken(pl *Payload) (string, error) {
    token, err := jwt.Sign(pl, s.hmacKey)
    if err != nil {
        logs.Error("Failed to create token: %s", err)
        return "", err
    }

    return string(token), nil
}
