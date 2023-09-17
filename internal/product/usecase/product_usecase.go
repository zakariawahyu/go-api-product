package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/zakariawahyu/go-api-product/internal/dto"
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

func (u *productUsecase) GetAll() ([]*entity.Product, error) {
	res, err := u.productRepo.GetAll()
	if err != nil {
		u.logger.Errorf("productRepo.GetAll %v", err)
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) Create(req dto.CreateProduct) (*entity.Product, error) {
	if err := u.validate.Struct(&req); err != nil {
		u.logger.Errorf("validate.Struct %v", err)
		return nil, err
	}

	_, err := u.productRepo.GetBySlug(slug.Make(req.Name))
	if err == nil {
		return nil, response.ErrConflict
	}

	product := entity.Product{
		Name:        req.Name,
		Slug:        slug.Make(req.Name),
		Description: req.Description,
		Price:       req.Price,
		Variety:     req.Variety,
		Rating:      req.Rating,
		Stock:       req.Stock,
		IsActive:    true,
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

func (u *productUsecase) Update(req dto.UpdateProduct, id int) (*entity.Product, error) {
	if err := u.validate.Struct(&req); err != nil {
		u.logger.Errorf("validate.Struct %v", err)
		return nil, err
	}

	product := entity.Product{
		ID:          int64(id),
		Name:        req.Name,
		Slug:        slug.Make(req.Name),
		Description: req.Description,
		Price:       req.Price,
		Variety:     req.Variety,
		Rating:      req.Rating,
		Stock:       req.Stock,
		IsActive:    true,
	}

	res, err := u.productRepo.Update(product)
	if err != nil {
		u.logger.Errorf("productRepo.Update %v", err)
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) SoftDelete(id int) (*entity.Product, error) {
	res, err := u.productRepo.SoftDelete(id)
	if err != nil {
		u.logger.Errorf("productRepo.SoftDelete")
		return nil, err
	}

	return res, nil
}

func (u *productUsecase) HardDelete(id int) (*entity.Product, error) {
	res, err := u.productRepo.HardDelete(id)
	if err != nil {
		u.logger.Errorf("productRepo.Delete")
		return nil, err
	}

	return res, nil
}
