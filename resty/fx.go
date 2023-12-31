package resty

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

// ModuleName.
const ModuleName = "resty"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	config.Provide(NewConfig),

	fx.Provide(
		NewClient,
	),

	fx.Invoke(
		RegisterMetrics,
	),

	// inject named logger
	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger {
			return logger.With(zap.Namespace(ModuleName))
		},
	),
)
