package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

const ModuleName = "logger"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	config.Provide(NewConfig),

	fx.Provide(
		fx.Annotate(
			NewLogger,
			fx.OnStop(
				func(logger *zap.Logger) { _ = logger.Sync() },
			),
		),
	),

	// force
	fx.Invoke(func(_ *zap.Logger) {}),
)

// Module provided to fx with nop logger, for testing.
var ModuleNop = fx.Module(
	ModuleName,

	fx.Provide(zap.NewNop),

	fx.NopLogger,
)

// FxLogger returns fx.Option with custom fxvent.Logger.
var FxLogger = fx.WithLogger(
	func(logger *zap.Logger) fxevent.Logger {
		var (
			fxlogger = &fxevent.ZapLogger{
				Logger: logger,
			}
		)

		fxlogger.UseLogLevel(zapcore.DebugLevel)

		return fxlogger
	},
)
