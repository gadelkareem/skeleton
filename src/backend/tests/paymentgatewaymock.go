package tests

import (
	"backend/di"
	"backend/models"
	"backend/payment_gateway"
	"backend/services"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
)

type paymentGatewayMock struct {
	mutex sync.Mutex
	mu    *mutexProps
}
type mutexProps struct {
	paymentMethods map[string][]*models.PaymentMethod
	customer       map[string]*models.Customer
	subscriptions  map[string]map[string]*models.Subscription
	invoices       map[string]map[string]*models.Invoice
}

func initPaymentService(c *di.Container) {
	p := &paymentGatewayMock{
		mu: &mutexProps{
			paymentMethods: make(map[string][]*models.PaymentMethod),
			customer:       make(map[string]*models.Customer),
			subscriptions:  make(map[string]map[string]*models.Subscription),
			invoices:       make(map[string]map[string]*models.Invoice),
		},
	}
	c.PaymentService = services.NewPaymentService(p, c.CacheService)
}

func (p *paymentGatewayMock) CreateCustomer(user *models.User) (*models.Customer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.mu.customer[user.CustomerID] = &models.Customer{ID: user.CustomerID}
	return p.mu.customer[user.CustomerID], nil
}

func (p *paymentGatewayMock) ListInvoices(s string) ([]*models.Invoice, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	invoices, ok := p.mu.invoices[s]
	if ok {
		var invoices1 []*models.Invoice
		for _, sub := range invoices {
			if sub != nil {
				invoices1 = append(invoices1, sub)
			}
		}
		return invoices1, nil
	}
	return nil, nil
}

func (p *paymentGatewayMock) ListProducts() ([]*models.Product, error) {
	return []*models.Product{
		{
			ID:          "prod_KEBfT",
			Name:        "Enterprise",
			Description: "1000 Customers, 1 TB Storage, 1000 Calls, Chat and Email support",
			Subheader:   "open",
			Prices: []*models.Price{
				{
					ID:         "price_1JZjIS",
					Recurring:  models.PriceRecurring{Interval: "year", IntervalCount: 1, TrialPeriodDays: 0},
					UnitAmount: 1000000,
				},
				{
					ID:         "price_1JZjIS",
					Recurring:  models.PriceRecurring{Interval: "month", IntervalCount: 1, TrialPeriodDays: 0},
					UnitAmount: 100000,
				},
			},
		},
		{
			ID:          "prod_KEBe6T",
			Name:        "Basic",
			Description: "500 Customers, 500 GB Storage, 250 Calls, Email support",
			Subheader:   "",
			Prices: []*models.Price{
				{
					ID:         "price_1JZjHX",
					Recurring:  models.PriceRecurring{Interval: "year", IntervalCount: 1, TrialPeriodDays: 0},
					UnitAmount: 100000,
				},
				{
					ID:         "price_1JZjHXB",
					Recurring:  models.PriceRecurring{Interval: "month", IntervalCount: 1, TrialPeriodDays: 0},
					UnitAmount: 10000,
				},
			},
		},
		{
			ID:          "prod_KEBdwpar3GscTI",
			Name:        "Free",
			Description: "100 Customers, 100 GB Storage, 100 Calls, Email support",
			Subheader:   "open",
			Prices: []*models.Price{
				{
					ID:         "price_1JZjGgB",
					Recurring:  models.PriceRecurring{Interval: "month", IntervalCount: 1, TrialPeriodDays: 0},
					UnitAmount: 0,
				},
			},
		},
	}, nil
}

func (p *paymentGatewayMock) Subscriptions(s string) []*models.Subscription {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	subs, ok := p.mu.subscriptions[s]
	if ok {
		var subs1 []*models.Subscription
		for _, sub := range subs {
			if sub != nil {
				subs1 = append(subs1, sub)
			}
		}
		return subs1
	}
	return nil
}

