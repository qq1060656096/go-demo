package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	s := `{{.test}}`
	data := map[string]interface{}{
		"test": 123,
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
