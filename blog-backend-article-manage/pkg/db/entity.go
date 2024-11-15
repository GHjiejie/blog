package db

import (
	articlepb "articleManage/pb/articleManage"
	"time"
)

// Comment 表示用户对文章的评论模型
type Comment struct {
	ID           int64      `gorm:"column:id;primaryKey" json:"id"`                              // 评论ID
	ArticleId    int64      `gorm:"column:article_id;index" json:"article_id"`                   // 关联的文章ID
	UserId       int64      `gorm:"column:user_id;index" json:"user_id"`                         // 评论者的用户ID
	Content      string     `gorm:"column:content;type:text" json:"content"`                     // 评论内容
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`          // 创建时间
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`          // 更新时间
	DeletedAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`         // 软删除时间
	ReviewStatus string     `gorm:"column:review_status;default:'pending'" json:"review_status"` // 评论审核状态
}

// 常量定义审核状态
const (
	ReviewPending  = "pending"  // 待审核
	ReviewApproved = "approved" // 已通过
	ReviewRejected = "rejected" // 已拒绝
)

func (c *Comment) TableName() string {
	return "comments"
}

// Article 结构体更新以包含评论
type Article struct {
	ID           int64            `gorm:"column:id;primaryKey" json:"id"`                      // 文章ID
	Title        string           `gorm:"column:title" json:"title"`                           // 文章标题
	Content      string           `gorm:"column:content" json:"content"`                       // 文章内容
	Summary      string           `gorm:"column:summary" json:"summary"`                       // 文章摘要
	Tag          string           `gorm:"column:tag" json:"tag"`                               // 文章标签
	Status       articlepb.Status `gorm:"column:status; default:3" json:"status"`              // 文章状态
	ViewCount    int              `gorm:"column:view_count;default:0" json:"view_count"`       // 文章浏览次数
	LikeCount    int              `gorm:"column:like_count;default:0" json:"like_count"`       // 文章点赞次数
	CommentCount int              `gorm:"column:comment_count;default:0" json:"comment_count"` // 文章评论次数
	ImageURL     string           `gorm:"column:image_url" json:"image_url"`                   // 新增：封面图 URL
	CategoryId   int64            `gorm:"column:category_id" json:"category_id"`               // 文章分类ID
	CreatedAt    time.Time        `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
	UpdatedAt    time.Time        `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
	AuthorId     int64            `gorm:"column:author_id" json:"author_id"`                   // 文章作者ID
	Comments     []Comment        `gorm:"foreignKey:ArticleId" json:"comments"`                // 新增：与评论的关系
	DeletedAt    *time.Time       `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
}

func (a *Article) TableName() string {
	return "articles"
}
