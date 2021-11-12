package main

import (
	"context"
	"errors"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/restful/account/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatalf("listen err: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterAccountServiceServer(grpcServer, new(AccountService))
	err = grpcServer.Serve(lis)
	log.Println("serve err: ", err)
}

type AccountService struct {
	v1.UnimplementedAccountServiceServer
}

func (s *AccountService) LoginAccount(ctx context.Context, in *v1.LoginAccountRequest) (*v1.Account, error) {
	if time.Now().Unix()%2 == 0 {
		return nil, errors.New("LoginAccount rand err")
	}
	return &v1.Account{
		Name: in.Name,
		Pass: in.Pass,
	}, nil
}