package main

import "fmt"

func main() {
	fmt.Println("func main run")

	a()

}

func a() {
	fmt.Println("func a run")
	panic("手动panic")
}
