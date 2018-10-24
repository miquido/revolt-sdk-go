package revolt

import "net/http"

// RoundTripFunc is a type that will implement http.RoundTripper interface
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip is a method on RoundTripFunc that implements http.RoundTripper interface
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewMockClient returns *http.Client with Transport replaced to avoid making real calls
func NewMockClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
