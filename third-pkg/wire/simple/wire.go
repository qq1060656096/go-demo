package main

import "github.com/google/wire"

func InitializeBusiness() *Business {
	wire.Build(NewDao, NewBusiness, NewDB)
	return nil
}


