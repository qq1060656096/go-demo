# 

# 1. 生成 wire文件
```sh
wire
```


## 2.运行
```sh
go run wire_gen.go main.go
```


### wire_gen文件内容如下
```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeBusiness() *Business {
	db := NewDB()
	dao := NewDao(db)
	business := NewBusiness(dao)
	return business
}

```