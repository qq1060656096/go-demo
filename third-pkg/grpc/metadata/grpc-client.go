package main

import (
	"context"
	"fmt"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/metadata/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

func main() {
	net, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	fmt.Println(err)
	client := v1.NewAccountClient(net)
	res, err := client.Login(context.Background(), &v1.LoginRequest{
		Name: "metadata.name",
		Pass: "123456",
	})
	if err != nil {
		log.Fatal("client.Login.fail", err)
	}
	fmt.Println("client login ok, ", res)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{
		"name": []string{res.Token},
	})
	res2, err := client.GetInfo(ctx, &v1.InfoRequest{

	})
	fmt.Println("GetInfo ok, ", res2)
	if err != nil {
		fmt.Println("GetInfo.Response", res2, err)
	}

}