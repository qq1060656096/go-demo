package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go start")
	go func() {
		fmt.Println("go func")
	}()
	fmt.Println("go end")
	time.Sleep(1 * time.Second)
}
