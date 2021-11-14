package biz

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParseString(t *testing.T) {
	text := `
{{.}}
`
	headers := *NewHeaders()
	headers["token"] = "123"

	data := &ApiObject{
		Name: "测试用例1",
		Method: "POST",
		Url: "{{.global.host}}",
		QueryParams: map[string]string{
			"age[0]": "1",
			"age[1]": "2",
			"name": "33",
		},
		Headers: headers,
		Vars: NewVars(),
		Client: http.DefaultClient,
		SetVar: SetVars{
			//"global.Name": "{{.responseJsonBody.name}}",
		},
	}
	(*data.Vars.Global)["host"] = "http://localhost:8080"
	data.Run()
	s, err := ParseString(data.Template, text, data.Data)
	fmt.Println(">>>>>><<<<<<<", s, err, data.Err())
	text = `
{{.}}
<<<<<<<
{{.responseJsonBody.name}}
{{.responseBody}}
{{.global}}
{{.session}}
{{.local}}

`

	s = ParseStringApiObject(data, text, data.Data)
	fmt.Println(s, "err: ", data.Err(), " ++++", data.Data)
	fmt.Printf("%v", data.Data)
}
