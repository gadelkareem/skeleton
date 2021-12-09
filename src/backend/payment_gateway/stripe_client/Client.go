package stripe_client

import (
	"encoding/json"
	"time"

	"backend/internal"
	"backend/kernel"
	"backend/models"
	"backend/payment_gateway"
	"backend/services"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
	"github.com/stripe/stripe-go/v72/refund"
	"github.com/stripe/stripe-go/v72/sub"
	"github.com/stripe/stripe-go/v72/webhook"
)

type Client struct {
}

func NewClient() *Client {
	stripe.Key = kernel.App.Config.String("stripe::stripeKey")
	return &Client{}
}

func (c *Client) ListProducts() (pl []*models.Product, err error) {
	pi := product.List(&stripe.ProductListParams{ListParams: stripe.ListParams{Single: true}, Active: stripe.Bool(true)})
	sprl := pi.ProductList().Data
	for _, spr := range sprl {
		spi := price.List(&stripe.PriceListParams{ListParams: stripe.ListParams{Single: true},
			Product: stripe.String(spr.ID), Active: stripe.Bool(true)})
		spl := spi.PriceList().Data
		var prl []*models.Price
		for _, sp := range spl {
			pr := &models.Price{
				ID:         sp.ID,
				UnitAmount: sp.UnitAmount,
				Recurring: &models.PriceRecurring{
					Interval:        string(sp.Recurring.Interval),
					IntervalCount:   sp.Recurring.IntervalCount,
					TrialPeriodDays: sp.Recurring.TrialPeriodDays,
				},
			}
			prl = append(prl, pr)
		}
		p := &models.Product{
			ID:          spr.ID,
			Name:        spr.Name,
			Description: spr.Description,
			Subheader:   spr.Metadata["subheader"],
			Prices:      prl,
		}
		pl = append(pl, p)
	}
	return
}

func (c *Client) CreateCustomer(u *models.User) (*models.Customer, error) {
	cr, err := customer.New(&stripe.CustomerParams{
		Email: stripe.String(u.Email),
		Name:  stripe.String(u.GetFullName()),
		Phone: stripe.String(u.Mobile),
	})
	if err != nil {
		return nil, err
	}

	return &models.Customer{ID: cr.ID}, err
}

func (c *Client) CreateSubscriptionIntent(priceID, customerID string) (*models.Subscription, error) {
	sps := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(priceID),
			},
		},
		PaymentBehavior: stripe.String("default_incomplete"),
	}
	sps.AddExpand("latest_invoice.payment_intent")
	ss, err := sub.New(sps)
	if err != nil {
		return nil, err
	}
	return &models.Subscription{
		ID:                        ss.ID,
		PaymentIntentClientSecret: ss.LatestInvoice.PaymentIntent.ClientSecret,
	}, nil
}

func (c *Client) CreateSubscription(s *models.Subscription) (*models.Subscription, error) {
	sps := &stripe.SubscriptionParams{
		Customer: stripe.String(s.CustomerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(s.PriceID),
			},
		},
	}
	if s.PaymentMethodId != "" {
		sps.DefaultPaymentMethod = stripe.String(s.PaymentMethodId)
	} else if s.CreatePaymentIntent {
		s.PaymentBehavior = "default_incomplete"
		sps.PaymentBehavior = stripe.String(s.PaymentBehavior)
		sps.AddExpand("latest_invoice.payment_intent")
	}
	ss, err := sub.New(sps)
	if err != nil {
		return nil, err
	}
	s.ID = ss.ID
	if s.PaymentMethodId == "" && ss.LatestInvoice != nil && ss.LatestInvoice.PaymentIntent != nil {
		s.PaymentIntentClientSecret = ss.LatestInvoice.PaymentIntent.ClientSecret
	}
	return s, nil
}

func (c *Client) UpdateSubscription(s *models.Subscription) (*models.Subscription, error) {
	sps := &stripe.SubscriptionParams{
		DefaultPaymentMethod: stripe.String(s.PaymentMethodId),
		CancelAtPeriodEnd:    stripe.Bool(false),
		ProrationBehavior:    stripe.String(string(stripe.SubscriptionProrationBehaviorAlwaysInvoice)),
	}
	if s.CreatePaymentIntent {
		s.PaymentBehavior = "allow_incomplete"
		sps.PaymentBehavior = stripe.String(s.PaymentBehavior)
		sps.AddExpand("latest_invoice.payment_intent")
	}
	if s.ItemID != "" {
		sps.Items = []*stripe.SubscriptionItemsParams{
			{Price: stripe.String(s.PriceID)},
			{ID: stripe.String(s.ItemID), Deleted: stripe.Bool(true)},
		}
	}
	ss, err := sub.Update(s.ID, sps)
	if ss.LatestInvoice != nil && ss.LatestInvoice.PaymentIntent != nil {
		s.PaymentIntentClientSecret = ss.LatestInvoice.PaymentIntent.ClientSecret
	}
	return s, err
}

