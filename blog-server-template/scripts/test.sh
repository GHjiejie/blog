  protoc --go_out=. ./proto/articleManage.proto
  protoc --go-grpc_out=. ./proto/articleManage.proto
  protoc --grpc-gateway_out=. ./proto/articleManage.proto
  protoc --swagger_out=./apis/swagger/ ./proto/articleManage.proto 