package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("时间戳")
	fmt.Println("当前时间戳  （秒）: ", time.Now().Unix())
	fmt.Println("当前时间戳（毫秒）: ", time.Now().UnixNano()/int64(time.Millisecond))
	fmt.Println("当前时间戳（微秒）: ", time.Now().UnixNano()/int64(time.Microsecond))
	fmt.Println("当前时间戳（纳秒）: ", time.Now().UnixNano())

	fmt.Println()
	fmt.Println("时间戳日期互转")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-10-31 20:23:01", time.Local)
	fmt.Println("(2021-10-31 20:23:01)转换成时间戳：", t.Unix(), err)
	t2 := time.Unix(1635711781, 0)
	fmt.Println("(2021-10-31 20:23:01)的时间戳转时间：", t2.UTC().Format("2006-01-02 15:04:05"))

	fmt.Println()
	fmt.Println("时间运算")
	fmt.Println("10分钟之前: ", time.Now().Add(-10 * time.Minute).Format("2006-01-02 15:04:05"))
	fmt.Println("10分钟之后: ", time.Now().Add(10 * time.Minute).Format("2006-01-02 15:04:05"))
	fmt.Println("1月之前: ", time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("1月之后: ", time.Now().AddDate(0, 1, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("1年之前: ", time.Now().AddDate(-1, 0, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("1年之后: ", time.Now().AddDate(1, 0, 0).Format("2006-01-02 15:04:05"))

	fmt.Println()
	fmt.Println("2021-10-31 20:23:01和 2021-10-31 20:26:01两个时间比较: ")
	t01, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-10-31 20:23:01", time.Local)
	t02, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-10-31 20:26:01", time.Local)

	fmt.Println("相差多少分钟: ", int64(t02.Sub(t01)/time.Minute))
	fmt.Println("多久过期(负数已经过期了): ", int64(time.Until(t01.UTC())/time.Minute))
	fmt.Println("2021-10-31 20:23:01到现在过去了多少分钟: ", int64(time.Since(t01)/time.Minute))


}