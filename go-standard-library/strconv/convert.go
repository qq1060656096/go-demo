package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("10")
	fmt.Println("Atoi: ", i, err)

	id, err := strconv.ParseInt("22", 10, 64)
	fmt.Println("ParseInt: ", id, err)
	b, err := strconv.ParseBool("true")
	fmt.Println("ParseBool: ", b, err)
}
