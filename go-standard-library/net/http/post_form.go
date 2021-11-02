package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Second
	form := url.Values{
		"name": {"test name"},
		"nickname": {
			"nickname 1",
			"nickname 2",
		},
	}
	resp, err := http.PostForm("http://localhost:8080", form)
	if err != nil {
		fmt.Printf("response err: %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(body))
	fmt.Printf("response body err: %v\n", err)

}