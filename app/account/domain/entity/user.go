package entity

type User struct {
	ID string
	Name string
}

func (e User) GetID() string {
	return e.ID
}