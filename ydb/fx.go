package ydb

import (
	"context"
	"database/sql"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"go.uber.org/fx"

	"github.com/kabelsea-sandbox/fx-bundles/config"
)

const ModuleName = "ydb"

var Module = func() fx.Option {
	return fx.Module(
		ModuleName,

		config.Provide(NewConfig),

		fx.Provide(
			fx.Annotate(
				NewDriver,

				fx.OnStop(
					func(ctx context.Context, db *ydb.Driver) error {
						return db.Close(ctx)
					},
				),
			),
		),

		fx.Provide(
			fx.Annotate(
				func(driver *ydb.Driver) *sql.DB {
					return sql.OpenDB(
						ydb.MustConnector(driver,
							ydb.WithAutoDeclare(),
						),
					)
				},

				fx.OnStop(
					func(db *sql.DB) error {
						return db.Close()
					},
				),
			),
		),
	)
}
