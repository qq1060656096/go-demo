package main

import "fmt"

type Person struct {
	name string
}

func main() {
	arr := []Person{
		Person{"小明"},
		Person{"小刚"},
	}
	var res []*Person

	for _, v := range arr {
		res = append(res, &v)
	}
	// 遍历查看结果集
	for _, person := range res{
		fmt.Println("name-->:", person.name)
	}
}