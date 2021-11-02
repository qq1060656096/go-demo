package main

import (
	"fmt"
	"strings"
)

func main()  {
	id := "1,2,3"
	fmt.Printf("字符串分割：%#v\n", strings.Split(id, ","))

	id2 := []string{"1", "2", "3"}
	fmt.Printf("字符串合并：%#v\n", strings.Join(id2, ","))
	fmt.Printf("字符串查找子字符串是否存在：%#v\n", strings.Contains(id, "1,2"))
	fmt.Printf("字符串查找子字符串的索引（不存在返回-1）：%#v\n", strings.Index(id, "1,2,3,4"))
}
