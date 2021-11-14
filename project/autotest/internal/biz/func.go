package biz

import (
	"net/http"
	"net/url"
)

func HeaderAppend(header , append http.Header) http.Header {
	for  hk, hv:= range append {
		for _, v := range hv {
			header.Add(hk, v)
		}
	}
	return header
}

func QueryParamsAppend(query QueryParams, values url.Values) *url.Values {
	for  p, v := range query {
		values.Add(p, v)
	}
	return &values
}

