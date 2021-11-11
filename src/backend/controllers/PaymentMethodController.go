package controllers

import (
    "net/http"

    "backend/internal"
    "backend/models"
)

type PaymentMethodController struct {
    ApiController
}

// @router /:id [delete]
func (c *PaymentMethodController) DeletePaymentMethod(id string) {
    c.AssertCustomerHasPaymentMethod(c.user.CustomerID, id)
    err := c.C.PaymentService.DeletePaymentMethod(id, c.user.CustomerID)
    c.handleError(err)

    c.SendStatus(http.StatusNoContent)
}

// @router / [post]
func (c *PaymentMethodController) CreatePaymentMethod() {
    r := new(models.PaymentMethod)
    c.parseRequest(r)
    c.validate(r)
    pm, err := c.C.PaymentService.AttachPaymentMethod(c.user.CustomerID, r)
    c.handleError(err)

    c.json(pm)
}

func (c *PaymentMethodController) AssertCustomerHasPaymentMethod(customerID string, paymentMethodID string) {
    if !c.C.PaymentService.CustomerHasPaymentMethod(customerID, paymentMethodID) {
        c.handleError(internal.ErrForbidden)
    }
}
