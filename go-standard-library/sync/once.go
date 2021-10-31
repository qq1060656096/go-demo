package main

import (
	"fmt"
	"sync"
	"time"
)

// once 无论调用多少次，都只执行一次
var once sync.Once
func main()  {
	go func() {
		for i := 1; i <= 5; i ++{
			go func(i int) {
				fmt.Printf("goroutine %d\n", i)
				once.Do(func() {
					fmt.Println("只执行一次")
				})
			}(i)
		}
	}()
	time.Sleep(1 * time.Second)
}
/**输出：
goroutine 1
goroutine 4
只执行一次
goroutine 5
goroutine 2
goroutine 3

 */