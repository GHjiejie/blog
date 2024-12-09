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
	err := r.ParseMultipartForm(10 << 20) // 限制最大上传10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// 获取formdata里面的username
	username := r.FormValue("username")
	// 获取formdata里面的email
	email := r.FormValue("email")
	// 获取formdata里面的phone
	phone := r.FormValue("phone")
	// 获取formdata里面的userId
	userId := r.FormValue("userId")

	userIdInt64, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.Errorf("Failed to convert userId to int64: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// 创建一个map来存储要更新的字段
	updateFields := map[string]interface{}{
		"username":   username,
		"email":      email,
		"phone":      phone,
		"updated_at": time.Now(),
	}

	// 获取上传的文件
	avatar, _, err := r.FormFile("avatar")
	if err == nil {
		defer avatar.Close()
		avatarBytes, err := io.ReadAll(avatar)
		if err != nil {
			logger.Errorf("Failed to read file: %v", err)
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
		updateFields["avatar"] = avatarBytes
	}

	// 更新用户的个人信息
	if err := s.DBEngine.UserUpdate(userIdInt64, updateFields); err != nil {
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
	userInfo := db.User{
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
