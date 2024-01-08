package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func NewEndpointConfig(c *Config) (aws.Config, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           c.AWS.Endpoint,
			SigningRegion: c.AWS.Region,
		}, nil
	})

	return config.LoadDefaultConfig(
		context.Background(),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AWS.AccessKeyID, c.AWS.SecretAccessKey, "")),
		config.WithDefaultRegion(c.AWS.Region),
	)
}
