package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zakariawahyu/go-api-product/docs"
	"github.com/zakariawahyu/go-api-product/pkg/response"
	"net/http"
	"strings"
	"time"
)

func (s *server) runHttpServer() error {
	s.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"version":       s.cfg.AppVersion,
			"development":   s.cfg.Server.Development,
			"read_timeout":  s.cfg.Server.ReadTimeout,
			"write_timeout": s.cfg.Server.WriteTimeout,
		})
	})
	s.mapRoutes()

	s.echo.Server.ReadTimeout = time.Second * s.cfg.Server.ReadTimeout
	s.echo.Server.WriteTimeout = time.Second * s.cfg.Server.WriteTimeout
	s.echo.Server.MaxHeaderBytes = maxHeaderBytes

	return s.echo.Start(s.cfg.Server.Port)
}

func (s *server) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Product Api Services"
	docs.SwaggerInfo.Description = "Product REST API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1/product"
	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	s.echo.Use(s.mw.RequestLoggerMiddleware)

	s.echo.HTTPErrorHandler = response.NewHttpErrorHandler(response.NewErrorStatusCodeMaps()).Handler
	s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(middleware.BodyLimit(bodyLimit))
}
