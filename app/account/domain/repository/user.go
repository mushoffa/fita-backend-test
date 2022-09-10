package domain

import (
	"fita-backend-test/app/account/domain/entity"
)

type UserRepository interface {
	AddUser(entity.User) error
	GetAllUsers() ([]entity.User, error)
	GetUserByID(string) (entity.User, error)
}