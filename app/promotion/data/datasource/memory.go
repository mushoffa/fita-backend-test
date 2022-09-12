package datasource

import (
	"fita-backend-test/app/promotion/domain/entity"
)

type InMemory interface {
	Add(entity.Promotion) error
	GetAll() ([]entity.Promotion, error)
	GetByID(string) (entity.Promotion, error)
}