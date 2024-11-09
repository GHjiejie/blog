package server

import (
	"blog-backend/pkg/casbinpermit"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/any"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// ErrorBody 用于格式化 API 返回的错误信息，包含错误、消息、状态码及详细信息。这使得客户端能够更好地理解发生的错误。
type ErrorBody struct {
	Error   string     `protobuf:"bytes,1,name=error" json:"error"`
	Message string     `protobuf:"bytes,1,name=message" json:"message"`
	Code    int32      `protobuf:"varint,2,name=code" json:"code"`
	Details []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

func (s *BlogServer) Tracing(nextHandle http.Handler, userPermit *casbinpermit.Permit) http.Handler {
	logger := log.WithFields(log.Fields{
		"module": "Tracing",
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求体中获取用户信息
		// username := r.Header.Get("username")
		// 判断当前接口是否需要进行鉴权验证
		if needAuthorizations(r.URL.Path, r.Method) {
			var (
				username  string
				userToken string
			)
			logger.Infof("当前接口需要进行鉴权验证,同时需要进行权限检查")
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
				logger.Error("token error: ", err.Error())
				s, ok := status.FromError(err)
				if !ok {
					s = status.New(codes.Unknown, err.Error())
				}
				body := &ErrorBody{
					Error:   s.Err().Error(),
					Message: s.Message(),
					Code:    int32(s.Code()),
					Details: s.Proto().GetDetails(),
				}
				msg, err := json.Marshal(body)
				if err != nil {
					logger.Error("marshal error body: ", err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, err = w.Write(msg)
				if err != nil {
					logger.Error("failed to write error msg: ", err.Error())
				}
				return
			}

			username = claims.Username
			logger.Infof("token解析成功, 输出获取的claims: %v", claims)

			// 我们直接进行权限验证
			permitted, err := userPermit.CheckPermission(username, r.URL.Path, r.Method)
			if err != nil {
				log.Infof("user(%s) permission 1st checking result, permit(%v)", username, permitted)
				newerr := status.New(codes.Unknown, err.Error())
				// 将错误信息返回给前端
				body := &ErrorBody{
					Error:   newerr.Message(),
					Message: newerr.Message(),
					Code:    int32(newerr.Code()),
					Details: newerr.Proto().GetDetails(),
				}
				msg, err := json.Marshal(body)
				if err != nil {
					log.Errorf("failed to marshal body with err(%s)", err.Error())

				}
				w.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderVal)
				w.WriteHeader(http.StatusInternalServerError)
				_, err = w.Write(msg)
				if err != nil {
					logger.Errorf("failed to write msg with err(%s)", err.Error())
				}
				return
			}
			if permitted {
				logger.Infof("user(%s) permission 1st checking result, permit(%v)", username, permitted)
				nextHandle.ServeHTTP(w, r)
			} else {
				if err := userPermit.Enforcer.LoadPolicy(); err != nil {
					isOk, _ := userPermit.CheckPermission(username, r.URL.Path, r.Method)
					logger.Debugf("user(%s) permission 2nd checking result, permit(%v)", username, isOk)
					if isOk {
						nextHandle.ServeHTTP(w, r)
					}
				}
				logger.Errorf("%s does not have permission to %s %s", username, r.Method, r.URL.Path)
				newerr := status.New(
					codes.PermissionDenied,
					fmt.Sprintf("%s does not have permission to %s %s", username, r.Method, r.URL.Path),
				)
				body := ErrorBody{
					Error:   newerr.Message(),
					Message: newerr.Message(),
					Code:    int32(newerr.Code()),
					Details: newerr.Proto().GetDetails(),
				}
				msg, err := json.Marshal(body)
				if err != nil {
					logger.Errorf("marshal error body failed with err(%s)", err.Error())
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				_, err = w.Write(msg)
				if err != nil {
					logger.Errorf("failed to write msg with err(%s)", err.Error())
				}
				return
			}
			// nextHandle.ServeHTTP(w, r)

			// 如果首次校验失败,则重新从数据库加载策略,然后再次进行校验
			// if !permitted {
			// 	if err := userPermit.Enforcer.LoadPolicy(); err != nil {
			// 		isOk, _ := userPermit.CheckPermission(username, r.URL.Path, r.Method)
			// 		logger.Debugf("user(%s) permission 2nd checking result, permit(%v)", username, isOk)
			// 		if isOk {
			// 			nextHandle.ServeHTTP(w, r)
			// 		}
			// 	}
			// 	logger.Errorf("%s does not have permission to %s %s", username, r.Method, r.URL.Path)
			// 	newerr := status.New(
			// 		codes.PermissionDenied,
			// 		fmt.Sprintf("%s does not have permission to %s %s", username, r.Method, r.URL.Path),
			// 	)
			// 	body := ErrorBody{
			// 		Error:   newerr.Message(),
			// 		Message: newerr.Message(),
			// 		Code:    int32(newerr.Code()),
			// 		Details: newerr.Proto().GetDetails(),
			// 	}
			// 	msg, err := json.Marshal(body)
			// 	if err != nil {
			// 		logger.Errorf("marshal error body failed with err(%s)", err.Error())
			// 	}
			// 	w.Header().Set("Content-Type", "application/json")
			// 	w.WriteHeader(http.StatusForbidden)
			// 	_, err = w.Write(msg)
			// 	if err != nil {
			// 		logger.Errorf("failed to write msg with err(%s)", err.Error())
			// 	}
			// 	return
			// }
		} else {
			log.Info("当前接口无需进行鉴权验证")
			nextHandle.ServeHTTP(w, r)
		}
	})
}

func needAuthorizations(urlPath, method string) bool {
	log.Info("判断当前接口是否需要进行鉴权验证")
	log.Infof("urlPath: %s, method: %s", urlPath, method)
	need := true
	for path := range NotNeedAuthorizationPaths {
		if path == urlPath {
			need = false
			break
		}
	}
	return need
}
