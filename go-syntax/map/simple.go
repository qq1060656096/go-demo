package main

import "fmt"

func main()  {
	m := make(map[int]string)
	fmt.Printf("%v\n", m)

	b := m
	fmt.Printf("b=%p, m=%p\n", b, m)

	b = make(map[int]string)
	fmt.Printf("b=%p, m=%p\n", b, m)

}