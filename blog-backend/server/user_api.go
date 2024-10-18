package server

import (
	"context"
	"log"

	pb "blog-backend/pb/user"
	"blog-backend/pkg/db"
)

func (s *BlogServer) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Received Test request: %v", req)
	return &pb.TestResponse{Message: "Hello, " + req.Name}, nil
}

// Register 实现用户注册逻辑(只有管理员才可以)
func (s *BlogServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received RegisterUser request: %v", req)
	// 输出一下这个ctx看看有什么信息
	// log.Printf("输出ctx: %v", ctx)
	// 获取用户名
	username := req.GetUsername()
	log.Printf("输出获取的用户名: %s", username)
	// 只允许admin用户去添加新用户
	if err := s.auth.IsAdmin(ctx); err != nil {
		log.Printf("failed to check user is admin with err(%s)", err.Error())
		return nil, err
	}
	user := db.User{
		Username: username,
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}
	log.Printf("输出user: %v", user)
	// 将用户信息存储到数据库中
	userID, err := s.DBEngine.Register(user)
	if err != nil {
		log.Printf("failed to register user with err(%s)", err.Error())
		return nil, err
	}
	log.Printf("注册用户成功, 用户ID: %d", userID)
	// TODO 将前端传递过来的密码进行解密
	// 然后对密码的有效性进行校验(前端已经校验了一次,这个是第二次(后端)校验)
	return &pb.RegisterResponse{}, nil

}

func (s *BlogServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received Login request: %v", req)
	username := req.GetUsername()
	password := req.GetPassword()
	log.Printf("username: %s", username)
	log.Printf("password:%s", password)
	// TODO 进行验证码校验

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
	// log.Printf("user: %v", user)

	// 服务端生成token(需要用户提供用户ID、用户名、角色等信息)
	token, err := s.auth.GenUserToken(user.ID, user.Username, user.Role.String())
	if err != nil {
		log.Printf("failed to generate token with err(%s)", err.Error())
		return nil, err
	}
	// log.Printf("输出生成的token: %s", token)

	// 然后我们需要更新当前用户的token信息

	err = s.DBEngine.UserUpdate(user.ID, map[string]interface{}{"token": token})
	if err != nil {
		log.Printf("failed to update user token with err(%s)", err.Error())
		return nil, err
	}

	// log.Printf("更新用户token成功")
	log.Printf("登录成功")
	// 返回登录成功的响应

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
