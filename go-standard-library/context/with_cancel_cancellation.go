package main

import (
	"context"
	"fmt"
	"time"
)

// 级联取消示例
func main() {
	ctx0, cancel0 := context.WithCancel(context.Background())

	fmt.Println("main: start")
	go goroutineLevel1(ctx0)
	fmt.Println("main: 手动取消goroutine")
	cancel0()
	// 演示为了简单使用sleep，但是你不应该模仿
	// 等待其他goroutine结束，业务逻辑请使用 sync.WaitGroup 或者 chan
	time.Sleep(1 * time.Second)
	/*输出
main: start
main: 手动取消goroutine
goroutineLevel1: start
goroutineLevel1: done
goroutineLevel2: start
goroutineLevel2: done
	 */
}



func goroutineLevel1(ctx0 context.Context) {
	ctx1, cancel1 := context.WithCancel(ctx0)
	defer func() {
		fmt.Println("goroutineLevel1: defer")
		cancel1()
	}()
	fmt.Println("goroutineLevel1: start")
	go goroutineLevel2(ctx1)
	select {
	case <-ctx1.Done():
		fmt.Println("goroutineLevel1: done")
	}
}

func goroutineLevel2(ctx1 context.Context) {
	fmt.Println("goroutineLevel2: start")
	select {
	case <-ctx1.Done():
		fmt.Println("goroutineLevel2: done")
	}
}
