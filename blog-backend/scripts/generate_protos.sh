
# 用户接口生成
  protoc --go_out=. ./proto/user.proto
  protoc --go-grpc_out=. ./proto/user.proto
  protoc --grpc-gateway_out=. ./proto/user.proto
  protoc --swagger_out=./apis/swagger ./proto/user.proto

#  文件管理接口生成
  protoc --go_out=. ./proto/fileManage/fileManage.proto
  protoc --go-grpc_out=. ./proto/fileManage/fileManage.proto
  protoc --grpc-gateway_out=. ./proto/fileManage/fileManage.proto
  protoc --swagger_out=./apis/swagger/ ./proto/fileManage/fileManage.proto 

