package biz

import (
	"net/http"
	"net/url"
)

func MergeHeaders(headers Headers, httpHeader http.Header) http.Header {
	for k, v := range headers {
		t := v
		httpHeader.Add(k, t)
	}
	return httpHeader
}

func MergeQueryParams(params QueryParams, values url.Values) *url.Values {
	for  p, v := range params {
		values.Add(p, v)
	}
	return &values
}
