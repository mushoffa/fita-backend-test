package datasource

import (
	"fita-backend-test/app/account/domain/entity"
)

type InMemory interface {
	Add(entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(string) (entity.User, error)
}