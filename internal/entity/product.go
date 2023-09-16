package entity

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name,omitempty" validate:"required,min=3,max=250"`
	Slug        string    `json:"slug,omitempty"`
	Description string    `json:"description,omitempty" validate:"required,min=3,max=500"`
	Price       float64   `json:"price,omitempty" validate:"required"`
	Variety     string    `json:"variety,omitempty" validate:"required"`
	Rating      int       `json:"rating,omitempty" validate:"required,min=0,max=10"`
	Stock       int64     `json:"stock,omitempty" validate:"required"`
	IsActive    bool      `json:"is_active,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
