package models

import (
	"errors"
	"fmt"
	"strings"

	"backend/internal"
	"backend/kernel"
	"github.com/astaxie/beego/logs"
	h "github.com/gadelkareem/go-helpers"
	"github.com/go-pg/pg/v9"
	"github.com/ttacon/libphonenumber"
)

type UserRepository struct {
	*BaseRepository
}

var (
	ErrNoResult = errors.New("query had no affected result")
)

func NewUserRepository(db *kernel.PgDb, poolSize int) *UserRepository {
	return &UserRepository{BaseRepository: NewBaseRepository(db, poolSize)}
}

func (r *UserRepository) FindUnfilteredById(id int64) (m *User, err error) {
	m = NewEmptyUser()
	q := r.Model(m).Limit(1).
		Where("id = ?", id)

	err = q.Select()
	m, err = r.cleanErr(m, err)
	return
}

func (r *UserRepository) FindById(id int64, checkActive bool) (m *User, err error) {
	m = NewEmptyUser()
	q := r.Model(m).Limit(1).
		Where("id = ?", id).Where("deleted_at IS NULL")
	if checkActive {
		q = q.Where("active = ?", true)
	}
	err = q.Select()
	m, err = r.cleanErr(m, err)
	return
}

func (r *UserRepository) FindByUsername(username string, checkActive bool) (m *User, err error) {
	m = NewEmptyUser()
	q := r.Model(m).Limit(1).
		Where("username = ?", username).Where("deleted_at IS NULL")

	if checkActive {
		q = q.Where("active = ?", true)
	}
	err = q.Select()
	m, err = r.cleanErr(m, err)
	return
}

func (r *UserRepository) ValidateUser(m *User) error {
	m.CleanStrings()

	v := kernel.Validation()
	b, err := v.Valid(m)
	if err != nil {
		return err
	}
	if !b {
		vErrs := make(map[string]interface{})
		for _, e := range v.Errors {
			vErrs[e.Key] = e.Error()
		}
		return internal.ValidationErrors(vErrs)
	}
	if m.Mobile != "" {
		n, err := libphonenumber.Parse(m.Mobile, "")
		if n != nil {
			b = libphonenumber.IsValidNumber(n)
		}
		if err != nil || !b {
			return internal.ValidationErrors(map[string]interface{}{"Mobile": "Invalid mobile number."})
		}
	}

	return nil
}

func (r *UserRepository) CheckPassword(username, password string) (u *User, err error) {
	username = strings.ToLower(h.TrimLine(username))
	if username == "" || password == "" {
		return nil, internal.ErrInvalidPass
	}
	u, _ = r.FindByUsername(username, true)

	if u == nil || !u.IsValidPass(password) {
		return nil, internal.ErrInvalidPass
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string, checkActive bool) (m *User, err error) {
	m = NewEmptyUser()
	q := r.Model(m).Limit(1).
		Where("email = ?", strings.ToLower(email)).Where("deleted_at IS NULL")
	if checkActive {
		q = q.Where("active = ?", true)
	}
	err = q.Select()
	m, err = r.cleanErr(m, err)
	return
}

func (r *UserRepository) cleanErr(m *User, err error) (*User, error) {
	if m == nil || m.ID == 0 {
		m = nil
	}
	if err != nil {
		if pgErr, ok := err.(pg.Error); ok {
			if pgErr.IntegrityViolation() {
				switch pgErr.Field('n') {
				case "s_users_unique_email_idx":
					return m, internal.ErrEmailExists
				case "s_users_unique_username_idx":
					return m, internal.ErrUsernameExists
				}
			}
		}
		if err == pg.ErrNoRows {
			return nil, internal.ErrNotFound
		}
		logs.Error("error operating on user: %s", err)
	}
	return m, err
}

func (r *UserRepository) Save(m BaseInterface, columns ...string) error {

	err := r.BaseRepository.Save(m, columns...)

	_, err = r.cleanErr(nil, err)

	return err
}

func (r *UserRepository) Paginate(keyword string, sort map[string]string, offset, limit int, columns []string) (users []*User, size int, err error) {
	if keyword != "" {
		return r.Search(keyword, sort, offset, limit, columns)
	}
	q := r.Model(&users).Column(columns...).Order(r.formatSort(sort)).Offset(offset)

	if limit > 0 {
		q = q.Limit(limit)
	}
	size, err = q.SelectAndCount()

	return
}

func (r *UserRepository) Search(keyword string, sort map[string]string, offset, limit int, columns []string) (users []*User, size int, err error) {
	q := r.Model(&users).Column(columns...)

	keyword = strings.ToLower(keyword)
	cl := "(username || ' ' || email || ' ' || coalesce(first_name,'') || ' ' || coalesce(last_name,''))"
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

	size = len(users)
	if size > kernel.ListLimit {
		size = 300
	}
	return
}

func (r *UserRepository) Count() (size int, err error) {
	return r.PgDb.Model((*User)(nil)).Count()
}
