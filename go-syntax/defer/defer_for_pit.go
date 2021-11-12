package main

import "fmt"

func main() {
	fmt.Println("for start")
	for i := 0; i < 5; i++ {
		// 这里会造成 defer 泄漏，因为 defer 在函数结束后才会执行
		defer fmt.Println("i index: ", i)
	}
	fmt.Println("for end")

}

