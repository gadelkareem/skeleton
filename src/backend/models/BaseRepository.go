package models

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"backend/kernel"
	"github.com/astaxie/beego/logs"
	"github.com/go-pg/pg/v9"
)

type BaseRepository struct {
	*kernel.PgDb

	PoolSize int
	poolMu   sync.RWMutex
	pool     map[interface{}]BaseInterface
}

func NewBaseRepository(db *kernel.PgDb, poolSize int) *BaseRepository {
	return &BaseRepository{
		PgDb:     db,
		PoolSize: poolSize,
		pool:     make(map[interface{}]BaseInterface),
	}
}

func (r *BaseRepository) multiInsert(pool map[interface{}]BaseInterface) error {
	var (
		updateDocs []BaseInterface
		newDocs    []BaseInterface
		err        error
	)
	for _, m := range pool {
		if m.IsNew() {
			newDocs = append(newDocs, m)
		} else {
			updateDocs = append(updateDocs, m)
		}
	}

	if len(newDocs) > 0 {
		bulkInsertResult, err1 := r.Model(&newDocs).OnConflict("DO NOTHING").Insert()
		if err1 != nil {
			logs.Error("Error inserting Models %v", err1)
			err = err1
		}
		logs.Alert("Bulk insert result: %+v", bulkInsertResult)
	}
	if len(updateDocs) > 0 {
		bulkUpdateResult, err2 := r.Model(&updateDocs).Update()
		if err2 != nil {
			logs.Error("Error updating Models %v", err2)
			err = err2
		}
		logs.Alert("Bulk update result: %+v", bulkUpdateResult)
	}

	return err
}

func (r *BaseRepository) Pool(m BaseInterface) (err error) {
	if r.PoolSize < 2 {
		return r.Save(m)
	}

	r.poolMu.Lock()
	defer r.poolMu.Unlock()
	if _, exists := r.pool[m.GetPoolIndex()]; exists {
		return nil
	}
	r.pool[m.GetPoolIndex()] = m
	if len(r.pool) > r.PoolSize {
		err = r.multiInsert(r.pool)
		r.pool = make(map[interface{}]BaseInterface)
	}
	return err
}

func (r *BaseRepository) IsInPool(index interface{}) bool {
	r.poolMu.RLock()
	defer r.poolMu.RUnlock()
	_, exists := r.pool[index]
	return exists
}

func (r *BaseRepository) Flush() error {
	var err error
	r.poolMu.RLock()
	defer r.poolMu.RUnlock()
	if len(r.pool) > 0 {
		logs.Debug("Flushing %d Models..", len(r.pool))
		err = r.multiInsert(r.pool)
		r.pool = make(map[interface{}]BaseInterface)
	}
	return err
}

func (r *BaseRepository) Save(m BaseInterface, columns ...string) error {

	m.SetUpdatedAt(time.Now())

	var err error
	if m.IsNew() {
		err = r.Insert(m)
		if err == nil {
			m.MakeOld()
		}
	} else {
		q := r.Model(m)
		if len(columns) > 0 {
			q = q.Column(columns...)
		}
		var rs pg.Result
		rs, err = q.Where("id = ?id").Update()
		if err == nil && rs.RowsAffected() == 0 {
			err = ErrNoResult
		}
	}

	return err
}

func (r *BaseRepository) formatSort(s map[string]string) string {
	t := ""
	for c, d := range s {
		t += fmt.Sprintf("%s %s,", c, d)
	}
	return strings.TrimRight(t, ",")
}
