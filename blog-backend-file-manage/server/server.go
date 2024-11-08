package server

import (
	filepb "fileManage/pb/fileManage"
	"fileManage/pkg/config"
	"fileManage/pkg/db"

	log "github.com/sirupsen/logrus"
)

type FileServer struct {
	*filepb.UnimplementedFileManageServiceServer                //grpc里面未实现的接口
	config                                       *config.Config //配置信息
	DBEngine                                     db.Handle      //数据库引擎
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
		config:   c,
		DBEngine: dbEngine,
	}

	return s, nil
}
