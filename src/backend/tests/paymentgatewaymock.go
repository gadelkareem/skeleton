package tests

import (
	"backend/di"
	"backend/models"
	"backend/payment_gateway"
	"backend/services"
)

type paymentGatewayMock struct{}

func initPaymentService(c *di.Container) {
	c.PaymentService = services.NewPaymentService(new(paymentGatewayMock), c.CacheService)
}

func (paymentGatewayMock) CreateCustomer(user *models.User) (*models.Customer, error) {
	return &models.Customer{
		ID: "123",
	}, nil
}

func (paymentGatewayMock) ListProducts() ([]*models.Product, error) {
	return []*models.Product{}, nil
}

func (paymentGatewayMock) Subscriptions(s string) []*models.Subscription {
	return []*models.Subscription{}
}

func (paymentGatewayMock) CreateSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	return subscription, nil
}

func (paymentGatewayMock) CancelSubscription(s string) (*models.Subscription, error) {
	return &models.Subscription{}, nil
}

func (paymentGatewayMock) Webhook(event *models.PaymentEvent, service payment_gateway.PaymentService) (*models.PaymentEvent, error) {
	return event, nil
}

func (paymentGatewayMock) UpcomingInvoice(subscription *models.Subscription) (*models.Invoice, error) {
	return &models.Invoice{}, nil
}

func (paymentGatewayMock) UpdateCustomer(customer *models.Customer, user *models.User) (*models.Customer, error) {
	return customer, nil
}

func (paymentGatewayMock) Customer(s string) (*models.Customer, error) {
	return &models.Customer{}, nil
}

func (paymentGatewayMock) ListPaymentMethods(s string) []*models.PaymentMethod {
	return []*models.PaymentMethod{}
}

func (paymentGatewayMock) DeletePaymentMethods(s string) error {
	return nil
}

func (paymentGatewayMock) AttachPaymentMethod(s string, method *models.PaymentMethod) (*models.PaymentMethod, error) {
	return method, nil
}

func (paymentGatewayMock) CreatePaymentMethod(method *models.PaymentMethod) (*models.PaymentMethod, error) {
	return method, nil
}

func (paymentGatewayMock) PaymentMethod(s string) (*models.PaymentMethod, error) {
	return &models.PaymentMethod{}, nil
}

func (paymentGatewayMock) UpdateSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	return subscription, nil
}
