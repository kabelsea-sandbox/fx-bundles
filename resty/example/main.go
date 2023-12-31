package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	rc "github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/build"
	"github.com/kabelsea-sandbox/fx-bundles/config"
	"github.com/kabelsea-sandbox/fx-bundles/http"
	"github.com/kabelsea-sandbox/fx-bundles/logger"
	"github.com/kabelsea-sandbox/fx-bundles/monitoring"
	"github.com/kabelsea-sandbox/fx-bundles/resty"
	"github.com/kabelsea-sandbox/fx-bundles/worker"
)

type (
	Config struct {
		Foo resty.Config `mapstructure:"foo"`
		Bar resty.Config `mapstructure:"bar"`
	}
)

func NewConfig() *Config { return &Config{} }

func main() {
	os.Setenv("DEBUG", "True")

	os.Setenv("FOO_USER_AGENT", "foo-ua")
	os.Setenv("FOO_TIMEOUT", "1s")

	os.Setenv("BAR_USER_AGENT", "bar-ua")
	os.Setenv("BAR_TIMEOUT", "30s")

	app := fx.New(
		build.Module,
		logger.Module(),
		config.Module(),
		monitoring.Module,
		worker.Module(),
		http.Module(),
		resty.Module,

		config.Provide(NewConfig),

		// use original resty client
		fx.Module(
			"original",

			fx.Invoke(
				func(client *rc.Client) {
					spew.Dump(
						// fmt.Sprintf("client: %v", client.GetClient()),
						client.R().Get("https://httpbin.org/uuid"),
					)
				},
			),
		),

		// use custom resty client, inherit from resty client
		fx.Module(
			"foo",

			fx.Decorate(
				func(config *Config, rc *rc.Client) *rc.Client {
					return resty.Prototype(config.Foo, rc)
				},
			),

			fx.Invoke(
				func(client *rc.Client) {
					spew.Dump(
						// fmt.Sprintf("client: %v", client.GetClient()),
						client.R().Get("https://httpbin.org/uuid"),
					)
				},
			),

			fx.Module(
				"foo_bar",

				fx.Decorate(
					func(config *Config, rc *rc.Client) *rc.Client {
						return resty.Prototype(config.Bar, rc)
					},
				),

				fx.Invoke(
					func(client *rc.Client) {
						spew.Dump(
							// fmt.Sprintf("client: %v", client.GetClient()),
							client.R().Get("https://httpbin.org/uuid"),
						)
					},
				),
			),
		),
	)

	// run
	app.Run()
}
