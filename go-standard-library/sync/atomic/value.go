package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 提供原子获取值和保存值，适用于进程缓存如配置文件加载、热点数据或者运营内数据、降级数据等
var value atomic.Value

func main() {
	for i := 1; i <= 3; i++{
		go func(i int) {
			value.Store(fmt.Sprintf("v%d", i))
			fmt.Printf("goroutine write %d\n", i)
		}(i)
		go func(i int) {
			time.Sleep(500 * time.Nanosecond)
			fmt.Printf("goroutine read %d %s\n", i, value.Load())
		}(i)
	}
	time.Sleep(1 * time.Second)
}
