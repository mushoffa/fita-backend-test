package order

import (
	"fita-backend-test/app/order/data/repository"
	"fita-backend-test/app/order/domain/usecase"
	"fita-backend-test/app/order/infrastructure"
)

type OrderService interface {
	usecase.OrderUsecase
}

func NewOrderService() OrderService {
	memory := infrastructure.NewInMemory()
	repository := repository.NewCartRepository(memory)

	return usecase.NewOrderUsecase(repository)
}