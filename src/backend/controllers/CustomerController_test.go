package controllers_test

import (
	"backend/models"
	"backend/services"
	"backend/tests"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCustomerController_UpdateCustomer(t *testing.T) {
	f := func(tk string, r *models.Customer, status int) {
		tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodPatch,
			URI:       "/customers/" + r.ID,
			Status:    status,
			Model:     r,
		})
	}
	r, _, tk := customer(t)
	f(tk, r, http.StatusOK)
	r.DefaultPaymentMethodID = gofakeit.UUID()
	f(tk, r, http.StatusOK)
	r1, err := tests.C.PaymentService.Customer(r.ID)
	tests.FailOnErr(t, err)
	assert.Equal(t, r1.DefaultPaymentMethodID, r.DefaultPaymentMethodID)
	_, tk = tests.UserWithToken(true)
	f(tk, r, http.StatusForbidden)
}

func TestCustomerController_CustomerSubscription(t *testing.T) {
	t.Parallel()
	f := func(tk, id string, status int) (w *httptest.ResponseRecorder, rq *http.Request) {
		return tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodGet,
			URI:       "/customers/" + id + "/subscription",
			Status:    status,
		})
	}
	r, u, tk := createPaymentIntent(t)
	r, pm := createSubscription(t, u, r)
	tests.WaitForCustomerCache(u.CustomerID)
	w, _ := f(tk, u.CustomerID, http.StatusOK)
	r1 := new(models.Subscription)
	tests.ReadModel(t, w.Body.Bytes(), r1)
	assert.Equal(t, r1.ID, r1.ID)
	assert.Equal(t, r1.CustomerID, r1.CustomerID)
	assert.Equal(t, r1.PriceID, r1.PriceID)
	assert.Equal(t, pm.ID, r1.PaymentMethodID)

	// test list subscription for user without subscription
	u1, tk2 := tests.UserWithToken(true)
	f(tk2, u1.CustomerID, http.StatusNotFound)

	// test list subscription with wrong customer id
	f(tk2, "wrong", http.StatusForbidden)
}

func TestCustomerController_ListPaymentMethods(t *testing.T) {
	t.Parallel()
	f := func(tk, id string, status int) ([]interface{}, *jsonapi.ManyPayload) {
		w, _ := tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodGet,
			URI:       "/customers/" + id + "/payment-methods",
			Status:    status,
		})
		return tests.ParseModels(t, reflect.TypeOf(new(models.PaymentMethod)), w.Body)
	}
	r, u, tk := customer(t)
	pm := paymentMethod()
	createPaymentMethod(t, u, pm)

	pms, _ := f(tk, r.ID, http.StatusOK)
	assert.Len(t, pms, 1)
	assert.Equal(t, pms[0].(*models.PaymentMethod).ID, pm.ID)
	_, tk = tests.UserWithToken(true)
	f(tk, pm.ID, http.StatusForbidden)
}

func TestCustomerController_CustomerInvoices(t *testing.T) {
	t.Parallel()
	f := func(tk, id string, status int) ([]interface{}, *jsonapi.ManyPayload) {
		w, _ := tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodGet,
			URI:       "/customers/" + id + "/invoices",
			Status:    status,
		})
		return tests.ParseModels(t, reflect.TypeOf(new(models.Invoice)), w.Body)
	}
	r, u, tk := createPaymentIntent(t)
	createSubscription(t, u, r)

	ms, _ := f(tk, u.CustomerID, http.StatusOK)
	assert.Len(t, ms, 1)
	assert.NotEmpty(t, ms[0].(*models.Invoice).ID)
	ps, err := tests.C.PaymentService.Products()
	tests.FailOnErr(t, err)
	assert.Equal(t, ms[0].(*models.Invoice).Total, services.ProductPrice(ps, r.PriceID))

	_, tk = tests.UserWithToken(true)
	f(tk, u.CustomerID, http.StatusForbidden)
}

func customer(t *testing.T) (r *models.Customer, u *models.User, tk string) {
	r = &models.Customer{}
	gofakeit.Struct(r)
	u, tk = tests.UserWithToken(true)
	u.CustomerID = r.ID
	err := tests.C.UserService.Save(u)
	tests.FailOnErr(t, err)
	return
}
