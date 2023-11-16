package debugService

import (
	"net/http"
	"net/netip"
	"net/url"
)

type Response struct {
	Method    string      `json:"method"`
	RequestIP netip.Addr  `json:"requestIP"`
	Body      string      `json:"body,omitempty"`
	Query     url.Values  `json:"query,omitempty"`
	Header    http.Header `json:"header,omitempty"`
}
