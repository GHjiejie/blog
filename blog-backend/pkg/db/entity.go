package db

import (
	pb "blog-backend/pb/user"
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;primaryKey"             json:"id"`
	Username  string    `gorm:"column:username;unique"     json:"username"`
	Password  string    `gorm:"column:password"            json:"password"`
	Role      pb.Role   `gorm:"column:role"                     json:"role"`
	Token     string    `gorm:"column:token"                     json:"token"`
	Email     string    `gorm:"column:email"                     json:"email"`
	Phone     string    `gorm:"column:phone"                     json:"phone"`
	Avatar    string    `gorm:"column:avatar"                    json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`
}

func (u User) TableName() string {
	return "users"
}

// casbin_rules 特殊处理, 因为无需额外定义 model
func GetCasbinRuleTableName() string {
	return "casbin_rules"
}
