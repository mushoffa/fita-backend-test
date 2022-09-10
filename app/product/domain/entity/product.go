package entity

type Product struct {
	SKU string
	Name string
	Price float64
	Qty int
}

func (e Product) GetID() string {
	return e.SKU
}