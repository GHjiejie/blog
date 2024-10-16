package server

import (
	"context"
	"log"

	pb "blog-backend/pb/user"
)

func (s *BlogServer) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Received Test request: %v", req)
	return &pb.TestResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

// Register 实现用户注册逻辑
func (s *BlogServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received RegisterUser request: %v", req)

	// 这里可以添加实际的注册逻辑，例如将用户信息保存到数据库
	newUser := &pb.User{
		UserId:    "generated_user_id", // 您可以使用 UUID 或其他方法生成用户 ID
		Username:  req.Username,
		Email:     "3426571530@qq.com",
		Phone:     "18196756670",
		Role:      pb.Role_USER,
		Status:    pb.Status_NORMAL,
		CreatedAt: "2023-10-15T12:00:00Z", // 使用当前时间戳
		UpdatedAt: "2023-10-15T12:00:00Z",
	}

	return &pb.RegisterResponse{User: newUser}, nil
	// return &pb.RegisterResponse{}, nil
}

func (s *BlogServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received Login request: %v", req)
	username := req.GetUsername()
	password := req.GetPassword()
	log.Printf("username: %s", username)
	log.Printf("password:%s", password)

	// 去数据库加载策略
	if err := s.casbinPermit.LoadPolicy(); err != nil {
		log.Printf("failed to load policy with err(%s)", err.Error())
		return nil, err
	}
	//从数据库中查询用户信息
	user, err := s.DBEngine.UserGetByName(username)
	if err != nil {
		log.Printf("failed to get user by name with err(%s)", err.Error())
		return nil, err
	}
	// TODO 将用户登录的密码与数据库中的密码进行对比（这个后面在支持，需要对密码进行加密与解密）
	log.Printf("user: %v", user)

	// 服务端生成token(需要用户提供用户ID、用户名、角色等信息)
	token, err := s.auth.GenUserToken(user.ID, user.Username, user.Role.String())
	if err != nil {
		log.Printf("failed to generate token with err(%s)", err.Error())
		return nil, err
	}
	log.Printf("输出生成的token: %s", token)

	// 然后我们需要更新当前用户的token信息

	err = s.DBEngine.UserUpdate(user.ID, map[string]interface{}{"token": token})
	if err != nil {
		log.Printf("failed to update user token with err(%s)", err.Error())
		return nil, err
	}

	log.Printf("更新用户token成功")
	log.Printf("登录成功")
	// 返回登录成功的响应

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
