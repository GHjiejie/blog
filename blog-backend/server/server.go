package server

import (
	"blog-backend/pkg/auth"
	"blog-backend/pkg/casbinpermit"
	"blog-backend/pkg/config"
	"blog-backend/pkg/db"
	"blog-backend/pkg/middleware"
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	pb "blog-backend/pb/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BlogServer struct {
	*pb.UnimplementedUserServiceServer
	DBEngine     db.Handler
	config       *config.Config //配置信息
	casbinPermit *casbinpermit.Permit
	httpServer   *http.Server
	grpcServer   *grpc.Server
	auth         *auth.Auth
}

// NewBlogServer 创建并返回一个新的 BlogServer 实例
func NewBlogServer(c *config.Config) (*BlogServer, error) {
	// 连接数据库
	dbEngine, err := db.NEWHandler(c.DBConfig)
	if err != nil {
		log.Errorf("failed to create db engine handler with err(%s)", err.Error())
		return nil, err
	}
	casbinPermit, err := casbinpermit.NewPermit(dbEngine.GetORMDB())
	if err != nil {
		log.Errorf("failed to create casbin permit with err(%s)", err.Error())
		return nil, err
	}

	// 接下来就是加载策略
	if err := casbinPermit.Enforcer.LoadPolicy(); err != nil {
		log.Infof("failed to load policy with err(%s)", err.Error())
		return nil, err
	}
	auth := auth.NewAuth(dbEngine, c, casbinPermit)
	// 初始化 BlogServer
	s := &BlogServer{
		DBEngine:                       dbEngine,
		config:                         c,
		casbinPermit:                   casbinPermit,
		auth:                           auth,
		UnimplementedUserServiceServer: &pb.UnimplementedUserServiceServer{},
	}

	// 准备 gRPC 和 HTTP 服务
	if err := s.prepareServer(); err != nil {
		log.Errorf("failed to prepare server with err(%s)", err.Error())
		return nil, err
	}
	return s, nil
}

// Start 启动 gRPC 和 HTTP 服务
func (s *BlogServer) Start() error {
	log.Info("BlogServer starting...")

	// 启动 gRPC 服务
	go func() {
		lis, err := net.Listen("tcp", s.config.GRPCEndpoint)
		// lis, err := net.Listen("tcp", s.config.GRPCEndpoint)
		if err != nil {
			log.Errorf("gRPC server failed to listen: %v", err)
			panic(err)
		}
		log.Info("gRPC server listening on :8081")

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
	log.Info("HTTP/JSON server listening on :8080")

	return nil
}

// prepareServer 准备 gRPC 和 HTTP 服务器
func (s *BlogServer) prepareServer() error {
	log.Info("Preparing gRPC and HTTP servers")

	// 创建 gRPC 服务器
	s.grpcServer = grpc.NewServer()
	pb.RegisterUserServiceServer(s.grpcServer, s)

	// 创建 HTTP 服务器并配置 grpc-gateway
	rmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// 注册 gRPC-Gateway，确保它能正确连接到 gRPC 服务器
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), rmux, s.config.GRPCEndpoint, opts)
	if err != nil {
		log.Errorf("Failed to register user service handler from endpoint: %v", err)
		return err
	}

	// 创建 HTTP 多路复用器
	httpmux := http.NewServeMux()

	target_files, _ := url.Parse(s.config.EngineEndpointFiles)
	target_articles, _ := url.Parse(s.config.EngineEndpointArticles)

	engine_files := middleware.ReverseProxy(target_files)
	engine_articles := middleware.ReverseProxy(target_articles)

	// // 判断当前接口是否需要进行远程调用
	for _, path := range s.config.EnginePaths {
		// 这里思考一个问题，就是这里需不需要进行中间件的执行
		// 如果path以/v1/files开头，那么就需要转发给engine_files
		// 如果path以/v1/articles开头，那么就需要转发给engine_articles
		if strings.HasPrefix(path, "/v1/files") {
			httpmux.Handle(path, s.Tracing(engine_files, s.casbinPermit))
		} else if strings.HasPrefix(path, "/v1/articles") {
			httpmux.Handle(path, s.Tracing(engine_articles, s.casbinPermit))
		}
	}

	httpmux.Handle("/v1/", s.Tracing(rmux, s.casbinPermit))

	// 创建 HTTP 服务器
	s.httpServer = &http.Server{
		Addr:    s.config.HTTPEndpoint,
		Handler: httpmux,
	}

	log.Info("gRPC and HTTP servers prepared successfully")
	return nil
}
