package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// 注意设置超时
	http.DefaultClient.Timeout = time.Millisecond * 1000
	resp, err := http.Get("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=USERID")
	defer resp.Body.Close()
	fmt.Printf("httpGetErr: %v\n", err)
	body, err2 := ioutil.ReadAll(resp.Body)
	fmt.Printf("body: %v\n", string(body))
	fmt.Printf("bodyReadErr: %v\n", err2)
}
