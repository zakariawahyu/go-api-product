package product

import "github.com/zakariawahyu/go-api-product/internal/entity"

type ProductUsecase interface {
	Create(product entity.Product) (*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	Update(product entity.Product) (*entity.Product, error)
	Delete(id int) (*entity.Product, error)
}
