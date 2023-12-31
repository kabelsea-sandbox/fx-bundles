package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapio"

	"github.com/kabelsea-sandbox/fx-bundles/build"
	"github.com/kabelsea-sandbox/fx-bundles/http/middlewares"
	"github.com/kabelsea-sandbox/fx-bundles/http/serializers"
	"github.com/kabelsea-sandbox/fx-bundles/http/validators"
)

// NewServerMux.
func NewServerMux(logger *zap.Logger, config *Config, buildContext *build.Context) *echo.Echo {
	e := echo.New()

	e.Debug = config.Debug

	e.HideBanner = true
	e.HidePort = true

	// setup zap logger as a echo logger backend
	e.Logger.SetOutput(
		&zapio.Writer{
			Log: logger,
		},
	)

	// custom json serializer via jsoniter
	e.JSONSerializer = serializers.NewJSONSerialzer(config.Debug)

	// custom validation
	e.Validator = validators.NewValidator()

	// custom error handler
	// e.HTTPErrorHandler = NewErrorHandler(config) // TODO

	var middlewares = []echo.MiddlewareFunc{
		// recovery
		middleware.Recover(),

		// request id
		middleware.RequestID(),

		// logger
		middlewares.LoggerMiddleware(logger),

		// dump
		middlewares.DumpMiddleware(logger, !config.Debug),

		// metrics
		echoprometheus.NewMiddleware("http"),

		// tracing
		middlewares.TraceMiddleware(buildContext.Service, buildContext.Version, buildContext.Environment),
	}

	e.Use(middlewares...)

	return e
}

// NewServer construct.
func NewServer(config *Config, logger *zap.Logger, e *echo.Echo) *http.Server {
	server := &http.Server{
		Addr:              config.HTTP.Bind,
		Handler:           e,
		ReadHeaderTimeout: 15 * time.Second,
	}

	server.SetKeepAlivesEnabled(config.HTTP.KeepAlive)

	return server
}
