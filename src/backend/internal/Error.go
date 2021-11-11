package internal

import (
    "fmt"
    "net/http"

    "github.com/google/jsonapi"
)

type Error interface {
    error
    Status() int
    Object() *jsonapi.ErrorObject
    Error() string
}

type err struct {
    jsonapi.ErrorObject
    status int
}

func newErr(code string, s int, t string, args ...interface{}) *err {
    return &err{
        ErrorObject: jsonapi.ErrorObject{
            Title:  fmt.Sprintf(t, args...),
            Status: fmt.Sprintf("%d", s),
            Code:   code,
        },
        status: s,
    }
}
func Errorf(s int, t string, args ...interface{}) Error {
    return newErr("", s, t, args...)
}

func ErrorWithCode(code string, s int, t string, args ...interface{}) Error {
    return newErr(code, s, t, args...)
}

func ValidationError(k, s string) Error {
    e := ErrValidationError.(*err)
    e.ErrorObject.Meta = &map[string]interface{}{k: s}
    return e
}

func ValidationErrors(v map[string]interface{}) Error {
    e := ErrValidationError.(*err)
    e.ErrorObject.Meta = &v
    return e
}

func (e err) Error() string {
    return e.ErrorObject.Title
}

func (e err) Status() int {
    return e.status
}

func (e err) Object() *jsonapi.ErrorObject {
    return &e.ErrorObject
}

var (
    ErrInternalError                   = Errorf(http.StatusInternalServerError, "Internal server error.")
    ErrEmailExists                     = Errorf(http.StatusBadRequest, "Email already exists in our database.")
    ErrPaymentMethodLimitExceeded      = Errorf(http.StatusBadRequest, "You have reached your payment methods limit.")
    ErrPaymentMethodExists             = Errorf(http.StatusBadRequest, "Payment Method already exists in our database.")
    ErrPaymentMethodDeletionNotAllowed = Errorf(http.StatusBadRequest, "Cannot delete primary payment method.")
    ErrSubscriptionExists              = Errorf(http.StatusBadRequest, "You are already subscribed, please cancel your current subscription to be able to change your subscription.")
    ErrEmailNotExist                   = Errorf(http.StatusBadRequest, "Email does not exists in our database.")
    ErrEmailRequired                   = Errorf(http.StatusBadRequest, "An email address is required.")
    ErrUsernameExists                  = Errorf(http.StatusBadRequest, "Username already exists in our database.")
    ErrValidationError                 = Errorf(http.StatusBadRequest, "Error validating your request.")
    ErrInvalidPass                     = Errorf(http.StatusUnauthorized, "Invalid username or password.")
    ErrInvalidRequest                  = Errorf(http.StatusBadRequest, "Invalid request specified.")
    ErrInvalidActivationCode           = Errorf(http.StatusBadRequest, "Invalid activation code.")
    ErrInvalidCSRFToken                = Errorf(http.StatusUnauthorized, "unauthorized request.")
    ErrForbidden                       = Errorf(http.StatusForbidden, "You do not have permission to perform this action.")
    ErrInvalidResetPassHash            = Errorf(http.StatusBadRequest, "Invalid or expired reset password hash.")
    ErrEmailAlreadyVerified            = Errorf(http.StatusBadRequest, "Email is already verified.")
    ErrMobileAlreadyVerified           = Errorf(http.StatusBadRequest, "Mobile is already verified.")
    ErrMobileRequired                  = Errorf(http.StatusBadRequest, "Mobile is not set for this account.")
    ErrInvalidSMSCode                  = Errorf(http.StatusBadRequest, "Invalid verification code.")
    ErrResetPasswordAlreadyGenerated   = Errorf(http.StatusBadRequest, "A reset password process was already started for this account.")
    ErrNotFound                        = Errorf(http.StatusNotFound, "The requested resource was not found.")
    ErrNotImplemented                  = Errorf(http.StatusNotImplemented, "The requested resource is not implemented.")
    ErrInvalidSocialProvider           = Errorf(http.StatusBadRequest, "Invalid Social network provided.")
    ErrAuthenticatorAlreadyEnabled     = Errorf(http.StatusBadRequest, "2-step verification is already enabled on your account.")
    ErrInvalidAuthenticatorCode        = Errorf(http.StatusUnauthorized, "Invalid 2-step verification code.")
    ErrRecoveryChangeDisallowed        = Errorf(http.StatusForbidden, "Recovery questions are already set for this account.")
    ErrRecoveryQuestionNum             = Errorf(http.StatusBadRequest, "Recovery questions should contain 3 questions.")
    ErrRecoveryQuestionNotSet          = Errorf(http.StatusBadRequest, "Please add your recovery questions to be able to use this feature.")
    ErrBadRecoveryAnswers              = Errorf(http.StatusUnauthorized, "The recovery questions and answers provided are not correct.")
    ErrAuthenticatorCodeMissing        = ErrorWithCode("MISSING_AUTHENTICATOR", http.StatusUnprocessableEntity, "No 2-step verification code provided.")
    ErrMobileCodeMissing               = ErrorWithCode("MISSING_MOBILE_CODE", http.StatusUnprocessableEntity, "No mobile verification code provided.")
    ErrInvalidJWTToken                 = ErrorWithCode("INVALID_TOKEN", http.StatusUnauthorized, "Invalid authentication token.")
    ErrTooManyRequests                 = Errorf(http.StatusTooManyRequests, "Too Many Requests, please try again in few minutes.")
)
