package main

import (
	"bytes"
	"fmt"
)

func main() {
	c1 := []byte("1")
	c2 := []byte("1")
	cr := bytes.Compare(c1, c2)
	fmt.Printf("c1、c2 Compare: %d \n", cr)
	fmt.Printf("c1、c2 Equal: %v \n", bytes.Equal(c1, c2))
	j := [][]byte{[]byte("hello go"), []byte("hello go2")}
	fmt.Printf("c1、c2 Join: %s \n", bytes.Join(j, []byte(", ")))
}
