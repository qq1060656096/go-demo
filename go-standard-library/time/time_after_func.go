package main

import (
	"fmt"
	"time"
)

func main()  {
	// 2秒后执行一次 fn方法
	fn := func() {
		fmt.Println("time.AfterFunc 2 seconds")
	}
	timer := time.AfterFunc(time.Duration(2) * time.Second, fn)
	defer timer.Stop()
	fmt.Println()
	time.Sleep(6 * time.Second)
}
