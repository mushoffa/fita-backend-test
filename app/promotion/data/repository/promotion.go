package repository

import (
	"errors"

	"fita-backend-test/app/promotion/data/datasource"
	"fita-backend-test/app/promotion/domain/entity"
	"fita-backend-test/app/promotion/domain/repository"

	"github.com/google/uuid"
)

const (
	BXGY = "BXGY"
	BXPY = "BXPY"
	BXDY = "BXDY"

	DescBXGY = "Buy Minimum Product X, Get Product Y"
	DescBXPY = "Buy Minimum Product Y, Only Pay Y"
	DescBXDY = "Buy Minimum Product Y, Get Discount of Y"
)
 
// Each sale of a MacBook Pro comes with a free Raspberry Pi B
// Buy 3 Google Homes for the price of 2
// Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

var (
	ErrPromotionNotFound = errors.New("The promotion was not found in the database")
)

type promotion struct {
	memory datasource.InMemory
}

func NewPromotionRepository(memory datasource.InMemory) domain.PromotionRepository {
	repository := &promotion{memory}

	repository.initialize()

	return repository
}

func (r *promotion) initialize() {
	promotions := []entity.Promotion {
		entity.Promotion{uuid.New().String(), "43N23P", 5399.99, BXGY, DescBXGY, 1, "234234", 0, 0}, 
		entity.Promotion{uuid.New().String(), "120P90", 49.99, BXPY, DescBXPY, 3, "", 2, 0}, 
		entity.Promotion{uuid.New().String(), "A304SD", 109.50, BXDY, DescBXDY, 3, "", 0, 10},
	}

	for _, promotion := range promotions {
		r.memory.Add(promotion)
	}
}

func (r *promotion) GetAllPromotions() ([]entity.Promotion, error) {
	promotions, err := r.memory.GetAll()
	if err != nil {
		return nil, err
	}

	return promotions, nil
}

func (r *promotion) GetPromotionBySKU(sku string) (entity.Promotion, error) {
	promotion, err := r.memory.GetByID(sku)
	if err != nil {
		return entity.Promotion{}, err
	}

	return promotion, nil
}