package models

import (
    "fmt"
    "strings"

    "backend/kernel"
)

type AuditLogRepository struct {
    *BaseRepository
}

func NewAuditLogRepository(db *kernel.PgDb, poolSize int) *AuditLogRepository {
    return &AuditLogRepository{BaseRepository: NewBaseRepository(db, poolSize)}
}

func (r *AuditLogRepository) Paginate(keyword string, sort map[string]string, offset, limit int, columns []string) (auditLogs []*AuditLog, size int, err error) {
    if keyword != "" {
        return r.Search(keyword, sort, offset, limit, columns)
    }
    q := r.Model(&auditLogs).Column(columns...).Order(r.formatSort(sort)).Offset(offset)

    if limit > 0 {
        q = q.Limit(limit)
    }
    size, err = q.SelectAndCount()

    return
}

func (r *AuditLogRepository) Search(keyword string, sort map[string]string, offset, limit int, columns []string) (auditLogs []*AuditLog, size int, err error) {
    q := r.Model(&auditLogs).Column(columns...)

    keyword = strings.ToLower(keyword)
    cl := "jsonb_values(log)"
    if _, k := sort["id"]; len(sort) < 2 && k {
        q = q.OrderExpr(fmt.Sprintf("%s <-> ? %s", cl, sort["id"]), keyword)
    } else {
        q = q.Order(r.formatSort(sort))
    }

    q = q.Where(fmt.Sprintf("%s ~* ?", cl), keyword).Offset(offset).Limit(limit)
    size, err = q.SelectAndCount()
    if err != nil {
        return
    }

    size = len(auditLogs)
    if size > kernel.ListLimit {
        size = 300
    }
    return
}

func (r *AuditLogRepository) Count() (size int, err error) {
    return r.PgDb.Model((*AuditLog)(nil)).Count()
}
