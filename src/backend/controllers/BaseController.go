package controllers

import (
    "net"
    "net/http"

    "backend/di"
    "backend/kernel"
    "backend/models"
    "github.com/astaxie/beego/logs"
    "github.com/gadelkareem/go-helpers"
)

type (
    BaseController struct {
        kernel.Controller
        C *di.Container

        user *models.User
    }
)

func (c *BaseController) Prepare() {
    domain := c.Ctx.Input.Domain()
    if !kernel.IsIPTrusted(c.requestIP()) {
        if !kernel.IsHostAllowed(domain) {
            c.Redirect(kernel.App.FrontEndURL+c.Ctx.Input.URI(), http.StatusMovedPermanently)
            return
        }
        c.Ctx.Request.Header.Del("X-Forwarded-For")
        c.Ctx.Request.Header.Del("X-Real-IP")
    }
}

func (c *BaseController) logOnError(err error) {
    if err != nil {
        logs.Error("Error: %s", err)
    }
}

func (c *BaseController) log(status int) {
    if b, _ := kernel.App.Config.Bool("AccessLogs"); !b {
        return
    }
    logs.AccessLog(&logs.AccessLogRecord{
        RemoteAddr:    c.requestIP(),
        RequestMethod: c.Ctx.Input.Method(),
        Request:       c.Ctx.Request.URL.String(),
        Host:          c.Ctx.Request.Host,
        HTTPReferrer:  c.Ctx.Request.Referer(),
        HTTPUserAgent: c.Ctx.Request.UserAgent(),
        Status:        status,
    }, "")
}

func (c *BaseController) readString(key string, def ...string) string {
    return h.CleanString(c.GetString(key, def...))
}

func (c *BaseController) requestIP() string {
    ips := c.Ctx.Input.Proxy()
    if len(ips) > 0 && ips[0] != "" {
        rip, _, err := net.SplitHostPort(ips[0])
        if err != nil {
            rip = ips[len(ips)-1]
        }
        return rip
    }
    if ip, _, err := net.SplitHostPort(c.Ctx.Input.Context.Request.RemoteAddr); err == nil {
        return ip
    }
    return c.Ctx.Input.Context.Request.RemoteAddr
}

func (c *BaseController) requestUserAgent() string {
    return c.Ctx.Request.UserAgent()
}
