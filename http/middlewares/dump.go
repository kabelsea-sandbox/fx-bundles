package middlewares

import (
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/http/controllers"
)

var (
	dumpMiddlewareSkippedLinks = []string{
		controllers.HealthCheckPath,
		"/metrics",
	}
)

var DumpMiddleware = func(logger *zap.Logger, skip bool) echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(
		middleware.BodyDumpConfig{
			Skipper: func(ctx echo.Context) bool {
				if skip {
					return true
				}

				if slices.Contains(dumpMiddlewareSkippedLinks, ctx.Request().URL.String()) {
					return false
				}

				return false
			},
			Handler: func(ctx echo.Context, reqBody, resBody []byte) {
				logger.Info("dump body",
					zap.String("path", ctx.Path()),

					// request
					zap.Dict("request",
						zap.String("id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
						zap.Any("headers", ctx.Request().Header),
						zap.ByteString("body", reqBody),
					),

					// response
					zap.Dict("response",
						zap.Any("headers", ctx.Response().Header()),
						zap.ByteString("body", resBody),
					),
				)
			},
		},
	)
}
