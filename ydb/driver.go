package ydb

import (
	"context"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3"
)

func NewDriver(config *Config) (*ydb.Driver, error) {
	var (
		db      *ydb.Driver
		options = []ydb.Option{}
		err     error
	)

	// auth type
	switch {

	case config.YDB.Credentials.AccessToken != "":
		options = append(options,
			ydb.WithAccessTokenCredentials(config.YDB.Credentials.AccessToken),
		)

	case config.YDB.Credentials.Login != "" && config.YDB.Credentials.Password != "":
		options = append(options,
			ydb.WithStaticCredentials(config.YDB.Credentials.Login, config.YDB.Credentials.Password),
		)

	default:
		options = append(options,
			ydb.WithAnonymousCredentials(),
		)
	}

	if err != nil {
		return nil, err
	}

	if db, err = ydb.Open(context.Background(), config.YDB.DSN, options...); err != nil {
		return nil, fmt.Errorf("create ydb driver failed: %w", err)
	}

	return db, nil
}
