package middlewares

import (
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/http/controllers"
)

var (
	loggerMiddlewareSkippedLinks = []string{
		controllers.HealthCheckPath,
		"/metrics",
	}
)

var LoggerMiddleware = func(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			Skipper: func(c echo.Context) bool {
				return slices.Contains(loggerMiddlewareSkippedLinks, c.Request().URL.String())
			},
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger.Debug("http_request",
					zap.String("uri", v.URI),
					zap.Int("status", v.Status),
				)
				return nil
			},
		},
	)
}
