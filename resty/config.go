package resty

import "time"

type Config struct {
	Debug bool `mapstructure:"debug"`

	Resty struct {
		Trace     bool              `mapstructure:"trace"`
		URL       string            `mapstructure:"url"`
		UserAgent string            `mapstructure:"user_agent"`
		Headers   map[string]string `mapstructure:"headers"`
		Timeout   time.Duration     `mapstructure:"timeout" validate:"required" default:"10s"`

		Auth *struct {
			Basic *struct {
				Username string `mapstructure:"username" validate:"required"`
				Password string `mapstructure:"password" validate:"required"`
			} `mapstructure:"basic"`

			Token *struct {
				Scheme string `mapstring:"scheme" validate:"required" default:"bearer"`
				Value  string `mapstructure:"value" validate:"required"`
			}
		} `mapstructure:"auth"`

		Retry struct {
			Count int `mapstructure:"count" default:"0"`
		} `mapstructure:"retry"`
	} `mapstructure:"resty"`
}

func NewConfig() *Config {
	return &Config{}
}
