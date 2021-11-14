package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = "one"
	m[1] = "tow"
	fmt.Printf("mapFunc %p\n", m)
	mapFunc(m)
	fmt.Printf("mapFunc %#v\n", m)

}

func mapFunc(m1 map[int]string)  {
	fmt.Printf("mapFunc %p\n", m1)
	m1 = make(map[int]string)
	m1[0] = "one.mapFunc"
	fmt.Printf("mapFunc make after %p\n", m1)
	fmt.Printf("mapFunc make after %#v\n", m1)

}
