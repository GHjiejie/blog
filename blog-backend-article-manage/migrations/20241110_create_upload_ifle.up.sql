-- // Comment 表示用户对文章的评论模型
-- type Comment struct {
-- 	ID           int64      `gorm:"column:id;primaryKey" json:"id"`                              // 评论ID
-- 	ArticleId    int64      `gorm:"column:article_id;index" json:"article_id"`                   // 关联的文章ID
-- 	UserId       int64      `gorm:"column:user_id;index" json:"user_id"`                         // 评论者的用户ID
-- 	Content      string     `gorm:"column:content;type:text" json:"content"`                     // 评论内容
-- 	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`          // 创建时间
-- 	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`          // 更新时间
-- 	DeletedAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`         // 软删除时间
-- 	ReviewStatus string     `gorm:"column:review_status;default:'pending'" json:"review_status"` // 评论审核状态
-- }
-- // 常量定义审核状态
-- const (
-- 	ReviewPending  = "pending"  // 待审核
-- 	ReviewApproved = "approved" // 已通过
-- 	ReviewRejected = "rejected" // 已拒绝
-- )
-- // Article 结构体更新以包含评论
-- type Article struct {
-- 	ID           int64      `gorm:"column:id;primaryKey" json:"id"`                      // 文章ID
-- 	Title        string     `gorm:"column:title" json:"title"`                           // 文章标题
-- 	Content      string     `gorm:"column:content" json:"content"`                       // 文章内容
-- 	Summary      string     `gorm:"column:summary" json:"summary"`                       // 文章摘要
-- 	Tag          string     `gorm:"column:tag;index" json:"tag"`                               // 文章标签
-- 	Status       string     `gorm:"column:status;default:'draft'" json:"status"`         // 文章状态
-- 	ViewCount    int        `gorm:"column:view_count;default:0" json:"view_count"`       // 文章浏览次数
-- 	LikeCount    int        `gorm:"column:like_count;default:0" json:"like_count"`       // 文章点赞次数
-- 	CommentCount int        `gorm:"column:comment_count;default:0" json:"comment_count"` // 文章评论次数
-- 	ImageURL     string     `gorm:"column:image_url" json:"image_url"`                   // 新增：封面图 URL
-- 	CategoryId   int64      `gorm:"column:category_id" json:"category_id"`               // 文章分类ID
-- 	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
-- 	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
-- 	AuthorId     int64      `gorm:"column:author_id" json:"author_id"`                   // 文章作者ID
-- 	Comments     []Comment  `gorm:"foreignKey:ArticleId" json:"comments"`                // 新增：与评论的关系
-- 	DeletedAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
-- }
-- 创建article表
CREATE TABLE IF NOT EXISTS `articles` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `content` TEXT NOT NULL,
    `summary` VARCHAR(255) NOT NULL,
    `tag` VARCHAR(255) NOT NULL,
    `status` TINYINT NOT NULL DEFAULT 3,
    `view_count` INT(11) NOT NULL DEFAULT 0,
    `like_count` INT(11) NOT NULL DEFAULT 0,
    `comment_count` INT(11) NOT NULL DEFAULT 0,
    `image_url` VARCHAR(255) DEFAULT NULL,
    -- 允许为 NULL
    `category_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `author_id` BIGINT(20) NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_article_tag` (`tag`),
    INDEX `idx_article_category_id` (`category_id`),
    INDEX `idx_article_author_id` (`author_id`),
    INDEX `idx_article_deleted_at` (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
-- 创建文章点赞记录表
CREATE TABLE IF NOT EXISTS `article_likes` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `article_id` BIGINT(20) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_article_like_article_id` (`article_id`),
    INDEX `idx_article_like_user_id` (`user_id`),
    CONSTRAINT `fk_article_like_article` FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_article_like_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
-- 创建comment表
CREATE TABLE IF NOT EXISTS `comments` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    -- 评论ID，主键，自增
    `article_id` BIGINT(20) NOT NULL,
    -- 关联的文章ID
    `user_id` BIGINT(20) NOT NULL,
    -- 发表评论的用户ID
    `content` TEXT NOT NULL,
    -- 评论内容
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 创建时间，默认为当前时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 更新时间，自动更新为当前时间
    `deleted_at` DATETIME DEFAULT NULL,
    -- 软删除时间（用于标记删除）
    `review_status` TINYINT NOT NULL DEFAULT 0,
    -- 审核状态，默认为0
    `like_count` INT(11) NOT NULL DEFAULT 0,
    -- 新增：评论点赞数量，默认为0
    PRIMARY KEY (`id`),
    -- 主键约束
    INDEX `idx_comment_article_id` (`article_id`),
    -- 针对 article_id 的索引
    INDEX `idx_comment_user_id` (`user_id`),
    -- 针对 user_id 的索引
    INDEX `idx_comment_deleted_at` (`deleted_at`),
    -- 针对 deleted_at 的索引
    CONSTRAINT `fk_comment_article` FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    -- 外键约束，关联 articles 表
    CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE -- 外键约束，关联 users 表
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
-- 创建comment_like表
CREATE TABLE IF NOT EXISTS `comment_likes` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    -- 点赞ID，主键，自增
    `article_id` BIGINT(20) NOT NULL,
    -- 关联的文章ID
    `comment_id` BIGINT(20) NOT NULL,
    -- 关联的评论ID
    `user_id` BIGINT(20) NOT NULL,
    -- 点赞的用户ID
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- 创建时间，默认为当前时间
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- 更新时间，自动更新为当前时间
    PRIMARY KEY (`id`),
    -- 主键约束
    INDEX `idx_comment_like_comment_id` (`comment_id`),
    -- 针对 comment_id 的索引
    INDEX `idx_comment_like_user_id` (`user_id`),
    -- 针对 user_id 的索引
    CONSTRAINT `fk_comment_like_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    -- 外键约束，关联 comments 表
    CONSTRAINT `fk_comment_like_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE -- 外键约束，关联 users 表
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 创建tag表
CREATE TABLE IF NOT EXISTS `tags` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `category_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_tag_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;