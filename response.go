package revolt

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

// ServiceResponse represents a response from Revolt service
type ServiceResponse struct {
	EventsAccepted int `json:"eventsAccepted"`
	EventError     *struct {
		EventOffset  int         `json:"eventOffset"`
		EventId      interface{} `json:"eventId"`
		ErrorCode    int         `json:"errorCode"`
		ErrorMessage string      `json:"errorMessage"`
	} `json:"eventError"`
}

// ServiceResponseFromHTTP extracts Revolt service struct response from an http response.
func ServiceResponseFromHTTP(response *http.Response) (*ServiceResponse, error) {
	var revoltResp ServiceResponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &revoltResp)
	if err != nil {
		return nil, err
	}

	return &revoltResp, nil
}