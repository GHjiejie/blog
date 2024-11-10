package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "blog-backend/pb/user"
	"blog-backend/pkg/auth"
	"blog-backend/pkg/db"
	"blog-backend/pkg/validate"
)

func (s *BlogServer) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	log.Infof("Received Test request: %v", req)
	return &pb.TestResponse{Message: "Hello, " + req.Name}, nil
}

// Register 实现用户注册逻辑(只有管理员才可以)
func (s *BlogServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	logger := log.WithFields(log.Fields{
		"api": "Register",
	})
	// 输出一下这个ctx看看有什么信息
	// log.Printf("输出ctx: %v", ctx)
	// 获取用户名
	username := req.GetUsername()
	// log.Infof("输出获取的用户名: %s", username)
	// 只允许admin用户去添加新用户
	if err := s.auth.IsAdmin(ctx); err != nil {
		logger.Errorf("failed to check user is admin with err(%s)", err.Error())
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	user := db.User{
		Username: username,
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}

	// log.Infof("输出user: %v", user)
	// 将用户信息存储到数据库中
	userId, err := s.DBEngine.Register(user)
	if err != nil {
		logger.Errorf("failed to register user with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "register user failed with err(%s)", err.Error())
	}
	// log.Infof("注册用户成功, 用户ID: %d", userId)
	// TODO 将前端传递过来的密码进行解密
	// 然后对密码的有效性进行校验(前端已经校验了一次,这个是第二次(后端)校验)
	return &pb.RegisterResponse{
		User: &pb.User{
			// 这里我们可以将我们接口涉及的字段一个一个编写映射返回给前端，但是考虑到使用的场景会比较多，所以我们就将其封装成一个函数，然后在函数中进行映射返回给前端
			UserId:   userId,
			Username: username,
			Role:     req.GetRole(),
		},
	}, nil

}

// 用户登录
func (s *BlogServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// log.Infof("Received Login request: %v", req)
	logger := log.WithFields(log.Fields{
		"api": "Login",
	})
	username := req.GetUsername()
	password := req.GetPassword()
	log.Infof("username: %s", username)
	log.Infof("password:%s", password)
	// TODO 进行验证码校验

	// 去数据库加载策略
	if err := s.casbinPermit.LoadPolicy(); err != nil {
		logger.Errorf("failed to load policy with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "load policy failed with err(%s)", err.Error())
	}
	//从数据库中查询用户信息
	user, err := s.DBEngine.UserGetByName(username)
	if err != nil {
		logger.Errorf("failed to get user by name with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
	}
	log.Infof("输出从数据库获取的密码:%v", user.Password)
	// TODO 将用户登录的密码与数据库中的密码进行对比（这个后面在支持，需要对密码进行加密与解密）
	if err = auth.ComparePassword(user.Password, password); err != nil {
		logger.Errorf("failed to compare password with err(%s)", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// 服务端生成token(需要用户提供用户ID、用户名、角色等信息)
	token, err := s.auth.GenUserToken(user.ID, user.Username, user.Role.String())
	if err != nil {
		logger.Errorf("failed to generate token with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "generate token failed with err(%s)", err.Error())
	}

	// 然后我们需要更新当前用户的token信息

	err = s.DBEngine.UserUpdate(user.ID, map[string]interface{}{"token": token})
	if err != nil {
		logger.Errorf("failed to update user token with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "update user token failed with err(%s)", err.Error())
	}

	// log.Printf("更新用户token成功")
	logger.Info("登录成功")
	// 返回登录成功的响应

	return &pb.LoginResponse{
		Token: token,
		User: &pb.User{
			Username: user.Username,
		},
	}, nil
}

// 用户登出
func (s *BlogServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "Logout",
	})
	// 首先我们要获取用户ID
	userID := req.GetUserId()
	// log.Infof("输出获取到的用户ID: %v", userID)
	// 然后我们需要去数据库中获取用户信息
	user, err := s.DBEngine.UserGetByID(userID)
	if err != nil {
		logger.Errorf("failed to get user by id with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
	}
	log.Infof("输出获取到的用户信息: %v", user)
	// 然后我们需要去更新用户的token信息
	err = s.DBEngine.UserUpdate(userID, map[string]interface{}{"token": ""})
	if err != nil {
		logger.Errorf("failed to update user token with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "update user token failed with err(%s)", err.Error())
	}
	// log.Info("用户登出成功")
	logger.Info("logout success")
	return &pb.LogoutResponse{
		Message: "logout success",
	}, nil
}

// 用户列表获取
func (s *BlogServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	// log.Infof("Received ListUser request: %v", req)
	logger := log.WithFields(log.Fields{
		"api": "ListUser",
	})

	// 首先要获取page
	page := req.GetPage()
	// 再获取每页的数量(默认是2吧)
	pageSize := req.GetPageSize()
	// log.Infof("输出获取到的page: %d", page)
	// log.Infof("输出获取到的pageSize: %d", pageSize)

	// 首先我们要获取用户总数
	userCount, err := s.DBEngine.UserCount()
	if err != nil {
		logger.Errorf("failed to get user count with err(%s)", err.Error())
		return nil, err
	}
	logger.Infof("获取到的用户总数: %d", userCount)
	users, err := s.DBEngine.UserList(int64(page), int64(pageSize))
	if err != nil {
		logger.Errorf("failed to get user list with err(%s)", err.Error())
		return nil, err
	}
	// 返回数据
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			UserId:    user.ID,
			Username:  user.Username,
			Role:      pb.Role(user.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	return &pb.ListUserResponse{
		Users: pbUsers,
		Total: int32(userCount),
	}, nil
}

// 删除用户（管理员）
func (s *BlogServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// 在输出日志的时候我们需要提供一些关键的信息来帮助我们更好的定位问题，比如我们需要输出当前的API名称，用户ID等信息
	logger := log.WithFields(log.Fields{
		"api": "DeleteUser",
	})
	if err := validate.UserDelete(req); err != nil {
		logger.Errorf("failed to delete user with err(%s)", err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	userID := req.GetUserId()

	logger = logger.WithFields(log.Fields{
		"UserID": userID,
	})

	// 然后判断是不是admin用户
	if err := s.auth.IsAdmin(ctx); err != nil {
		logger.Errorf("failed to check user is admin with err(%s)", err.Error())
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	// logger.Info("用户是系统管理员,可以进行用户删除操作")

	// 接下来就是去根据id去这个用户是否存在
	user, err := s.DBEngine.UserGetByID(userID)
	if err != nil {
		logger.Errorf("failed to get user by id with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
	}
	// logger.Infof("输出在删除用户的时候获取到的用户信息: %v", user)
	// 然后就进行删除操作
	err = s.DBEngine.UserDelete(userID)
	if err != nil {
		logger.Errorf("failed to delete user with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "delete user failed with err(%s)", err.Error())
	}

	// logger.Info("删除用户成功")
	return &pb.DeleteUserResponse{
		User: &pb.User{
			UserId: user.ID,
		},
	}, nil
}

// 更新用户信息
// func (s *BlogServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
// 	logger := log.WithFields(log.Fields{
// 		"api": "UpdateUser",
// 	})
// 	// 首先我们要判断用户是否是管理员
// 	if err := s.auth.IsAdmin(ctx); err != nil {
// 		logger.Errorf("failed to check user is admin with err(%s)", err.Error())
// 		return nil, status.Errorf(codes.PermissionDenied, err.Error())
// 	}
// 	// logger.Info("用户是系统管理员,可以进行用户更新操作")
// 	// 获取用户ID
// 	userID := req.GetUserId()
// 	// logger.Infof("输出获取到的用户ID: %v", userID)
// 	// 获取用户信息
// 	user, err := s.DBEngine.UserGetByID(userID)
// 	if err != nil {
// 		logger.Errorf("failed to get user by id with err(%s)", err.Error())
// 		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
// 	}
// 	// logger.Infof("输出在更新用户的时候获取到的用户信息: %v", user)
// 	// 更新用户信息
// 	err = s.DBEngine.UserUpdate(userID, map[string]interface{}{
// 		"username": req.GetUsername(),
// 		"role":     req.GetRole(),
// 	})
// 	if err != nil {
// 		logger.Errorf("failed to update user with err(%s)", err.Error())
// 		return nil, status.Errorf(codes.Internal, "update user failed with err(%s)", err.Error())
// 	}
// 	// logger.Info("更新用户成功")
// 	return &pb.UpdateUserResponse{
// 		User: &pb.User{
// 			UserId:   user.ID,
// 			Username: req.GetUsername(),
// 			Role:     req.GetRole(),
// 		},
// 	}, nil
// }

// 用户密码重置
func (s *BlogServer) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "ResetPassword",
	})
	// 首先我们要判断用户是否是管理员
	if err := s.auth.IsAdmin(ctx); err != nil {
		logger.Errorf("failed to check user is admin with err(%s)", err.Error())
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	// 获取用户ID
	userID := req.GetUserId()
	// logger.Infof("输出获取到的用户ID: %v", userID)
	// 获取用户信息
	user, err := s.DBEngine.UserGetByID(userID)

	if err != nil {
		logger.Errorf("failed to get user by id with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
	}
	// 判断操作的对象的角色是否是admin（因为我们是不允许对admin进行重置密码操作的）
	// logger.Infof("输出用户的角色: %v", user.Role)
	// logger.Infof("输出用户的角色: %v", *pb.Role_ADMIN.Enum())
	if user.Role == *pb.Role_ADMIN.Enum() {
		logger.Errorf("failed to reset password with err(%s)", "admin user can not reset password")
		return nil, status.Errorf(codes.PermissionDenied, "admin user can not reset password")
	}
	newPassword := "12345" //所有人的初始化密码都是12345
	// 更新用户信息
	err = s.DBEngine.UserUpdate(userID, map[string]interface{}{
		"password": newPassword,
	})
	if err != nil {
		logger.Errorf("failed to update user with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "update user failed with err(%s)", err.Error())
	}
	logger.Info("user password reset success")
	return &pb.ResetPasswordResponse{
		Message: "user password reset success",
	}, nil
}

// 根据用户ID获取用户信息
func (s *BlogServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	logger := log.WithFields(log.Fields{
		"api": "GetUserByID",
	})
	// 首先我们要获取用户ID
	userID := req.GetUserId()
	// logger.Infof("输出获取到的用户ID: %v", userID)
	// 获取用户信息
	user, err := s.DBEngine.UserGetByID(userID)
	if err != nil {
		logger.Errorf("failed to get user by id with err(%s)", err.Error())
		return nil, status.Errorf(codes.Internal, "get user from db failed with err(%s)", err.Error())
	}
	// logger.Infof("输出获取到的用户信息: %v", user)
	return &pb.GetUserResponse{
		User: &pb.User{
			UserId:    user.ID,
			Username:  user.Username,
			Role:      pb.Role(user.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
