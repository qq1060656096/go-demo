package main

import (
	"context"
	"fmt"
)

//  WithValue 父子获取值互不受到影响
func main() {
	ctx := context.WithValue(context.Background(), "key1", "key1.isParent")
	ctxChild := context.WithValue(ctx, "key1", "key1.isChild")
	fmt.Println(ctx.Value("key1"))
	fmt.Println(ctxChild.Value("key1"))
}
/**输出
key1.isParent
key1.isChild
 */