package auth

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 对前端传递的密码和数据库中的密码进行比较
func ComparePassword(dbPassword, rePasswd string) error {
	log.Infof("对前端传递的密码和数据库中的密码进行比较")
	log.Infof("dbPassword: %s, rePasswd: %s", dbPassword, rePasswd)
	// 直接比较
	if dbPassword != rePasswd {
		return status.Errorf(codes.InvalidArgument, "password is not correct")
	}
	log.Info("密码校验正确")
	// aesPW, err := utils.AESDecrypt(rePasswd)
	// if err != nil {
	// 	log.Errorf("failed to decrypt password: %v", err)
	// 	return status.Errorf(codes.Internal, "failed to decrypt password: %v", err)
	// }
	// log.Info("解密后的密码: ", aesPW)
	return nil
}
