package auth

import (
	"blog-backend/pkg/casbinpermit"
	"blog-backend/pkg/config"
	"blog-backend/pkg/db"
	"time"
)

const (
	DefaultTokenExpireDuration = "24h"
	DefaultAuthSecretKey       = "blog"
	MaxTokenExpireDuration     = 36 * time.Hour // 36h0m0s
)

type Auth struct {
	db                  db.Handler
	config              *config.Config
	tokenExpireDuration time.Duration
	casbinPermit        *casbinpermit.Permit
}

func NewAuth(db db.Handler, config *config.Config, casbinPermit *casbinpermit.Permit) *Auth {
	if config.UserAuth.TokenExpireDuration == "" {
		config.UserAuth.TokenExpireDuration = DefaultTokenExpireDuration
	}
	// 用于加密生成token, 前端或者其他接口用户使用
	if config.UserAuth.SecretKey == "" {
		config.UserAuth.SecretKey = DefaultAuthSecretKey
	}
	return &Auth{
		db:                  db,
		config:              config,
		tokenExpireDuration: MaxTokenExpireDuration,
		casbinPermit:        casbinPermit,
	}
}
