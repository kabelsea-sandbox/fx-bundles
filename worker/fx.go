package worker

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "worker"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	fx.Provide(
		fx.Annotate(
			NewDispatcher,

			// on start fx hook
			fx.OnStart(func(ctx context.Context, d Dispatcher) error {
				return d.Run(ctx)
			}),

			// on stop fx hook
			fx.OnStop(func(ctx context.Context, d Dispatcher) error {
				return d.Shutdown(ctx)
			}),
		),
	),

	// inject named logger
	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger {
			return logger.Named(ModuleName)
		},
	),
)

// Provide service as a Worker.
func Provide[T Worker](constructor interface{}) fx.Option {
	return fx.Module(
		ModuleName,

		fx.Provide(
			fx.Annotate(
				constructor,
				fx.As(new(T)),
			),
		),

		fx.Invoke(
			func(d Dispatcher, worker T) {
				d.Register(worker)
			},
		),
	)
}
