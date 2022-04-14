package services

import (
	"backend/internal"
	"backend/models"
	"backend/payment_gateway"
	"backend/utils/paginator"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type (
	PaymentService struct {
		p payment_gateway.PaymentGateway
		c *CacheService
	}
)

const (
	PaymentMethodCountLimit = 5
	PaymentMethodListLimit  = 100
)

func NewPaymentService(p payment_gateway.PaymentGateway, c *CacheService) *PaymentService {
	return &PaymentService{p: p, c: c}
}

func (s *PaymentService) CreateCustomer(m *models.User) (*models.Customer, error) {
	return s.p.CreateCustomer(m)
}

func (s *PaymentService) CreateSubscription(sub *models.Subscription) (m *models.Subscription, err error) {
	if s.CustomerHasActiveSubscriptions(sub.CustomerID) {
		return nil, internal.ErrSubscriptionExists
	}

	m, err = s.p.CreateSubscription(sub)

	go s.InvalidateCustomer(sub.CustomerID)

	return
}

func (s *PaymentService) Products() (ms []*models.Product, err error) {
	k, _ := s.c.Get(&ms, nil)
	if len(ms) > 0 {
		return
	}

	ms, err = s.p.ListProducts()
	if err != nil {
		return nil, err
	}
	go s.c.Put(&ms, k, nil, s.c.interfaceType(ms))

	return
}

func (s *PaymentService) InvalidateCacheModel(m models.BaseInterface) {
	s.c.InvalidateModel(m)
}

func (s *PaymentService) InvalidateCustomer(customerID string) {
	_ = s.c.InvalidateTags(s.CustomerCacheTag(customerID))
}

func (s *PaymentService) ActiveSubscription(customerID string) (*models.Subscription, error) {
	ms := s.ActiveSubscriptions(customerID)
	l := len(ms)
	if l > 1 {
		logs.Alert("Customer %s has multiple subscriptions", customerID)
		return nil, internal.ErrInternalError
	}
	if l == 0 {
		return nil, internal.ErrNotFound
	}

	return ms[0], nil
}

func (s *PaymentService) ActiveSubscriptions(customerID string) (ms []*models.Subscription) {
	ss := s.Subscriptions(customerID)
	for _, m := range ss {
		if m.Status == models.SubscriptionStatusActive {
			ms = append(ms, m)
		}
	}
	return
}

func (s *PaymentService) Subscriptions(customerID string) (ms []*models.Subscription) {
	k, _ := s.c.Get(&ms, nil, customerID)
	if len(ms) > 0 {
		return
	}
	ms = s.p.Subscriptions(customerID)

	if len(ms) > 0 {
		go s.c.Put(&ms, k, nil, s.CustomerCacheTag(customerID))
	}

	return
}

func (s *PaymentService) subscription(id, customerID string) (m *models.Subscription) {
	ms := s.Subscriptions(customerID)
	for _, m := range ms {
		if m.ID == id {
			return m
		}
	}

	return
}

func (s *PaymentService) CustomerHasSubscription(customerID string, subscriptionID string) bool {
	ms := s.Subscriptions(customerID)
	for _, m := range ms {
		if m.ID == subscriptionID {
			return true
		}
	}

	return false
}

func (s *PaymentService) CustomerHasActiveSubscriptions(customerID string) bool {
	return customerID != "" && len(s.ActiveSubscriptions(customerID)) > 0
}

func (s PaymentService) CustomerCacheTag(customerID string) string {
	return fmt.Sprintf("CustomerID_%s", customerID)
}

func (s *PaymentService) UpdateSubscription(sub *models.Subscription) (m *models.Subscription, err error) {
	oldSub := s.subscription(sub.ID, sub.CustomerID)
	if sub.ItemID != "" && sub.PriceID == oldSub.PriceID {
		return nil, internal.ErrSubscriptionNotUpdated
	}
	m, err = s.p.UpdateSubscription(sub)
	go s.InvalidateCustomer(sub.CustomerID)

	return m, err
}

func (s *PaymentService) CancelSubscription(subscriptionID, customerID string) error {
	_, err := s.p.CancelSubscription(subscriptionID)
	go s.InvalidateCustomer(customerID)

	return err
}

func (s *PaymentService) UpdateDefaultPaymentMethod(paymentMethodID string, subscriptionID string, customerID string) error {
	_, err := s.UpdateCustomer(&models.Customer{ID: customerID, DefaultPaymentMethodID: paymentMethodID}, nil)
	if err != nil {
		return err
	}
	_, err = s.UpdateSubscription(&models.Subscription{ID: subscriptionID, PaymentMethodID: paymentMethodID})
	return err
}

func (s *PaymentService) Webhook(e *models.PaymentEvent) error {
	_, err := s.p.Webhook(e, s)
	return err
}

// func (s *PaymentService) SetupIntent(customerID string) (*models.SetupIntent, error) {
//     return s.p.SetupIntent(customerID)
// }

func (s *PaymentService) UpdateCustomer(cus *models.Customer, u *models.User) (*models.Customer, error) {
	m, err := s.p.UpdateCustomer(cus, u)
	if err != nil {
		return nil, err
	}
	go s.InvalidateCustomer(m.ID)
	return m, nil
}

func (s *PaymentService) Customer(id string) (*models.Customer, error) {
	m := &models.Customer{}
	cID, _ := s.c.Get(m, nil, id)
	if m.ID != "" {
		return m, nil
	}
	m, err := s.p.Customer(id)
	if err != nil {
		return nil, err
	}
	go s.c.Put(m, cID, nil, s.CustomerCacheTag(m.ID))

	return m, nil
}

func (s *PaymentService) PaymentMethods(customerID string, resetCache bool) (ms []*models.PaymentMethod) {
	k, _ := s.c.Key(&ms, nil, customerID)
	if !resetCache {
		_, _ = s.c.Get(&ms, nil, customerID)
		if len(ms) > 0 {
			return
		}
	}

	ms = s.p.ListPaymentMethods(customerID)
	if len(ms) < 1 {
		return
	}
	customer, _ := s.Customer(customerID)
	if customer != nil {
		for _, m := range ms {
			if m.ID == customer.DefaultPaymentMethodID {
				m.IsDefault = true
				break
			}
		}
	}

	go s.c.Put(&ms, k, nil, s.CustomerCacheTag(customerID))

	return
}

func (s *PaymentService) CustomerInvoices(customerID string) (ms []*models.Invoice, err error) {
	k, _ := s.c.Get(&ms, nil, customerID)
	if len(ms) > 0 {
		return
	}

	ms, err = s.p.ListInvoices(customerID)
	if err != nil || len(ms) < 1 {
		return
	}

	go s.c.Put(&ms, k, nil, s.CustomerCacheTag(customerID))

	return
}

func (s *PaymentService) PaginateInvoices(customerID string, p *paginator.Paginator) (*paginator.Paginator, error) {
	ms, err := s.CustomerInvoices(customerID)
	if err != nil {
		return nil, err
	}

	for _, m := range ms {
		p.Models = append(p.Models, m)
	}
	p.Slice()

	return p, nil
}

func (s *PaymentService) PaginateProducts(p *paginator.Paginator) (*paginator.Paginator, error) {
	ms, err := s.Products()
	if err != nil {
		return nil, err
	}

	for _, m := range ms {
		p.Models = append(p.Models, m)
	}
	p.Slice()

	return p, nil
}

func (s *PaymentService) PaginatePaymentMethods(customerID string, resetCache bool, p *paginator.Paginator) *paginator.Paginator {
	ms := s.PaymentMethods(customerID, resetCache)

	for _, m := range ms {
		p.Models = append(p.Models, m)
	}
	p.Slice()

	return p
}

func (s *PaymentService) CustomerHasPaymentMethod(customerID string, paymentMethodID string) bool {
	ms := s.PaymentMethods(customerID, false)
	for _, m := range ms {
		if m.ID == paymentMethodID {
			return true
		}
	}

	return false
}

func (s *PaymentService) subscriptionHasPaymentMethod(customerID string, paymentMethodID string) bool {
	ms := s.Subscriptions(customerID)
	for _, m := range ms {
		if m.PaymentMethodID == paymentMethodID {
			return true
		}
	}

	return false
}

func (s *PaymentService) DeletePaymentMethod(id, customerID string) (err error) {
	ms := s.PaymentMethods(customerID, true)
	for _, m := range ms {
		if m.ID == id {
			if s.subscriptionHasPaymentMethod(customerID, id) {
				return internal.ErrPaymentMethodDeletionNotAllowed
			}
			err = s.p.DeletePaymentMethods(id)
			go s.InvalidateCustomer(customerID)

			return
		}
	}
	return nil
}

func (s *PaymentService) AttachPaymentMethod(customerID string, om *models.PaymentMethod) (m *models.PaymentMethod, err error) {
	ms := s.PaymentMethods(customerID, true)
	if len(ms) >= PaymentMethodCountLimit {
		return nil, internal.ErrPaymentMethodLimitExceeded
	}
	for _, pm := range ms {
		if pm.Card.Fingerprint == om.Card.Fingerprint &&
			pm.Card.ExpYear == om.Card.ExpYear &&
			pm.Card.ExpMonth == om.Card.ExpMonth {
			return nil, internal.ErrPaymentMethodExists
		}
	}
	if len(ms) == 0 {
		om.IsDefault = true
	}
	m, err = s.p.AttachPaymentMethod(customerID, om)
	if len(ms) == 0 {
		_, err = s.UpdateCustomer(&models.Customer{ID: customerID, DefaultPaymentMethodID: om.ID}, nil)
		if err != nil {
			return
		}
	}

	go s.InvalidateCustomer(customerID)

	return
}

func (s *PaymentService) paymentMethodPaginator() *paginator.Paginator {
	return paginator.NewPaginator(map[string]int{"size": PaymentMethodListLimit, "after": 1}, "", "", "")
}

func (s *PaymentService) UpcomingInvoice(sub *models.Subscription) (*models.Invoice, error) {
	return s.p.UpcomingInvoice(sub)

}

func ProductPrice(ps []*models.Product, priceID string) (total int64) {
	for _, p := range ps {
		for _, price := range p.Prices {
			if price.ID == priceID {
				total = price.UnitAmount
				return
			}
		}
	}
	return
}
