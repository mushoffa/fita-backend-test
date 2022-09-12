package promotion

import (
	"fita-backend-test/app/promotion/data/repository"
	"fita-backend-test/app/promotion/domain/usecase"
	"fita-backend-test/app/promotion/infrastructure"
)

type PromotionService interface {
	usecase.PromotionUsecase
}

func NewPromotionService() PromotionService {
	memory := infrastructure.NewInMemory()
	repository := repository.NewPromotionRepository(memory)

	return usecase.NewPromotionUsecase(repository)
}