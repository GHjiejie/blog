package db

import (
	"fileManage/pkg/config"
	"fmt"
	"strings"
)

type Handle interface {
	// 上传文件
	UploadFile(fileInfo UploadFile) (UploadFile, error)
	// 根据文件ID获取文件信息
	GetFileByID(fileID int64) (UploadFile, error)
	// 删除文件
	DeleteFile(fileID int64) error
	// 获取文件列表
	GetFileList(page, pageSize int64) ([]UploadFile, error)
	// 获取文件总数
	GetFileTotal() (int64, error)
}

func NewHandler(c *config.DBConfig) (Handle, error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		return NewSQLDB(c.SQLPara)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.Type)
	}

}
