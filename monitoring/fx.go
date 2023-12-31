package monitoring

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/kabelsea-sandbox/fx-bundles/monitoring/controllers"
)

var ModuleName = "monitoring"

var Module = fx.Module(
	ModuleName,

	fx.Provide(
		fx.Annotate(
			NewExporter,

			fx.OnStop(func(ctx context.Context, provider *metric.MeterProvider) {
				_ = provider.Shutdown(ctx)
			}),
		),
	),

	fx.Invoke(
		controllers.NewMetrics,
	),

	// force
	fx.Invoke(func(_ *prometheus.Exporter) {}),

	// inject named logger
	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger {
			return logger.Named(ModuleName)
		},
	),
)
