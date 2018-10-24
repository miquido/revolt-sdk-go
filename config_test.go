package revolt

import (
	"testing"

	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestRevoltConfig(t *testing.T) {
	tests := []struct {
		name      string
		expClient *Client
		expError  error
		opts      []ClientConfig
	}{
		{
			name: "Default config",
			expClient: &Client{
				client:               http.DefaultClient,
				endpoint:             defaultRevoltEndpoint,
				maxBatchSize:         defaultMaxBatchSize,
				retryIntervalSecs:    defaultRetryIntervalSecs,
				maxRetryIntervalSecs: defaultMaxRetryIntervalSecs,
				disabled:             false,
				secretKey:            "secret",
				trackingId:           "trackingId",
				appCode:              "appCode",
			},
		},
		{
			name: "Invalid Endpoint",
			opts: []ClientConfig{
				WithEndpoint(""),
			},
			expError: ErrApiEndpoint,
		},
		{
			name: "Config with options",
			expClient: &Client{
				client:               http.DefaultClient,
				endpoint:             `https://api-random/api/test`,
				maxBatchSize:         defaultMaxBatchSize,
				retryIntervalSecs:    defaultRetryIntervalSecs,
				maxRetryIntervalSecs: defaultMaxRetryIntervalSecs,
				disabled:             false,
				secretKey:            "secret",
				trackingId:           "trackingId",
				appCode:              "appCode",
			},
			opts: []ClientConfig{
				WithEndpoint("https://api-random/api/test"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg, err := newConfigClient("trackingId", "appCode", "secret", test.opts...)

			if test.expError != nil {
				assert.Equal(t, test.expError, err)
			} else {
				assert.Equal(t, test.expClient, cfg)
			}
		})
	}
}
