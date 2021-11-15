package biz

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type ApiObject struct {
	Client *http.Client
	Name string
	Method string
	Url string
	QueryParams map[string]string
	Vars *Vars
	SetVar SetVars
	Assert Assert
	AssertResult bool
	Template *template.Template
	Headers Headers
	Request *http.Request
	Response *http.Response
	ResponseJsonBody interface{}
	ResponseBody string
	TemplateData map[string]interface{}
	err error
	errs []error
}
func (obj *ApiObject) AutoNewTemplate() {
	if obj.Err() != nil {
		return
	}
	//obj.Template = nil
	if obj.Template == nil {
		obj.Template = template.New("123")
	}
}


func (obj *ApiObject) Before()  {
	if obj.Err() != nil {
		return
	}
	data, err := ApiObjectConvertParseData(obj)
	obj.TemplateData = data
	err = errors.Wrap(err, "Before ApiObjectConvertParseData fail")
	obj.SetErr(err)
	// 1. 创建 request
	// 2. 解析 url，合并 url query
	// 3. 解析 queryParams， 合并 queryParams
	// 4. 解析headers， 合并 headers
	obj.Url = ParseStringApiObject("BeforeParseUrlVar", obj, obj.Url, data)
	obj.AutoNewRequest()

	queryParams := ParseQueryParamsApiObject(obj, data)
	values := MergeQueryParams(queryParams, obj.Request.URL.Query())
	obj.Request.URL.RawQuery = values.Encode()

	headers := ParseHeadersParamsApiObject(obj, data)
	obj.Request.Header = MergeHeaders(headers, obj.Request.Header)
}

func (obj *ApiObject) Execute()  {
	if obj.Err() != nil {
		return
	}
	response, err := obj.Client.Do(obj.Request)
	obj.Response = response
	err = errors.Wrap(err, "Execute Client.Do fail")
	obj.SetErr(err)

}

func (obj *ApiObject) After()  {
	if obj.Err() != nil {
		return
	}

}

func (obj *ApiObject) ParseVars()  {
	if obj.Err() != nil {
		return
	}
	data, err := ApiObjectConvertParseData(obj)
	err = errors.Wrap(err, "ParseVars ApiObjectConvertParseData fail")
	obj.SetErr(err)
	for k, v := range obj.SetVar {
		if strings.HasPrefix(k, "global") {
			t := strings.TrimPrefix(k, "global")
			obj.Vars.Global[t] = ParseStringApiObject("ParseVarsGlobal", obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".global") {
			t := strings.TrimPrefix(k, ".global")
			obj.Vars.Global[t] = ParseStringApiObject("ParseVars.Global", obj, v, data)
			continue
		}

		if strings.HasPrefix(k, "session") {
			t := strings.TrimPrefix(k, "session")
			obj.Vars.Session[t] = ParseStringApiObject("ParseVarsSession", obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".session") {
			t := strings.TrimPrefix(k, ".session")
			obj.Vars.Session[t] = ParseStringApiObject("ParseVars.Session", obj, v, data)
			continue
		}

		if strings.HasPrefix(k, "local") {
			t := strings.TrimPrefix(k, "local")
			obj.Vars.Local[t] = ParseStringApiObject("ParseVarsLocal", obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".local") {
			t := strings.TrimPrefix(k, ".local")
			obj.Vars.Local[t] = ParseStringApiObject("ParseVars.local", obj, v, data)
			continue
		}
	}
	obj.TemplateData = data
}

func (obj *ApiObject) assert() {
	data, err := ApiObjectConvertParseData(obj)
	obj.TemplateData = data
	err = errors.Wrap(err, "assert ApiObjectConvertParseData fail")
	obj.SetErr(err)
	assertResult := true
	for _, v := range obj.Assert {
		result := ParseStringApiObject("assert", obj, v, data)
		resultBool, err := strconv.ParseBool(result)
		errors.Wrapf(err, "assert fail: %s", v)
		obj.SetErr(err)
		if !resultBool {
			assertResult = false
		}
	}
	obj.AssertResult = assertResult
}

func (obj *ApiObject) Run()  {
	obj.Before()
	obj.Execute()
	obj.After()
	obj.ParseVars()
	obj.assert()
}

func (obj *ApiObject) Err() error {
	return obj.err
}

func (obj *ApiObject) SetErr(err error) {
	obj.SetErrs(err)
	if obj.err != nil {
		return
	}
	obj.err = err
}

func (obj *ApiObject) Errs() []error {
	return obj.errs
}

func (obj *ApiObject) SetErrs(err error) {
	if (err == nil) {
		return
	}

	if obj.errs == nil {
		obj.errs = make([]error, 0)
		return
	}
	obj.errs = append(obj.errs, err)
}

func (obj *ApiObject) AutoNewRequest() {
	if obj.Err() != nil {
		return
	}
	reqBody := strings.NewReader(fmt.Sprintf(`{"name": "%s", "age": %d}`, "张三", 10))
	obj.Request, obj.err = http.NewRequest(obj.Method, obj.Url, reqBody)
}

func NewApiObject() *ApiObject{
	obj := &ApiObject{}
	return obj
}

func ApiObjectConvertParseData(obj *ApiObject) (map[string]interface{}, error) {
	if obj.Err() != nil {
		return nil, obj.Err()
	}
	var data = make(map[string]interface{})

	var err error
	data["global"] = obj.Vars.Global
	data["session"] = obj.Vars.Session
	data["local"] = obj.Vars.Local
	data["request"] = obj.Request

	var body []byte
	data["response"] = obj.Response
	data["responseBody"] = obj.ResponseBody
	data["responseJsonBody"] = obj.ResponseJsonBody
	if obj.Response == nil {
		return data, errors.Wrap(err, "Response is nil")
	}
	if  obj.Response.Body == nil{
		return data, errors.Wrap(err, "Response.Body is nil")
	}

	if  obj.ResponseBody == "" {
		body, err = ioutil.ReadAll(obj.Response.Body)
		obj.ResponseBody = string(body)
		data["responseBody"] = obj.ResponseBody
		if err != nil {
			return data, errors.Wrap(err, "Response.Body read all fail")
		}
	}

	if obj.ResponseJsonBody == nil {
		obj.ResponseJsonBody = make(map[string]interface{}, 0)
		err = json.Unmarshal(body, &obj.ResponseJsonBody)
		data["responseJsonBody"] = obj.ResponseJsonBody
		if err != nil {
			return data, errors.Wrap(err, "Response.Body json Unmarshal fail")
		}
	}


	return data, err
}