package main

import "fmt"

// 复制 slice 到目标 slice, 复制成功返回值大于0
func main() {
	src := []string{"go", "hello", "!"}
	dst := make([]string, 0, 3)
	i := copy(dst, src)
	fmt.Printf("copy: \n")
	fmt.Printf("i = %d\n", i)
	fmt.Printf("src = %#v\n", src)
	fmt.Printf("dst = %#v\n", dst)

}