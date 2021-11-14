package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	s := `{{if eq .userName "test1"}} true {{else}} false {{end}}`
	data := map[string]interface{}{
		"userName": "test1",
	}
	t := template.New("test")
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
