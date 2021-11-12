package main

import "fmt"

func main() {
	a()
}

func a() int {
	defer func() {
		fmt.Println("defer a1")
	}()
	defer func() {
		fmt.Println("defer a2")
	}()
	fmt.Println("func a run")
	return func() int {
		fmt.Println("a return func")
		return 1
	}()
}
