package main

import "github.com/google/wire"

func InitializeBusiness() (*Business, error) {
	wire.Build(NewBusiness, NewDao, NewDB)
	return &Business{}, nil
}


