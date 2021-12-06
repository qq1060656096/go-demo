package main

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/metadata/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type MetadataService struct {
	v1.UnimplementedAccountServer
}

func (m MetadataService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if req.Pass != "123456"{
		return nil, errors.New("出错了")
	}
	return &v1.LoginResponse{
		Token: fmt.Sprintf("%s:%s", req.Name, req.Pass),
	}, nil
}

func (m MetadataService) GetInfo(ctx context.Context, req *v1.InfoRequest) (*v1.InfoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("md: ", md["name"], ok)
	if !ok {
		fmt.Println("login fail ", md)
	}
	return &v1.InfoResponse{
		Name: fmt.Sprintf("GetInfoPrefix: %s", md["name"][0]),
	}, nil
}

func main() {
	l, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("server: ", err)
	}
	server := grpc.NewServer()
	v1.RegisterAccountServer(server, new(MetadataService))
	err = server.Serve(l)
	if err != nil {
		fmt.Println("server: ", err)
	}
}