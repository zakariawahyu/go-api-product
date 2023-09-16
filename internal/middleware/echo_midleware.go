package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-product/config"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"time"
)

type EchoMiddleware interface {
	RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type echoMiddleware struct {
	log logger.Logger
	cfg *config.Config
}

func NewEchoMiddleware(log logger.Logger, cfg *config.Config) *echoMiddleware {
	return &echoMiddleware{log: log, cfg: cfg}
}

func (m *echoMiddleware) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		err := next(ctx)
		if err != nil {
			ctx.Error(err)
		}

		req := ctx.Request()
		res := ctx.Response()
		status := res.Status
		size := res.Size
		s := time.Since(start).String()
		requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)

		m.log.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			requestID, req.Method, req.URL, status, size, s)

		return err
	}
}
