package product

import (
	"fita-backend-test/app/product/data/repository"
	"fita-backend-test/app/product/domain/usecase"
	"fita-backend-test/app/product/infrastructure"
)

type ProductService interface {
	usecase.ProductUsecase
}

func NewProductService() ProductService {
	memory := infrastructure.NewInMemory()
	repository := repository.NewProductRepository(memory)

	return usecase.NewProductUsecase(repository)
}