package controllers

import "backend/kernel"

type ProductController struct {
	ApiController
}

// @router / [get]
func (c *ProductController) GetProducts() {
	p, err := c.C.PaymentService.PaginateProducts(c.paginator(kernel.ListLimit))
	c.handleError(err)

	c.jsonMany(p)
}
