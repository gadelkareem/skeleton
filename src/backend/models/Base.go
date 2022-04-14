package models

import (
	"time"
)

type (
	Base struct {
		CreatedAt time.Time `pg:"created_at,type:TIMESTAMPTZ" json:"-"`
		UpdatedAt time.Time `pg:"updated_at,type:TIMESTAMPTZ" json:"-"`
		new       bool      `pg:"-"`
	}
	BaseInterface interface {
		IsNew() bool
		MakeOld()
		SetUpdatedAt(v time.Time)
		GetPoolIndex() interface{}
		Sanitize()
		GetID() string
	}
)

func NewBaseModel() Base {
	timeNow := time.Now()
	return Base{
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		new:       true,
	}
}

func (m *Base) IsNew() bool {
	return m.new
}

func (m *Base) MakeNew() {
	m.new = true
}

func (m *Base) MakeOld() {
	m.new = false
}

func (m *Base) GetPoolIndex() interface{} {
	return nil
}

func (m *Base) SetUpdatedAt(v time.Time) {
	m.UpdatedAt = v
}
