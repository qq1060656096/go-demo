package main

import (
	"fmt"
	"time"
)

// 周期定时器：每隔1秒钟处理一次
func main() {
	t := time.Tick(1 * time.Second)
	for i:=1; i <= 5; i ++ {
		select {
		case <-t:
			fmt.Println("run ", i, ": sleep 1 seconds")
		}
	}
}
/**输出
run  1 : sleep 1 seconds
run  2 : sleep 1 seconds
run  3 : sleep 1 seconds
run  4 : sleep 1 seconds
run  5 : sleep 1 seconds

 */