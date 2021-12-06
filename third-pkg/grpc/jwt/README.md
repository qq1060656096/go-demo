
```sh
protoc --go_out=./ account/v1/account.proto

protoc --go-grpc_out=. account/v1/account.proto
```