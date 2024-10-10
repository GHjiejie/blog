package server

import (
	"blog-backend/pkg/db"
	"log"
)

type BlogServer struct {
	DBEngine db.Handle
}

// Start 函数启动 HTTP 服务器
func NewBlogServer() (*BlogServer, error) {
	// 首先要进行数据库的连接,使用的是
	DBEngine, err := db.NEWHandler()
	if err != nil {
		log.Printf("failed to create db engine handler with err(%s)", err.Error())
		return nil, err
	}

}
