package server

import (
	"blog-backend/pkg/db"
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "blog-backend/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BlogServer struct {
	*pb.UnimplementedUserServiceServer
	DBEngine   db.Handle
	httpServer *http.Server
	grpcServer *grpc.Server
}

// NewBlogServer 创建并返回一个新的 BlogServer 实例
func NewBlogServer() (*BlogServer, error) {
	// 连接数据库
	dbEngine, err := db.NEWHandler()
	if err != nil {
		log.Printf("failed to create db engine handler with err(%s)", err.Error())
		return nil, err
	}

	// 初始化 BlogServer
	s := &BlogServer{
		DBEngine:                       dbEngine,
		UnimplementedUserServiceServer: &pb.UnimplementedUserServiceServer{},
	}

	// 准备 gRPC 和 HTTP 服务
	if err := s.prepareServer(); err != nil {
		log.Printf("failed to prepare server with err(%s)", err.Error())
		return nil, err
	}
	return s, nil
}

// Start 启动 gRPC 和 HTTP 服务
func (s *BlogServer) Start() error {
	log.Printf("BlogServer starting...")

	// 启动 gRPC 服务
	go func() {
		lis, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Printf("gRPC server failed to listen: %v", err)
			panic(err)
		}
		log.Printf("gRPC server listening on :8081")

		if err := s.grpcServer.Serve(lis); err != nil {
			log.Printf("gRPC server failed to serve: %v", err)
			panic(err)
		}
	}()
	// 启动 HTTP 服务
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Printf("HTTP/JSON server failed to start: %v", err)
			panic(err)
		}
	}()
	log.Printf("HTTP/JSON server listening on :8080")

	return nil
}

// prepareServer 准备 gRPC 和 HTTP 服务器
func (s *BlogServer) prepareServer() error {
	log.Printf("Preparing gRPC and HTTP servers")

	// 创建 gRPC 服务器
	s.grpcServer = grpc.NewServer()
	pb.RegisterUserServiceServer(s.grpcServer, s)

	// 创建 HTTP 服务器并配置 grpc-gateway
	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// 注册 gRPC-Gateway，确保它能正确连接到 gRPC 服务器
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), rmux, "localhost:8081", opts)
	if err != nil {
		log.Printf("Failed to register user service handler from endpoint: %v", err)
		return err
	}

	// 创建 HTTP 多路复用器
	httpmux := http.NewServeMux()
	httpmux.Handle("/v1/", rmux)

	// 创建 HTTP 服务器
	s.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: httpmux,
	}

	log.Printf("gRPC and HTTP servers prepared successfully")
	return nil
}
