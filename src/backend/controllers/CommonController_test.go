package controllers_test

import (
    "fmt"
    "net/http"
    "testing"

    "backend/models"
    "backend/tests"
    "github.com/brianvoe/gofakeit/v6"
)

func TestCommonController_Contact(t *testing.T) {
    r := &models.Contact{Email: gofakeit.Email(), Message: gofakeit.HackerPhrase(), Name: gofakeit.Name()}
    sub := fmt.Sprintf("[Contact] Message from %s", r.Name)
    tests.ExpectEmail(t,
        "",
        []string{"sender@skeleton-gadelkareem.herokuapp.com"},
        sub,
        "")

    tests.ApiRequest(&tests.Request{T: t,
        Method: http.MethodPost,
        URI:    "/common/contact",
        Model:  r,
        Status: http.StatusOK})
    tests.CheckEmailRetry(t, sub)
}
