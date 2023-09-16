package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-product/internal/entity"
	"github.com/zakariawahyu/go-api-product/internal/product"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"github.com/zakariawahyu/go-api-product/pkg/response"
	"net/http"
	"strconv"
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

func (p *productHandlers) CreateProduct(ctx echo.Context) error {
	product := entity.Product{}

	if err := ctx.Bind(&product); err != nil {
		p.logger.Errorf("ctx.Bind: %v", err)
		panic(err)
	}

	res, err := p.productUsecase.Create(product)
	if err != nil {
		p.logger.Errorf("productUsecase.Create: %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (p *productHandlers) GetProductByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := p.productUsecase.GetByID(id)
	if err != nil {
		p.logger.Errorf("productUsecase.GetByID: %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (p *productHandlers) UpdateProduct(ctx echo.Context) error {
	product := entity.Product{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&product); err != nil {
		p.logger.Errorf("ctx.Bind: %v", err)
		panic(err)
	}

	product.ID = int64(id)
	res, err := p.productUsecase.Update(product)
	if err != nil {
		p.logger.Errorf("productUsecase.Update %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

func (p *productHandlers) DeleteProduct(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := p.productUsecase.Delete(id)
	if err != nil {
		p.logger.Errorf("productUsecase.Delete %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse("Success delete item"))
}
