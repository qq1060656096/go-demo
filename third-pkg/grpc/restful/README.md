

```
# 生成 go pb 代码
protoc --proto_path=./ --go_out=./ account/v1/account.proto  

# 生成 grpc.pb.go
protoc --proto_path=./ --go-grpc_out=. account/v1/account.proto
```