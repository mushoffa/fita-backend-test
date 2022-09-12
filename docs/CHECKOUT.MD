# Checkout Flow
Please add item to cart as a prerequisite condition to perform checkout process, otherwise the system will give error response. The checkout API has four usecases
1. Successfully checkout order item(s)
2. Insufficient promotion product quantity in stock (success usecase with condition)
3. No cart to checkout
4. Insufficient order item quantity in stock

## Success Checkout
![](checkout/flow_checkout_success.jpeg)

## Success Checkout with Condition
![](checkout/flow_checkout_success_insufficient_promotion_stock.jpeg)

## No Cart to Checkout
![](checkout/flow_checkout_error_cart_not_found.jpeg)

## Insufficient Order Item Stock
![](checkout/flow_checkout_error_insufficient_product_stock.jpeg)