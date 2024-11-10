package db

import (
	"fileManage/pkg/config"
	"fmt"
	"strings"
)

type Handle interface {
	// 上传文件
	UploadFile(fileInfo UploadFile) (UploadFile, error)
}

func NewHandler(c *config.DBConfig) (Handle, error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		return NewSQLDB(c.SQLPara)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.Type)
	}

}
