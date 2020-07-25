package services

import (
    "strings"

    "backend/models"
    "backend/utils/paginator"
    "github.com/astaxie/beego/logs"
)

type (
    AuditLogService struct {
        r *models.AuditLogRepository
    }
)

func NewAuditLogService(r *models.AuditLogRepository) *AuditLogService {
    return &AuditLogService{r: r}
}

func (s *AuditLogService) AddAuditLog(m models.Log) {
    a := models.NewAuditLog(m)
    err := s.r.Save(a)
    if err != nil {
        logs.Error("Could not add %s audit log: %s", m.Action, err)
    }
}

func (s *AuditLogService) PaginateAuditLogs(p *paginator.Paginator) (*paginator.Paginator, error) {
    var err error
    p.Sort = s.sanitizeSort(p.Sort)
    var ms []*models.AuditLog
    ms, p.Size, err = s.r.Paginate(p.Filter, p.Sort, p.Offset, p.Limit, []string{
        "id", "log", "created_at",
    })
    if err != nil {
        return p, err
    }
    for _, m := range ms {
        p.Models = append(p.Models, m)
    }

    return p, nil
}

func (s *AuditLogService) sanitizeSort(sort map[string]string) map[string]string {
    sort2 := make(map[string]string)
    for c, d := range sort {
        c = strings.ToLower(c)
        if c != "id" && c != "created_at"{
            c = "id"
        }
        sort2[c] = d
    }

    return sort2
}
