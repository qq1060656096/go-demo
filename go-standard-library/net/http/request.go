package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second,
	}
	reqBody := strings.NewReader(fmt.Sprintf(`{"book_name": "%s"}`, "go 入门"))
	req, err := http.NewRequest("post", "http://localhost:8080", reqBody)
	if err != nil {
		log.Fatalf("request new fail: %s", err)
	}

	resp, err := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("request body: %v", string(b))
	fmt.Printf("request body err: %v", err)

}