
# 安装protoc编译插件和grpc生成插件
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



```
# 生成go pb 文件
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
# 生成go grpc service
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1


# 生成 go pb 代码
protoc --go_out=./ account/v1/account.proto
# 生成 grpc.pb.go
protoc --go-grpc_out=. account/v1/account.proto

```