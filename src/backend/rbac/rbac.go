package rbac

import (
    "fmt"
    "strings"

    "backend/models"
    "backend/utils"
    "github.com/astaxie/beego/logs"
)

const (
    RoleGuest = "guest"
    RoleUser  = "user"
    RoleAdmin = "admin"
)

var routes = map[string][][]string{
    RoleUser: {
        {"GET|POST|PATCH|DELETE", "/api/v1/users/{userID}*"},
        {"POST", "/api/v1/subscriptions"},
        {"PATCH|DELETE", "/api/v1/subscriptions/*"},
        {"GET|POST", "/api/v1/invoices*"},
        {"POST", "/api/v1/payment-methods"},
        {"DELETE", "/api/v1/payment-methods/*"},
        {"GET|PATCH", "/api/v1/customers/{customerID}*"},
        {"DELETE", "/api/v1/customers/{customerID}/subscriptions*"},
        {"GET|POST|DELETE", "/api/v1/customers/{customerID}/payment-methods*"},
    },
    RoleGuest: {
        {"GET|POST", "/api/v1/auth/*"},
        {"POST", "/api/v1/users"},
        {"POST", "/api/v1/users/verify-email"},
        {"POST", "/api/v1/users/reset-password"},
        {"POST", "/api/v1/users/disable-authenticator"},
        {"POST", "/api/v1/users/forgot-password"},
        {"POST", "/api/v1/users/recovery-questions"},
        {"POST", "/api/v1/users/disable-mfa"},
        {"POST", "/api/v1/common/contact"},
        {"GET", "/api/v1/products"},
        {"POST", "/api/v1/payments/webhook"},
    },
    RoleAdmin: {
        {"*", "/api/v1/*"},
    },
}

var permissions = map[string][]string{
    RoleUser:  {},
    RoleGuest: {},
    RoleAdmin: {"*"},
}

type RBAC struct {
    routes      map[string][][]string
    permissions map[string][]string
    debug       bool
}

func New(b bool) *RBAC {
    return &RBAC{routes: routes, permissions: permissions, debug: b}
}

func (r *RBAC) CanAccessRoute(u *models.User, route, method string) bool {
    roles := roles(u)
    for _, role := range roles {
        for _, rt := range r.routes[role] {
            if utils.HasMethod(rt[0], method) && utils.HasRoute(u, rt[1], route) {
                r.log("Access granted for %s %s %s", method, route, role)
                return true
            }
        }
    }
    r.log("Access denied for %s %s %s", method, route, roles)
    return false
}

func (r *RBAC) HasPermission(u *models.User, permission string) bool {
    roles := roles(u)
    for _, role := range roles {
        for _, rule := range r.permissions[role] {
            if hasPermission(u, rule, permission) {
                r.log("Access granted for %s %s", permission, role)
                return true
            }
        }
    }
    r.log("Access denied for %s %v", permission, roles)
    return false
}

func (r *RBAC) log(f interface{}, v ...interface{}) {
    if r.debug {
        logs.Debug(f, v...)
    }
}

func hasPermission(u *models.User, rule, p string) bool {
    if rule == "*" {
        return true
    }
    if u != nil {
        rule = strings.Replace(rule, "{userID}", fmt.Sprintf("%d", u.ID), 1)
        rule = strings.Replace(rule, "{customerID}", fmt.Sprintf("%s", u.CustomerID), 1)
    }
    if rule == p {
        return true
    }

    if strings.HasSuffix(rule, "*") {
        rule = strings.TrimRight(rule, "*")
        if strings.HasPrefix(p, rule) {
            return true
        }
    }
    return false
}

func roles(u *models.User) []string {
    r := []string{RoleGuest}
    if u == nil {
        return r
    }
    r = append(r, RoleUser)
    r = append(r, u.Roles...)
    return r
}

var Debug = true
