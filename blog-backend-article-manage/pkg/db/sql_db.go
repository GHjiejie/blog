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

	// 进行数据库的迁移操作
	if err := gormDB.AutoMigrate(&Article{}, &Comment{}, &CommentLike{}, &ArticleLike{}, &Tag{}); err != nil {
		logger.Errorf("failed to migrate schema: %v", err)
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

// 获取文章列表（管理员）
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

// 获取文章列表（用户）
func (s *SQLDB) GetArticleListByAuthor(userId, page, pageSize int32) ([]Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleListByAuthor",
	})
	var articleList []Article
	offset := (page - 1) * pageSize
	if err := s.db.Where("author_id = ?", userId).Offset(int(offset)).Limit(int(pageSize)).Find(&articleList).Error; err != nil {
		logger.Errorf("failed to get article list by author: %v", err)
		return nil, err
	}
	return articleList, nil
}

// 根据keyword查询文章
func (s *SQLDB) GetArticleListByKeyWord(tag string) ([]Article, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleListByTag",
	})
	var articleList []Article

	// 使用 OR 条件查询，tag 匹配或者 title 包含 keyword
	if err := s.db.Where("tag = ? OR title LIKE ?", tag, "%"+tag+"%").Find(&articleList).Error; err != nil {
		logger.Errorf("failed to get article list by tag or keyword: %v", err)
		return nil, err
	}

	return articleList, nil
}

// func (s *SQLDB) GetArticleListByTag(tag string) ([]Article, error) {
// 	logger := log.WithFields(log.Fields{
// 		"module": "GetArticleListByTag",
// 	})
// 	var articleList []Article
// 	if err := s.db.Where("tag = ?", tag).Find(&articleList).Error; err != nil {
// 		logger.Errorf("failed to get article list by tag: %v", err)
// 		return nil, err
// 	}
// 	return articleList, nil
// }

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

// 更新文章评论数
func (s *SQLDB) UpdateArticleCommentCount(articleId, updateCount int64) error {
	logger := log.WithFields(log.Fields{
		"module": "UpdateArticleCommentCount",
	})
	if err := s.db.Model(&Article{}).Where("id = ?", articleId).Update("comment_count", gorm.Expr("comment_count + ?", updateCount)).Error; err != nil {
		logger.Errorf("failed to update article comment count: %v", err)
		return err
	}
	return nil
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

// 创建文章评论
func (s *SQLDB) CreateArticleComment(articleId, userId int64, content string) (Comment, error) {
	logger := log.WithFields(log.Fields{
		"module": "CreateArticleComment",
	})
	comment := Comment{
		ArticleId: articleId,
		UserId:    userId,
		Content:   content,
	}
	if err := s.db.Create(&comment).Error; err != nil {
		logger.Errorf("failed to create article comment: %v", err)
		return Comment{}, err
	}
	return comment, nil
}

// 获取文章评论总数
func (s *SQLDB) GetArticleCommentCount(articleId int64) (int64, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleCommentCount",
	})
	var count int64
	if err := s.db.Model(&Comment{}).Where("article_id = ?", articleId).Count(&count).Error; err != nil {
		logger.Errorf("failed to get article comment count: %v", err)
		return 0, err
	}
	return count, nil
}

// 获取文章评论列表（按时间排列）
func (s *SQLDB) GetArticleCommentList(articleId int64, page, pageSize int32) ([]Comment, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleCommentList",
	})
	var commentList []Comment
	offset := (page - 1) * pageSize
	if err := s.db.Where("article_id = ?", articleId).Order("created_at DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&commentList).Error; err != nil {
		logger.Errorf("failed to get article comment list: %v", err)
		return nil, err
	}
	return commentList, nil
}

// 更新评论点赞数
func (s *SQLDB) UpdateCommentLikeCount(commentId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "UpdateCommentLikeCount",
	})
	if err := s.db.Model(&Comment{}).Where("id = ?", commentId).Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		logger.Errorf("failed to update comment like count: %v", err)
		return err
	}
	return nil
}

// 更新评论来源(comment_likes)表
func (s *SQLDB) UpDateCommentLike(commentId, userId, articleId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "UpDateCommentLike",
	})
	commentLike := CommentLike{
		CommentId: commentId,
		UserId:    userId,
		ArticleId: articleId,
	}
	if err := s.db.Create(&commentLike).Error; err != nil {
		logger.Errorf("failed to update comment like: %v", err)
		return err
	}
	return nil
}

// 删除文章评论
func (s *SQLDB) DeleteArticleComment(commentId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "DeleteArticleComment",
	})
	if err := s.db.Delete(&Comment{}, commentId).Error; err != nil {
		logger.Errorf("failed to delete article comment: %v", err)
		return err
	}
	return nil

}

