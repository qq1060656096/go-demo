package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁：同一时刻，只能有一个读或者写，其他读写会堵塞
var mutex sync.Mutex

func main() {
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d lock before\n", i)
			mutex.Lock()
			fmt.Printf("goroutine %d locking\n", i)
			time.Sleep(3 * time.Second)
			mutex.Unlock()
			fmt.Printf("goroutine %d unlock\n", i)
		}(i)
	}
	time.Sleep(10 * time.Second)
}
/**输出：
goroutine 1 lock before
goroutine 1 locking
goroutine 3 lock before
goroutine 2 lock before
goroutine 1 unlock
goroutine 3 locking
goroutine 3 unlock
goroutine 2 locking
goroutine 2 unlock

 */