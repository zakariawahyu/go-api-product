package product

import (
	"github.com/zakariawahyu/go-api-product/internal/dto"
	"github.com/zakariawahyu/go-api-product/internal/entity"
)

type ProductUsecase interface {
	GetAll() ([]*entity.Product, error)
	Create(req dto.CreateProduct) (*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	Update(req dto.UpdateProduct, id int) (*entity.Product, error)
	SoftDelete(id int) (*entity.Product, error)
	HardDelete(id int) (*entity.Product, error)
}
