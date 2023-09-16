package product

import "github.com/zakariawahyu/go-api-product/internal/entity"

type ProductRepository interface {
	Create(product entity.Product) (*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	GetBySlug(slug string) (*entity.Product, error)
}
