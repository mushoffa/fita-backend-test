package usecase

import (
	"errors"

	"fita-backend-test/app/promotion/domain/entity"
	"fita-backend-test/app/promotion/domain/repository"
)

var (
	ErrPromotionNotFound = errors.New("")
)

type PromotionUsecase interface {
	CheckPromotion(string) (entity.Promotion, error)
	GetAllPromotions() ([]entity.Promotion, error)
}

type promotion struct {
	repository domain.PromotionRepository
}

func NewPromotionUsecase(repository domain.PromotionRepository) PromotionUsecase {
	return &promotion{repository}
}

func (u *promotion) CheckPromotion(sku string) (entity.Promotion, error) {
	promotion, err := u.repository.GetPromotionBySKU(sku)
	if err != nil {
		return entity.Promotion{}, err
	}

	return promotion, nil
}

func (u *promotion) GetAllPromotions() ([]entity.Promotion, error) {
	return u.repository.GetAllPromotions()
}