package main

import (
	"blog-backend/pkg/config"
	"blog-backend/server"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

var (
	// 定义一个名为 --config 的命令行标志，用于指定配置文件的位置
	// 	如果在命令行中没有提供该标志，程序将默认使用 /config/console.toml 作为配置文件路径。
	// 用户可以通过在运行程序时传递 --config 参数来自定义配置文件的位置，
	configFile = flag.String("config", "../conf/config.local.toml", "")
	// configFile = flag.String("config", "./conf/config.toml", "")
)

func processConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Errorf("Failed to load config(%v) and master service is exiting with err(%v):", *configFile, err)
		return nil, err
	}
	return cfg, nil
}

func main() {
	cfg, err := processConfig()
	if err != nil {
		log.Errorf("Failed to load config(%v) and master service is exiting with err(%v):", *configFile, err)
		return
	}
	log.Infof("config: %v", cfg.DBConfig)
	log.Infof("config: %v", cfg.GRPCEndpoint)
	log.Infof("config: %v", cfg.HTTPEndpoint)

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
	svr, err := server.NewBlogServer(cfg)
	if err != nil {
		log.Error("failed to start server:", err)
	}
	if err := svr.Start(); err != nil {
		log.Error("Failed to start server:", err)
		return
	}
	<-shutdown

}
