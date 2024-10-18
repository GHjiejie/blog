package server

import (
	"blog-backend/pkg/casbinpermit"
	"log"
	"net/http"
)

// 设置无需验证的路径
var NotNeedAuthorizationPaths = map[string]bool{
	"/v1/user/login":    true,
	"/v1/user/register": true,
	"/v1/captchas":      true,
}

const (
	AuthorizationHeaderKey = "Authorization"
	AuthorizationQueryKey  = "authorization"
	APIKeyHeaderKey        = "APIKey"
	ContentTypeHeaderKey   = "Content-Type"
	ContentTypeHeaderVal   = "application/json"
)

func (s *BlogServer) Tracing(nextHandle http.Handler, userPermit *casbinpermit.Permit) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求体中获取用户信息
		username := r.Header.Get("username")
		log.Printf("username: %s", username)
		// 判断当前接口是否需要进行鉴权验证
		if needAuthorizations(r.URL.Path, r.Method) {
			var (
				username  string
				userToken string
			)
			log.Printf("当前接口需要进行鉴权验证,同时需要进行权限检查")
			// 前端在登录成功后，会在下一次请求的时候将token设置在请求头里面，所以我们需要从请求头中获取token
			authHeader := r.Header.Get(AuthorizationHeaderKey)
			// 也可以从请求参数中获取token（这个主要是用于调试？）
			authQuery := r.URL.Query().Get(AuthorizationQueryKey)
			if authHeader != "" {
				userToken = authHeader
			} else if authQuery != "" {
				userToken = authQuery
			}
			// 接下来我们就需要将token解析
			// claims, err := s.auth.ParseUserToken(userToken)
			claims, err := s.auth.ParseUserToken(userToken)
			if err != nil {
				log.Printf("token解析失败, err: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if claims == nil {
				log.Printf("token解析失败, claims为空")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			username = claims.Username
			log.Printf("token解析成功, 输出获取的claims: %v", claims)

			// 我们直接进行权限验证
			permitted, err := userPermit.CheckPermission(username, r.URL.Path, r.Method)
			if err != nil {
				log.Printf("权限验证失败, err: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// 如果首次校验失败,则重新从数据库加载策略,然后再次进行校验
			if !permitted {
				if err := userPermit.Enforcer.LoadPolicy(); err != nil {
					log.Printf("failed to load policy with err(%s)", err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				permitted, err = userPermit.CheckPermission(username, r.URL.Path, r.Method)
				if err != nil {
					log.Printf("权限验证失败, err: %v", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			log.Printf("user(%s) permission 1st checking result, permit(%v)", username, permitted)

			// 进行鉴权验证
			// TODO
		} else {

			log.Printf("当前接口无需进行鉴权验证")
		}
		log.Printf("接下来将请求转发给下一个处理函数")
		nextHandle.ServeHTTP(w, r)
	})
}

func needAuthorizations(urlPath, method string) bool {
	log.Printf("判断当前接口是否需要进行鉴权验证")
	log.Printf("urlPath: %s, method: %s", urlPath, method)
	need := true
	for path := range NotNeedAuthorizationPaths {
		if path == urlPath {
			need = false
			break
		}
	}
	return need
}
