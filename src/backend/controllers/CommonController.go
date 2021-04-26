package controllers

import (
    "net/http"

    "backend/models"
)

type CommonController struct {
    ApiController
}

// @router /contact [post]
func (c *CommonController) Contact() {
    r := new(models.Contact)
    c.parseRequest(r)
    c.validate(r)

    err := c.C.EmailService.Contact(r.Name, r.Email, r.Message, c.requestIP())
    c.handleError(err)

    c.SendStatus(http.StatusOK)
}
