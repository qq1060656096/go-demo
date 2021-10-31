package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.WithValue(context.Background(), "key1", "value1")
	ctx2 := context.WithValue(ctx1, "key2", "value2")

	// ctx1.key2:  value1
	fmt.Println("ctx1.key2: ", ctx1.Value("key1"))
	// ctx1.key2:  <nil>
	fmt.Println("ctx1.key2: ", ctx1.Value("key2"))
	// ctx2.key1:  value1
	fmt.Println("ctx2.key1: ", ctx2.Value("key1"))
	// ctx2.key1:  value2
	fmt.Println("ctx2.key1: ", ctx2.Value("key2"))
}
