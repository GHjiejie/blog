package server

import (
	"context"

	filepb "fileManage/pb/fileManage"
	"fileManage/validate"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//	func (UnimplementedFileManageServiceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
//	}
//
//	func (UnimplementedFileManageServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
//	}
//
// // 文件模糊查询
//
//	func (UnimplementedFileManageServiceServer) QueryFile(context.Context, *QueryFileRequest) (*QueryFileResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method QueryFile not implemented")
//	}
func (s *FileServer) QueryFileById(ctx context.Context, req *filepb.QueryFileByIdRequest) (*filepb.QueryFileByIdResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "QueryFileById",
	})
	fileId := req.GetFileId()
	logger.Infof("Received file id: %d", fileId)

	if err := validate.FileQueryById(req); err != nil {
		logger.Errorf("Failed to validate file id: %v", err)
		return nil, status.Error(codes.InvalidArgument, "invalid fileId")
	}

	// 接下来从数据库中查询文件信息
	file, err := s.DBEngine.GetFileByID(fileId)
	if err != nil {
		logger.Errorf("Failed to query file by id: %v", err)
		return nil, status.Error(codes.Internal, "failed to query file by id")
	}

	fileInfoProto := &filepb.FileInfo{
		FileId:   file.ID,
		FileName: file.FileName,
		Bytes:    file.FileSize,
		FileType: file.FileType,
		Tag:      file.Tag,
		Content:  file.FileContent,
	}
	// 将获取的文件信息转换为 proto 格式
	return &filepb.QueryFileByIdResponse{
		FileInfo: fileInfoProto,
	}, nil

}
