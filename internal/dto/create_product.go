package dto

type CreateProduct struct {
	Name        string  `json:"name,omitempty" validate:"required,min=3,max=250"`
	Description string  `json:"description,omitempty" validate:"required,min=3,max=500"`
	Price       float64 `json:"price,omitempty" validate:"required"`
	Variety     string  `json:"variety,omitempty" validate:"required"`
	Rating      int     `json:"rating,omitempty" validate:"required,min=0,max=5"`
	Stock       int64   `json:"stock,omitempty" validate:"required"`
}
