package controllers_test

import (
    "net/http"
    "reflect"
    "testing"

    "backend/kernel"
    "backend/models"
    "backend/tests"
    "github.com/brianvoe/gofakeit/v6"
    "github.com/stretchr/testify/assert"
)

func TestAuditLogController_GetAuditLogs(t *testing.T) {
    l := models.Log{}
    gofakeit.Struct(&l)
    tests.C.AuditLogService.AddAuditLog(l)
    gofakeit.Struct(&l)
    tests.C.AuditLogService.AddAuditLog(l)

    _, tk := tests.AdminWithToken(t)

    w, _ := tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       "/audit-logs",
        AuthToken: tk,
        Status:    http.StatusOK})
    auditLogs, payload := tests.ParseModels(t, reflect.TypeOf(new(models.AuditLog)), w.Body)
    assert.NotEmpty(t, auditLogs)

    totalAuditLogs, err := tests.C.AuditLogRepository.Count()
    tests.FailOnErr(t, err)
    count := totalAuditLogs
    if totalAuditLogs > kernel.ListLimit {
        count = kernel.ListLimit
    }
    assert.Len(t, auditLogs, count)
    assert.NotEmpty(t, auditLogs[0].(*models.AuditLog).ID)
    assert.NotEmpty(t, auditLogs[1].(*models.AuditLog).Log.Action)

    m := *payload.Meta
    assert.Equal(t, totalAuditLogs, int(m["page"].(map[string]interface{})["total"].(float64)))
    assert.Equal(t, kernel.ListLimit, int(m["page"].(map[string]interface{})["size"].(float64)))

    // test admin access only
    _, tk = tests.UserWithToken(true)
    w, _ = tests.ApiRequest(&tests.Request{T: t,
        Method:    http.MethodGet,
        URI:       "/audit-logs",
        AuthToken: tk,
        Status:    http.StatusForbidden})
}
