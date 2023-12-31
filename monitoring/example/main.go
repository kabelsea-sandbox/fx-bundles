package main

import (
	"os"

	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/build"
	"github.com/kabelsea-sandbox/fx-bundles/config"
	"github.com/kabelsea-sandbox/fx-bundles/http"
	"github.com/kabelsea-sandbox/fx-bundles/logger"
	"github.com/kabelsea-sandbox/fx-bundles/monitoring"
	"github.com/kabelsea-sandbox/fx-bundles/worker"
)

func main() {
	os.Setenv("DEBUG", "True")

	app := fx.New(
		build.Module,
		logger.Module(),
		config.Module(),
		worker.Module(),
		monitoring.Module,
		http.Module(),
	)

	// run
	app.Run()
}
