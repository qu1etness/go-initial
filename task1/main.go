package main

import (
	"go-initial/task1/order"
	"go-initial/task1/payment"
)

func main() {

	stripeClient := payment.NewStripeClient("")
	orderService := order.NewOrderService(stripeClient)
	orderService.PlaceOrder(100)

}
