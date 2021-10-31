package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁：读写互斥，有读锁时，写锁堵塞。有写锁时，读锁会堵塞
var rwMutex sync.RWMutex

func main() {
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("goroutine %d lock before\n", i)
			rwMutex.Lock()
			fmt.Printf("goroutine %d locking\n", i)
			time.Sleep(3 * time.Second)
			rwMutex.Unlock()
			fmt.Printf("goroutine %d unlock\n", i)
		}(i)

		go func(i int) {
			fmt.Printf("goroutine read %d RLock before\n", i)
			rwMutex.RLock()
			fmt.Printf("goroutine read %d RLocking\n", i)
			time.Sleep(3 * time.Second)
			rwMutex.RUnlock()
			fmt.Printf("goroutine read %d RUnlock\n", i)
		}(i)
	}

	time.Sleep(20 * time.Second)
}
/**
goroutine 1 lock before
goroutine 1 locking
goroutine read 3 RLock before
goroutine 2 lock before
goroutine read 1 RLock before
goroutine read 2 RLock before
goroutine 3 lock before
goroutine 1 unlock
goroutine read 1 RLocking
goroutine read 3 RLocking
goroutine read 2 RLocking
goroutine read 3 RUnlock
goroutine read 1 RUnlock
goroutine read 2 RUnlock
goroutine 2 locking
goroutine 2 unlock
goroutine 3 locking
goroutine 3 unlock

 */