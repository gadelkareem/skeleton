package controllers

import (
    "net/http"

    "backend/models"
)

type PaymentController struct {
    ApiController
}

// @router /webhook/ [post]
func (c *PaymentController) Webhook() {
    r := &models.PaymentEvent{
        Payload:   c.Ctx.Input.RequestBody,
        Signature: c.Ctx.Request.Header.Get("Stripe-Signature"),
    }
    c.validate(r)

    err := c.C.PaymentService.Webhook(r)
    c.handleError(err)

    c.SendStatus(http.StatusOK)
}