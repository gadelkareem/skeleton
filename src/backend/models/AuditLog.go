package models

import (
    "fmt"
    "time"
)

type (
    AuditLog struct {
        Base      `pg:" inherit"`
        tableName struct{} `pg:"audit_log,alias:al" pg:" discard_unknown_columns"`

        ID        int64     `pg:"id,pk" jsonapi:"primary,audit_logs"`
        Log       Log       `pg:"log" jsonapi:"attr,log"`
        CreatedAt time.Time `pg:"created_at,type:TIMESTAMPTZ" jsonapi:"attr,created_at"`
    }
    Log struct {
        UserId    int64  `json:"user_id" jsonapi:"attr,user_id,omitempty"`
        AdminId   int64  `json:"admin_id" jsonapi:"attr,admin_id,omitempty"`
        UserEmail string `json:"user_email" jsonapi:"attr,user_email,omitempty" fake:"{person.first}####@{person.last}.{internet.domain_suffix}"`
        Username  string `json:"username" jsonapi:"attr,username,omitempty" jsonapi:"attr,username" fake:"{person.last}-###"`
        Action    string `json:"action" jsonapi:"attr,action,omitempty"`
        IP        string `json:"ip" jsonapi:"attr,ip,omitempty" fake:"##.##.##.#"`
        UserAgent string `json:"user_agent" jsonapi:"attr,user_agent,omitempty"`
        Request   string `json:"request" jsonapi:"attr,request,omitempty"`
    }
)

func NewAuditLog(l Log) *AuditLog {
    return &AuditLog{
        Base:      NewBaseModel(),
        Log:       l,
        CreatedAt: time.Now(),
    }
}

func (m *AuditLog) GetID() string {
    return fmt.Sprintf("%d", m.ID)
}

func (m *AuditLog) Sanitize() {
}
