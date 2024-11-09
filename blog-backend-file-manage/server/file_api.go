package server

import (
	"context"
	filepb "fileManage/pb/fileManage"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 文件上传接口
func (s *FileServer) UploadFile(ctx context.Context, req *filepb.UploadFileRequest) (*filepb.UploadFileResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "UploadFile",
	})

	logger.Info("UploadFile called")

	// 输出上传文件的信息
	logger.Infof("UploadFile called with %v", req.FileName)
	logger.Infof("UploadFile called with %v", req.FileTag)
	logger.Infof("UploadFile called with %v", req.Content)

	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
