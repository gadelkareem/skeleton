package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"backend/di"
	"backend/internal"
	"backend/kernel"
	"backend/limiter"
	"backend/models"
	"backend/rbac"
	"backend/routers"
	"backend/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gadelkareem/cachita"
	h "github.com/gadelkareem/go-helpers"
	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
)

var (
	C     *di.Container
	debug = flag.Int("debug", 0, "Show more debug info.")
)

func Bootstrap() {
	flag.Parse()
	kernel.Bootstrap()
	C = container()
	routers.InitRouters(C)
	gofakeit.Seed(time.Now().UnixNano())
}

func container() *di.Container {
	c := new(di.Container)
	initSMSService(c)
	initEmailService(c)
	initCacheService(c)
	initPaymentService(c)

	c.RateLimiter = limiter.NewRateLimiter(c.Cache, []limiter.Rate{{
		L:      &limiter.Limiter{Name: "test_unlimited", Limit: int64(kernel.MaxInt), TTL: time.Hour},
		Routes: [][]string{{"*", "/api/*"}},
	}})
	c.RBAC = rbac.New(false)
	c.InitTest()
	initSocialAuthService(c)
	// c.DB.EnableLogging()
	return c
}

func initCacheService(c *di.Container) {
	c.Cache = cachita.Memory()
	c.CacheService = services.NewCacheService(c.Cache, true)
}

func FailOnErr(t *testing.T, err error) {
	assert.Empty(t, err)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func AssertValidationError(t *testing.T, b *bytes.Buffer, k, v string) {
	r := ParseErrors(t, b)
	m := *r.Errors[0].Meta
	if m[k] == nil {
		FailOnErr(t, errors.New("invalid validation error"))
	}
	assert.Equal(t, v, m[k].(string))
}

func AssertError(t *testing.T, b *bytes.Buffer, e string) {
	r := ParseErrors(t, b)
	assert.Equal(t, e, r.Errors[0].Title)
}

func AssertErrorObj(t *testing.T, b *bytes.Buffer, err internal.Error) {
	r := ParseErrors(t, b)
	assert.Equal(t, r.Errors[0], err.Object())
}

func ParseErrors(t *testing.T, b *bytes.Buffer) *jsonapi.ErrorsPayload {
	r := new(jsonapi.ErrorsPayload)
	err := json.Unmarshal(b.Bytes(), &r)
	FailOnErr(t, err)
	return r
}

type Request struct {
	T         *testing.T
	Method    string
	URI       string
	Model     interface{}
	Query     url.Values
	AuthToken string
	Headers   map[string]string
	Cookies   []*http.Cookie
	Status    int
}

func ApiRequest(r *Request) (w *httptest.ResponseRecorder, rq *http.Request) {
	var (
		b   = bytes.NewBuffer(nil)
		err error
	)
	if r.Model != nil {
		err = jsonapi.MarshalPayload(b, r.Model)
		FailOnErr(r.T, err)
	}
	rq, err = http.NewRequest(r.Method, filepath.Join("/api/v1/", r.URI), b)
	FailOnErr(r.T, err)
	rq.Header.Set("Content-Type", jsonapi.MediaType)
	if r.Query != nil {
		rq.URL.RawQuery = r.Query.Encode()
	}
	if r.AuthToken != "" {
		rq.Header.Set("Authorization", C.JWTService.Header(r.AuthToken))
	}
	for k, v := range r.Headers {
		rq.Header.Set(k, v)
	}
	for _, c := range r.Cookies {
		rq.AddCookie(c)
	}

	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, rq)
	if *debug > 0 {
		b, err := ioutil.ReadAll(rq.Body)
		FailOnErr(r.T, err)
		logs.Trace("\n%s %s Code[%d]\nReq:%s\nRes:%s\n", rq.Method, rq.URL, w.Code, b, w.Body)
	}
	if r.Status != 0 {
		assert.Equal(r.T, r.Status, w.Code)
	}

	return
}

func User(active bool) *models.User {
	u := models.NewUser()
	gofakeit.Struct(&u)

	pass := u.Password
	u.HashPass()
	u.CleanStrings()
	u.GenerateEmailVerificationHash()
	if active {
		u.Activate()
	}
	u.MakeNew()
	err := C.UserService.Save(u)
	h.PanicOnError(err)

	u.Password = pass

	return u
}

func Token(u *models.User) string {
	tk, err := C.JWTService.Token(u)
	h.PanicOnError(err)

	return tk
}

func UserWithToken(active bool) (*models.User, string) {
	u := User(active)
	return u, Token(u)
}

func AdminWithToken(t *testing.T) (*models.User, string) {
	u, tk := UserWithToken(true)
	u.MakeAdmin()
	SaveUser(t, u)
	return u, tk
}

func ReadToken(t *testing.T, b []byte) *models.AuthToken {
	tk := new(models.AuthToken)
	err := jsonapi.UnmarshalPayload(bytes.NewBuffer(b), tk)
	FailOnErr(t, err)
	assert.NotEmpty(t, tk.Token)
	return tk
}

func Login(t *testing.T, username, password string) *httptest.ResponseRecorder {
	r := &models.Login{Username: username, Password: password}

	w, _ := ApiRequest(&Request{T: t, Method: http.MethodPost, URI: "/auth/token", Model: r})

	return w
}

func ParseModel(t *testing.T, m interface{}, b *bytes.Buffer) {
	err := jsonapi.UnmarshalPayload(b, m)
	FailOnErr(t, err)
}

func ParseModels(t *testing.T, r reflect.Type, bf *bytes.Buffer) ([]interface{}, *jsonapi.ManyPayload) {
	b := bf.Bytes()
	ms, err := jsonapi.UnmarshalManyPayload(bytes.NewBuffer(b), r)
	FailOnErr(t, err)

	payload := new(jsonapi.ManyPayload)
	err = json.Unmarshal(b, payload)
	FailOnErr(t, err)

	return ms, payload
}

func RefreshUser(t *testing.T, u *models.User, active bool) *models.User {
	nu, err := C.UserService.FindUser(u, active)
	FailOnErr(t, err)
	return nu
}

func SaveUser(t *testing.T, u *models.User) {
	err := C.UserService.Save(u)
	FailOnErr(t, err)
	C.CacheService.InvalidateModel(u)
}
