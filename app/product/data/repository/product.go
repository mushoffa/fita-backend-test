package repository

import (
	"errors"
	"sync"

	"fita-backend-test/app/product/data/datasource"
	"fita-backend-test/app/product/domain/entity"
	"fita-backend-test/app/product/domain/repository"
)

var (
	ErrProductNotFound = errors.New("The product was not found in the database")
	ErrProductAlreadyExist = errors.New("The product with given SKU already exists in database")
)

type product struct {
	sync.Mutex
	memory datasource.InMemory
}

func NewProductRepository(memory datasource.InMemory) domain.ProductRepository {
	repository := &product{memory: memory}

	repository.initialize()

	return repository
}

// 120P90
// Google Home
// $49.99
// 10


// 43N23P
// MacBook Pro
// $5,399.99
// 5


// A304SD
// Alexa Speaker
// $109.50
// 10


// 234234
// Raspberry Pi B
// $30.00
// 2
func (r *product) initialize() {
	products := []entity.Product{
		entity.Product{"120P90", "Google Home", 49.99, 10},
		entity.Product{"43N23P", "MacBook Pro", 5399.99, 5},
		entity.Product{"A304SD", "Alexa Speaker", 109.50, 10},
		entity.Product{"234234", "Raspberry Pi B", 30.00, 2},
	}

	for _, product := range products {
		r.memory.Add(product)
	}
}

func (r *product) AddProduct(product entity.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, err := r.memory.GetByID(product.SKU); err == nil {
		return ErrProductAlreadyExist
	}

	if err := r.memory.Add(product); err != nil {
		return err
	}

	return nil
}

func (r *product) GetAllProducts() ([]entity.Product, error) {

	products, err := r.memory.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *product) GetProductBySKU(sku string) (*entity.Product, error) {

	result, err := r.memory.GetByID(sku)
	if err != nil {
		return nil, err
	}

	product := entity.Product{
		SKU: result.SKU,
		Name: result.Name,
		Price: result.Price,
		Qty: result.Qty,
	}

	return &product, nil
}

func (r *product) UpdateProduct(product entity.Product) error {
	r.Lock()
	defer r.Unlock()

	if err := r.memory.Update(product); err != nil {
		return err
	}

	return nil
}