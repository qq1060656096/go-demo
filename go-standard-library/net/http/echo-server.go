package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.DefaultServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		b , _ := ioutil.ReadAll(request.Body)
		fmt.Fprintln(writer, string(b))
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}