func (c *Client) CancelSubscription(id string) (*models.Subscription, error) {
	subscription, err := c.subscription(id)
	if err != nil {
		return nil, err
	}
	cancelTime := time.Now().Unix()
	_, err = sub.Cancel(subscription.ID, &stripe.SubscriptionCancelParams{
		Prorate:    stripe.Bool(true),
		InvoiceNow: stripe.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	invItr := invoice.List(&stripe.InvoiceListParams{
		Customer:     stripe.String(subscription.Customer.ID),
		Subscription: stripe.String(subscription.ID),
	})
	var refundAmount int64
	for _, i := range invItr.InvoiceList().Data {
		if i.Created >= cancelTime {
			refundAmount = i.Total * -1
			break
		}
	}
	if refundAmount > 0 {
		latestInvoice, err := invoice.Get(subscription.LatestInvoice.ID, nil)
		if err != nil {
			return nil, err
		}
		_, err = refund.New(&stripe.RefundParams{
			Charge: stripe.String(latestInvoice.Charge.ID),
			Amount: stripe.Int64(refundAmount),
		})
		if err != nil {
			return nil, err
		}
	}

	return &models.Subscription{
		ID: subscription.ID,
	}, err
}

func (c *Client) subscription(id string) (*stripe.Subscription, error) {
	return sub.Get(id, nil)
}

// func (c *Client) CustomerSubscription(customerID string) (*models.Subscription, error) {
//     sps := &stripe.SubscriptionListParams{
//         Customer: customerID,
//         Status:   string(stripe.SubscriptionStatusActive),
//     }
//     var s *stripe.Subscription
//     sl := sub.List(sps)
//     for _, ss := range sl.SubscriptionList().Data {
//         // if ss.DefaultPaymentMethod != nil {
//         // if s != nil {
//         //     logs.Alert("Customer %s has multiple subscriptions", customerID)
//         //     return nil, internal.ErrInternalError
//         // }
//         s = ss
//         // }
//     }
//     if s == nil {
//         return nil, internal.ErrNotFound
//     }
//
//     paymentMethodID := ""
//     if s.DefaultPaymentMethod != nil {
//         paymentMethodID = s.DefaultPaymentMethod.ID
//     }
//
//     i := s.Items.Data[0]
//
//     return &models.Subscription{
//         ID:              s.ID,
//         PriceID:         i.Price.ID,
//         ItemID:          i.ID,
//         CustomerID:      s.Customer.ID,
//         PaymentMethodId: paymentMethodID,
//     }, nil
// }

func (c *Client) Subscriptions(customerID string) (ss []*models.Subscription) {
	sps := &stripe.SubscriptionListParams{
		Customer: customerID,
	}
	sl := sub.List(sps)
	for _, s := range sl.SubscriptionList().Data {
		m := &models.Subscription{
			ID:         s.ID,
			CustomerID: s.Customer.ID,
			Status:     models.SubscriptionStatus(s.Status),
		}
		i := s.Items.Data[0]
		if i != nil {
			if i.Price != nil {
				m.PriceID = i.Price.ID
			}
			m.ItemID = i.ID
		}
		if s.DefaultPaymentMethod != nil {
			m.PaymentMethodId = s.DefaultPaymentMethod.ID
		}
		ss = append(ss, m)
	}
	return
}

// Webhook @todo turn webhooks into ques
func (c *Client) Webhook(m *models.PaymentEvent, s payment_gateway.PaymentService) (*models.PaymentEvent, error) {
	e, err := webhook.ConstructEvent(m.Payload, m.Signature, kernel.App.Config.String("stripe::webhookSecret"))
	if err != nil {
		return nil, err
	}
	switch e.Type {
	case "invoice.payment_succeeded":
		var invoice stripe.Invoice
		err = json.Unmarshal(e.Data.Raw, &invoice)
		if err != nil {
			return nil, err
		}

		pi, _ := paymentintent.Get(
			invoice.PaymentIntent.ID,
			nil,
		)
		err = s.UpdateDefaultPaymentMethod(pi.PaymentMethod.ID, invoice.Subscription.ID, invoice.Customer.ID)
		if err != nil {
			return nil, err
		}
		return m, nil
	case "product.updated":
		var p stripe.Product
		err = json.Unmarshal(e.Data.Raw, &p)
		if err != nil {
			return nil, err
		}
		s.InvalidateCacheModel(&models.Product{ID: p.ID}, "")
		return nil, err
	default:
		return nil, internal.ErrNotImplemented
	}

}

//
// func (c *Client) SetupIntent(customerID string) (*models.SetupIntent, error) {
//     ps := &stripe.PaymentIntentParams{
//         PaymentMethodTypes: stripe.StringSlice([]string{
//             "card",
//         }),
//     }
//     si, err := paymentintent.New(ps)
//     if err != nil {
//         return nil, err
//     }
//
//     return &models.SetupIntent{ClientSecret: si.ClientSecret}, nil
// }

func (c *Client) UpcomingInvoice(s *models.Subscription) (*models.Invoice, error) {
	ps := &stripe.InvoiceParams{
		Customer:     stripe.String(s.CustomerID),
		Subscription: stripe.String(s.ID),
		SubscriptionItems: []*stripe.SubscriptionItemsParams{
			{Price: stripe.String(s.PriceID)},
			{ID: stripe.String(s.ItemID), Deleted: stripe.Bool(true)},
		},
		// SubscriptionBillingCycleAnchorUnchanged: stripe.Bool(true),
		// SubscriptionCancelNow:                   stripe.Bool(true),
		SubscriptionProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorAlwaysInvoice)),
	}
	in, err := invoice.GetNext(ps)
	if err != nil {
		return nil, err
	}
	return &models.Invoice{ID: in.ID, Total: in.Total}, nil
}

