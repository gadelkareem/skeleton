package controllers

import (
    "backend/internal"
    "backend/kernel"
)

type (
    ErrorController struct {
        ApiController
    }
)

func (c *ErrorController) Prepare() {
    c.EnableRender = false
    c.setUser()
    kernel.SetCORS(c.Ctx)
}

func (c *ErrorController) Error401() {
    c.handleError(internal.ErrForbidden)
}

func (c *ErrorController) Error400() {
    c.handleError(internal.ErrInvalidRequest)
}

func (c *ErrorController) Error403() {
    c.handleError(internal.ErrForbidden)
}

func (c *ErrorController) Error404() {
    c.handleError(internal.ErrNotFound)
}

func (c *ErrorController) Error429() {
    c.handleError(internal.ErrTooManyRequests)
}

func (c *ErrorController) Error500() {
    c.handleError(internal.ErrInternalError)
}

func (c *ErrorController) Error503() {
    c.handleError(internal.ErrInternalError)
}

func (c *ErrorController) ErrorDb() {
    c.handleError(internal.ErrInternalError)
}
