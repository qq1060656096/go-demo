package main

import (
	"fmt"
	"time"
)

func main() {

	select {
	// 2秒以后收到返回, 然后打印"sleep 2 seconds"
	case <-time.After(2 * time.Second):
		fmt.Println("sleep 2 seconds")
	}
}