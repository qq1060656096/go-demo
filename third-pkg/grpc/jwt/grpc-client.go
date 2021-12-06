package main

import (
	"context"
	"fmt"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/jwt/account/v1"
	"google.golang.org/grpc"
	"log"
)


type AuthToken struct {
	Token string
}

func (c AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c AuthToken) RequireTransportSecurity() bool {
	return false
}

func main() {
	con, err := grpc.Dial(":8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s", err)
	}
	server := v1.NewAccountServiceClient(con)
	req := &v1.LoginAccountRequest{
		Name: "demo66",
		Pass: "123456",
	}
	res,err := server.LoginAccount(context.Background(), req)
	if err != nil {
		log.Fatal("LoginAccount: ", err)
	}
	fmt.Printf("loginResponseToken: %v\n", res)
	requestToken := AuthToken{}
	requestToken.Token = res.Token
	ctx := context.Background()
	con2, err := grpc.Dial(":8089", grpc.WithInsecure(), grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		log.Fatalf("%s", err)
	}
	server2 := v1.NewAccountServiceClient(con2)
	res2,err := server2.SayHello(ctx, &v1.PingRequest{})
	if err != nil {
		log.Fatal("LoginAccount: ", err)
	}
	fmt.Printf("SayHelloResponseCheckToken: %v", res2.Msg)

}