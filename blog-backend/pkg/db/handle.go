package db

import "gorm.io/gorm"

type Handler interface {
	// 根据用户名获取用户信息
	UserGetByName(username string) (*User, error)
	// 更新用户的消息
	UserUpdate(userID int64, fields map[string]interface{}) error
	// 注册用户
	Register(user User) (int64, error)
	GetORMDB() *gorm.DB
}

func NEWHandler() (Handler, error) {
	return NEWSQLDB()
}
