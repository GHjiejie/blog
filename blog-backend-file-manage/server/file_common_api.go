package server

import (
	"context"

	filepb "fileManage/pb/fileManage"
	"fileManage/validate"

	"fileManage/pkg/auth"
	"fileManage/pkg/db"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//	func (UnimplementedFileManageServiceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
//		return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
//	}
func (s *FileServer) DeleteFile(ctx context.Context, req *filepb.DeleteFileRequest) (*filepb.DeleteFileResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "DeleteFile",
	})
	fileId := req.GetFileId()
	// logger.Infof("Received file id: %d", fileId)

	// 首先进行参数的校验
	if err := validate.FileDelete(req); err != nil {
		logger.Errorf("Failed to validate file id: %v", err)
		return nil, status.Error(codes.InvalidArgument, "invalid fileId")
	}

	// 然后验证这个文件是否存在
	_, err := s.DBEngine.GetFileByID(fileId)
	if err != nil {
		logger.Errorf("Failed to query file by id: %v", err)
		return nil, status.Error(codes.Internal, "failed to query file by id")
	}

	// 接下来进行文件的删除操作
	if err := s.DBEngine.DeleteFile(fileId); err != nil {
		logger.Errorf("Failed to delete file by id: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete file by id")
	}

	return &filepb.DeleteFileResponse{
		Message: "Delete file successfully",
	}, nil
}

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
	if err := validate.FileQueryById(req); err != nil {
		logger.Errorf("Failed to validate file id: %v", err)
		return nil, status.Error(codes.InvalidArgument, "invalid fileId")
	}

	// 接下来从数据库中查询文件信息
	file, err := s.DBEngine.GetFileByID(fileId)
	// logger.Infof("file: %v", file)
	// logger.Infof("file content: %x", file.FileContent)
	if err != nil {
		logger.Errorf("Failed to query file by id: %v", err)
		return nil, status.Error(codes.Internal, "failed to query file by id")
	}

	fileInfoProto := &filepb.FileInfo{
		FileId:     file.ID,
		FileName:   file.FileName,
		Bytes:      file.FileSize,
		FileType:   file.FileType,
		Tag:        file.Tag,
		UploaderId: file.UploaderId,
		Content:    file.FileContent,
		CreatedAt:  timestamppb.New(file.CreatedAt),
		UpdatedAt:  timestamppb.New(file.UpdatedAt),
	}
	// 将获取的文件信息转换为 proto 格式
	return &filepb.QueryFileByIdResponse{
		FileInfo: fileInfoProto,
	}, nil

}

// 获取文件列表
func (s *FileServer) GetFileList(ctx context.Context, req *filepb.GetFileListRequest) (*filepb.GetFileListResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "GetFileList",
	})
	// 获取请求参数
	page := req.GetPage()
	logger.Infof("Received page: %d", page)
	pageSize := req.GetPageSize()
	logger.Infof("Received pageSize: %d", pageSize)

	// 获取请求头里面的auth
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Errorf("Failed to get metadata from context")
		return nil, status.Error(codes.Internal, "failed to get metadata from context")
	}
	authInfo := md.Get("authorization")
	if len(authInfo) == 0 {
		logger.Errorf("Failed to get authorization from context")
		return nil, status.Error(codes.Internal, "failed to get authorization from context")
	}
	// logger.Infof("Received authorization: %s", authInfo[0])

	// 解析token
	claims, err := auth.ParseUserToken(authInfo[0])
	if err != nil {
		logger.Errorf("Failed to parse token: %v", err)
		return nil, status.Error(codes.Internal, "failed to parse token")
	}
	logger.Infof("Received claims: %v", claims)

	// 先获取文件总数
	total, err := s.DBEngine.GetFileTotal()
	if err != nil {
		logger.Errorf("Failed to query file total: %v", err)
		return nil, status.Error(codes.Internal, "failed to query file total")
	}

	// 如果是管理员，可以查看所有文件，否则只能查看自己上传的文件
	var files []db.UploadFile
	if claims.Role == "ADMIN" {
		files, err = s.DBEngine.GetFileList(page, pageSize)
	} else {
		files, err = s.DBEngine.GetFileListByUserId(claims.UserId, page, pageSize)
	}
	if err != nil {
		logger.Errorf("Failed to query file list: %v", err)
		return nil, status.Error(codes.Internal, "failed to query file list")
	}

	// // 将获取的文件信息转换为 proto 格式
	var fileListInfoProtos []*filepb.FileListInfo
	for _, file := range files {
		fileListInfoProtos = append(fileListInfoProtos, &filepb.FileListInfo{
			FileId:    file.ID,
			FileName:  file.FileName,
			Bytes:     file.FileSize,
			FileType:  file.FileType,
			Tag:       file.Tag,
			CreatedAt: timestamppb.New(file.CreatedAt),
			UpdatedAt: timestamppb.New(file.UpdatedAt),
		})

	}

	return &filepb.GetFileListResponse{
		FileInfos: fileListInfoProtos,
		Total:     total,
	}, nil

}

// // 文件下载
// func (s *FileServer) DownloadFile(ctx context.Context, req *filepb.DownloadFileRequest) (*filepb.DownloadFileResponse, error) {
// 	logger := log.WithFields(log.Fields{
// 		"api": "DownloadFile",
// 	})
// 	fileId := req.GetFileId()
// 	// logger.Infof("Received file id: %d", fileId)
// 	// 查询文件信息
// 	file, err := s.DBEngine.GetFileByID(fileId)
// 	if err != nil {
// 		logger.Errorf("Failed to query file by id: %v", err)
// 		return nil, status.Error(codes.Internal, "failed to query file by id")
// 	}

// 	// 返回文件信息
// 	return &filepb.DownloadFileResponse{
// 		Content: file.FileContent,
// 	}, nil

// }
