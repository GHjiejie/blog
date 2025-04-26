package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"articleManage/pkg/config"
	"articleManage/server"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

var (
	// 定义一个名为 --config 的命令行标志，用于指定配置文件的位置
	// 	如果在命令行中没有提供该标志，程序将默认使用 /config/console.toml 作为配置文件路径。
	// 用户可以通过在运行程序时传递 --config 参数来自定义配置文件的位置，
	// configFile = flag.String("config", "./conf/config.toml", "")
	configFile = flag.String("config", "./conf/config.toml", "")
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

	// 首先设置日志级别
	log.AddHook(filename.NewHook())
	switch strings.ToLower(cfg.LogLevel) {
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

	// 创建一个信号量，用于优雅关闭服务
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 创建并启动服务
	svr, err := server.NewArticleServer(cfg)
	if err != nil {
		log.Error("failed to start server:", err)
	}

	if err := svr.Start(); err != nil {
		log.Error("failed to start server:", err)
		return
	}
	// log.Info("server started successfully")
	<-shutdown
}
