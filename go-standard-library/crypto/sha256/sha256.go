package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := sha256.New()
	fmt.Printf("123456的sha256: %x\n", s.Sum(nil))
}