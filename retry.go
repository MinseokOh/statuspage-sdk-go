package statuspage

import (
	"context"
	"net/http"
	"time"

	"github.com/avast/retry-go/v4"
)

type RetryOption func([]retry.Option) []retry.Option

type RequestOption func(*RequestConfig)

type RequestConfig struct {
	RetryOptions []retry.Option
}

func WithRetry(retryOpts ...RetryOption) RequestOption {
	return func(rc *RequestConfig) {
		opts := []retry.Option{}
		for _, opt := range retryOpts {
			opts = opt(opts)
		}
		rc.RetryOptions = opts
	}
}

func WithDefaultRetryOption() RequestOption {
	return WithRetry(
		WithAttempts(4),
		WithDelay(retry.BackOffDelay),
		WithDelayType(retry.BackOffDelay),
	)
}

func WithAttempts(attempts uint) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.Attempts(attempts))
	}
}

func WithDelay(delayType retry.DelayTypeFunc) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.DelayType(delayType))
	}
}

func WithDelayType(delayType retry.DelayTypeFunc) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.DelayType(delayType))
	}
}

func WithFixedDelay(delay time.Duration) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.Delay(delay))
	}
}

func WithBackoffDelay(initialDelay time.Duration, maxDelay time.Duration) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts,
			retry.DelayType(retry.BackOffDelay),
			retry.Delay(initialDelay),
			retry.MaxDelay(maxDelay),
		)
	}
}

func WithContext(ctx context.Context) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.Context(ctx))
	}
}

func WithOnRetry(onRetryFunc func(n uint, err error)) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.OnRetry(onRetryFunc))
	}
}

// WithRetryIf sets a custom retry condition function
func WithRetryIf(retryFunc func(*http.Response, error) bool) RetryOption {
	return func(opts []retry.Option) []retry.Option {
		return append(opts, retry.RetryIf(func(err error) bool {
			if httpErr, ok := err.(*HTTPError); ok {
				return retryFunc(httpErr.Response, httpErr.Err)
			}
			return retryFunc(nil, err)
		}))
	}
}

// HTTPError represents an HTTP error with response details and underlying error
type HTTPError struct {
	Response *http.Response
	Err      error
}

// Error implements the error interface for HTTPError
func (e *HTTPError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	if e.Response != nil {
		return e.Response.Status
	}
	return "unknown HTTP error"
}

func DefaultRetryableFunc(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}

	if resp == nil {
		return false
	}

	switch resp.StatusCode {
	case http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		StatusTooManyRequestsEnhanceYourCalm,
		http.StatusTooManyRequests:
		return true
	default:
		return false
	}
}

func (c *Client) DoWithOptions(ctx context.Context, req *http.Request, v interface{}, opts ...RequestOption) (*Response, error) {
	config := &RequestConfig{}
	for _, opt := range opts {
		opt(config)
	}

	if len(config.RetryOptions) == 0 {
		return c.Do(ctx, req, v)
	}

	return c.doWithRetry(ctx, req, v, config.RetryOptions)
}

func (c *Client) doWithRetry(ctx context.Context, req *http.Request, v interface{}, retryOptions []retry.Option) (*Response, error) {
	var lastResp *Response
	var lastErr error

	err := retry.Do(func() error {
		reqCopy := c.cloneRequest(ctx, req)
		resp, err := c.doRequest(ctx, reqCopy, v)
		lastResp = resp
		lastErr = err

		if err != nil {
			return &HTTPError{Response: resp.Response, Err: err}
		}

		return nil
	}, retryOptions...)

	if err != nil {
		return lastResp, lastErr
	}

	return lastResp, nil
}

func (c *Client) cloneRequest(ctx context.Context, req *http.Request) *http.Request {
	reqCopy := req.Clone(ctx)

	if req.Body != nil && req.GetBody != nil {
		body, err := req.GetBody()
		if err == nil {
			reqCopy.Body = body
		}
	}

	return reqCopy
}
