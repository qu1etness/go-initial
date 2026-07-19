package payment

import "log/slog"

type StripeClient struct {
	apiKey string
}

func NewStripeClient(apiKey string) *StripeClient {

	return &StripeClient{apiKey: apiKey}
}

func (s *StripeClient) Pay(amount float64) error {
	slog.Info("Processing payment using Stripe", slog.Float64("amount", amount))
	return nil
}

func (s *StripeClient) Refund(amount float64) error {
	return nil
}
