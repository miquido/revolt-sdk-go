package revolt

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("trackingId", "appCode", "secret")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestClientSendMethod(t *testing.T) {
	mockClient := NewMockClient(
		func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
				Header:     make(http.Header),
			}
		},
	)

	client, err := NewClient("trackingId", "appCode", "secret", WithCustomClient(mockClient))
	assert.Nil(t, err)
	assert.NotNil(t, client)

	event, err := NewEvent("test.event.type", struct {
		Value1 string
		Value2 string
	}{
		Value1: "string",
		Value2: "string2",
	})
	assert.Nil(t, err)

	resp, err := client.SendEvent(event)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
