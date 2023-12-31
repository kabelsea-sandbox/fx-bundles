package sentry

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/build"
)

type FxClientParam struct {
	fx.In

	Config *Config
	Build  *build.Context `optional:"true"`
}

// NewClient construct.
func NewClient(p FxClientParam) (*sentry.Client, error) {
	var (
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: p.Config.Sentry.SkipSSLVerification, //nolint:gosec
				},
			},
		}

		opts = sentry.ClientOptions{
			Debug:            p.Config.Debug,
			Dsn:              p.Config.Sentry.DSN,
			SampleRate:       p.Config.Sentry.SampleRate,
			Release:          p.Config.Sentry.Release.Service,
			Dist:             p.Config.Sentry.Release.Version,
			Environment:      p.Config.Sentry.Release.Environment,
			AttachStacktrace: true,
			HTTPClient:       httpClient,
		}

		client *sentry.Client
		err    error
	)

	if serverName, err := os.Hostname(); err == nil {
		opts.ServerName = serverName
	}

	if b := p.Build; b != nil {
		opts.Release = b.Service
		opts.Dist = b.Version
		opts.Environment = b.Environment
	}

	hub := sentry.CurrentHub()

	if client, err = sentry.NewClient(opts); err != nil {
		return nil, err
	}

	hub.BindClient(client)

	return client, nil
}
