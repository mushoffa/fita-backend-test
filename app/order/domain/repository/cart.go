package domain

import (
	"fita-backend-test/app/order/domain/entity"
)

type CartRepository interface {
	AddCart(*entity.Cart) error
	GetAllCarts() ([]*entity.Cart, error)
	GetCartByUserID(string) (*entity.Cart, error)
	UpdateCart(*entity.Cart) error
	DeleteCart(string) error
}