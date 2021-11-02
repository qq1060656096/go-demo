package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Second

	s := strings.NewReader(fmt.Sprintf(`{"name": "%s", "age": %d}`, "张三", 10))
	resp, err := http.Post("http://localhost:8080", "application/json", s)
	if err != nil {
		fmt.Printf("response err: %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response body: %s\n", string(body))
	fmt.Printf("response body err: %v\n", err)

}