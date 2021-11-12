package main

import (
	"fmt"
	"time"
)

func main()  {
	go func() {
		for i:=0; ;i++{
			time.Sleep(time.Second)
			fmt.Println("goroutine var: i = ", i)
			panic("goroutine panic")
		}
	}()
	for i:=0;;i++{
		time.Sleep(time.Second)
		fmt.Println("i = ", i)
	}
}