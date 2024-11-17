package constant

import (
	pb "blog-backend/pb/user"
)

// 用户权限
const (
	DefaultUserAdminPolicy  = "user-admin-policy"  // 管理员
	DefaultUserNormalPolicy = "user-normal-policy" // 普通用户
)

var DefaultUserRole = map[pb.Role]string{
	pb.Role_ADMIN: DefaultUserAdminPolicy,
	pb.Role_USER:  DefaultUserNormalPolicy,
}
