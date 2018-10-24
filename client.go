package revolt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Client is an Revolt API client. It can be configured in NewClient via ClientConfig fun
type Client struct {
	client               *http.Client
	endpoint             string
	appCode              string
	trackingId           string
	secretKey            string
	maxBatchSize         uint
	eventDelay           uint
	offlineQueue         uint
	retryIntervalSecs    uint
	maxRetryIntervalSecs uint
	disabled             bool
}

// NewClient returns an Revolt Client for use communicating with a Revolt service API configured with the given ClientConfig.
func NewClient(trackingId, appCode, secret string, opts ...ClientConfig) (*Client, error) {
	client, err := newConfigClient(trackingId, appCode, secret, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot initialize revolt Client config")
	}

	return client, nil
}

// SendEvent Sends an Event to Revolt API.
func (c *Client) SendEvent(event Event) (*http.Response, error) {
	payload, err := json.Marshal([]Event{event})
	if err != nil {
		return nil, err
	}

	req, err := c.createEventRequest(payload)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewEvent is a clients method wrapper around NewEvent function. It creates a new Event instance.
func (c *Client) NewEvent(eventType string, data interface{}) (Event, error) {
	return NewEvent(eventType, data)
}

// ServiceResponseFromHTTP is a clients method wrapper around ServiceResponseFromHTTP function.
// It extracts Revolt service response from an http response.
func (c *Client) ServiceResponseFromHTTP(response *http.Response) (*ServiceResponse, error) {
	return ServiceResponseFromHTTP(response)
}

// createEventRequest creates event from given payload and applies required request headers.
func (c *Client) createEventRequest(payload []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s/events", c.endpoint, c.trackingId, c.appCode), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.trackingId+":"+c.secretKey)))

	return req, nil
}
