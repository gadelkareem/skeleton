package utils

import (
    "fmt"
    "strings"

    "backend/models"
    "github.com/mitchellh/hashstructure"
)

func HasRoute(u *models.User, rule, route string) bool {
    route = strings.TrimRight(route, "/")
    if u != nil {
        rule = strings.Replace(rule, "{userId}", fmt.Sprintf("%d", u.ID), 1)
    }
    if rule == route {
        return true
    }

    if strings.HasSuffix(rule, "*") {
        rule = strings.TrimRight(rule, "*")
        if strings.HasPrefix(route, rule) {
            return true
        }
    }
    return false
}

func HasMethod(methods, method string) bool {
    return methods == "*" || strings.Contains(methods, method)
}

func Hash(i ...interface{}) (string, error) {
    h, err := hashstructure.Hash(i, nil)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("%d", h), nil
}
