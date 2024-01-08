package resty

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

// ModuleName.
const ModuleName = "resty"

// Module provided to fx.
var Module = func() fx.Option {
	return fx.Module(
		ModuleName,

		config.Provide(NewConfig),

		fx.Provide(
			NewClient,
		),

		// inject named logger
		fx.Decorate(
			func(logger *zap.Logger) *zap.Logger {
				return logger.Named(ModuleName)
			},
		),
	)
}
