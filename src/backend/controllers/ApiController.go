package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"backend/di"
	"backend/internal"
	"backend/kernel"
	"backend/models"
	"backend/utils/paginator"
	"github.com/astaxie/beego/logs"
	"github.com/google/jsonapi"
)

type (
	ApiController struct {
		BaseController
	}
)

func NewApiController(c *di.Container) ApiController {
	return ApiController{BaseController: BaseController{C: c}}
}

func (c *ApiController) Prepare() {
	c.EnableRender = false
	c.BaseController.Prepare()
	kernel.SetCORS(c.Ctx)
	c.setUser()

	u := c.Ctx.Request.URL.String()
	method := c.Ctx.Request.Method
	c.rbac(u, method)
	c.rateLimit(u, method)
}

func (c *ApiController) rateLimit(u, method string) {
	ip := c.requestIP()
	b, err := c.C.RateLimiter.IsRateExceeded(c.user, ip, u, method)
	c.handleError(err)
	if b {
		logs.Error("rate exceeded for %s %s %s", ip, method, c.Ctx.Request.URL)
		c.handleError(internal.ErrTooManyRequests)
	}
}

func (c *ApiController) rbac(u, method string) {
	if !c.C.RBAC.CanAccessRoute(c.user, u, method) {
		if c.user == nil {
			c.handleError(internal.ErrInvalidJWTToken)
		}
		c.handleError(internal.ErrForbidden)
	}
}

func (c *ApiController) parseRequest(m interface{}) {
	b := bytes.NewBuffer(c.Ctx.Input.RequestBody)
	if err := jsonapi.UnmarshalPayload(b, m); err != nil {
		logs.Error("Error parsing request %s err: %s", c.Ctx.Request.URL, err)
		c.handleError(internal.ErrInvalidRequest)
	}
}

func (c *ApiController) validate(m interface{}) {
	v := kernel.Validation()
	b, err := v.Valid(m)
	if err != nil {
		c.handleError(err)
	}
	if !b {
		vErrs := make(map[string]interface{})
		for _, e := range v.Errors {
			vErrs[e.Key] = e.Error()
		}
		c.handleError(internal.ValidationErrors(vErrs))
	}

	return
}

func (c *ApiController) parseJSONRequest(m interface{}) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, m); err != nil {
		logs.Error("Error parsing request %s err: %s", c.Ctx.Request.URL, err)
		c.handleError(internal.ErrInvalidRequest)
	}
}

func (c *ApiController) jsonMany(p *paginator.Paginator) {
	c.Ctx.Output.Header("Content-Type", jsonapi.MediaType)
	c.Ctx.Output.SetStatus(http.StatusOK)

	for _, m := range p.Models {
		if ml, k := m.(models.BaseInterface); k {
			ml.Sanitize()
		}
	}
	pl, err := jsonapi.Marshal(p.Models)
	c.handleError(err)
	py := pl.(*jsonapi.ManyPayload)
	p.Optimize()
	py.Links = p.Links()
	py.Meta = p.Meta()

	b := bytes.NewBuffer(nil)
	err = json.NewEncoder(b).Encode(py)
	c.handleError(err)

	err = c.Ctx.Output.Body(b.Bytes())
	c.handleError(err)

	c.StopRun()
}

func (c *ApiController) json(m interface{}) {
	c.Ctx.Output.Header("Content-Type", jsonapi.MediaType)
	c.Ctx.Output.SetStatus(http.StatusOK)

	if ml, k := m.(models.BaseInterface); k {
		ml.Sanitize()
	}

	b := bytes.NewBuffer(nil)
	err := jsonapi.MarshalPayload(b, m)
	c.handleError(err)

	err = c.Ctx.Output.Body(b.Bytes())
	c.handleError(err)

	c.StopRun()
}

func (c *ApiController) handleError(err error) {
	if err == nil {
		return
	}
	internalErr, k := err.(internal.Error)
	if !k || internalErr == nil {
		logs.Error("Error: %s", err)
		internalErr = internal.ErrInternalError
	}
	c.Ctx.Output.Header("Content-Type", jsonapi.MediaType)
	c.Ctx.Output.SetStatus(internalErr.Status())
	b := bytes.NewBuffer(nil)
	e := jsonapi.MarshalErrors(b, []*jsonapi.ErrorObject{internalErr.Object()})
	if e != nil {
		logs.Error("Error unmarshal errors: %s", e)
	}
	e = c.Ctx.Output.Body(b.Bytes())
	if e != nil {
		logs.Error("Error writing body: %s", e)
	}
	go c.log(internalErr.Status())
	c.StopRun()
}

func (c *ApiController) SendStatus(s int) {
	c.Ctx.ResponseWriter.WriteHeader(s)
	go c.log(s)
	c.StopRun()
}

func (c *ApiController) setUser() {
	tk := c.Ctx.Input.Header("Authorization")
	if tk == "" {
		return
	}

	p := c.C.JWTService.ParseHeader(tk)
	if p == nil {
		return
	}

	var err error
	c.user, err = c.C.UserService.UserById(p.UserId)
	if err != nil {
		logs.Error("Could not find user: %s", err)
	}
}

func (c *ApiController) auditLog(l models.Log) {
	l.IP = c.requestIP()
	l.UserAgent = c.requestUserAgent()
	if l.Request == "" {
		l.Request = string(c.Ctx.Input.RequestBody)
	}
	if l.UserID != 0 && c.user != nil && l.UserID != c.user.ID {
		l.AdminID = c.user.ID
	}
	if l.UserID == 0 && c.user != nil {
		l.UserID = c.user.ID
	}
	c.C.AuditLogService.AddAuditLog(l)
}

func (c *ApiController) paginator(limit int) *paginator.Paginator {
	page := map[string]int{"size": limit, "after": 1}
	err := c.Ctx.Input.Bind(&page, "page")
	c.logOnError(err)
	return paginator.NewPaginator(page, c.readString("sort"), c.readString("filter"), c.Ctx.Input.URI())
}

func (c *ApiController) AssertCustomer(id string) {
	if id != c.user.CustomerID {
		c.handleError(internal.ErrForbidden)
	}
}
