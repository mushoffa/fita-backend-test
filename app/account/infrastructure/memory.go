package infrastructure

import (
	"fita-backend-test/app/account/data/datasource"
	"fita-backend-test/app/account/data/repository"
	"fita-backend-test/app/account/domain/entity"
)

type memory struct {
	data map[string] entity.User
}

func NewInMemory() datasource.InMemory {
	return &memory{data: make(map[string] entity.User)}
}

func (m *memory) Add(user entity.User) error {

	m.data[user.GetID()] = user

	return nil
}

func (m *memory) GetAll() ([]entity.User, error) {
	var users []entity.User

	for _, user := range m.data {
		users = append(users, user)
	}

	return users, nil
}

func (m *memory) GetByID(id string) (entity.User, error) {
	if user, exist := m.data[id]; exist {
		return user, nil
	}

	return entity.User{}, repository.ErrUserNotFound
}