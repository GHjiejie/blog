package main

import (
	"blog-backend/server"
	"log"
)

func main() {
	// 启动服务
	svc, err := server.NewBlogServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
