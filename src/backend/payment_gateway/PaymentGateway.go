package payment_gateway

import (
	"backend/models"
)

type (
	PaymentGateway interface {
		CreateCustomer(*models.User) (*models.Customer, error)
		ListProducts() ([]*models.Product, error)
		Subscriptions(string) []*models.Subscription
		CreateSubscription(*models.Subscription) (*models.Subscription, error)
		CancelSubscription(string) (*models.Subscription, error)
		Webhook(*models.PaymentEvent, PaymentService) (*models.PaymentEvent, error)
		// SetupIntent(string) (*models.SetupIntent, error)
		UpcomingInvoice(*models.Subscription) (*models.Invoice, error)
		UpdateCustomer(*models.Customer, *models.User) (*models.Customer, error)
		Customer(string) (*models.Customer, error)
		ListPaymentMethods(string) []*models.PaymentMethod
		DeletePaymentMethods(string) error
		AttachPaymentMethod(string, *models.PaymentMethod) (*models.PaymentMethod, error)
		CreatePaymentMethod(*models.PaymentMethod) (*models.PaymentMethod, error)
		PaymentMethod(string) (*models.PaymentMethod, error)
		UpdateSubscription(*models.Subscription) (*models.Subscription, error)
	}
	PaymentService interface {
		UpdateDefaultPaymentMethod(string, string, string) error
		InvalidateCacheModel(models.BaseInterface, string)
	}
)
