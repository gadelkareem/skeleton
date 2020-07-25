package tests

import (
    "net/http"

    "backend/di"
    "backend/services"
)

type smsCl struct {
}

func initSMSService(c *di.Container) {
    c.SMSService = services.NewSMSService(new(smsCl))
}

func (c *smsCl) Do(*http.Request) (*http.Response, error) {
    r := new(http.Response)
    r.StatusCode = http.StatusOK
    return r, nil
}
