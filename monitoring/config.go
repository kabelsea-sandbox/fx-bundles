package monitoring

// Config type.
type Config struct {
	Monitoring struct{} `mapstructure:"monitoring"`
}

// NewConfig construct.
func NewConfig() *Config {
	return &Config{}
}
