package controllers

import (
    "net/http"

    "backend/internal"
    "backend/models"
)

type SubscriptionController struct {
    ApiController
}

// @router / [post]
func (c *SubscriptionController) CreateSubscription() {
    r := new(models.Subscription)
    c.parseRequest(r)
    c.AssertCustomer(r.CustomerID)
    c.validate(r)

    s, err := c.C.PaymentService.CreateSubscription(r)
    c.handleError(err)

    c.json(s)
}

// @router /:id [patch]
func (c *SubscriptionController) UpdateSubscription() {
    r := new(models.Subscription)
    c.parseRequest(r)
    c.AssertCustomerHasSubscription(r.CustomerID, r.ID)
    c.validate(r)

    s, err := c.C.PaymentService.UpdateSubscription(r)
    c.handleError(err)

    c.json(s)
}

// @router /:id [delete]
func (c *SubscriptionController) CancelSubscription(id string) {
    c.AssertCustomerHasSubscription(c.user.CustomerID, id)
    err := c.C.PaymentService.CancelSubscription(id, c.user.CustomerID)
    c.handleError(err)

    c.SendStatus(http.StatusNoContent)
}

func (c *SubscriptionController) AssertCustomerHasSubscription(customerID string, subscriptionID string) {
    c.AssertCustomer(customerID)
    if !c.C.PaymentService.CustomerHasSubscription(customerID, subscriptionID) {
        c.handleError(internal.ErrForbidden)
    }
}
