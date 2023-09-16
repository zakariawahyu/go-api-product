package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-product/config"
	"github.com/zakariawahyu/go-api-product/internal/middleware"
	productsHandlerV1 "github.com/zakariawahyu/go-api-product/internal/product/delivery/http/v1"
	"github.com/zakariawahyu/go-api-product/internal/product/repository"
	"github.com/zakariawahyu/go-api-product/internal/product/usecase"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"gorm.io/gorm"
)

const (
	maxHeaderBytes = 1 << 20 // 2 KB
	gzipLevel      = 5
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
)

type server struct {
	log   logger.Logger
	cfg   *config.Config
	echo  *echo.Echo
	mysql *gorm.DB
	mw    middleware.EchoMiddleware
}

func NewServer(log logger.Logger, cfg *config.Config, mysql *gorm.DB) *server {
	return &server{
		log:   log,
		cfg:   cfg,
		echo:  echo.New(),
		mysql: mysql,
	}
}

func (s *server) Run() error {
	s.mw = middleware.NewEchoMiddleware(s.log, s.cfg)

	validate := validator.New()
	productRepository := repository.NewProductRepository(s.mysql)
	prouctUsecase := usecase.NewProductUsecase(productRepository, s.log, validate)
	v1 := s.echo.Group("api/v1")
	productHandlers := productsHandlerV1.NewProductHandlers(prouctUsecase, s.log, v1.Group("/product"))
	productHandlers.MapRoutes()

	if err := s.runHttpServer(); err != nil {
		s.log.Errorf(" s.runHttpServer: %v", err)
	}

	return nil
}
