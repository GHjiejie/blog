  protoc --go_out=. ./proto/fileManage.proto
  protoc --go-grpc_out=. ./proto/fileManage.proto
  protoc --grpc-gateway_out=. ./proto/fileManage.proto
  protoc --swagger_out=./apis/swagger/ ./proto/fileManage.proto 