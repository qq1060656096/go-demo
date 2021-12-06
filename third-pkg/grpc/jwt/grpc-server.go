package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	v1 "github.com/qq1060656096/go-demo/third-pkg/grpc/jwt/account/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer
}

func (a *AccountService) LoginAccount(ctx context.Context, req *v1.LoginAccountRequest) (*v1.LoginAccountResponse, error) {
	if req.Pass != "123456" {
		return nil, status.Errorf(codes.Unimplemented, "method LoginAccount not implemented")
	}
	fmt.Println("LoginAccount: ", req.Name)
	res := &v1.LoginAccountResponse{
		Name: req.Name,
		Token: createToken(req.Name),
	}
	return res, nil
}

func (a *AccountService) SayHello(ctx context.Context, req *v1.PingRequest) (*v1.PingResponse, error) {
	name := checkToken(ctx)
	fmt.Println("cc, ", name)
	return &v1.PingResponse{
		Msg: name,
	}, nil
}

func createToken(username string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: int64(time.Now().Unix() + 100000000),
		Issuer:    username,
	}
	mySigningKey := []byte("123456")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal("createToken fail")
	}
	return s
}

func checkToken(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("get metadata fail")
	}
	token, ok := md["authorization"]
	if !ok  || len(token) < 1 {
		log.Fatal("authorization metadata error")
	}
	tokenString := token[0]
	log.Println("\nauthorization: ", tokenString)
	var claims jwt.StandardClaims
	pt, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456"), nil
	})
	if err != nil {
		return ""
	}
	if pt.Valid {
		return claims.Issuer
	}
	return ""
}
func main() {
	lis, err := net.Listen("tcp", "localhost:8089")
	if err != nil {
		log.Fatalf("tcp listent fail: %s", err)
	}
	s := grpc.NewServer()
	v1.RegisterAccountServiceServer(s, new(AccountService))
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("tcp listent fail: %s", err)
	}
}