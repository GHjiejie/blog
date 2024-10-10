package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:12345@tcp(127.0.0.1:3306)/" // 请替换为你的数据库信息
	dbName := "blog"

	// 连接到MySQL服务器
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// 检查数据库是否存在，如果不存在则创建
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	// 连接到具体的数据库
	dsn = fmt.Sprintf("root:12345@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	} else {
		log.Println("连接数据库成功")
	}

	// 设置连接池
	if err := setupConnectionPool(DB); err != nil {
		log.Fatalf("failed to setup connection pool: %v", err)
	}
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
