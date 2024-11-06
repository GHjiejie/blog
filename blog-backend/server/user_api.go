package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "blog-backend/pb/user"
	"blog-backend/pkg/db"
	"blog-backend/pkg/validate"
)

func (s *BlogServer) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	log.Infof("Received Test request: %v", req)
	return &pb.TestResponse{Message: "Hello, " + req.Name}, nil
}

// Register 实现用户注册逻辑(只有管理员才可以)
func (s *BlogServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Infof("Received RegisterUser request: %v", req)
	// 输出一下这个ctx看看有什么信息
	// log.Printf("输出ctx: %v", ctx)
	// 获取用户名
	username := req.GetUsername()
	// log.Infof("输出获取的用户名: %s", username)
	// 只允许admin用户去添加新用户
	if err := s.auth.IsAdmin(ctx); err != nil {
		log.Errorf("failed to check user is admin with err(%s)", err.Error())
		return nil, err
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
		log.Errorf("failed to register user with err(%s)", err.Error())
		return nil, err
	}
	log.Infof("注册用户成功, 用户ID: %d", userId)
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
	log.Infof("Received Login request: %v", req)
	username := req.GetUsername()
	password := req.GetPassword()
	log.Infof("username: %s", username)
	log.Infof("password:%s", password)
	// TODO 进行验证码校验

	// 去数据库加载策略
	if err := s.casbinPermit.LoadPolicy(); err != nil {
		log.Errorf("failed to load policy with err(%s)", err.Error())
		return nil, err
	}
	//从数据库中查询用户信息
	user, err := s.DBEngine.UserGetByName(username)
	if err != nil {
		log.Errorf("failed to get user by name with err(%s)", err.Error())
		return nil, err
	}
	// TODO 将用户登录的密码与数据库中的密码进行对比（这个后面在支持，需要对密码进行加密与解密）
	// log.Printf("user: %v", user)

	// 服务端生成token(需要用户提供用户ID、用户名、角色等信息)
	token, err := s.auth.GenUserToken(user.ID, user.Username, user.Role.String())
	if err != nil {
		log.Errorf("failed to generate token with err(%s)", err.Error())
		return nil, err
	}
	// log.Printf("输出生成的token: %s", token)

	// 然后我们需要更新当前用户的token信息

	err = s.DBEngine.UserUpdate(user.ID, map[string]interface{}{"token": token})
	if err != nil {
		log.Errorf("failed to update user token with err(%s)", err.Error())
		return nil, err
	}

	// log.Printf("更新用户token成功")
	log.Info("登录成功")
	// 返回登录成功的响应

	return &pb.LoginResponse{
		Token: token,
		User: &pb.User{
			Username: user.Username,
		},
	}, nil
}

// 用户列表获取
func (s *BlogServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.Infof("Received ListUser request: %v", req)

	// 首先要获取page
	page := req.GetPage()
	// 再获取每页的数量(默认是2吧)
	pageSize := req.GetPageSize()
	log.Infof("输出获取到的page: %d", page)
	log.Infof("输出获取到的pageSize: %d", pageSize)

	// 首先我们要获取用户总数
	userCount, err := s.DBEngine.UserCount()
	if err != nil {
		log.Info("获取用户总数失败！")
		return nil, err
	}
	log.Infof("输出获取到的用户总数: %d", userCount)
	users, err := s.DBEngine.UserList(int64(page), int64(pageSize))
	if err != nil {
		log.Info("获取用户列表失败！")
		return nil, err
	}
	log.Infof("输出获取到的用户列表: %v", users)
	// 返回数据
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			UserId:   user.ID,
			Username: user.Username,
			Role:     pb.Role(user.Role),
		})
	}
	log.Infof("输出转换后的用户列表: %v", pbUsers)
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
	// logger.Infof("输出获取的用户ID: %v", userID)

	// 然后判断是不死admin用户
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
