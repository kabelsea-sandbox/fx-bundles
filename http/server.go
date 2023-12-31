package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapio"

	"github.com/kabelsea-sandbox/fx-bundles/build"
	"github.com/kabelsea-sandbox/fx-bundles/http/middlewares"
	"github.com/kabelsea-sandbox/fx-bundles/http/serializers"
	"github.com/kabelsea-sandbox/fx-bundles/http/validators"
)

type ServerMuxParam struct {
	fx.In

	Logger       *zap.Logger
	Config       *Config
	BuildContext *build.Context `optional:"true"`
}

// NewServerMux.
func NewServerMux(param ServerMuxParam) *echo.Echo {
	e := echo.New()

	e.Debug = param.Config.Debug

	e.HideBanner = true
	e.HidePort = true

	// setup zap logger as a echo logger backend
	e.Logger.SetOutput(
		&zapio.Writer{
			Log: param.Logger,
		},
	)

	// custom json serializer via jsoniter
	e.JSONSerializer = serializers.NewJSONSerialzer(param.Config.Debug)

	// custom validation
	e.Validator = validators.NewValidator()

	// custom error handler
	// e.HTTPErrorHandler = NewErrorHandler(config) // TODO

	var middlewaresList = []echo.MiddlewareFunc{
		// recovery
		middleware.Recover(),

		// request id
		middleware.RequestID(),

		// cors
		middleware.CORS(),

		// logger
		middlewares.LoggerMiddleware(param.Logger),

		// dump
		middlewares.DumpMiddleware(param.Logger, !param.Config.Debug),
	}

	// metrics
	if param.Config.HTTP.Metrics.Enabled {
		middlewaresList = append(middlewaresList,
			echoprometheus.NewMiddleware("http"),
		)
	}

	// tracing
	if param.Config.HTTP.Tracing.Enabled {
		var (
			name, version, environment string = "unknown", "unknown", "unknown"
		)

		if param.BuildContext != nil {
			name = param.BuildContext.Service
			version = param.BuildContext.Version
			environment = param.BuildContext.Environment
		}

		middlewaresList = append(middlewaresList,
			middlewares.TraceMiddleware(name, version, environment),
		)
	}

	e.Use(middlewaresList...)

	return e
}

// NewServer construct.
func NewServer(config *Config, logger *zap.Logger, e *echo.Echo) *http.Server {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Port),
		Handler:           e,
		ReadHeaderTimeout: 15 * time.Second,
	}

	server.SetKeepAlivesEnabled(config.HTTP.KeepAlive)

	return server
}
