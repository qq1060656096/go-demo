package main

import "fmt"

func main() {
	// a = [5]int{1, 2, 3, 4, 5}
	a := [...]int{ 1, 2, 3, 4, 5}
	fmt.Printf("%#v \n", a)

	// a1 = []int{0, 0, 0, 0, 0}
	a1 := make([]int, 5, 5)
	fmt.Printf("%#v \n", a1)
	for k, v := range a1 {
		v = k
		a1[k] = v
	}
	a1[0] = 0
	fmt.Printf("%#v \n", a1)
	fmt.Printf("a1[0:1]=%#v \n", a1[0:1])
	fmt.Printf("a1[2:]=%#v \n", a1[2:])
	fmt.Printf("a1[:3]=%#v \n", a1[:1])
	fmt.Printf("a1[2:2]=%#v \n", a1[3:4])


	// a2 = []*int{0, 0, 0, 0, 0}
	a2 := make([]*int, 5, 5)
	fmt.Printf("%#v \n", a2)
	for k, _ := range a2 {
		t := k
		a2[k] = &t
	}
	fmt.Printf("a2=%#v \n", a2)

	for k, v := range a2 {
		*v =  k + 1
		fmt.Printf("a2=%#v \n", *v)
	}
	for k, v := range a2 {
		fmt.Printf("k=%d, v=%#v \n", k, *v)
	}
}