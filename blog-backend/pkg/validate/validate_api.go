package validate

import (
	pb "blog-backend/pb/user"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UserDelete(req *pb.DeleteUserRequest) error {
	userID := req.GetUserId()
	log.Infof("Received DeleteUser request: %v", userID)
	if req.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user_id")
	}
	return nil
}
