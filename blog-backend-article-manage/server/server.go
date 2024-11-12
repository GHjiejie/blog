package server

import (
	articlepb "articleManage/pb/articleManage"
	"articleManage/pkg/config"
	"articleManage/pkg/db"
	"articleManage/pkg/middleware"
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ArticleServer struct {
	*articlepb.UnimplementedArticleManageServiceServer                //grpc里面未实现的接口
	config                                             *config.Config //配置信息
	DBEngine                                           db.Handle      //数据库引擎
	grpcServer                                         *grpc.Server
	httpServer                                         *http.Server
}

// Start 启动 gRPC 和 HTTP 服务
func (s *ArticleServer) Start() error {
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
	log.Infof("http server listening on %s", s.config.HTTPEndpoint)

	return nil
}

func NewArticleServer(c *config.Config) (*ArticleServer, error) {
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
	s := &ArticleServer{
		config:                                  c,
		DBEngine:                                dbEngine,
		UnimplementedArticleManageServiceServer: &articlepb.UnimplementedArticleManageServiceServer{},
	}
	logger.Info("file server created successfully")
	if err := s.prepareServer(); err != nil {
		log.Error("Failed to prepare server:", err)
		return nil, err
	}

	return s, nil
}

func (s *ArticleServer) prepareServer() error {
	return s.prepareNetServer()
}

func (s *ArticleServer) prepareNetServer() error {
	logger := log.WithFields(log.Fields{
		"module": "prepareNetServer",
	})

	// 创建GRPC服务器
	s.grpcServer = grpc.NewServer()
	articlepb.RegisterArticleManageServiceServer(s.grpcServer, s)
	logger.Info("grpc server prepared successfully")

	// 创建http到grpc的多路复用器
	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// 注册 gRPC-Gateway，确保它能正确连接到 gRPC 服务器
	err := articlepb.RegisterArticleManageServiceHandlerFromEndpoint(context.Background(), rmux, s.config.GRPCEndpoint, opts)
	if err != nil {
		log.Errorf("Failed to register user service handler from endpoint: %v", err)
		return err
	}

	httpmux := http.NewServeMux()
	// 对于特殊的接口不需要进行grpc调用，直接在httpmux中注册
	target_files, _ := url.Parse(s.config.EngineEndpointFiles)
	target_users, _ := url.Parse(s.config.EngineEndpointUsers)

	engine_files := middleware.ReverseProxy(target_files)
	engine_users := middleware.ReverseProxy(target_users)

	// // 判断当前接口是否需要进行远程调用
	for _, path := range s.config.EnginePaths {
		// 这里思考一个问题，就是这里需不需要进行中间件的执行
		// 如果path以/v1/files开头，那么就需要转发给engine_files
		// 如果path以/v1/articles开头，那么就需要转发给engine_articles
		if strings.HasPrefix(path, "/v1/files") {
			httpmux.Handle(path, engine_files)
		} else if strings.HasPrefix(path, "/v1/users") {
			httpmux.Handle(path, engine_users)
		}
	}

	httpmux.Handle("/v1/", rmux)

	s.httpServer = &http.Server{
		Addr:    s.config.HTTPEndpoint,
		Handler: httpmux,
	}
	log.Info("gRPC and HTTP servers prepared successfully")
	return nil
}
