package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-product/internal/product"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
)

type productHandlers struct {
	productUsecase product.ProductUsecase
	logger         logger.Logger
	group          *echo.Group
}

func NewProductHandlers(productUsecase product.ProductUsecase, logger logger.Logger, group *echo.Group) *productHandlers {
	return &productHandlers{
		productUsecase: productUsecase,
		logger:         logger,
		group:          group,
	}
}
