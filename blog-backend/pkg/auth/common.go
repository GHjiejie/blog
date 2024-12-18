package auth

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	pb "blog-backend/pb/user"

	jwt "github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
		log.Info("failed to sign token with err:", err.Error())
		return "", err
	}
	return tokenString, nil
}

// 解析用户的token
func (au *Auth) ParseUserToken(tokenStr string) (*UserClaims, error) {
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

	// 校验用户名和角色
	if userClaims.Username == " " || userClaims.Role == " " {
		log.Error("username or role is empty")
		return nil, err
	}
	// 校验用户是否存在
	user, err := au.db.UserGetByName(userClaims.Username)
	if err != nil {
		log.Errorf("failed to get user by username with err(%s)", err.Error())
		return nil, err
	}
	// log.Info("获取用户声明成功,同时在数据库查询到用户消息，现在输出user: %v", user)

	// 判断从数据库获取的用户的消息里面是否有token
	if user.Token == "" {
		log.Error("token is empty")
		return nil, status.Errorf(codes.Unauthenticated, "token is empty")
	}
	// 限制只可以单用户登录(这里面可以细分,比如说是开发环境还是生产环境,这个我们可以放在配置文件中)(TODO)
	if user.Token != tokenStr {
		log.Error("token is invalid")
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid")
	}

	return userClaims, nil
}

// 获取用户的声明
func (au *Auth) getUserClaims(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not found")
	}
	// log.Info("输出获取的metadata: %v", md)

	// 获取token
	token, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "token is not found")
	}
	// log.Info("输出获取的token: %v", token)

	// 解析token
	claims, err := au.ParseUserToken(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid")
	}

	return claims, nil
}

// 判断用户是不是admin用户
func (au *Auth) IsAdmin(ctx context.Context) error {
	claims, err := au.getUserClaims(ctx)
	if err != nil {
		return err
	}
	log.Info("输出获取的claims: ", claims)
	// 注意,我们需要从定义生成的user.pb.go文件里面去找用户的身份
	if claims.Role != pb.Role_ADMIN.String() {
		return status.Errorf(codes.PermissionDenied, "user is not admin")
	} else {
		log.Info("用户是系统管理员,可以进行用户添加操作")
	}
	return nil

}
