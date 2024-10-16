package auth

import (
	"blog-backend/pkg/casbinpermit"
	"blog-backend/pkg/db"
	"time"
)

const (
	DefaultTokenExpireDuration = "12h"
	DefaultAuthSecretKey       = "blog"
	MaxTokenExpireDuration     = 36 * time.Hour // 36h0m0s
)

type Auth struct {
	db                  db.Handler
	tokenExpireDuration time.Duration
	casbinPermit        *casbinpermit.Permit
}

func NewAuth(db db.Handler, casbinPermit *casbinpermit.Permit) *Auth {
	return &Auth{
		db:                  db,
		tokenExpireDuration: MaxTokenExpireDuration,
		casbinPermit:        casbinPermit,
	}
}
