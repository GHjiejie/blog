package server

import (
	"log"
	"net/http"
)

// 设置无需验证的路径
var NotNeedAuthorizationPaths = map[string]bool{
	"/v1/users/login":    true,
	"/v1/users/register": true,
	"/v1/captchas":       true,
}

func (s *BlogServer) Tracing(nextHandle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 判断当前接口是否需要进行鉴权验证
		if needAuthorizations(r.URL.Path, r.Method) {
			log.Printf("当前接口需要进行鉴权验证")
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
