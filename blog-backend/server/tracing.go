package server

import (
	"blog-backend/pkg/casbinpermit"
	"bytes"
	"io"
	"log"
	"net/http"
)

// 设置无需验证的路径
var NotNeedAuthorizationPaths = map[string]bool{
	"/v1/users/login":    true,
	"/v1/users/register": true,
	"/v1/captchas":       true,
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
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重新设置 r.Body

		log.Printf("request body: %s", string(bodyBytes)) // 输出请求体内容

		log.Print("请求体", r.Body) // 输出所有请求头的信息

		// 从请求体中获取用户信息
		username := r.Header.Get("username")
		log.Printf("username: %s", username)
		// 判断当前接口是否需要进行鉴权验证
		if needAuthorizations(r.URL.Path, r.Method) {
			log.Printf("当前接口需要进行鉴权验证,同时需要进行权限检查")
			// 我们直接进行权限验证
			permitted, err := userPermit.CheckPermission(username, r.URL.Path, r.Method)
			if err != nil {
				log.Printf("权限验证失败, err: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			log.Printf("user(%s) permission 1st checking result, permit(%v)", username, permitted)

			// 进行鉴权验证
			// TODO
		} else {
			log.Printf("当前接口无需进行鉴权验证")
		}
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
