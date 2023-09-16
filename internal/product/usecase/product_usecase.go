package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/zakariawahyu/go-api-product/internal/entity"
	"github.com/zakariawahyu/go-api-product/internal/product"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"github.com/zakariawahyu/go-api-product/pkg/response"
)

type productUsecase struct {
	productRepo product.ProductRepository
	logger      logger.Logger
	validate    *validator.Validate
}

func NewProductUsecase(productRepo product.ProductRepository, logger logger.Logger, validate *validator.Validate) *productUsecase {
	return &productUsecase{
		productRepo: productRepo,
		logger:      logger,
		validate:    validate,
	}
}

func (u *productUsecase) Create(product entity.Product) (*entity.Product, error) {
	if err := u.validate.Struct(&product); err != nil {
		u.logger.Errorf("validate.Struct %v", err)
		return nil, err
	}

	_, err := u.productRepo.GetBySlug(slug.Make(product.Name))
	if err == nil {
		return nil, response.ErrConflict
	}

	res, err := u.productRepo.Create(product)
	if err != nil {
		u.logger.Errorf("productRepo.Create %v", err)
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) GetByID(id int) (*entity.Product, error) {
	res, err := u.productRepo.GetByID(id)
	if err != nil {
		u.logger.Errorf("productRepo.GetByID %v", err)
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) Update(product entity.Product) (*entity.Product, error) {
	if err := u.validate.Struct(&product); err != nil {
		u.logger.Errorf("validate.Struct %v", err)
		return nil, err
	}

	res, err := u.productRepo.Update(product)
	if err != nil {
		u.logger.Errorf("productRepo.Update %v", err)
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) Delete(id int) (*entity.Product, error) {
	res, err := u.productRepo.Delete(id)
	if err != nil {
		u.logger.Errorf("productRepo.Delete")
		return nil, err
	}

	return res, nil
}
