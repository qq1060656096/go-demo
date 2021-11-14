package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.DefaultServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		b , _ := ioutil.ReadAll(request.Body)
		fmt.Fprintln(writer, string(b))
		log.Println("url: ", request.URL)
		log.Println("header: ", request.Header)
		log.Println("body: ", string(b))
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}