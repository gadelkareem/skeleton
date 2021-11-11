package controllers

import (
    "backend/models"
)

type InvoiceController struct {
    ApiController
}

// @router /upcoming [get]
func (c *InvoiceController) UpcomingInvoice() {

    r := &models.Subscription{
        ID:         c.GetString("id"),
        CustomerID: c.GetString("customer_id"),
        PriceID:    c.GetString("price_id"),
        ItemID:     c.GetString("item_id"),
    }
    c.AssertCustomer(r.CustomerID)
    c.validate(r)

    in, err := c.C.PaymentService.UpcomingInvoice(r)
    c.handleError(err)

    c.json(in)
}
