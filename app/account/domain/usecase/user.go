package usecase

import (
	"fita-backend-test/app/account/domain/entity"
	"fita-backend-test/app/account/domain/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(string) (entity.User, error)
}

type user struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) UserUsecase {
	return &user{repository}
}

func (u *user) GetAllUsers() ([]entity.User, error) {
	return u.repository.GetAllUsers()	
}

func (u *user) GetUserByID(id string) (entity.User, error) {
	user, err := u.repository.GetUserByID(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}