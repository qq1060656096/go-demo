package main

import (
	"context"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/simple/account/v1"
	"google.golang.org/grpc"
	"log"
)

func main() {
	con, err := grpc.Dial(":8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s", err)
	}
	req := &v1.LoginAccountRequest{
		Name: "test",
		Pass: "123456",
	}
	client := v1.NewAccountServiceClient(con)
	account, err := client.LoginAccount(context.Background(), req)
	log.Printf("account: %v \n", account)
	log.Printf("err: %v \n", err)
}