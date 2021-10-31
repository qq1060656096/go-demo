package main

import (
	"context"
	"fmt"
	"time"
)

// 手动取消
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	go func() {
		fmt.Println("手动取消")
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("goroutine done")
	}
}
/**
手动取消
goroutine done

 */