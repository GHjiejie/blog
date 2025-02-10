package db

import (
	articlepb "articleManage/pb/articleManage"
	"time"
)

// Article 结构体更新以包含评论
type Article struct {
	ID           int64                   `gorm:"column:id;primaryKey" json:"id"`                      // 文章ID
	Title        string                  `gorm:"column:title" json:"title"`                           // 文章标题
	Content      string                  `gorm:"column:content" json:"content"`                       // 文章内容
	Summary      string                  `gorm:"column:summary" json:"summary"`                       // 文章摘要
	Tag          string                  `gorm:"column:tag" json:"tag"`                               // 文章标签
	Status       articlepb.ArticleStatus `gorm:"column:status; default:3" json:"status"`              // 文章状态
	ViewCount    int                     `gorm:"column:view_count;default:0" json:"view_count"`       // 文章浏览次数
	LikeCount    int                     `gorm:"column:like_count;default:0" json:"like_count"`       // 文章点赞次数
	CommentCount int                     `gorm:"column:comment_count;default:0" json:"comment_count"` // 文章评论次数
	ImageURL     string                  `gorm:"column:image_url" json:"image_url"`                   // 新增：封面图 URL
	CategoryId   int64                   `gorm:"column:category_id" json:"category_id"`               // 文章分类ID
	CreatedAt    time.Time               `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
	UpdatedAt    time.Time               `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
	AuthorId     int64                   `gorm:"column:author_id" json:"author_id"`                   // 文章作者ID
	Comments     []Comment               `gorm:"foreignKey:ArticleId" json:"comments"`                // 新增：与评论的关系
	DeletedAt    *time.Time              `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
}

func (a *Article) TableName() string {
	return "articles"
}

// 创建文章点赞记录模型
type ArticleLike struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`                     // 点赞ID
	ArticleId int64     `gorm:"column:article_id;index" json:"article_id"`          // 关联的文章ID
	UserId    int64     `gorm:"column:user_id;index" json:"user_id"`                // 点赞者的用户ID
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
}

func (l *ArticleLike) TableName() string {
	return "article_likes"
}

// Comment 表示用户对文章的评论模型
type Comment struct {
	ID           int64                   `gorm:"column:id;primaryKey" json:"id"`                      // 评论ID
	ArticleId    int64                   `gorm:"column:article_id;index" json:"article_id"`           // 关联的文章ID
	UserId       int64                   `gorm:"column:user_id;index" json:"user_id"`                 // 评论者的用户ID
	Content      string                  `gorm:"column:content;type:text" json:"content"`             // 评论内容
	LikeCount    int                     `gorm:"column:like_count;default:0" json:"like_count"`       // 评论点赞次数
	CreatedAt    time.Time               `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
	UpdatedAt    time.Time               `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
	DeletedAt    *time.Time              `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
	ReviewStatus articlepb.CommentStatus `gorm:"column:review_status;default:0" json:"review_status"` // 评论审核状态
}

func (c *Comment) TableName() string {
	return "comments"
}

// likes表示用户对文章的点赞模型
type CommentLike struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`                     // 点赞ID
	ArticleId int64     `gorm:"column:article_id;index" json:"article_id"`          // 关联的文章ID
	UserId    int64     `gorm:"column:user_id;index" json:"user_id"`                // 点赞者的用户ID
	CommentId int64     `gorm:"column:comment_id;index" json:"comment_id"`          // 关联的评论ID
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
}

func (l *CommentLike) TableName() string {
	return "comment_likes"
}

// 创建tag模型
type Tag struct {
	ID         int64      `gorm:"column:id;primaryKey" json:"id"`                      // 标签ID
	Name       string     `gorm:"column:name;index" json:"name"`                       // 标签名称
	CategoryId int64      `gorm:"column:category_id;index" json:"category_id"`         // 分类ID
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
	UpdatedAt  time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
	DeleteAt   *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
}

func (t *Tag) TableName() string {
	return "tags"
}
