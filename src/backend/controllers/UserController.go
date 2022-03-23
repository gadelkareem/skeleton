package controllers

import (
	"net/http"

	"backend/kernel"
	"backend/models"
)

type UserController struct {
	ApiController
}

// @router / [post]
func (c *UserController) SignUp() {
	u := models.NewUser()
	c.parseRequest(u)

	err := c.C.UserService.SignUp(u)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: u.ID, Action: "SignUp", Request: "SKIP"})

	c.SendStatus(http.StatusCreated)
}

// @router /verify-email [post]
func (c *UserController) VerifyEmail() {
	r := new(models.VerifyEmail)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.VerifyEmail(r.Email, r.Token)
	c.handleError(err)

	go c.auditLog(models.Log{UserEmail: r.Email, Action: "VerifyEmail"})

	c.SendStatus(http.StatusOK)
}

// @router /forgot-password [post]
func (c *UserController) ForgotPassword() {
	r := new(models.ResetPassword)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.ForgotPassword(r.Email, r.Username)
	c.handleError(err)

	go c.auditLog(models.Log{UserEmail: r.Email, Username: r.Username, Action: "ForgotPassword"})

	c.SendStatus(http.StatusOK)
}

// @router /reset-password [post]
func (c *UserController) ResetPassword() {
	r := new(models.ResetPassword)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.ResetPassword(r.Email, r.Token, r.Password)
	c.handleError(err)

	go c.auditLog(models.Log{UserEmail: r.Email, Action: "ResetPassword", Request: "SKIP"})

	c.SendStatus(http.StatusOK)
}

// @router /:id [patch]
func (c *UserController) UpdateUser() {
	u := new(models.User)
	c.parseRequest(u)

	var err error
	if c.user.ID == u.ID {
		u, err = c.C.UserService.UpdateProfile(u)
	} else {
		u, err = c.C.UserService.UpdateUser(u, c.user)
	}
	c.handleError(err)

	go c.auditLog(models.Log{UserID: u.ID, Action: "Update", Request: "SKIP"})

	c.json(u)
}

// @router /:id/send-verify-sms [patch]
func (c *UserController) SendVerifySMS() {
	err := c.C.UserService.SendVerifySMS(c.user)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "SendVerifySMS"})

	c.SendStatus(http.StatusOK)
}

// @router /:id/verify-mobile [patch]
func (c *UserController) VerifyMobile() {
	r := new(models.VerifyMobile)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.VerifyMobile(r.Code, c.user)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "VerifyMobile"})

	c.SendStatus(http.StatusOK)
}

// @router /:id/password [patch]
func (c *UserController) UpdatePassword() {
	r := new(models.UpdatePassword)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.UpdatePassword(c.user, r.OldPassword, r.Password)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "UpdatePassword", Request: "SKIP"})

	c.SendStatus(http.StatusOK)
}

// @router /:id [get]
func (c *UserController) GetUser(id int64) {
	u, err := c.C.UserService.FindUser(&models.User{ID: id}, false)
	c.handleError(err)

	c.json(u)
}

// @router /:id/generate-auth-code [patch]
func (c *UserController) GenerateAuthenticator() {
	r, err := c.C.AuthenticatorService.Generate(c.user)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "GenerateAuthenticator"})

	c.json(r)
}

// @router /:id/authenticator [patch]
func (c *UserController) Authenticator() {
	r := new(models.Authenticator)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.AuthenticatorService.Process(c.user, r)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "Authenticator", Request: "SKIP"})

	c.SendStatus(http.StatusOK)
}

// @router /:id/recovery-questions [patch]
func (c *UserController) RecoveryQuestions() {
	r := new(models.RecoveryQuestions)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.SaveRecoveryQuestions(c.user, r)
	c.handleError(err)

	go c.auditLog(models.Log{UserID: c.user.ID, Action: "Questions", Request: "SKIP"})

	c.SendStatus(http.StatusOK)
}

// @router /:id/notifications [patch]
func (c *UserController) ReadNotification() {
	r := new(models.Notification)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.ReadNotification(c.user, r)
	c.handleError(err)

	c.SendStatus(http.StatusOK)
}

// @router /recovery-questions [post]
func (c *UserController) GetRecoveryQuestions() {
	r := new(models.Login)
	c.parseRequest(r)
	c.validate(r)

	rc, err := c.C.UserService.RecoveryQuestions(r)
	c.handleError(err)

	go c.auditLog(models.Log{Username: r.Username, Action: "GetRecoveryQuestions", Request: "SKIP"})

	c.json(rc)
}

// @router /disable-mfa [post]
func (c *UserController) DisableMFA() {
	r := new(models.DisableMFA)
	c.parseRequest(r)
	c.validate(r)

	err := c.C.UserService.DisableMFA(r)
	c.handleError(err)

	go c.auditLog(models.Log{Username: r.Username, Action: "DisableMFA", Request: "SKIP"})

	c.SendStatus(http.StatusOK)
}

// @router / [get]
func (c *UserController) GetUsers() {
	p, err := c.C.UserService.PaginateUsers(c.paginator(kernel.ListLimit))
	c.handleError(err)

	c.jsonMany(p)
}

// @router /:id [delete]
func (c *UserController) DeleteUser(id int64) {
	err := c.C.UserService.DeleteUser(id)
	c.handleError(err)

	c.SendStatus(http.StatusNoContent)
}
