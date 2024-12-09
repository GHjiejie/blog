package server

import (
	"blog-backend/pkg/db"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func (s *BlogServer) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{
		"api": "updateUser",
	})
	// 解析表单数据
	err := r.ParseMultipartForm(10 << 20) //限制最大上传10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	avatar, _, err := r.FormFile("avatar") // 获取上传的文件
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	defer avatar.Close()

	avatarBytes, err := io.ReadAll(avatar)
	if err != nil {
		logger.Errorf("Failed to read file: %v", err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	// 获取formdata里面的username
	username := r.FormValue("username")
	// 获取formdata里面的password
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	userId := r.FormValue("userId")

	userInfo := db.User{
		Username: username,
		Email:    email,
		Phone:    phone,
		Avatar:   avatarBytes,
	}
	// 然后更新用户的个人消息
	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.Errorf("Failed to convert userId to int64: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	if err := s.DBEngine.UserUpdate(userIdInt64, map[string]interface{}{
		"username":   userInfo.Username,
		"email":      userInfo.Email,
		"phone":      userInfo.Phone,
		"avatar":     userInfo.Avatar,
		"updated_at": time.Now(),
	}); err != nil {
		logger.Errorf("Failed to update user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// 根据用户id查询用户信息（最后我们要返回用户的最新消息）
	user, err := s.DBEngine.UserGetByID(userIdInt64)
	if err != nil {
		logger.Errorf("Failed to get user: %v", err)
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	// 创建一个新的user对象，用于返回给前端
	userInfo = db.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
	}

	// 返回用户信息
	response := map[string]interface{}{
		"username": userInfo.Username,
		"email":    userInfo.Email,
		"phone":    userInfo.Phone,
		"avatar":   userInfo.Avatar,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回JSON响应
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Errorf("Failed to create response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Errorf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