func (c *Client) ListPaymentMethods(customerID string) (cs []*models.PaymentMethod) {
	ps := &stripe.PaymentMethodListParams{
		Customer:   stripe.String(customerID),
		Type:       stripe.String("card"),
		ListParams: stripe.ListParams{Limit: stripe.Int64(services.PaymentMethodListLimit)},
	}
	i := paymentmethod.List(ps)
	// main:
	for i.Next() {
		spm := i.PaymentMethod()
		sc := spm.Card
		cr := &models.PaymentMethod{
			ID:   spm.ID,
			Type: string(spm.Type),
			Card: &models.PaymentMethodCard{
				Brand:       string(sc.Brand),
				Fingerprint: sc.Fingerprint,
				ExpMonth:    sc.ExpMonth,
				ExpYear:     sc.ExpYear,
				Last4:       sc.Last4,
			},
			BillingDetails: &models.BillingDetails{
				Name: spm.BillingDetails.Name,
			},
		}
		// // Workaround Stripe payment methods are not unique
		// // https://github.com/stripe/stripe-payments-demo/issues/45
		// for _, cc := range cs {
		//     if cc.Card.Fingerprint == cr.Card.Fingerprint &&
		//         cc.Card.ExpMonth == cr.Card.ExpMonth &&
		//         cc.Card.ExpYear == cr.Card.ExpYear &&
		//         cc.BillingDetails.Name == cr.BillingDetails.Name {
		//         continue main
		//     }
		// }
		cs = append(cs, cr)
	}
	return
}

func (c *Client) DeletePaymentMethods(id string) error {
	_, err := paymentmethod.Detach(id, nil)
	return err
}

func (c *Client) UpdateCustomer(m *models.Customer, u *models.User) (*models.Customer, error) {
	ps := &stripe.CustomerParams{}
	if u != nil {
		ps.Email = stripe.String(u.Email)
		ps.Name = stripe.String(u.GetFullName())
		ps.Phone = stripe.String(u.Mobile)
	}
	if m.DefaultPaymentMethodID != "" {
		ps.InvoiceSettings = &stripe.CustomerInvoiceSettingsParams{DefaultPaymentMethod: stripe.String(m.DefaultPaymentMethodID)}
	}
	_, err := customer.Update(m.ID, ps)
	return m, err
}

func (c *Client) Customer(id string) (*models.Customer, error) {
	m, err := customer.Get(id, nil)
	if err != nil {
		return nil, err
	}
	var pmID string
	if m.InvoiceSettings.DefaultPaymentMethod != nil {
		pmID = m.InvoiceSettings.DefaultPaymentMethod.ID
	}
	return &models.Customer{
		ID:                     id,
		DefaultPaymentMethodID: pmID,
	}, nil
}

func (c *Client) AttachPaymentMethod(customerID string, pm *models.PaymentMethod) (*models.PaymentMethod, error) {
	spm, err := paymentmethod.Attach(
		pm.ID,
		&stripe.PaymentMethodAttachParams{
			Customer: stripe.String(customerID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.PaymentMethod{ID: spm.ID}, nil
}

func (c *Client) CreatePaymentMethod(pm *models.PaymentMethod) (*models.PaymentMethod, error) {
	spm, err := paymentmethod.New(&stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(pm.Card.Token),
		},
		Type: stripe.String("card"),
	})
	if err != nil {
		return nil, err
	}
	sc := spm.Card
	cr := &models.PaymentMethod{
		ID:   spm.ID,
		Type: string(spm.Type),
		Card: &models.PaymentMethodCard{
			Brand:       string(sc.Brand),
			Fingerprint: sc.Fingerprint,
			ExpMonth:    sc.ExpMonth,
			ExpYear:     sc.ExpYear,
			Last4:       sc.Last4,
		},
		BillingDetails: &models.BillingDetails{
			Name: spm.BillingDetails.Name,
		},
	}
	return cr, nil
}

func (c *Client) PaymentMethod(id string) (*models.PaymentMethod, error) {
	spm, err := paymentmethod.Get(id, nil)
	if err != nil {
		return nil, err
	}
	sc := spm.Card
	cr := &models.PaymentMethod{
		ID:   spm.ID,
		Type: string(spm.Type),
		Card: &models.PaymentMethodCard{
			Brand:       string(sc.Brand),
			Fingerprint: sc.Fingerprint,
			ExpMonth:    sc.ExpMonth,
			ExpYear:     sc.ExpYear,
			Last4:       sc.Last4,
		},
		BillingDetails: &models.BillingDetails{
			Name: spm.BillingDetails.Name,
		},
	}
	return cr, nil
}
