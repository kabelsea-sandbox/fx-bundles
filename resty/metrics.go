package resty

import (
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdk "go.opentelemetry.io/otel/sdk/metric"
)

var (
	responseTimeoutHistogram metric.Int64Histogram
	responseStatusCounter    metric.Int64Counter
)

// RegisterMetrics to prometheus.
func RegisterMetrics(provider *sdk.MeterProvider) error {
	var (
		metric = provider.Meter("gatefi/resty")
		err    error
	)

	// register response timeout histrogram
	responseTimeoutHistogram, err = metric.Int64Histogram(
		"resty_response_timeout",
	)
	if err != nil {
		return err
	}

	// register response status counter
	responseStatusCounter, err = metric.Int64Counter(
		"resty_status",
	)
	if err != nil {
		return err
	}

	return nil
}

// Collect metrics from resty.Response.
func collect(r *resty.Response) error {
	var (
		ctx = r.Request.Context()

		opt = metric.WithAttributes(
			attribute.Key("domain").String(r.Request.RawRequest.Host),
			attribute.Key("path").String(r.Request.RawRequest.URL.Path),
			attribute.Key("method").String(r.Request.Method),
			attribute.Key("status").Int(r.RawResponse.StatusCode),
		)
	)

	// response counter
	if responseStatusCounter != nil {
		responseStatusCounter.Add(ctx, 1, opt)
	}

	// response timeout hidtogram
	if responseTimeoutHistogram != nil {
		responseTimeoutHistogram.Record(ctx, r.Time().Milliseconds(), opt)
	}

	return nil
}
