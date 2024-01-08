package resty

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// NewClient constructor.
func NewClient(config *Config, logger *zap.Logger) (*resty.Client, error) {
	var (
		client = resty.New()
	)

	// set custom logger
	client.SetLogger(
		NewLoggerAdapter(logger),
	)

	// debug mode
	if config.Debug {
		client = client.SetDebug(config.Debug)
	}

	// enable tracing
	if config.Resty.Trace {
		client = client.EnableTrace()
	}

	// set base url
	if u := config.Resty.URL; u != "" {
		client = client.SetBaseURL(u)
	}

	// authentication
	if auth := config.Resty.Auth; auth != nil {
		if basic := auth.Basic; basic != nil {
			client = client.SetBasicAuth(basic.Username, basic.Password)
		}

		if token := auth.Token; token != nil {
			client = client.SetAuthScheme(token.Scheme)
			client = client.SetAuthToken(token.Value)
		}
	}

	// user-agent
	if ua := config.Resty.UserAgent; ua != "" {
		client = client.SetHeader("User-Agent", ua)
	}

	// headers
	if headers := config.Resty.Headers; len(headers) > 0 {
		client = client.SetHeaders(headers)
	}

	// timeout for request
	if t := config.Resty.Timeout; t.Seconds() > 0 {
		client = client.SetTimeout(t)
	}

	// retry count
	if c := config.Resty.Retry.Count; c > 0 {
		client = client.SetRetryCount(c)
	}

	return client, nil
}
