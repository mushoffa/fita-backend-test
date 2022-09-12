package domain

import (
	"fita-backend-test/app/promotion/domain/entity"
)

type PromotionRepository interface {
	GetAllPromotions() ([]entity.Promotion, error)
	GetPromotionBySKU(string) (entity.Promotion, error)
}