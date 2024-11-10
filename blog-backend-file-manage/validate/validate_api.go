package validate

import (
	filepb "fileManage/pb/fileManage"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FileQueryById(req *filepb.QueryFileByIdRequest) error {
	fileId := req.GetFileId()
	log.Infof("Received DeleteUser request: %v", fileId)
	if req.GetFileId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user_id")
	}
	return nil
}

func FileDelete(req *filepb.DeleteFileRequest) error {
	fileId := req.GetFileId()
	log.Infof("Received DeleteUser request: %v", fileId)
	if req.GetFileId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user_id")
	}
	return nil
}
