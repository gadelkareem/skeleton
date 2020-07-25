package controllers

import (
    "fmt"
    "net/http"

    "backend/internal"
    "backend/kernel"
    "backend/models"
    "backend/services"
)

type AuthController struct {
    ApiController
}

// @router /token [post]
func (c *AuthController) Token() {
    r := new(models.Login)
    c.parseRequest(r)

    a, err := c.C.JWTService.Authenticate(r)
    c.handleError(err)

    if c.IsAjax() {
        c.addCookie(a.RefreshToken, services.RefreshCookieMaxAge)
    }

    c.json(a)
}

// @router /refresh-token [post]
func (c *AuthController) RefreshToken() {
    r := new(models.AuthToken)
    c.parseRequest(r)

    a, err := c.C.JWTService.AuthenticateRefreshToken(r.RefreshToken)
    c.handleError(err)

    c.json(a)
}

// @router /refresh-cookie [post]
func (c *AuthController) RefreshCookie() {
    if !c.IsAjax() {
        c.handleError(internal.ErrForbidden)
    }
    rt := c.Ctx.GetCookie(services.RefreshTokenCookieName)
    a, err := c.C.JWTService.AuthenticateRefreshToken(rt)
    c.handleError(err)

    c.json(a)
}

// @router /logout [get]
func (c *AuthController) Logout() {
    if !c.IsAjax() {
        c.handleError(internal.ErrForbidden)
    }
    c.addCookie("", -1)

    c.SendStatus(http.StatusOK)
}

// @router /social/redirect [post]
func (c *AuthController) SocialRedirect() {
    r := new(models.SocialAuth)
    c.parseRequest(r)

    a, err := c.C.SocialAuthService.Redirect(r.Provider)
    c.handleError(err)

    c.json(a)
}

// @router /social/callback [post]
func (c *AuthController) SocialCallback() {
    r := new(models.SocialAuth)
    c.parseRequest(r)

    a, err := c.C.SocialAuthService.Authenticate(r)
    c.handleError(err)

    if c.IsAjax() {
        c.addCookie(a.RefreshToken, services.RefreshCookieMaxAge)
    }

    c.json(a)
}

func (c *AuthController) addCookie(v string, exp int) {
    // max age time, path,domain, secure and httponly
    c.Ctx.SetCookie(services.RefreshTokenCookieName,
        v,
        exp,
        fmt.Sprintf("%s/auth/refresh-cookie", kernel.App.APIPath),
        kernel.App.Host,
        !kernel.IsDev(),
        true)
}
