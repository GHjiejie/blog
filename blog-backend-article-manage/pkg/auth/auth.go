package auth

import "time"

const (
	DefaultTokenExpireDuration = "24h"
	DefaultAuthSecretKey       = "blog"
	MaxTokenExpireDuration     = 36 * time.Hour // 36h0m0s
)
