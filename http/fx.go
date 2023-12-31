package http

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/config"
	"github.com/kabelsea-sandbox/fx-bundles/http/controllers"
	"github.com/kabelsea-sandbox/fx-bundles/worker"
)

const ModuleName = "http"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	config.Provide(NewConfig),

	fx.Provide(
		NewServer,
		NewServerMux,
	),

	worker.Provide[ServerWorker](NewServerWorker),

	// controllers
	fx.Invoke(
		controllers.NewHealthCheckController,
	),

	// inject named logger
	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger {
			return logger.Named(ModuleName)
		},
	),
)
