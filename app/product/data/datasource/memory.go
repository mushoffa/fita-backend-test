package datasource

import (
	"fita-backend-test/app/product/domain/entity"
)

type InMemory interface {
	Add(entity.Product) error
	GetAll() ([]entity.Product, error)
	GetByID(string) (entity.Product, error)
	Update(entity.Product) error
}