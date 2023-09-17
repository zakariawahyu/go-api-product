package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-product/internal/dto"
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

// GetAllProduct Get all product
// @Tags Products
// @Summary Get all product
// @Description Get all product
// @Produce json
// @Success 200 {object} entity.Product
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product [get]
func (p *productHandlers) GetAllProduct(ctx echo.Context) error {
	res, err := p.productUsecase.GetAll()
	if err != nil {
		p.logger.Errorf("productUsecase.GetAll: %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

// CreateProduct Create product
// @Tags Products
// @Summary Create new product
// @Description Create new single product
// @Accept json
// @Produce json
// @Param body body dto.CreateProduct true "Create Product"
// @Success 200 {object} entity.Product
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product [post]
func (p *productHandlers) CreateProduct(ctx echo.Context) error {
	product := dto.CreateProduct{}

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

// GetProductByID Get product by id
// @Tags Products
// @Summary Get product by id
// @Description Get single product by id
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} entity.Product
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{id} [get]
func (p *productHandlers) GetProductByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := p.productUsecase.GetByID(id)
	if err != nil {
		p.logger.Errorf("productUsecase.GetByID: %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

// UpdateProduct Update product
// @Tags Products
// @Summary Update single product
// @Description Update single product by id
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Param body body dto.UpdateProduct true "Create Product"
// @Success 200 {object} entity.Product
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{id} [put]
func (p *productHandlers) UpdateProduct(ctx echo.Context) error {
	product := dto.UpdateProduct{}
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&product); err != nil {
		p.logger.Errorf("ctx.Bind: %v", err)
		panic(err)
	}

	res, err := p.productUsecase.Update(product, id)
	if err != nil {
		p.logger.Errorf("productUsecase.Update %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse(res))
}

// DeleteProduct Delete product by id
// @Tags Products
// @Summary Delete product by id
// @Description Delete product by id
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{id} [delete]
func (p *productHandlers) DeleteProduct(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := p.productUsecase.SoftDelete(id)
	if err != nil {
		p.logger.Errorf("productUsecase.Delete %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse("Success delete item"))
}

// HardDeleteProduct Hard Delete product by id
// @Tags Products
// @Summary Hard Delete product by id
// @Description Hard Delete product by id
// @Produce json
// @Param id path string true "product id"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{id}/hard-delete [delete]
func (p *productHandlers) HardDeleteProduct(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	p.logger.Info("AAAA")
	_, err := p.productUsecase.HardDelete(id)
	if err != nil {
		p.logger.Errorf("productUsecase.Delete %v", err)
		panic(err)
	}

	return ctx.JSON(http.StatusOK, response.NewSuccessResponse("Success delete item"))
}
