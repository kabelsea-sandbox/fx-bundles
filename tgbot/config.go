package tgbot

type Config struct {
	Debug bool `mapstructure:"debug"`

	Telegram struct {
		Bot struct {
			Token string `mapstructure:"token" validate:"required"`
		} `mapstructure:"bot"`
	} `mapstructure:"telegram"`
}

func NewConfig() *Config {
	return &Config{}
}
