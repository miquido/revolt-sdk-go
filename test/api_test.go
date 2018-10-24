package test

import (
	"net/http"
	"testing"

	"encoding/json"
	"io/ioutil"

	"bitbucket.org/miquido/revolt-client-go"
	"github.com/stretchr/testify/assert"
)

func TestSendEvent(t *testing.T) {
	tests := []struct {
		name  string
		event interface{}
	}{
		{
			name: "Struct event",
			event: struct {
				UserId       int    `json:"userId"`
				CreationType string `json:"creationType"`
				TaskId       int    `json:"taskId"`
				TaskDay      string `json:"taskDay"`
				Description  string `json:"description"`
				ProjectId    int    `json:"projectId"`
			}{
				UserId:       1,
				CreationType: "webbapp.test",
				TaskId:       1,
				TaskDay:      "2018-01-01",
				Description:  "short description",
				ProjectId:    1,
			},
		},
		{
			name: "map event",
			event: map[string]interface{}{
				"userId":      5,
				"description": "short description",
			},
		},
	}

	client, err := revolt.NewClient("revolttest", "com.quidlo.timesheets.backend.test", "ZjdMyTrmjVDC8Wr8")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event, err := revolt.NewEvent("test.event.type", test.event)
			assert.Nil(t, err)
			assert.NotNil(t, event)

			resp, err := client.SendEvent(event)
			assert.Nil(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var revoltResp revolt.ServiceResponse
			extractServiceResponse(t, resp, &revoltResp)

			assert.Equal(t, 1, revoltResp.EventsAccepted)
			assert.Nil(t, revoltResp.EventError)
		})
	}
}

func extractServiceResponse(t *testing.T, resp *http.Response, revoltResp *revolt.ServiceResponse) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(body, revoltResp)
	if err != nil {
		t.Error(err)
	}
}
