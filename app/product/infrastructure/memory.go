package infrastructure

import (
	"fita-backend-test/app/product/data/datasource"
	"fita-backend-test/app/product/data/repository"
	"fita-backend-test/app/product/domain/entity"
)

type memory struct {
	data map[string] entity.Product
}

func NewInMemory() datasource.InMemory {
	return &memory{data: make(map[string] entity.Product)}
}

func (m *memory) Add(product entity.Product) error {

	m.data[product.GetID()] = product

	return nil
}

func (m *memory) GetAll() ([]entity.Product, error) {
	var products []entity.Product

	for _, product := range m.data {
		products = append(products, product)
	}

	return products, nil
}

func (m *memory) GetByID(id string) (entity.Product, error) {
	if product, exist := m.data[id]; exist {
		return product, nil
	}

	return entity.Product{}, repository.ErrProductNotFound
}

func (m *memory) Update(product entity.Product) error {

	m.data[product.GetID()] = product

	return nil
}