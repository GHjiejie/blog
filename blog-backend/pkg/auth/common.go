package auth

import (
	"log"
	"time"

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

// 生成token
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

func (au *Auth) ParseUserToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// SigningMethodHMAC
		// 其中 SigningMethodHS256 为其中一种
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Fatalf("Unexpected signing method: %v", token.Header["alg"])
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return []byte(DefaultAuthSecretKey), nil
	})
	if err != nil {
		log.Fatalf("failed to parse token with err(%s)", err.Error())
		return nil, err
	}

	// 	 提取用户声明，如果我们的token存储了一些用户信息，我们可以通过下面的方法将用户的信息提取出来
	// 参考我们上面token生成的代码,我们token的生成使用了用户的ID 用户名和用户身份,所以我们下面代码解析后的token也会获取这个消息
	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Fatalf("convert token claims to UserClaims failed")
		return nil, err
	}
	if !token.Valid {
		log.Fatalf("token is invalid")
		return nil, err
	}
	// log.Printf("用户token鉴权成功，现在输出UserClaims: %v", userClaims)

	// 校验用户名和角色
	if userClaims.Username == " " || userClaims.Role == " " {
		log.Fatalf("username or role is empty")
		return nil, err
	}
	// 校验用户是否存在
	user, err := au.db.UserGetByName(userClaims.Username)
	if err != nil {
		log.Fatalf("failed to get user by username with err(%s)", err.Error())
		return nil, err
	}
	// log.Printf("获取用户声明成功,同时在数据库查询到用户消息，现在输出user: %v", user)

	// 判断从数据库获取的用户的消息里面是否有token
	if user.Token == "" {
		log.Fatalf("token is empty")
		return nil, err
	}
	// 限制只可以单用户登录(这里面可以细分,比如说是开发环境还是生产环境,这个我们可以放在配置文件中)(TODO)
	if user.Token != tokenStr {
		log.Fatalf("token is invalid")
		return nil, err
	}

	return userClaims, nil

}
