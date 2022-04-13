package controllers_test

import (
	"backend/models"
	"backend/services"
	"backend/tests"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPaymentMethodController_CreatePaymentMethod(t *testing.T) {
	t.Parallel()
	r := paymentMethod()
	f := func(tk string, r *models.PaymentMethod, status int) {
		tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodPost,
			URI:       "/payment-methods",
			Model:     r,
			Status:    status})
	}
	u, tk := tests.UserWithToken(true)
	f(tk, r, http.StatusCreated)

	pms := tests.C.PaymentService.PaymentMethods(u.CustomerID, true)
	assert.Greater(t, len(pms), 0)
	assert.Equal(t, r.ID, pms[0].ID)
	assert.True(t, pms[0].IsDefault)
	f(tk, r, http.StatusBadRequest)

	u, tk = tests.UserWithToken(true)
	for i := 0; i < services.PaymentMethodCountLimit; i++ {
		f(tk, paymentMethod(), http.StatusCreated)
	}
	f(tk, paymentMethod(), http.StatusBadRequest)
}

func TestPaymentMethodController_DeletePaymentMethod(t *testing.T) {
	f := func(tk string, id string, status int) {
		tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodDelete,
			URI:       "/payment-methods/" + id,
			Status:    status,
		})
	}
	r := paymentMethod()
	u, tk := tests.UserWithToken(true)
	createPaymentMethod(t, u, r)

	f(tk, r.ID, http.StatusNoContent)

	pms := tests.C.PaymentService.PaymentMethods(u.CustomerID, true)
	assert.Empty(t, pms)

	r = paymentMethod()
	u2, _ := tests.UserWithToken(true)
	createPaymentMethod(t, u2, r)
	f(tk, r.ID, http.StatusForbidden)

	//@todo: test delete payment method with subscription is bad request
}

func paymentMethod() (r *models.PaymentMethod) {
	r = &models.PaymentMethod{}
	gofakeit.Struct(r)
	return
}

func createPaymentMethod(t *testing.T, u *models.User, r *models.PaymentMethod) {
	_, err := tests.C.PaymentService.AttachPaymentMethod(u.CustomerID, r)
	tests.FailOnErr(t, err)
}
