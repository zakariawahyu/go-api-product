package usecase

import (
	"github.com/zakariawahyu/go-api-product/internal/product"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
)

type productUsecase struct {
	productRepo product.ProductRepository
	logger      logger.Logger
}

func NewProductUsecase(productRepo product.ProductRepository, logger logger.Logger) *productUsecase {
	return &productUsecase{
		productRepo: productRepo,
		logger:      logger,
	}
}
