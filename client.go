package statuspage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
)

const (
	defaultBaseURL = "https://api.statuspage.io/v1"
	userAgent      = "statuspage-sdk-go/1.0.0"

	// StatusTooManyRequestsEnhanceYourCalm is a non-standard HTTP status code (420)
	// used by Twitter and Statuspage for rate limiting
	StatusTooManyRequestsEnhanceYourCalm = 420
)

// Client is the main HTTP client for interacting with the Statuspage API
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	apiKey     string
	userAgent  string

	// Default retry options for all requests
	defaultRetryOptions []retry.Option

	Pages             *PagesService
	Components        *ComponentsService
	ComponentGroups   *ComponentGroupsService
	Incidents         *IncidentsService
	IncidentUpdates   *IncidentUpdatesService
	Subscribers       *SubscribersService
	Metrics           *MetricsService
	PageAccessUsers   *PageAccessUsersService
	PageAccessGroups  *PageAccessGroupsService
	Templates         *TemplatesService
	StatusEmbedConfig *StatusEmbedConfigService
}

// ClientOption is a functional option for configuring the Client
type ClientOption func(*Client)

// service is a base struct for all API service implementations
type service struct {
	client *Client
}

// NewClient creates a new Statuspage API client with the provided API key and optional configuration
func NewClient(apiKey string, opts ...ClientOption) *Client {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	baseURL, _ := url.Parse("https://api.statuspage.io/v1/")

	c := &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		apiKey:     apiKey,
		userAgent:  userAgent,
	}

	// Apply client options
	for _, opt := range opts {
		opt(c)
	}

	c.Pages = &PagesService{client: c}
	c.Components = &ComponentsService{client: c}
	c.ComponentGroups = &ComponentGroupsService{client: c}
	c.Incidents = &IncidentsService{client: c}
	c.IncidentUpdates = &IncidentUpdatesService{client: c}
	c.Subscribers = &SubscribersService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.PageAccessUsers = &PageAccessUsersService{client: c}
	c.PageAccessGroups = &PageAccessGroupsService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.StatusEmbedConfig = &StatusEmbedConfigService{client: c}

	return c
}

// WithHTTPClient sets a custom HTTP client for the Statuspage client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets a custom base URL for the Statuspage API
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if parsedURL, err := url.Parse(baseURL); err == nil {
			c.baseURL = parsedURL
		}
	}
}

// WithUserAgent sets a custom user agent string for API requests
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// WithRetryOptions configures retry behavior for failed requests
func WithRetryOptions(retryOpts ...RetryOption) ClientOption {
	return func(c *Client) {
		opts := []retry.Option{}
		for _, opt := range retryOpts {
			opts = opt(opts)
		}
		c.defaultRetryOptions = opts
	}
}

// WithDefaultRetryConfig applies sensible default retry settings for production use
func WithDefaultRetryConfig() ClientOption {
	return WithRetryOptions(
		WithAttempts(4),
		WithDelayType(retry.BackOffDelay),
		WithRetryIf(DefaultRetryableFunc),
	)
}

// SetHTTPClient updates the underlying HTTP client used for API requests
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// SetBaseURL updates the base URL for API requests with validation
func (c *Client) SetBaseURL(baseURL string) error {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	c.baseURL = parsedURL
	return nil
}

// NewRequest creates a new HTTP request with proper headers and authentication for the Statuspage API
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", c.apiKey))

	return req, nil
}

// Do executes an HTTP request and decodes the response, with optional retry logic
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	// Use default retry options if configured
	if len(c.defaultRetryOptions) > 0 {
		return c.doWithRetry(ctx, req, v, c.defaultRetryOptions)
	}

	return c.doRequest(ctx, req, v)
}

// doRequest performs the actual HTTP request without retry logic
func (c *Client) doRequest(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

// Response wraps the standard HTTP response with additional Statuspage-specific functionality
type Response struct {
	*http.Response
}

// ErrorResponse represents an API error with HTTP response details and descriptive message
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"message"`
}

// Error implements the error interface for ErrorResponse
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// CheckResponse validates an HTTP response and returns an appropriate error for non-2xx status codes
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; http.StatusOK <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	switch r.StatusCode {
	case http.StatusBadRequest:
		errorResponse.Message = "Bad request"
	case http.StatusUnauthorized:
		errorResponse.Message = "Could not authenticate"
	case http.StatusForbidden:
		errorResponse.Message = "You are not authorized to access this resource"
	case http.StatusNotFound:
		errorResponse.Message = "The requested resource could not be found"
	case http.StatusUnprocessableEntity:
		errorResponse.Message = "Unprocessable entity"
	case StatusTooManyRequestsEnhanceYourCalm, http.StatusTooManyRequests:
		errorResponse.Message = "Rate limit exceeded"
	default:
		if errorResponse.Message == "" {
			errorResponse.Message = r.Status
		}
	}

	return errorResponse
}

// addOptions appends query parameters to a URL string based on the provided options struct
func addOptions(s string, opts interface{}) (string, error) {
	v, err := query(opts)
	if err != nil {
		return s, err
	}
	if v.Encode() == "" {
		return s, nil
	}
	if strings.Contains(s, "?") {
		s += "&" + v.Encode()
	} else {
		s += "?" + v.Encode()
	}
	return s, nil
}

// query converts various option structs to URL query parameters using reflection
func query(opts interface{}) (url.Values, error) {
	v := url.Values{}
	if opts == nil {
		return v, nil
	}

	// Simple reflection-based query parameter encoding
	// For production use, consider using github.com/google/go-querystring
	switch o := opts.(type) {
	case *ListOptions:
		if o.Page > 0 {
			v.Set("page", fmt.Sprintf("%d", o.Page))
		}
		if o.PerPage > 0 {
			v.Set("per_page", fmt.Sprintf("%d", o.PerPage))
		}
	case *IncidentListOptions:
		if o.Q != "" {
			v.Set("q", o.Q)
		}
		if o.Impact != "" {
			v.Set("impact", o.Impact)
		}
		if o.Status != "" {
			v.Set("status", o.Status)
		}
		if o.Page > 0 {
			v.Set("page", fmt.Sprintf("%d", o.Page))
		}
		if o.PerPage > 0 {
			v.Set("per_page", fmt.Sprintf("%d", o.PerPage))
		}
	case *SubscriberListOptions:
		if o.Q != "" {
			v.Set("q", o.Q)
		}
		if o.Sort != "" {
			v.Set("sort", o.Sort)
		}
		if o.Page > 0 {
			v.Set("page", fmt.Sprintf("%d", o.Page))
		}
		if o.PerPage > 0 {
			v.Set("per_page", fmt.Sprintf("%d", o.PerPage))
		}
	case *MetricDataListOptions:
		if o.From != nil {
			v.Set("from", o.From.Format(time.RFC3339))
		}
		if o.To != nil {
			v.Set("to", o.To.Format(time.RFC3339))
		}
	}

	return v, nil
}
