package aws

import (
	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

const ModuleName = "aws"

var Module = func() fx.Option {
	return fx.Module(
		ModuleName,

		config.Provide(NewConfig),

		fx.Provide(
			NewEndpointConfig,
		),

		fx.Provide(
			NewSQS,
		),
	)
}
