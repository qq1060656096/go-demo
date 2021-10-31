package main

import (
	"context"
	"fmt"
)
// WithCancel会递归检测，如果父级是 cancelCtx， 就把当前context cancelFunc挂到最近的父级 cancelCtx
func main() {
	level1(context.Background())
}
/*
Background
  level1: cancelCtx
    children[1]
		level1Child2: cancelCtx
			level2: valueCtx
				level3: cancelCtx
					level3Child1: cancelCtx
*/

func level1(ctx context.Context)  {
	fmt.Println("level1 start")
	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	//go level1Child1(ctx1)
	level1Child2(ctx1)
}

func level1Child2(ctx1 context.Context) {
	fmt.Println("level1Child2 start")
	ctx12, cancel12 := context.WithCancel(ctx1)
	defer cancel12()
	level2(ctx12)
}

func level2(ctx1 context.Context) {
	fmt.Println("level2 start")
	ctx2 := context.WithValue(ctx1, "level2Key", "level2Value")
	level3(ctx2)
}

func level3(ctx2 context.Context) {
	fmt.Println("level3 start")
	ctx3, cancel3 := context.WithCancel(ctx2)
	defer cancel3()
	level3Child1(ctx3)
}

func level3Child1(ctx3 context.Context) {
	fmt.Println("level3Child1 start")
	ctx31, cancel31 := context.WithCancel(ctx3)
	defer cancel31()
	fmt.Println(ctx31.Value("level2Key"))
}