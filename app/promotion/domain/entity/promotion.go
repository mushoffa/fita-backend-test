package entity

// Promotion Type
// BXGY = Buy X Get Y
// BXPY = Buy X Pay Y
// BXDY = Buy X Discount Y

type Promotion struct {
	ID string
	SKU string
	Price float64
	Type string
	Description string
	MinQty int
	FreeProduct string
	PriceOffset int
	Discount float64
}

func (e Promotion) GetID() string {
	return e.SKU
}