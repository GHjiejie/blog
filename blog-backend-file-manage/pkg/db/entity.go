package db

import "time"

type UploadFile struct {
	// 文件ID
	ID int64 `gorm:"column:id;primaryKey"  json:"id"`
	// 文件名
	FileName string `gorm:"column:file_name" json:"file_name"`
	// 文件大小
	FileSize int64 `gorm:"column:file_size" json:"file_size"`
	// 文件类型
	FileType string `gorm:"column:file_type" json:"file_type"`
	// 文件内容
	FileContent []byte `gorm:"column:file_content" json:"file_content"`
	// 文件标签
	Tag string `gorm:"column:tag" json:"tag"`
	// 文件上传时间
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	// 文件更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`
	// 文件上传者ID
	UploaderId int64 `gorm:"column:uploader_id" json:"uploader_id"`
}

func (u UploadFile) TableName() string {
	return "files"
}
