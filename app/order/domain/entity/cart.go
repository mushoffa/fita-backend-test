package entity


type Cart struct {
	ID string
	UserID string
	Status string
	Items []*OrderItem
	TotalAmount float64
}