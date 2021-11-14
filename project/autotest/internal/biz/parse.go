package biz

import (
	"bytes"
	"log"
	"text/template"
)

func Parse(t *template.Template, text string, data interface{}) (bytes.Buffer, error) {
	var b bytes.Buffer
	t1, err := t.Parse(text)
	log.Println(">>>bbbbbbb: ", t, err)
	err = t1.Execute(&b, data)
	return b, err
}


func ParseString(t *template.Template, text string, data interface{}) (string, error) {
	b, err := Parse(t, text, data)
	return b.String(), err
}

func ParseHeaders(headers Headers, parse func(text string) string) Headers {
	for k, v := range headers {
		t := v
		t = parse(t)
		headers[k] = t
	}
	return headers
}

func ParseQueryParams(params QueryParams, parse func(text string) string) QueryParams {
	for k, v := range params {
		t := v
		t = parse(t)
		params[k] = t
	}
	return params
}


func ParseQueryParamsApiObject(obj *ApiObject, data interface{}) QueryParams {
	return ParseQueryParams(obj.QueryParams, func(text string) string {
		return ParseStringApiObject(obj, text, data)
	})
}

func ParseHeadersParamsApiObject(obj *ApiObject, data interface{}) Headers {
	return ParseHeaders(obj.Headers, func(text string) string {
		return ParseStringApiObject(obj, text, data)
	})
}

func ParseStringApiObject(obj *ApiObject, text string, data interface{}) string {
	if obj.Err() != nil {
		return text
	}
	obj.AutoNewTemplate()
	text, err := ParseString(obj.Template, text, data)
	obj.SetErr(err)
	return text
}
