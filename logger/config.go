package logger

type Config struct {
	Debug  bool `mapstructure:"debug"`
	Logger struct {
		Level string `mapstructure:"level" validate:"required"` // TODO: enums ...
	} `mapstructure:"logger"`
}

func NewConfig() *Config {
	c := &Config{}

	c.Logger.Level = "info"

	return c
}
