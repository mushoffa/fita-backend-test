package usecase

import (
	"errors"

	"fita-backend-test/app/product/domain/entity"
	"fita-backend-test/app/product/domain/repository"
)

var (
	ErrInvalidQuantity = errors.New("The product quantity must be a number greater than zero")
	ErrInsufficientQuantity = errors.New("Insufficient product quantity in stock")
)

type ProductUsecase interface {
	IncreaseProductQty(string, int) (entity.Product, error)
	CheckProduct(string, int) (entity.Product, error)
	CheckoutProduct(string, int) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}

type product struct {
	repository domain.ProductRepository
}

func NewProductUsecase(repository domain.ProductRepository) ProductUsecase {
	return &product{repository}
}

func (u *product) IncreaseProductQty(sku string, qty int) (entity.Product, error) {
	if(qty <= 0) {
		return entity.Product{}, ErrInvalidQuantity
	}

	product, err := u.repository.GetProductBySKU(sku)
	if err != nil {
		return entity.Product{}, err
	}

	product.Qty += qty

	if err := u.repository.UpdateProduct(*product); err != nil {
		return *product, err
	}

	return *product, nil
}

func (u *product) CheckProduct(sku string, qty int) (entity.Product, error) {
	if(qty <= 0) {
		return entity.Product{}, ErrInvalidQuantity
	}

	product, err := u.repository.GetProductBySKU(sku)
	if err != nil {
		return entity.Product{}, err
	}

	if(product.Qty < qty) {
		// return entity.Product{}, ErrInsufficientQuantity
		return *product, ErrInsufficientQuantity
	}

	return *product, nil
}

func (u *product) CheckoutProduct(sku string, qty int) (entity.Product, error) {
	if(qty <= 0) {
		return entity.Product{}, ErrInvalidQuantity
	}

	product, err := u.repository.GetProductBySKU(sku)
	if err != nil {
		return *product, err
	}

	
	if(product.Qty < qty) {
		return *product, ErrInsufficientQuantity
	}

	product.Qty -= qty

	if err := u.repository.UpdateProduct(*product); err != nil {
		return *product, err
	}

	return *product, nil
}

func (u *product) GetAllProducts() ([]entity.Product, error) {
	return u.repository.GetAllProducts()
}