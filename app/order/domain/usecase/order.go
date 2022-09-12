package usecase

import (
	"errors"
	"math"

	"fita-backend-test/app/order/domain/entity"
	"fita-backend-test/app/order/domain/repository"

	"github.com/google/uuid"
)

var (
	ErrOrderItemExceedsStock = errors.New("New order item exceeds stock limit")
)

type OrderUsecase interface {
	CreateNewCart(string, string, string, float64, int) (entity.Cart, error)
	AddOrderItem(string, int, string, string, float64, int) (entity.Cart, error)
	GetCart(string) (*entity.Cart, error)
	GetAllCarts() ([]*entity.Cart, error)
	DeleteCart(string)
	RoundOrderAmount(float64) float64
}

type order struct {
	repository domain.CartRepository
}

func NewOrderUsecase(repository domain.CartRepository) OrderUsecase {
	return &order{repository}
}

func (u *order) CreateNewCart(userID, sku, name string, price float64, qty int) (entity.Cart, error) {
	var items []*entity.OrderItem

	totalAmount := u.RoundOrderAmount(price * float64(qty))

	items = append(items, &entity.OrderItem{sku,name,price,qty})

	cart := entity.Cart {
		ID: uuid.New().String(),
		UserID: userID,
		Status: "OPEN",
		Items: items,
		TotalAmount: totalAmount,
	}

	if err := u.repository.AddCart(&cart); err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

func (u *order) AddOrderItem(userID string, orderQty int, sku, name string, price float64, qtyInStock int) (entity.Cart, error) {
	isNewOrderItem := true

	cart, err := u.repository.GetCartByUserID(userID)
	// The user doesn't have a cart, then create a new one
	if err != nil {
		return u.CreateNewCart(userID, sku, name, price, orderQty)
	}

	totalOrderPrice := (price * float64(orderQty))
	

	// Check the order item list whether similar product is exist
	for _, item := range cart.Items {
		if(item.SKU == sku) {
			newOrderQty := item.Qty + orderQty
			if(newOrderQty > qtyInStock) {
				return entity.Cart{}, ErrOrderItemExceedsStock
			}
			item.Qty = newOrderQty
			// cart.TotalAmount += totalOrderPrice
			cart.TotalAmount = u.RoundOrderAmount(cart.TotalAmount + totalOrderPrice)
			isNewOrderItem = false
		}
	}

	if(isNewOrderItem) {
		orderItem := &entity.OrderItem {
			SKU: sku,
			Name: name,
			Price: price,
			Qty: orderQty,
		}	

		cart.Items = append(cart.Items, orderItem)
		// cart.TotalAmount += totalOrderPrice
		cart.TotalAmount = u.RoundOrderAmount(cart.TotalAmount + totalOrderPrice)
	}


	return *cart, nil
}

func (u *order) GetCart(userID string) (*entity.Cart, error) {
	cart, err := u.repository.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (u *order) GetAllCarts() ([]*entity.Cart, error) {
	carts, err := u.repository.GetAllCarts()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (u *order) DeleteCart(userID string) {
	u.repository.DeleteCart(userID)
}

func (u *order) RoundOrderAmount(totalOrderAmount float64) float64 {
	ratio := math.Pow(10, float64(2))
	totalRoundAmount := math.Round(totalOrderAmount*ratio) / ratio
	return totalRoundAmount
}