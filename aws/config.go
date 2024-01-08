package aws

type Config struct {
	AWS struct {
		Endpoint string `mapstructure:"endpoint"`
		Region   string `mapstructure:"region" validate:"required" default:"us-west-1"`

		// Auth
		AccessKeyID     string `mapstructure:"access_key_id" validate:"required"`
		SecretAccessKey string `mapstructure:"secret_access_key" validate:"required"`

		// SQS
		SQS struct{} `mapstructure:"sqs"`
	} `mapstructure:"aws"`
}

func NewConfig() *Config {
	return &Config{}
}
