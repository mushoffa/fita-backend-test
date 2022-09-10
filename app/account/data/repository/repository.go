package repository

import (
	"errors"
	"sync"

	"fita-backend-test/app/account/data/datasource"
	"fita-backend-test/app/account/domain/entity"
	"fita-backend-test/app/account/domain/repository"
)

var (
	ErrUserNotFound = errors.New("The user was not found in the database")
	ErrUserAlreadyExist = errors.New("The user with given ID already exists in database")
)

type user struct {
	sync.Mutex
	memory datasource.InMemory
}

func NewUserRepository(memory datasource.InMemory) domain.UserRepository {
	repository := &user{memory: memory}

	repository.initialize()

	return repository
}

func (r *user) initialize() {
	users := []entity.User {
		entity.User{"bac6dd0d-eaaf-43d6-9d09-da99a2b05f14", "UserA"},
		entity.User{"f749e3d0-afde-4617-b9c7-16ccbc1a56c1", "UserB"},
		entity.User{"acfdffb5-852e-4773-b3bf-7095db19dd06", "UserC"},
	}

	for _, user := range users {
		r.memory.Add(user)
	}
}

func (r *user) AddUser(user entity.User) error {
	r.Lock()
	defer r.Unlock()

	if _, err := r.memory.GetByID(user.GetID()); err == nil {
		return ErrUserAlreadyExist
	}

	if err := r.memory.Add(user); err != nil {
		return err
	}

	return nil
}

func (r *user) GetAllUsers() ([]entity.User, error) {
	users, err := r.memory.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *user) GetUserByID(id string) (entity.User, error) {
	user, err := r.memory.GetByID(id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}