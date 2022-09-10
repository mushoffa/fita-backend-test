package datasource

import (
	"fita-backend-test/app/order/domain/entity"
)

type InMemory interface {
	Add(*entity.Cart) error
	GetAll() ([]*entity.Cart, error)
	GetByID(string) (*entity.Cart, error)
	Update(*entity.Cart) error
}