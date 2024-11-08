package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	filepb "blog-backend/pb/fileManage"
)

// 文件上传接口
func (s *BlogServer) UploadFile(ctx context.Context, req *filepb.UploadFileRequest) (*filepb.UploadFileResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "UploadFile",
	})
	logger.Info("UploadFile")

	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
