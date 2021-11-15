package biz

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestParseString(t *testing.T) {

	headers := *NewHeaders()
	headers["token"] = "123"

	apiObj := &ApiObject{
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
		Assert: Assert{
			"{{eq .response.StatusCode 200}}",
		},
	}
	apiObj.Vars.Global["host"] = "http://localhost:8080"
	apiObj.Run()
	text := `
{{.}}
{{.global}}
{{.session}}
{{.local}}
{{.request}}
{{.response}}
{{.response.StatusCode}}
{{.responseJsonBody.name}}
{{.responseBody}}
`
	textResult := ParseStringApiObject("testLog", apiObj, text, apiObj.TemplateData)
	log.Println(apiObj.Errs())
	log.Println("textResult: ", textResult)
	assert.Equal(t, true, apiObj.AssertResult)
}
