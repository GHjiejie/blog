package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	db "fileManage/pkg/db"

	log "github.com/sirupsen/logrus"
)

// 处理文件上传
func (s *FileServer) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{
		"api": "uploadFile",
	})

	// 解析表单数据
	err := r.ParseMultipartForm(10 << 20) //限制最大上传10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file") // 获取上传的文件
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	logger.Infof("文件的内容是：%s", fileBytes)
	logger.Infof("文件内容的十六进制表示：%x", fileBytes)
	if err != nil {
		logger.Errorf("Failed to read file: %v", err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	tag := r.FormValue("tag")                  // 获取文件标签
	uploaderIdStr := r.FormValue("uploaderId") // 获取上传者ID

	// 转换 uploaderIdStr 为 int64
	uploaderId, err := strconv.ParseInt(uploaderIdStr, 10, 64)
	if err != nil {
		// 处理转换错误，例如返回400 Bad Request
		logger.Errorf("Invalid uploader ID: %v", err)
		http.Error(w, "Invalid uploader ID", http.StatusBadRequest)
		return
	}
	fileInfo := db.UploadFile{
		FileName:    fileHeader.Filename,
		FileSize:    fileHeader.Size,
		FileType:    fileHeader.Header.Get("Content-Type"),
		Tag:         tag,
		FileContent: fileBytes,
		UploaderId:  uploaderId,
	}
	// 接下来将数据写入数据库
	uploadedFile, err := s.DBEngine.UploadFile(fileInfo)
	if err != nil {
		logger.Errorf("Failed to upload file: %v", err)
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	// 自定义响应，返回上传成功的文件信息
	response := map[string]interface{}{
		"fileName": uploadedFile.FileName,
		"bytes":    uploadedFile.FileSize,
		"tag":      tag,
		"message":  "File uploaded successfully",
	}

	// 设置响应头为 JSON 格式
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// 返回 JSON 响应
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
