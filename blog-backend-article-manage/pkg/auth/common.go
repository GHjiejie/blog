package auth

import (
	log "github.com/sirupsen/logrus"

	jwt "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserClaims struct {
	UserId   int64
	Username string
	Role     string
	jwt.RegisteredClaims
}

// 解析用户的token
func ParseUserToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// SigningMethodHMAC
		// 其中 SigningMethodHS256 为其中一种
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return []byte(DefaultAuthSecretKey), nil
	})
	if err != nil {
		log.Errorf("failed to parse token with err(%s)", err.Error())
		return nil, err
	}

	// 	 提取用户声明，如果我们的token存储了一些用户信息，我们可以通过下面的方法将用户的信息提取出来
	// 参考我们上面token生成的代码,我们token的生成使用了用户的ID 用户名和用户身份,所以我们下面代码解析后的token也会获取这个消息
	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Error("convert token claims to UserClaims failed")
		return nil, err
	}
	if !token.Valid {
		log.Error("token is invalid")
		return nil, err
	}
	// log.Info("用户token鉴权成功，现在输出UserClaims: %v", userClaims)

	return userClaims, nil
}

// 获取用户的声明
