package app

import (
	"fita-backend-test/app/account"
	"fita-backend-test/app/order"
	"fita-backend-test/app/product"
)

type Gateway struct {
	AccountService account.AccountService
	OrderService order.OrderService
	ProductService product.ProductService
}

func NewGateway() *Gateway{
	return &Gateway{
		AccountService: account.NewAccountService(),
		OrderService: order.NewOrderService(),
		ProductService: product.NewProductService(),
	}
}