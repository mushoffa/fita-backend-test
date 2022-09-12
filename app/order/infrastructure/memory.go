package infrastructure

import (
	"fita-backend-test/app/order/data/datasource"
	"fita-backend-test/app/order/data/repository"
	"fita-backend-test/app/order/domain/entity"
)

type memory struct {
	data map[string] *entity.Cart
}

func NewInMemory() datasource.InMemory {
	return &memory{data: make(map[string] *entity.Cart)}
}

func (m *memory) Add(cart *entity.Cart) error {

	m.data[cart.UserID] = cart

	return nil
}

func (m *memory) GetAll() ([]*entity.Cart, error) {
	var carts []*entity.Cart

	for _, cart := range m.data {
		carts = append(carts, cart)
	}

	return carts, nil
}

func (m *memory) GetByID(id string) (*entity.Cart, error) {
	if cart, exist := m.data[id]; exist {
		return cart, nil
	}

	return nil, repository.ErrCartNotFound
}

func (m *memory) Update(cart *entity.Cart) error {

	m.data[cart.UserID] = cart

	return nil
}

func (m *memory) Delete(id string) error {
	delete(m.data, id)
	return nil
}