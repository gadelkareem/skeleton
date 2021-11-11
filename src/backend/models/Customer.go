package models

type SubscriptionStatus string

const (
    SubscriptionStatusActive            SubscriptionStatus = "active"
    SubscriptionStatusAll               SubscriptionStatus = "all"
    SubscriptionStatusCanceled          SubscriptionStatus = "canceled"
    SubscriptionStatusIncomplete        SubscriptionStatus = "incomplete"
    SubscriptionStatusIncompleteExpired SubscriptionStatus = "incomplete_expired"
    SubscriptionStatusPastDue           SubscriptionStatus = "past_due"
    SubscriptionStatusTrialing          SubscriptionStatus = "trialing"
    SubscriptionStatusUnpaid            SubscriptionStatus = "unpaid"
)

type (
    Customer struct {
        Base
        ID                     string `json:"id"  jsonapi:"primary,customers"`
        DefaultPaymentMethodID string `json:"default_payment_method_id"  jsonapi:"attr,default_payment_method_id"`
    }
    Product struct {
        Base
        ID          string   `json:"id"  jsonapi:"primary,products"`
        Name        string   `json:"name"  jsonapi:"attr,name"`
        Description string   `json:"description" jsonapi:"attr,description"`
        Subheader   string   `json:"subheader" jsonapi:"attr,subheader"`
        Prices      []*Price `jsonapi:"attr,prices"`
    }
    Price struct {
        ID         string          `json:"id" jsonapi:"primary,prices"`
        Recurring  *PriceRecurring `json:"recurring" jsonapi:"attr,recurring"`
        UnitAmount int64           `json:"unit_amount" jsonapi:"attr,unit_amount"`
    }
    PriceRecurring struct {
        Interval        string `json:"interval" jsonapi:"attr,interval"`
        IntervalCount   int64  `json:"interval_count" jsonapi:"attr,interval_count"`
        TrialPeriodDays int64  `json:"trial_period_days" jsonapi:"attr,trial_period_days"`
    }
    Subscription struct {
        Base
        ID                        string             `json:"id" jsonapi:"primary,subscriptions"`
        PaymentIntentClientSecret string             `json:"payment_intent_client_secret" jsonapi:"attr,payment_intent_client_secret"`
        PriceID                   string             `json:"price_id" jsonapi:"attr,price_id" valid:"Required;MaxSize(100)"`
        ItemID                    string             `json:"item_id" jsonapi:"attr,item_id"`
        CustomerID                string             `json:"customer_id" jsonapi:"attr,customer_id" valid:"Required;MaxSize(100)"`
        CreatePaymentIntent       bool               `json:"create_payment_intent" jsonapi:"attr,create_payment_intent"`
        PaymentMethodId           string             `json:"payment_method_id" jsonapi:"attr,payment_method_id"`
        PaymentBehavior           string             `json:"payment_behavior" jsonapi:"attr,payment_behavior"`
        Status                    SubscriptionStatus `json:"status" jsonapi:"attr,status"`
    }
    PaymentEvent struct {
        Payload   []byte `json:"payload" valid:"Required"`
        Signature string `json:"signature" valid:"Required"`
    }
    SetupIntent struct {
        ID           string `json:"id" jsonapi:"primary,setup-intents"`
        ClientSecret string `json:"client_secret" jsonapi:"attr,client_secret"`
    }
    PaymentMethod struct {
        Base
        ID   string             `json:"id" jsonapi:"primary,payment-methods" valid:"Required;MaxSize(100)"`
        Card *PaymentMethodCard `json:"card" jsonapi:"attr,card" valid:"Required"`
        // Ideal            *PaymentMethodIdeal            `json:"ideal"`
        Type           string          `json:"type" jsonapi:"attr,type"`
        IsDefault      bool            `json:"is_default" jsonapi:"attr,is_default"`
        BillingDetails *BillingDetails `json:"billing_details" jsonapi:"attr,billing_details"`
    }
    Token struct {
        ID string `json:"id" jsonapi:"primary,tokens"`
    }
    PaymentMethodCard struct {
        Brand string `json:"brand" jsonapi:"attr,brand"`
        Token string `json:"token" jsonapi:"attr,token"`
        // Checks            *PaymentMethodCardChecks            `json:"checks"`
        ExpMonth    uint64 `json:"exp_month"`
        ExpYear     uint64 `json:"exp_year"`
        Fingerprint string `json:"fingerprint"`
        // Funding           CardFunding                         `json:"funding"`
        Last4 string `json:"last4"`
        // Networks          *PaymentMethodCardNetworks          `json:"networks"`
        // ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
        // Wallet            *PaymentMethodCardWallet            `json:"wallet"`
    }
    BillingDetails struct {
        Name string `json:"name" jsonapi:"attr,name"`
    }
    Invoice struct {
        ID    string `json:"id" jsonapi:"primary,invoices"`
        Total int64  `json:"total" jsonapi:"attr,total"`
    }
)

func (m *Customer) GetID() string      { return m.ID }
func (m *Customer) Sanitize()          {}
func (m *Product) GetID() string       { return m.ID }
func (m *Product) Sanitize()           {}
func (m *PaymentMethod) GetID() string { return m.ID }
func (m *PaymentMethod) Sanitize()     {}
func (m *Subscription) GetID() string  { return m.ID }
func (m *Subscription) Sanitize()      {}
