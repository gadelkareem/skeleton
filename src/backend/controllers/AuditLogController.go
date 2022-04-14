package controllers

import (
	"backend/kernel"
)

type AuditLogController struct {
	ApiController
}

// @router / [get]
func (c *AuditLogController) GetAuditLogs() {
	p, err := c.C.AuditLogService.PaginateAuditLogs(c.paginator(kernel.ListLimit))
	c.handleError(err)

	c.jsonMany(p)
}
