package ydb

type Config struct {
	YDB struct {
		DSN string `mapstructure:"dsn" validate:"required" default:"grpc://localhost:2136/local"`

		Credentials struct {
			AccessToken string `mapstructure:"access_token"`
			Login       string `mapstructure:"login"`
			Password    string `mapstructure:"password"`
		} `mapstructure:"credentials"`
	} `mapstructure:"ydb"`
}

func NewConfig() *Config {
	return &Config{}
}
