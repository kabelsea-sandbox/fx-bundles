package http

type Config struct {
	Debug bool `mapstructure:"debug"`

	Port int `mapstructure:"port" default:"8080"`

	HTTP struct {
		KeepAlive bool `mapstructure:"keep_alive" default:"true"`

		Tracing struct {
			Enabled bool `mapstructure:"enabled"`
		} `mapstructure:"tracing"`

		Metrics struct {
			Enabled bool `mapstructure:"enabled"`
		} `mapstructure:"metrics"`

		Router struct {
			Prefix string `mapstructure:"prefix"`
		} `mapstructure:"router"`
	} `mapstructure:"http"`
}

func NewConfig() *Config {
	return &Config{}
}
