package order

import "log/slog"

type PaymentSender interface {
	Pay(amount float64) error
}

type OrderService struct {
	paymentSender PaymentSender
}

func NewOrderService(paymentSender PaymentSender) *OrderService {
	return &OrderService{
		paymentSender: paymentSender,
	}
}

func (o *OrderService) PlaceOrder(amount float64) error {
	slog.Info("Processing order", slog.Float64("amount", amount))
	o.paymentSender.Pay(amount)
	return nil
}
