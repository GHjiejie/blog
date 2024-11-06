package db

import "gorm.io/gorm"

type Handler interface {
	// 根据用户名获取用户信息
	UserGetByName(username string) (*User, error)
	// 根据用户ID获取用户信息
	UserGetByID(userID int64) (*User, error)
	// 更新用户的消息
	UserUpdate(userID int64, fields map[string]interface{}) error
	// 注册用户
	Register(user User) (int64, error)
	// 获取用户总数
	UserCount() (int64, error)
	// 获取用户列表
	UserList(page int64, pageSize int64) ([]User, error)
	// 删除用户（管理员）
	UserDelete(userID int64) error
	GetORMDB() *gorm.DB
}

func NEWHandler() (Handler, error) {
	return NEWSQLDB()
}
