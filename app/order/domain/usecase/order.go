package usecase

import (
	// "fita-backend-test/app/order/domain/entity"
	"fita-backend-test/app/order/domain/repository"
)

var (

)

type OrderUsecase interface {
	AddOrderItem(string, string, int)
}

type order struct {
	repository domain.CartRepository
}

func NewOrderUsecase(repository domain.CartRepository) OrderUsecase {
	return &order{repository}
}

func (u *order) AddOrderItem(userID, sku string, qty int) {

}