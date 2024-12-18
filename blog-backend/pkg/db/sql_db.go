package db

import (
	"blog-backend/pkg/config"
	"database/sql"
	"fmt"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var DefaultSQLDriverName = "mysql"

type SQLDB struct {
	db *gorm.DB
}

func NewSQLDB(c *config.SQLPara) (*SQLDB, error) {
	logger := log.WithFields(log.Fields{
		"module": "NewSQLDB",
	})
	// dsn := "root:12345@tcp(127.0.0.1:3306)/" // 请替换为你的数据库信息
	// dbName := "blog"
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
	// 连接到MySQL服务器
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Errorf("failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// 连接到具体的数据库

	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			DriverName: driverName,
			DSN:        dsn,
		}),
		&gorm.Config{
			QueryFields: true, //避免查询所有字段
		})
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
	// log.Infof("先检查当前用户是否在数据库里面存在: %s", username)
	var user User
	res := s.db.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// 根据用户ID获取用户信息
func (s *SQLDB) UserGetByID(userID int64) (*User, error) {
	// log.Infof("获取用户信息: %d", userID)
	var user User
	res := s.db.Where("id = ?", userID).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// 更新用户的信息
func (s *SQLDB) UserUpdate(userID int64, fields map[string]interface{}) error {
	// log.Infof("更新用户信息: %d", userID)
	res := s.db.Model(&User{}).Where("id = ?", userID).Updates(fields)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 注册用户
func (s *SQLDB) Register(user User) (int64, error) {

	res := s.db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

// 获取用户总数
func (s *SQLDB) UserCount() (int64, error) {
	// log.Infof("获取用户总数")
	var count int64
	res := s.db.Model(&User{}).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

// 获取用户列表
func (s *SQLDB) UserList(page, pageSize int64) ([]User, error) {
	// log.Infof("获取用户列表")
	// log.Infof("page: %d, pageSize: %d", page, pageSize)
	// 跳过的用户数
	offset := (page - 1) * pageSize
	var users []User
	res := s.db.Offset(int(offset)).Limit(int(pageSize)).Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

// 删除用户（管理员）
func (s *SQLDB) UserDelete(userID int64) error {
	log.Infof("删除用户: %d", userID)
	res := s.db.Where("id = ?", userID).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
