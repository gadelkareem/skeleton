package controllers_test

import (
	"backend/models"
	"backend/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSubscriptionController_UpdateSubscription(t *testing.T) {
	f := func(tk string, r *models.Subscription, status int) (w *httptest.ResponseRecorder, rq *http.Request) {
		return tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodPatch,
			URI:       "/subscriptions/" + r.ID,
			Status:    status,
			Model:     r,
		})
	}
	r, u, tk := createPaymentIntent(t)
	pm := paymentMethod()
	createPaymentMethod(t, u, pm)
	ps, err := tests.C.PaymentService.Products()
	tests.FailOnErr(t, err)
	r2 := &models.Subscription{
		ID:                  r.ID,
		CustomerID:          u.CustomerID,
		PriceID:             ps[0].Prices[0].ID,
		CreatePaymentIntent: false,
		PaymentMethodID:     pm.ID,
	}
	w, _ := f(tk, r2, http.StatusOK)
	r3 := new(models.Subscription)
	tests.ReadModel(t, w.Body.Bytes(), r3)
	assert.Equal(t, r2.ID, r3.ID)
	assert.Equal(t, r2.CustomerID, r3.CustomerID)
	assert.Equal(t, r2.PriceID, r3.PriceID)
	assert.Equal(t, pm.ID, r3.PaymentMethodID)
	assert.True(t, tests.C.PaymentService.CustomerHasSubscription(u.CustomerID, r3.ID))
	subs, _ := tests.C.PaymentService.ActiveSubscription(u.CustomerID)
	assert.NotEmpty(t, subs)

	// test update subscription that is not owned by user
	_, tk2 := tests.UserWithToken(true)
	f(tk2, r, http.StatusForbidden)
}

func TestSubscriptionController_CancelSubscription(t *testing.T) {
	t.Parallel()
	// test cancel subscription
	f := func(tk string, id string, status int) {
		tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			URI:       "/subscriptions/" + id,
			Status:    status,
			Method:    http.MethodDelete,
		})
	}
	r, u, tk := createPaymentIntent(t)
	r, _ = createSubscription(t, u, r)
	f(tk, r.ID, http.StatusNoContent)
	subs, _ := tests.C.PaymentService.ActiveSubscription(u.CustomerID)
	assert.Empty(t, subs)

	// test cancel subscription that is not owned by user
	_, tk2 := tests.UserWithToken(true)
	f(tk2, r.ID, http.StatusForbidden)
}
func TestSubscriptionController_CreateSubscription(t *testing.T) {
	t.Parallel()
	f := func(tk string, r *models.Subscription, status int) (w *httptest.ResponseRecorder, rq *http.Request) {
		return tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodPost,
			URI:       "/subscriptions",
			Status:    status,
			Model:     r,
		})
	}
	u, tk := tests.UserWithToken(true)
	ps, err := tests.C.PaymentService.Products()
	tests.FailOnErr(t, err)
	// test payment intent
	r := &models.Subscription{
		CustomerID:          u.CustomerID,
		PriceID:             ps[0].Prices[0].ID,
		CreatePaymentIntent: true,
	}
	w, _ := f(tk, r, http.StatusOK)
	r2 := new(models.Subscription)
	tests.ReadModel(t, w.Body.Bytes(), r2)
	assert.NotEmpty(t, r2.ID)
	assert.NotEmpty(t, r2.PaymentIntentClientSecret)
	assert.Equal(t, r2.CustomerID, r.CustomerID)
	// @todo check if price exists
	assert.Equal(t, r2.PriceID, r.PriceID)
	assert.True(t, r2.CreatePaymentIntent)
	assert.Equal(t, r2.PaymentBehavior, "default_incomplete")

	// test create payment intent when customer has subscription
	createSubscription(t, u, r2)
	tests.WaitForCustomerCache(u.CustomerID)
	f(tk, r, http.StatusBadRequest)

	// test wrong customer id
	r.CustomerID = "wrong"
	f(tk, r, http.StatusForbidden)
}

func createPaymentIntent(t *testing.T) (r *models.Subscription, u *models.User, tk string) {
	u, tk = tests.UserWithToken(true)
	ps, err := tests.C.PaymentService.Products()
	tests.FailOnErr(t, err)
	r = &models.Subscription{
		CustomerID:          u.CustomerID,
		PriceID:             ps[0].Prices[0].ID,
		CreatePaymentIntent: true,
	}
	r, err = tests.C.PaymentService.CreateSubscription(r)
	tests.FailOnErr(t, err)
	return
}

func createSubscription(t *testing.T, u *models.User, r *models.Subscription) (*models.Subscription, *models.PaymentMethod) {
	pm := paymentMethod()
	createPaymentMethod(t, u, pm)
	r.CreatePaymentIntent = false
	r.PaymentMethodID = pm.ID
	r2, err := tests.C.PaymentService.UpdateSubscription(r)
	tests.FailOnErr(t, err)
	return r2, pm
}
