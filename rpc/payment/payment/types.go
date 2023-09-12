package payment

type StripeCheckoutCredential struct {
	PubKey        string `json:"pub_key"`
	SecKey        string `json:"sec_key"`
	WebhookSecKey string `json:"webhook_sec_key"`
}
