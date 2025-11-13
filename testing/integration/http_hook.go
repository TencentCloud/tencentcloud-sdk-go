package integration

import (
	"net/http"
)

type httpHook struct {
	// all requests sent by this client
	Reqs []*http.Request

	// real RoundTripper to send request
	RoundTripper http.RoundTripper

	backup http.RoundTripper
}

func (h *httpHook) RoundTrip(request *http.Request) (*http.Response, error) {
	h.Reqs = append(h.Reqs, request)
	if h.RoundTripper != nil {
		return h.RoundTripper.RoundTrip(request)
	}
	return http.DefaultTransport.RoundTrip(request)
}

func installHttpHook() *httpHook {
	hh := &httpHook{
		RoundTripper: http.DefaultTransport,
		backup:       http.DefaultTransport,
	}
	http.DefaultTransport = hh
	return hh
}

func (h *httpHook) Uninstall() {
	http.DefaultTransport = h.backup
}
