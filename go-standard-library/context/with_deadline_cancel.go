package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("自动取消")
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("超时取消")
	}
}
/**输出
超时取消
 */