package db

import (
	"articleManage/pkg/config"
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

// 创建文章
func (s *SQLDB) CreateArticle(articleInfo Article) (Article, error) {

	logger := log.WithFields(log.Fields{
		"module": "CreateArticle",
	})
	logger.Infof("输出传递的文章消息: %v", articleInfo)
	// 创建文章
	if err := s.db.Create(&articleInfo).Error; err != nil {
		logger.Errorf("failed to create article: %v", err)
		return Article{}, err
	}
	return articleInfo, nil

}

// 获取文章的总数
func (s *SQLDB) GetArticleCount() (int64, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleCount",
	})
	var count int64
	if err := s.db.Model(&Article{}).Count(&count).Error; err != nil {
		logger.Errorf("failed to get article count: %v", err)
		return 0, err
	}
	return count, nil
}

// 删除文章
func (s *SQLDB) DeleteArticle(articleId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "DeleteArticle",
	})
	if err := s.db.Delete(&Article{}, articleId).Error; err != nil {
		logger.Errorf("failed to delete article: %v", err)
		return err
	}
	return nil
}

// 获取文章列表
func (s *SQLDB) GetArticleList(page, pageSize int32) ([]Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleList",
	})
	var articleList []Article
	offset := (page - 1) * pageSize
	if err := s.db.Offset(int(offset)).Limit(int(pageSize)).Find(&articleList).Error; err != nil {
		logger.Errorf("failed to get article list: %v", err)
		return nil, err
	}
	return articleList, nil
}

// 根据tag查询文章
func (s *SQLDB) GetArticleListByTag(tag string) ([]Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleListByTag",
	})
	var articleList []Article
	if err := s.db.Where("tag = ?", tag).Find(&articleList).Error; err != nil {
		logger.Errorf("failed to get article list by tag: %v", err)
		return nil, err
	}
	return articleList, nil
}

// 更新文章
func (s *SQLDB) UpdateArticle(articleInfo Article) error {
	logger := log.WithFields(log.Fields{
		"module": "UpdateArticle",
	})
	if err := s.db.Save(&articleInfo).Error; err != nil {
		logger.Errorf("failed to update article: %v", err)
		return err
	}
	return nil
}

// 根据文章ID获取文章详情
func (s *SQLDB) GetArticleDetail(articleId int64) (Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleDetail",
	})
	var article Article
	if err := s.db.First(&article, articleId).Error; err != nil {
		logger.Errorf("failed to get article detail: %v", err)
		return Article{}, err
	}
	return article, nil
}

/*
api for web
*/

// 获取已发布的文章列表
func (s *SQLDB) GetPublishedArticleList(page, pageSize int32) ([]Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetPublishedArticleList",
	})
	var articleList []Article
	offset := (page - 1) * pageSize
	if err := s.db.Where("status = ?", 1).Offset(int(offset)).Limit(int(pageSize)).Find(&articleList).Error; err != nil {
		logger.Errorf("failed to get published article list: %v", err)
		return nil, err
	}
	return articleList, nil
}
