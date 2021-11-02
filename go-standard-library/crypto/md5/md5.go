package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	m := md5.New()
	m.Write([]byte("123456"))
	fmt.Printf("123456的md5: %x\n", m.Sum(nil))
	fmt.Printf("123456的md5: %X\n", m.Sum(nil))
}
