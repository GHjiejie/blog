package server

import (
	"context"
	filepb "fileManage/pb/fileManage"
	"fileManage/pkg/config"
	"fileManage/pkg/db"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type FileServer struct {
	*filepb.UnimplementedFileManageServiceServer                //grpc里面未实现的接口
	config                                       *config.Config //配置信息
	DBEngine                                     db.Handle      //数据库引擎
	grpcServer                                   *grpc.Server
	httpServer                                   *http.Server
}

// Start 启动 gRPC 和 HTTP 服务
func (s *FileServer) Start() error {
	log.Info("FileServer starting...")

	// 启动 gRPC 服务
	go func() {
		lis, err := net.Listen("tcp", s.config.GRPCEndpoint)
		if err != nil {
			log.Errorf("gRPC server failed to listen: %v", err)
			panic(err)
		}
		log.Infof("gRPC server listening on %s", s.config.GRPCEndpoint)

		if err := s.grpcServer.Serve(lis); err != nil {
			log.Errorf("gRPC server failed to serve: %v", err)
			panic(err)
		}
	}()
	// 启动 HTTP 服务
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Errorf("HTTP/JSON server failed to start: %v", err)
			panic(err)
		}
	}()
	log.Info("HTTP/JSON server listening on :8090")

	return nil
}

func NewFileServer(c *config.Config) (*FileServer, error) {
	logger := log.WithFields(log.Fields{
		"module": "startServer",
	})
	logger.Info("start to create file server")

	// 首先创建数据库引擎
	dbEngine, err := db.NewHandler(c.DBConfig)
	if err != nil {
		logger.Error("failed to create database engine:", err)
		return nil, err
	}
	s := &FileServer{
		config:                               c,
		DBEngine:                             dbEngine,
		UnimplementedFileManageServiceServer: &filepb.UnimplementedFileManageServiceServer{},
	}
	logger.Info("file server created successfully")
	if err := s.prepareServer(); err != nil {
		log.Error("Failed to prepare server:", err)
		return nil, err
	}

	return s, nil
}

func (s *FileServer) prepareServer() error {
	return s.prepareNetServer()
}

func (s *FileServer) prepareNetServer() error {
	logger := log.WithFields(log.Fields{
		"module": "prepareNetServer",
	})

	// 创建GRPC服务器
	s.grpcServer = grpc.NewServer()
	filepb.RegisterFileManageServiceServer(s.grpcServer, s)
	logger.Info("grpc server prepared successfully")

	// 创建http到grpc的多路复用器
	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// 注册 gRPC-Gateway，确保它能正确连接到 gRPC 服务器
	err := filepb.RegisterFileManageServiceHandlerFromEndpoint(context.Background(), rmux, s.config.GRPCEndpoint, opts)
	if err != nil {
		log.Errorf("Failed to register user service handler from endpoint: %v", err)
		return err
	}

	httpmux := http.NewServeMux()
	// 对于特殊的接口不需要进行grpc调用，直接在httpmux中注册
	httpmux.HandleFunc("/v1/files/upload", s.uploadFileHandler)
	httpmux.Handle("/v1/", rmux)

	s.httpServer = &http.Server{
		Addr:    s.config.HTTPEndpoint,
		Handler: httpmux,
	}
	log.Info("gRPC and HTTP servers prepared successfully")
	return nil
}
