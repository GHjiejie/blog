package db

import (
	"articleManage/pkg/config"
	"fmt"
	"strings"
)

type Handle interface {
	// 创建文章
	CreateArticle(articleInfo Article) (Article, error)
	// 获取文章的总数
	GetArticleCount() (int64, error)
	// 删除文章
	DeleteArticle(articleId int64) error
	// 获取文章列表
	GetArticleList(page, pageSize int32) ([]Article, error)
	// 根据tag查询文章
	GetArticleListByTag(tag string) ([]Article, error)
	// 更新文章
	UpdateArticle(articleInfo Article) error
	// 根据文章ID获取文章详情
	GetArticleDetail(articleId int64) (Article, error)

	/*
		api for web
	*/

	// 获取已发布的文章列表
	GetPublishedArticleList(page, pageSize int32) ([]Article, error)
}

func NewHandler(c *config.DBConfig) (Handle, error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		return NewSQLDB(c.SQLPara)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.Type)
	}

}