func (p *paymentGatewayMock) CreateSubscription(sub *models.Subscription) (*models.Subscription, error) {
	sub = &models.Subscription{
		ID:                        gofakeit.UUID(),
		CustomerID:                sub.CustomerID,
		CreatePaymentIntent:       true,
		PaymentIntentClientSecret: gofakeit.UUID(),
		PriceID:                   sub.PriceID,
		PaymentMethodID:           "",
		PaymentBehavior:           "default_incomplete",
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.initSubscriptions(sub.CustomerID)
	p.mu.subscriptions[sub.CustomerID][sub.ID] = sub
	return sub, nil
}

func (p *paymentGatewayMock) initSubscriptions(s string) {
	if _, ok := p.mu.subscriptions[s]; !ok {
		p.mu.subscriptions[s] = make(map[string]*models.Subscription)
	}
}

func (p *paymentGatewayMock) initInvoices(s string) {
	if _, ok := p.mu.invoices[s]; !ok {
		p.mu.invoices[s] = make(map[string]*models.Invoice)
	}
}

func (p *paymentGatewayMock) CancelSubscription(s string) (*models.Subscription, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, subs := range p.mu.subscriptions {
		for j, sub := range subs {
			if sub.ID == s {
				p.mu.subscriptions[sub.CustomerID][j].Status = models.SubscriptionStatusCanceled
				break
			}
		}
	}
	return nil, nil
}

func (p *paymentGatewayMock) Webhook(event *models.PaymentEvent, service payment_gateway.PaymentService) (*models.PaymentEvent, error) {
	return event, nil
}

func (p *paymentGatewayMock) UpcomingInvoice(sub *models.Subscription) (*models.Invoice, error) {
	return &models.Invoice{}, nil
}

func (p *paymentGatewayMock) UpdateCustomer(customer *models.Customer, user *models.User) (*models.Customer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.mu.customer[customer.ID] = customer
	return customer, nil
}

func (p *paymentGatewayMock) Customer(s string) (*models.Customer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	m, _ := p.mu.customer[s]
	return m, nil
}

func (p *paymentGatewayMock) ListPaymentMethods(s string) []*models.PaymentMethod {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	ms, _ := p.mu.paymentMethods[s]
	return ms
}

func (p *paymentGatewayMock) DeletePaymentMethods(s string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for customerID, ms := range p.mu.paymentMethods {
		for i, m := range ms {
			if m.ID == s {
				p.mu.paymentMethods[customerID] = append(p.mu.paymentMethods[customerID][:i], p.mu.paymentMethods[customerID][i+1:]...)
				return nil
			}
		}

	}
	return nil
}

func (p *paymentGatewayMock) AttachPaymentMethod(s string, m *models.PaymentMethod) (*models.PaymentMethod, error) {
	pms := p.ListPaymentMethods(s)
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.mu.paymentMethods[s] = append(pms, m)
	return m, nil
}

func (p *paymentGatewayMock) CreatePaymentMethod(method *models.PaymentMethod) (*models.PaymentMethod, error) {
	return method, nil
}

func (p *paymentGatewayMock) PaymentMethod(s string) (*models.PaymentMethod, error) {
	return &models.PaymentMethod{}, nil
}

func (p *paymentGatewayMock) UpdateSubscription(sub *models.Subscription) (*models.Subscription, error) {
	ps, _ := p.ListProducts()
	p.mutex.Lock()
	defer p.mutex.Unlock()
	sub = &models.Subscription{
		ID:                        sub.ID,
		CustomerID:                sub.CustomerID,
		CreatePaymentIntent:       false,
		PaymentIntentClientSecret: "",
		PriceID:                   sub.PriceID,
		PaymentMethodID:           sub.PaymentMethodID,
		Status:                    models.SubscriptionStatusActive,
		PaymentBehavior:           "",
	}
	p.initSubscriptions(sub.CustomerID)
	p.initInvoices(sub.CustomerID)
	p.mu.subscriptions[sub.CustomerID][sub.ID] = sub
	invoice := &models.Invoice{
		ID:         gofakeit.UUID(),
		Total:      services.ProductPrice(ps, sub.PriceID),
		InvoicePDF: gofakeit.URL(),
		Status:     gofakeit.Word(),
		Created:    gofakeit.Date().Unix(),
	}
	p.mu.invoices[sub.CustomerID][invoice.ID] = invoice
	return sub, nil
}
