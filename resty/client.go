package resty

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// NewClient constructor.
func NewClient(config *GlobalConfig, logger *zap.Logger) (*resty.Client, error) {
	var (
		client = Prototype(config.Resty, resty.New())
	)

	// set custom logger
	client.SetLogger(
		NewLoggerAdapter(logger),
	)

	// debug mode
	if config.Debug {
		client = client.SetDebug(config.Debug)
	}

	return client, nil
}

// Prototype resty client with custom configuration.
func Prototype[T *resty.Client](config Config, rc *resty.Client) T {
	var (
		client = resty.New()
	)

	if rc != nil {
		client = Clone(rc)
	}

	// enable tracing
	if config.Trace {
		client = client.EnableTrace()
	}

	// set base url
	if u := config.URL; u != "" {
		client = client.SetBaseURL(u)
	}

	// authentication
	if auth := config.Auth; auth != nil {
		if basic := auth.Basic; basic != nil {
			client = client.SetBasicAuth(basic.Username, basic.Password)
		}

		if token := auth.Token; token != nil {
			client = client.SetAuthScheme(token.Scheme)
			client = client.SetAuthToken(token.Value)
		}
	}

	// user-agent
	if ua := config.UserAgent; ua != "" {
		client = client.SetHeader("User-Agent", ua)
	}

	// headers
	if headers := config.Headers; len(headers) > 0 {
		client = client.SetHeaders(headers)
	}

	// timeout for request
	if t := config.Timeout; t.Seconds() > 0 {
		client = client.SetTimeout(t)
	}

	// retry count
	if c := config.Retry.Count; c > 0 {
		client = client.SetRetryCount(c)
	}

	// collect metrics
	client = client.OnAfterResponse(func(_ *resty.Client, r *resty.Response) error {
		return collect(r)
	})

	return client
}

// Clone and return resty client.
func Clone(orig *resty.Client) *resty.Client {
	var (
		client = resty.New()
	)

	// debug
	client = client.SetDebug(orig.Debug)

	// url
	client = client.SetBaseURL(orig.BaseURL)

	// timeout
	client = client.SetTimeout(orig.GetClient().Timeout)

	// retry
	client = client.SetRetryCount(orig.RetryCount)

	// copy basic auth
	if u := orig.UserInfo; u != nil {
		client = client.SetBasicAuth(
			u.Username,
			u.Password,
		)
	}

	// copy token auth
	client = client.
		SetAuthScheme(orig.AuthScheme).
		SetAuthToken(orig.Token)

	// headers
	if h := orig.Header.Clone(); len(h) > 0 {
		for k, v := range h {
			if len(v) > 0 {
				client = client.SetHeader(k, v[0])
			}
		}
	}

	return client
}
