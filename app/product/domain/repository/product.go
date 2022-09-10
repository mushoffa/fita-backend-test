package domain

import (
	"fita-backend-test/app/product/domain/entity"
)

type ProductRepository interface {
	AddProduct(entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	GetProductBySKU(string) (*entity.Product, error)
	UpdateProduct(entity.Product) error
}