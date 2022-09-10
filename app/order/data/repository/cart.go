package repository

import (
	"errors"
	"sync"

	"fita-backend-test/app/order/data/datasource"
	"fita-backend-test/app/order/domain/entity"
	"fita-backend-test/app/order/domain/repository"
)

var (
	ErrCartNotFound = errors.New("The cart was not found in the database")
)

type cart struct {
	sync.Mutex
	memory datasource.InMemory
}

func NewCartRepository(memory datasource.InMemory) domain.CartRepository {
	repository := &cart{memory:memory}

	repository.initialize()

	return repository
}

func (r *cart) initialize() {

}

func (r *cart) AddCart(cart *entity.Cart) error {
	r.Lock()
	defer r.Unlock()

	return nil
}

func (r *cart) GetAllCarts() ([]*entity.Cart, error) {
	carts, err := r.memory.GetAll()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *cart) GetCartByUserID(userID string) (*entity.Cart, error) {
	return nil, nil
}

func (r *cart) UpdateCart(cart *entity.Cart) error {
	r.Lock()
	defer r.Unlock()

	if err := r.memory.Update(cart); err != nil {
		return err
	}

	return nil
}