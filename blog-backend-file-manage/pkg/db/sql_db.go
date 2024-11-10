package db

import (
	"fileManage/pkg/config"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultSQLDriverName = "mysql" //设置默认的数据库驱动
)

// 设置数据库实例
type SQLDB struct {
	db *gorm.DB
}

func setupDbConnPool(gormDB *gorm.DB) error {
	// 设置连接池
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(2000)
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetConnMaxLifetime(time.Hour * 12)

	// 校验连接
	if err = sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}

func NewSQLDB(c *config.SQLPara) (Handle, error) {
	logger := log.WithFields(log.Fields{
		"module": "NewSQLDB",
	})
	dsn := fmt.Sprintf("%s@%s(%s)/%s?%s",
		c.User,
		c.NetMode,
		c.Addr,
		c.DBName,
		c.Para,
	)
	logger.Infof("DSN: %s", dsn)
	driverName := c.DriverName
	if driverName == "" {
		driverName = DefaultSQLDriverName
	}
	// 进行数据库了解操作
	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			DriverName: driverName,
			DSN:        dsn,
		}),
		&gorm.Config{
			QueryFields: true, //避免查询所有字段
		},
	)
	if err != nil {
		logger.Errorf("failed to open database: %v", err)
		return nil, err
	}
	// 设置数据库连接池
	if err := setupDbConnPool(gormDB); err != nil {
		logger.Errorf("failed to setup db connection pool: %v", err)
		return nil, err
	}

	log.Debug("database connection established") //数据库连接建立成功

	fileBackend := &SQLDB{
		db: gormDB,
	}
	return fileBackend, nil
}

// 上传文件
func (s *SQLDB) UploadFile(fileInfo UploadFile) (UploadFile, error) {
	log.Info("现在进行文件上传操作")
	// 创建文件
	if err := s.db.Create(&fileInfo).Error; err != nil {
		log.Errorf("failed to create file: %v", err)
		return UploadFile{}, err
	}
	return fileInfo, nil
}
