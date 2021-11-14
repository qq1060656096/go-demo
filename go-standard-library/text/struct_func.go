package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)
type User struct {
	Name string
	Pass string
}

func (u User) Login() bool {
	if u.Name == "" {
		return false
	}
	if u.Pass != "123456" {
		return false
	}

	return true
}
func main() {
	s := `
userName: {{if userLogin .userName "123456"}} success {{else}} fail {{end}}
user struct: {{if .user.Login}} success {{else}} fail {{end}}
`
	data := map[string]interface{}{
		"userName": "test1",
		"user": &User{
			"test2",
			"1234567",
		},
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
