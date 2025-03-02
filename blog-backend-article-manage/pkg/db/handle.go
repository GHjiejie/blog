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
	// 获取文章列表(管理员)
	GetArticleList(page, pageSize int32) ([]Article, error)
	// 获取文章列表(用户)
	GetArticleListByAuthor(userId, page, pageSize int32) ([]Article, error)
	// 根据tag查询文章
	GetArticleListByKeyWord(tag string) ([]Article, error)
	// 更新文章
	UpdateArticle(articleInfo Article) error
	// 根据文章ID获取文章详情
	GetArticleDetail(articleId int64) (Article, error)
	// 更新文章评论数
	UpdateArticleCommentCount(articleId, updateCount int64) error

	/*
		api for web
	*/

	// 获取已发布的文章列表
	GetPublishedArticleList(page, pageSize int32) ([]Article, error)

	/*
		文章评论
	*/
	// 创建文章评论
	CreateArticleComment(articleId, userId int64, content string) (Comment, error)

	// 获取文章评论总数
	GetArticleCommentCount(articleId int64) (int64, error)

	// 获取文章评论列表
	GetArticleCommentList(articleId int64, page, pageSize int32) ([]Comment, error)

	// 更新评论点赞数
	UpdateCommentLikeCount(commentId int64) error

	// 更新点赞记录
	UpDateCommentLike(commentId, userId, articleId int64) error

	// 删除文章评论
	DeleteArticleComment(commentId int64) error

	// 删除点赞记录
	DeleteCommentLike(articleId, commentId int64) error

	//查询用户是否点赞
	GetArticleLike(articleId, userId int64) (ArticleLike, error)

	//创建点赞记录
	CreateArticleLike(articleId, userId int64) error

	// 更新文章点赞数
	UpdateArticleLikeCount(articleId, updateCount int64) error

	// 删除文章点赞记录
	DeleteArticleLike(articleId, userId int64) error

	// 更新文章浏览次数
	UpdateArticleViewCount(articleId, updateCount int64) error

	// 查询Tag是否存在
	GetTagByName(tagName string) error

	// 创建Tag
	CreateTag(tagName string, categoryId int64) error

	// 获取Tag列表
	GetTagList(page, pageSize int32) ([]Tag, error)

	// 获取Tag总数
	GetTagCount() (int64, error)

	// 根据id查询tag
	GetTagById(tagId int64) (Tag, error)

	// 根据tag_id删除tag
	DeleteTag(tagId int64) error
}

func NewHandler(c *config.DBConfig) (Handle, error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		return NewSQLDB(c.SQLPara)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.Type)
	}

}
