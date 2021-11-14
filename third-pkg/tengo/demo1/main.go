package main

import (
	"fmt"
	"github.com/d5/tengo/v2"
	"log"
)

func main() {
	request := &Request{
		Name: "test",
		Age: 18,
	}
	s := `
b := request
`
	script := tengo.NewScript([]byte(s))
	err := script.Add("request", request.Name)
	if err != nil {
		log.Fatal("add ", err)
	}
	compiled, err := script.Run()
	if err != nil {
		log.Fatal("compare", err)
	}
	b := compiled.Get("b")
	fmt.Println(b)
}

type Request struct {
	Name string
	Age int
}

