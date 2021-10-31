package main

import (
	"context"
	"fmt"
)

// 简单示例
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	select {
	case <-ctx.Done():
		fmt.Println("手动取消")
	}
}
