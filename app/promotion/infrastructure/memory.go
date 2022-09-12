package infrastructure

import (
	"fita-backend-test/app/promotion/data/datasource"
	"fita-backend-test/app/promotion/data/repository"
	"fita-backend-test/app/promotion/domain/entity"
)

type memory struct {
	data map[string] entity.Promotion
}

func NewInMemory() datasource.InMemory {
	return &memory{data: make(map[string] entity.Promotion)}
}

func (m *memory) Add(promotion entity.Promotion) error {

	m.data[promotion.GetID()] = promotion

	return nil
}

func (m *memory) GetAll() ([]entity.Promotion, error) {
	var promotions []entity.Promotion

	for _, promotion := range m.data {
		promotions = append(promotions, promotion)
	}

	return promotions, nil
}

func (m *memory) GetByID(id string) (entity.Promotion, error) {
	if promotion, exist := m.data[id]; exist {
		return promotion, nil
	}

	return entity.Promotion{}, repository.ErrPromotionNotFound
}