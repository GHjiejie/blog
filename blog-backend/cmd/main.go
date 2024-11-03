package main

import (
	"blog-backend/server"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func main() {

	// 设置日志的输出格式
	log.AddHook(filename.NewHook())
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	// 设置信号量，用于优雅关闭服务
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// 启动服务
	svr, err := server.NewBlogServer()
	if err != nil {
		log.Error("failed to start server:", err)
	}
	if err := svr.Start(); err != nil {
		log.Error("Failed to start server:", err)
		return
	}
	<-shutdown

}
