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
		Password:  req.Password, // 注意：在实际应用中，请不要明文存储密码
		Nickname:  req.Nickname,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      pb.Role_USER,
		Status:    pb.Status_NORMAL,
		CreatedAt: "2023-10-15T12:00:00Z", // 使用当前时间戳
		UpdatedAt: "2023-10-15T12:00:00Z",
	}

	return &pb.RegisterResponse{User: newUser}, nil
	// return &pb.RegisterResponse{}, nil
}

// func (s *BlogServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
// 	log.Printf("Received Login request: %v", req)
// 	return &pb.LoginResponse{}, nil
// }
