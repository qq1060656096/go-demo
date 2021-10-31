package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
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