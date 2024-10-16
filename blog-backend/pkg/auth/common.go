package auth

import (
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserId   int64
	Username string
	Role     string
	jwt.RegisteredClaims
}

func (au *Auth) GenUserToken(userId int64, username, role string) (string, error) {

	userClaims := &UserClaims{
		UserId:   userId,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(au.tokenExpireDuration)),
		},
	}

	// 使用指定的签名方法创建一个新的令牌生成器。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// 使用指定的密钥签名并返回令牌字符串。
	tokenString, err := token.SignedString([]byte(DefaultAuthSecretKey))
	if err != nil {
		log.Printf("failed to sign token with err(%s)", err.Error())
		return "", err
	}
	return tokenString, nil
}
