package db

import (
	"database/sql"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

type SQLDB struct {
	db *gorm.DB
}

func NEWSQLDB() (*SQLDB, error) {
	var err error
	dsn := "root:12345@tcp(127.0.0.1:3306)/" // 请替换为你的数据库信息
	dbName := "blog"

	// 连接到MySQL服务器
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// 检查数据库是否存在，如果不存在则创建
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Errorf("failed to create database: %v", err)
	}

	// 连接到具体的数据库
	dsn = fmt.Sprintf("root:12345@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to connect database: %v", err)
	} else {
		log.Infof("连接数据库成功")
	}

	// 迁移 schema
	if err := gormDB.AutoMigrate(&User{}); err != nil {
		log.Errorf("failed to migrate schema: %v", err)
	} else {
		log.Infof("迁移 schema 成功")
	}

	// 设置连接池
	if err := setupConnectionPool(gormDB); err != nil {
		log.Errorf("failed to setup connection pool: %v", err)
	}
	backendDB := &SQLDB{
		db: gormDB,
	}
	return backendDB, nil
}
func (s *SQLDB) GetORMDB() *gorm.DB {
	return s.db
}

func setupConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 12)
	return nil
}

// 根据用户名获取用户信息
func (s *SQLDB) UserGetByName(username string) (*User, error) {
	log.Infof("先检查当前用户是否在数据库里面存在: %s", username)
	var user User
	res := s.db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil

}

// 更新用户的信息
func (s *SQLDB) UserUpdate(userID int64, fields map[string]interface{}) error {
	log.Infof("更新用户信息: %d", userID)
	res := s.db.Model(&User{}).Where("id = ?", userID).Updates(fields)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 注册用户
func (s *SQLDB) Register(user User) (int64, error) {
	log.Infof("注册用户: %v", user)
	res := s.db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}
