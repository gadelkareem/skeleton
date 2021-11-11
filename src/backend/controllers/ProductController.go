package controllers

type ProductController struct {
    ApiController
}

// @router / [get]
func (c *ProductController) GetProducts() {
    ms, err := c.C.PaymentService.Products()
    c.handleError(err)

    c.json(ms)
}
