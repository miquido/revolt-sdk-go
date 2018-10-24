package revolt

import (
	"net/http"

	"github.com/pkg/errors"
)

const defaultMaxBatchSize = 20
const defaultRetryIntervalSecs = 5
const defaultMaxRetryIntervalSecs = 300
const defaultRevoltEndpoint = `https://api.revolt.rocks/api/v1`

var ErrApiEndpoint = errors.New("API endpoint cannot be empty")
var ErrClientCode = errors.New("Client code cannot be empty")
var ErrTrackingId = errors.New("Tracking id cannot be empty")
var ErrSecretKey = errors.New("Secret key cannot be empty")
var ErrMaxRetryInterval = errors.New("Max retry interval cannot exceed 300 seconds")
var ErrRetryInterval = errors.New("Retry interval cannot bet set to 0 seconds")
var ErrMaxBatchSize = errors.New("Max batch size cannot bet set to 0 seconds")

func newConfigClient(trackingId, appCode, secret string, opts ...ClientConfig) (*Client, error) {
	c := Client{
		trackingId:           trackingId,
		appCode:              appCode,
		secretKey:            secret,
		client:               http.DefaultClient,
		endpoint:             defaultRevoltEndpoint,
		maxBatchSize:         defaultMaxBatchSize,
		retryIntervalSecs:    defaultRetryIntervalSecs,
		maxRetryIntervalSecs: defaultMaxRetryIntervalSecs,
		disabled:             false,
	}

	for _, opt := range opts {
		opt(&c)
	}

	err := c.validate()
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c Client) validate() error {
	if c.endpoint == "" {
		return ErrApiEndpoint
	}

	if c.trackingId == "" {
		return ErrTrackingId
	}

	if c.appCode == "" {
		return ErrClientCode
	}

	if c.secretKey == "" {
		return ErrSecretKey
	}

	if c.maxRetryIntervalSecs > 300 {
		return ErrMaxRetryInterval
	}

	if c.retryIntervalSecs == 0 {
		return ErrRetryInterval
	}

	if c.maxBatchSize == 0 {
		return ErrMaxBatchSize
	}

	return nil
}

type ClientConfig func(c *Client)

// Specifies Client used to communication with Revolt service.
func WithCustomClient(httpClient *http.Client) ClientConfig {
	return func(c *Client) {
		c.client = httpClient
	}
}

// Specifies endpoint on which communication with Revolt service should take place.
func WithEndpoint(endpoint string) ClientConfig {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// Specifies maximum batch size of events that can be sent to Revolt API.
func withMaxBatchSize(maxBatchSize uint) ClientConfig {
	return func(c *Client) {
		c.maxBatchSize = maxBatchSize
	}
}

// Specifies first time interval in seconds to retry sending batch of events when any error occurs.
func withRetryIntervalSecs(retryIntervalSecs uint) ClientConfig {
	return func(c *Client) {
		c.retryIntervalSecs = retryIntervalSecs
	}
}

// Specifies maximum time interval in seconds to retry sending batch of events when any error occurs.
func withMaxRetryIntervalSecs(maxRetryIntervalSecs uint) ClientConfig {
	return func(c *Client) {
		c.maxRetryIntervalSecs = maxRetryIntervalSecs
	}
}

// Specifies maximum size of events that can be stored in queue when service does not respond.
func withOfflineQueue(offlineQueueSize uint) ClientConfig {
	return func(c *Client) {
		c.offlineQueue = offlineQueueSize
	}
}

// Specifies maximum number of seconds for Event to be stored in queue.
// After delay is up, all events in queue will be sent automatically.
func withEventDelay(eventDelay uint) ClientConfig {
	return func(c *Client) {
		c.eventDelay = eventDelay
	}
}
