package repository

import (
	"github.com/gosimple/slug"
	"github.com/zakariawahyu/go-api-product/internal/entity"
	"github.com/zakariawahyu/go-api-product/pkg/response"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product entity.Product) (*entity.Product, error) {
	product.Slug = slug.Make(product.Name)
	product.IsActive = true
	if err := r.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetByID(id int) (*entity.Product, error) {
	product := entity.Product{}

	err := r.db.Where("id = ?", id).Take(&product).Error
	if err != nil {
		return nil, response.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) GetBySlug(slug string) (*entity.Product, error) {
	product := entity.Product{}

	err := r.db.Where("slug = ?", slug).Take(&product).Error
	if err != nil {
		return nil, response.ErrNotFound
	}

	return &product, nil
}
