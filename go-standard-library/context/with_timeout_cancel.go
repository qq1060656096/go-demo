package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("超时取消")
	}
	defer func() {
		fmt.Println("cancel func done")
		cancel()
	}()
}
/**
超时取消
cancel func done

 */