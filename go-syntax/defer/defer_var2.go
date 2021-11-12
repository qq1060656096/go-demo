package main

import "fmt"

func main() {
	i := 5
	defer func() {
		fmt.Println("defer: ", i)
	}()
	i ++
	fmt.Println("man end: ", i)
}