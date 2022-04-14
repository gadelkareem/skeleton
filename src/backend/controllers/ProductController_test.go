package controllers_test

import (
	"backend/models"
	"backend/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestProductController_GetProducts(t *testing.T) {
	w, _ := tests.ApiRequest(&tests.Request{T: t,
		Method: http.MethodGet,
		URI:    "/products",
		Status: http.StatusOK})
	products, _ := tests.ParseModels(t, reflect.TypeOf(new(models.Product)), w.Body)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 3)
	assert.NotEmpty(t, products[0].(*models.Product).Prices)
}
