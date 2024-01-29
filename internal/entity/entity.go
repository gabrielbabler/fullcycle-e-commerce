package entity

import "github.com/google/uuid"

type Category struct {
	ID string
	Name string
}

func NewCategory(name: string) *Category {
	return &Category {
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID string,
	Name string,
	Descriptio string,
	Price float64,
	CategoryID string,
	ImageURL string
}