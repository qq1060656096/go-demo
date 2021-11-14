package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	s := `{{if userLogin .userName "123456"}} success {{else}} fail {{end}}`
	data := map[string]interface{}{
		"userName": "test1",
	}
	fmap := make(template.FuncMap)
	fmap["userLogin"] = func(userName, password string) bool {
		if userName == "" {
			return false
		}
		if password != "123456" {
			return false
		}

		return true
	}
	t := template.New("test")
	t.Funcs(fmap)
	t, err := t.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer

	err = t.Execute(&b, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.String())
}
