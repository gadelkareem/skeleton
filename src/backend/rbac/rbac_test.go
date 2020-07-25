package rbac_test

import (
    "fmt"
    "net/http"
    "testing"

    "backend/models"
    "backend/rbac"
    "github.com/stretchr/testify/assert"
)

func TestCanAccessRoute(t *testing.T) {
    r := rbac.New(false)
    f := func(u *models.User, route, method string, result bool) {
        b := r.CanAccessRoute(u, route, method)
        assert.Equal(t, result, b)
    }
    a := models.NewUser()
    a.MakeAdmin()
    u := models.NewUser()
    f(a, "/api/v1/users", http.MethodGet, true)
    f(u, "/api/v1/users", http.MethodGet, false)
    f(u, fmt.Sprintf("/api/v1/users/%d", u.ID), http.MethodGet, true)
    f(a, fmt.Sprintf("/api/v1/users/%d", u.ID), http.MethodGet, true)
}
