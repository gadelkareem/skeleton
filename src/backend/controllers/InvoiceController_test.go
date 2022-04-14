package controllers_test

import (
	"backend/models"
	"backend/tests"
	"github.com/brianvoe/gofakeit/v6"
	"net/http"
	"net/url"
	"testing"
)

func TestInvoiceController_UpcomingInvoice(t *testing.T) {
	t.Parallel()
	f := func(tk string, r *models.Subscription, status int) {
		tests.ApiRequest(&tests.Request{
			T:         t,
			AuthToken: tk,
			Method:    http.MethodGet,
			URI:       "/invoices/upcoming",
			Status:    status,
			Query:     url.Values{"id": {r.ID}, "customer_id": {r.CustomerID}, "price_id": {r.PriceID}, "item_id": {r.ItemID}},
		})
	}
	u, tk := tests.UserWithToken(true)
	r := &models.Subscription{
		ID:         gofakeit.UUID(),
		CustomerID: u.CustomerID,
		PriceID:    gofakeit.UUID(),
		ItemID:     gofakeit.UUID(),
	}

	f(tk, r, http.StatusOK)

	// test wrong user
	_, tk2 := tests.UserWithToken(true)
	f(tk2, r, http.StatusForbidden)

}
