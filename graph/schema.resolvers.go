package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fita-backend-test/graph/generated"
	"fita-backend-test/graph/model"
)

// Checkout is the resolver for the checkout field.
func (r *mutationResolver) Checkout(ctx context.Context, input *model.CheckoutRequest) (*model.Cart, error) {
	userCart, err := r.Gateway.OrderService.GetCart(input.UserID)
	if err != nil {
		return nil, err
	}

	var orderItems []*model.Product

	for _, item := range userCart.Items {
		product := &model.Product{
			Sku:   item.SKU,
			Name:  item.Name,
			Price: item.Price,
			Qty:   item.Qty,
		}

		orderItems = append(orderItems, product)
	}

	cart := model.Cart{
		ID:          userCart.ID,
		UserID:      userCart.UserID,
		TotalAmount: userCart.TotalAmount,
		OrderItems:  orderItems,
	}

	var freeProducts []*model.FreeItem

	for _, cartItem := range cart.OrderItems {
		// Checkout product in cart, mutates product quantity in product stock
		_, err = r.Gateway.ProductService.CheckoutProduct(cartItem.Sku, cartItem.Qty)
		// Decline checkout process where order quantity is greater than product in stock
		if err != nil {
			return nil, err
		}

		promotion, err := r.Gateway.PromotionService.CheckPromotion(cartItem.Sku)

		// Proceed when promotion is available in selected order item
		if err == nil {
			switch promotion.Type {
			// Buy 1 Get 1
			case "BXGY":
				if cartItem.Qty >= promotion.MinQty {

					// Check whether product stock is available prior to the promotion quantity
					promotionProduct, err := r.Gateway.ProductService.CheckoutProduct(promotion.FreeProduct, cartItem.Qty)
					if err != nil {
						// Limit the free item in line with product in stock
						freeProducts = append(freeProducts, &model.FreeItem{
							Sku:  promotionProduct.SKU,
							Name: promotionProduct.Name,
							Qty:  promotionProduct.Qty,
						})
					}

					freeProducts = append(freeProducts, &model.FreeItem{
						Sku:  promotionProduct.SKU,
						Name: promotionProduct.Name,
						Qty:  cartItem.Qty,
					})
				}

			// Buy 3 Pay 2
			case "BXPY":
				if cartItem.Qty >= promotion.MinQty {
					division := cartItem.Qty / promotion.MinQty
					remainder := cartItem.Qty % promotion.MinQty

					if division >= 1 {
						paidQty := (division * promotion.PriceOffset) + remainder
						deductionQty := cartItem.Qty - paidQty
						// deduction := r.Gateway.OrderService.RoundOrderAmount(promotion.Price * float64(deductionQty))						
						deduction := promotion.Price * float64(deductionQty)
						cart.Discount += deduction
					}
				}

			// Buy 4 Discount 10%
			case "BXDY":
				if cartItem.Qty > promotion.MinQty {
					discountPercentage := promotion.Discount / 100
					// discount := r.Gateway.OrderService.RoundOrderAmount(promotion.Price * float64(cartItem.Qty) * discountPercentage)
					discount := promotion.Price * float64(cartItem.Qty) * discountPercentage
					cart.Discount += discount
				}
			}
		}
	}

	cart.FreeItems = freeProducts
	cart.Discount = r.Gateway.OrderService.RoundOrderAmount(cart.Discount)
	cart.GrandTotal = r.Gateway.OrderService.RoundOrderAmount(cart.TotalAmount - cart.Discount)

	// Remove the card data when checkout is succeed
	// Normally you want to save this data in database
	r.Gateway.OrderService.DeleteCart(cart.UserID)

	return &cart, nil
}

// AddCart is the resolver for the addCart field.
func (r *mutationResolver) AddCart(ctx context.Context, input *model.AddCart) (*model.Cart, error) {
	orderQty := input.Qty

	user, err := r.Gateway.AccountService.GetUserByID(input.UserID)
	if err != nil {
		return nil, err
	}

	product, err := r.Gateway.ProductService.CheckProduct(input.Sku, orderQty)
	if err != nil {
		return nil, err
	}

	order, err := r.Gateway.OrderService.AddOrderItem(user.ID, orderQty, product.SKU, product.Name, product.Price, product.Qty)
	if err != nil {
		return nil, err
	}

	var orderItems []*model.Product

	for _, item := range order.Items {
		product := &model.Product{
			Sku:   item.SKU,
			Name:  item.Name,
			Price: item.Price,
			Qty:   item.Qty,
		}

		orderItems = append(orderItems, product)
	}

	cart := model.Cart{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		OrderItems:  orderItems,
	}

	return &cart, nil
}

// AddProductStock is the resolver for the addProductStock field.
func (r *mutationResolver) AddProductStock(ctx context.Context, input *model.AddProductStock) (*model.Product, error) {
	product := &model.Product{}

	result, err := r.Gateway.ProductService.IncreaseProductQty(input.Sku, input.Qty)
	if err != nil {
		return nil, err
	}

	product.Sku = result.SKU
	product.Name = result.Name
	product.Price = result.Price
	product.Qty = result.Qty
	return product, nil
}

// Carts is the resolver for the carts field.
func (r *queryResolver) Carts(ctx context.Context) ([]*model.Cart, error) {
	var carts []*model.Cart

	cartList, err := r.Gateway.OrderService.GetAllCarts()
	if err != nil {
		return nil, err
	}

	if len(cartList) > 0 {
		for _, cart := range cartList {
			_cart := &model.Cart{
				ID:          cart.ID,
				UserID:      cart.UserID,
				TotalAmount: cart.TotalAmount,
			}

			carts = append(carts, _cart)
		}
	}

	return carts, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	productList, err := r.Gateway.ProductService.GetAllProducts()
	if err != nil {
		return products, err
	}

	if len(productList) > 0 {
		for _, product := range productList {
			_product := &model.Product{
				Sku:   product.SKU,
				Name:  product.Name,
				Price: product.Price,
				Qty:   product.Qty,
			}

			products = append(products, _product)
		}
	}

	return products, nil
}

// Promotions is the resolver for the promotions field.
func (r *queryResolver) Promotions(ctx context.Context) ([]*model.Promotion, error) {
	var promotions []*model.Promotion

	promotionList, err := r.Gateway.PromotionService.GetAllPromotions()
	if err != nil {
		return nil, err
	}

	if len(promotionList) > 0 {
		for _, promotion := range promotionList {
			_promotion := &model.Promotion{
				ID:          promotion.ID,
				Sku:         promotion.SKU,
				Type:        promotion.Type,
				Description: promotion.Description,
				MinQty:      promotion.MinQty,
				FreeProduct: promotion.FreeProduct,
				PriceOffset: promotion.PriceOffset,
				Discount:    promotion.Discount,
			}

			promotions = append(promotions, _promotion)
		}
	}

	return promotions, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	userList, err := r.Gateway.AccountService.GetAllUsers()
	if err != nil {
		return users, err
	}

	if len(userList) > 0 {
		for _, user := range userList {
			_user := &model.User{
				ID:   user.ID,
				Name: user.Name,
			}

			users = append(users, _user)
		}
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	return nil, nil
}
