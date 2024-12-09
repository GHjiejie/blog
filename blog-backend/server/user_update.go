package server

import (
	"blog-backend/pkg/db"
	"io"
	"net/http"
	"strconv"

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
		"username": userInfo.Username,
		"email":    userInfo.Email,
		"phone":    userInfo.Phone,
		"avatar":   userInfo.Avatar,
	}); err != nil {
		logger.Errorf("Failed to update user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	


}
