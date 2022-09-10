package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fita-backend-test/graph/generated"
	"fita-backend-test/graph/model"
	// "fmt"
)

// AddCart is the resolver for the addCart field.
func (r *mutationResolver) AddCart(ctx context.Context, input *model.AddCart) (*model.Cart, error) {
	_, err := r.Gateway.AccountService.GetUserByID(input.UserID)
	if err != nil {
		return nil, err
	}

	

	cart := model.Cart{}

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

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	return nil, nil
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
