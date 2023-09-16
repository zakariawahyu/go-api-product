package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
