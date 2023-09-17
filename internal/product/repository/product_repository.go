package repository

import (
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

func (r *productRepository) GetAll() ([]*entity.Product, error) {
	product := []*entity.Product{}
	if err := r.db.Where("is_active = ?", true).Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) Create(product entity.Product) (*entity.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetByID(id int) (*entity.Product, error) {
	product := entity.Product{}

	err := r.db.Where("id = ? and is_active = ?", id, true).Take(&product).Error
	if err != nil {
		return nil, response.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) GetBySlug(slug string) (*entity.Product, error) {
	product := entity.Product{}

	err := r.db.Where("slug = ? and is_active = ?", slug, true).Take(&product).Error
	if err != nil {
		return nil, response.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) Update(product entity.Product) (*entity.Product, error) {
	if err := r.db.Where("id = ? and is_active = ?", product.ID, true).Updates(&product).First(&product); err.RowsAffected == 0 {
		return nil, response.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) SoftDelete(id int) (*entity.Product, error) {
	product := entity.Product{}

	if err := r.db.Model(&product).Where("id = ? and is_active = ?", id, true).Update("is_active", false); err.RowsAffected == 0 {
		return nil, response.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) HardDelete(id int) (*entity.Product, error) {
	product := entity.Product{}

	if err := r.db.Where("id = ? and is_active = ?", id, true).Delete(&product); err.RowsAffected == 0 {
		return nil, response.ErrNotFound
	}

	return &product, nil
}
