package biz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Template *template.Template
	Headers Headers
	Request *http.Request
	Response *http.Response
	ResponseJsonBody interface{}
	ResponseBody string
	Data map[string]interface{}
	err error
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
	if err != nil {
		obj.err = err
	}
	// 1. 创建 request
	// 2. 解析 url，合并 url query
	// 3. 解析 queryParams， 合并 queryParams
	// 4. 解析headers， 合并 headers
	obj.Url = ParseStringApiObject(obj, obj.Url, data)
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
	obj.Response, obj.err = obj.Client.Do(obj.Request)
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
	obj.SetErr(err)
	for k, v := range obj.SetVar {
		if strings.HasPrefix(k, "global") {
			t := strings.TrimPrefix(k, "global")
			(*obj.Vars.Global)[t] = ParseStringApiObject(obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".global") {
			t := strings.TrimPrefix(k, ".global")
			(*obj.Vars.Global)[t] = ParseStringApiObject(obj, v, data)
			continue
		}

		if strings.HasPrefix(k, "session") {
			t := strings.TrimPrefix(k, "session")
			(*obj.Vars.Session)[t] = ParseStringApiObject(obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".session") {
			t := strings.TrimPrefix(k, ".session")
			(*obj.Vars.Session)[t] = ParseStringApiObject(obj, v, data)
			continue
		}

		if strings.HasPrefix(k, "local") {
			t := strings.TrimPrefix(k, "local")
			(*obj.Vars.Local)[t] = ParseStringApiObject(obj, v, data)
			continue
		}
		if strings.HasPrefix(k, ".local") {
			t := strings.TrimPrefix(k, ".local")
			(*obj.Vars.Local)[t] = ParseStringApiObject(obj, v, data)
			continue
		}
	}
	obj.Data = data
}

func (obj *ApiObject) Run()  {
	obj.Before()
	obj.Execute()
	obj.After()
	obj.ParseVars()
}

func (obj *ApiObject) Err() error {
	return obj.err
}

func (obj *ApiObject) SetErr(err error) {
	if obj.err != nil {
		return
	}
	obj.err = err
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
		return data, err
	}
	if  obj.Response.Body == nil{
		return data, err
	}

	if  obj.ResponseBody != "" || obj.ResponseJsonBody != nil {
		//return data, err
	}
	body, err = ioutil.ReadAll(obj.Response.Body)
	obj.ResponseBody = string(body)
	data["responseBody"] = obj.ResponseBody
	if err != nil {
		return data, err
	}

	if obj.ResponseJsonBody == nil {
		obj.ResponseJsonBody = make(map[string]interface{}, 0)
	}
	err = json.Unmarshal(body, &obj.ResponseJsonBody)
	data["responseJsonBody"] = obj.ResponseJsonBody
	if err != nil {
		return data, err
	}

	return data, err
}