// 删除点赞记录
func (s *SQLDB) DeleteCommentLike(articleId, commentId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "DeleteCommentLike",
	})
	// 根据articleId和commentId删除点赞记录
	if err := s.db.Where("article_id = ? AND comment_id = ?", articleId, commentId).Delete(&CommentLike{}).Error; err != nil {
		logger.Errorf("failed to delete comment like: %v", err)
		return err
	}
	return nil
}

// 查询用户是否点赞
func (s *SQLDB) GetArticleLike(articleId, userId int64) (ArticleLike, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetArticleLike",
	})
	var like ArticleLike

	if err := s.db.Where("article_id = ? AND user_id = ?", articleId, userId).First(&like).Error; err != nil {
		logger.Errorf("failed to get article like: %v", err)
		return ArticleLike{}, err
	}
	return like, nil
}

// 创建点赞记录
func (s *SQLDB) CreateArticleLike(articleId, userId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "CreateArticleLike",
	})
	like := ArticleLike{
		ArticleId: articleId,
		UserId:    userId,
	}
	if err := s.db.Create(&like).Error; err != nil {
		logger.Errorf("failed to create article like: %v", err)
		return err
	}
	return nil
}

// 更新文章点赞数（思考怎么做限流？？？）
func (s *SQLDB) UpdateArticleLikeCount(articleId, updateCount int64) error {
	logger := log.WithFields(log.Fields{
		"module": "UpdateArticleLikeCount",
	})
	if err := s.db.Model(&Article{}).Where("id = ?", articleId).Update("like_count", gorm.Expr("like_count + ?", updateCount)).Error; err != nil {
		logger.Errorf("failed to update article like count: %v", err)
		return err
	}
	return nil
}

// 删除文章点赞记录
func (s *SQLDB) DeleteArticleLike(articleId, userId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "DeleteArticleLike",
	})
	if err := s.db.Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&ArticleLike{}).Error; err != nil {
		logger.Errorf("failed to delete article like: %v", err)
		return err
	}
	return nil
}

// 更新文章浏览次数
func (s *SQLDB) UpdateArticleViewCount(articleId, updateCount int64) error {
	logger := log.WithFields(log.Fields{
		"module": "UpdateArticleViewCount",
	})
	if err := s.db.Model(&Article{}).Where("id = ?", articleId).Update("view_count", gorm.Expr("view_count + ?", updateCount)).Error; err != nil {
		logger.Errorf("failed to update article view count: %v", err)
		return err
	}
	return nil
}

// 查看是否已经存在该tag
func (s *SQLDB) GetTagByName(tagName string) error {
	logger := log.WithFields(log.Fields{
		"module": "GetTagByName",
	})
	var tag Tag
	if err := s.db.Where("name = ?", tagName).First(&tag).Error; err != nil {
		logger.Errorf("failed to get tag by name: %v", err)
		return err
	}
	return nil
}

// 创建Tag
func (s *SQLDB) CreateTag(tagName string, categoryId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "CreateTag",
	})
	tag := Tag{
		Name:       tagName,
		CategoryId: categoryId,
	}
	if err := s.db.Create(&tag).Error; err != nil {
		logger.Errorf("failed to create tag: %v", err)
		return err
	}
	return nil
}

// 获取tag列表
func (s *SQLDB) GetTagList(page, pageSize int32) ([]Tag, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetTagList",
	})
	var tagList []Tag
	offset := (page - 1) * pageSize
	if err := s.db.Offset(int(offset)).Limit(int(pageSize)).Find(&tagList).Error; err != nil {
		logger.Errorf("failed to get tag list: %v", err)
		return nil, err
	}
	return tagList, nil
}

// 获取tag总数
func (s *SQLDB) GetTagCount() (int64, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetTagCount",
	})
	var count int64
	if err := s.db.Model(&Tag{}).Count(&count).Error; err != nil {
		logger.Errorf("failed to get tag count: %v", err)
		return 0, err
	}
	return count, nil
}

// 根据id查询tag
func (s *SQLDB) GetTagById(tagId int64) (Tag, error) {
	logger := log.WithFields(log.Fields{
		"module": "GetTagById",
	})
	var tag Tag
	if err := s.db.First(&tag, tagId).Error; err != nil {
		logger.Errorf("failed to get tag by id: %v", err)
		return Tag{}, err
	}
	return tag, nil
}

// 根据tag_id删除tag
func (s *SQLDB) DeleteTag(tagId int64) error {
	logger := log.WithFields(log.Fields{
		"module": "DeleteTag",
	})
	if err := s.db.Delete(&Tag{}, tagId).Error; err != nil {
		logger.Errorf("failed to delete tag: %v", err)
		return err
	}
	return nil

}
