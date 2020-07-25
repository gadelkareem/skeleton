package controllers

import (
    "backend/utils/paginator"
)

type AuditLogController struct {
    ApiController
}

// @router / [get]
func (c *AuditLogController) GetAuditLogs() {
    page := map[string]int{"size": 10, "after": 1}
    err := c.Ctx.Input.Bind(&page, "page")
    c.logOnError(err)
    sort, filter := c.readString("sort"), c.readString("filter")

    p, err := c.C.AuditLogService.PaginateAuditLogs(paginator.NewPaginator(page, sort, filter))
    c.handleError(err)

    c.jsonMany(p)
}
