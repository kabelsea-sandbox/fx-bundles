package resty

import "time"

// Config config type.
type Config struct {
	Trace     bool              `mapstructure:"trace"`
	URL       string            `mapstructure:"url"`
	UserAgent string            `mapstructure:"user_agent" default:"gatefi/1.0.0"`
	Headers   map[string]string `mapstructure:"headers"`
	Timeout   time.Duration     `mapstructure:"timeout" default:"15s"`

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
}

// GlobalConfig type.
type GlobalConfig struct {
	Debug bool   `mapstructure:"debug"`
	Resty Config `mapstructure:"resty"`
}

// NewConfig construct.
func NewConfig() *GlobalConfig { return &GlobalConfig{} }
