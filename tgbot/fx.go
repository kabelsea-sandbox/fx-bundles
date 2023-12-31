package tgbot

import (
	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

const ModuleName = "tgbot"

var Module = func() fx.Option {
	return fx.Module(
		ModuleName,

		config.Provide(NewConfig),

		fx.Provide(
			NewBotAPI,
		),
	)
}
