package main

import (
	"fmt"
	"sync"
	"time"
)

// 等待多个 goroutine 完成操作
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine 1: sleep 1 seconds")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine 2: sleep 2 seconds")
	}()
	wg.Wait()
}
