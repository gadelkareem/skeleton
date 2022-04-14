package controllers

import (
	"backend/kernel"
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
	c.AssertCustomer(r.ID)

	r, err := c.C.PaymentService.UpdateCustomer(r, nil)
	c.handleError(err)

	c.json(r)
}

// @router /:id/payment-methods [get]
func (c *CustomerController) ListPaymentMethods(id string) {
	c.AssertCustomer(id)
	b, _ := c.GetBool("resetCache")
	p := c.C.PaymentService.PaginatePaymentMethods(id, b, c.paginator(kernel.ListLimit))

	c.jsonMany(p)
}

// @router /:id/subscription [get]
func (c *CustomerController) CustomerSubscription(id string) {
	c.AssertCustomer(id)
	s, err := c.C.PaymentService.ActiveSubscription(id)
	c.handleError(err)

	c.json(s)
}

// @router /:id/invoices [get]
func (c *CustomerController) CustomerInvoices(id string) {
	c.AssertCustomer(id)
	p, err := c.C.PaymentService.PaginateInvoices(id, c.paginator(kernel.ListLimit))
	c.handleError(err)

	c.jsonMany(p)
}
