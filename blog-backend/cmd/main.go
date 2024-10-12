package main

import (
	"blog-backend/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 设置信号量，用于优雅关闭服务
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// 启动服务
	svr, err := server.NewBlogServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	if err := svr.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
		return
	}
	<-shutdown

}
