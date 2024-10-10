package server

import "blog-backend/pkg/db"

// Start 函数启动 HTTP 服务器
func NewBlogServer() {
	// 首先要进行数据库的连接,使用的是
	db.ConnectDatabase()
	// 然后初始化路由
}
