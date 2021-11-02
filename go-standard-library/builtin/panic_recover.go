package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover: %#v \n", err)
		}
	}()
	panic("手动panic")
}