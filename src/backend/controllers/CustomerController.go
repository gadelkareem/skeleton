package controllers

import (
	"backend/models"
)

type CustomerController struct {
	ApiController
}

//
// // @router /:id/setup-intent [get]
// func (c *CustomerController) SetupIntent(id string) {
//     r, err := c.C.PaymentService.SetupIntent(id)
//     c.handleError(err)
//
//     c.json(r)
// }

// @router /:id [PATCH]
func (c *CustomerController) UpdateCustomer() {
	r := new(models.Customer)
	c.parseRequest(r)

	r, err := c.C.PaymentService.UpdateCustomer(r, nil)
	c.handleError(err)

	c.json(r)
}

// @router /:id/payment-methods [get]
func (c *CustomerController) ListPaymentMethods(id string) {
	b, _ := c.GetBool("resetCache")
	ms := c.C.PaymentService.PaymentMethods(id, b)

	c.json(ms)
}

// @router /:id/subscription [get]
func (c *CustomerController) CustomerSubscription(id string) {
	s, err := c.C.PaymentService.ActiveSubscription(id)
	c.handleError(err)

	c.json(s)
}
