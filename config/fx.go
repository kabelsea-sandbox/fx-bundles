package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleName.
const ModuleName = "config"

// Module provided to fx.
var Module = func(opts ...Option) fx.Option {
	options := &options{}

	// load functional options
	for _, opt := range opts {
		opt(options)
	}

	return fx.Module(
		ModuleName,

		// public usage
		fx.Provide(
			NewLoader(options.optionalPrefix),
		),

		// private usage
		fx.Provide(
			viper.New,
			fx.Private,
		),
	)
}

// Provide config to fx.
func Provide[C Config](constructor func() *C) fx.Option {
	return fx.Options(
		fx.Provide(
			constructor,
		),

		fx.Decorate(
			func(logger *zap.Logger, loader Loader, config *C) *C {
				if err := loader.Load(config); err != nil {
					logger.Fatal("config load failed", zap.Error(err))
				}
				return config
			},
		),
	)
}
