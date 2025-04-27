-- Active: 1745598180270@@127.0.0.1@3306@blog
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
    `category_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `author_id` BIGINT(20) NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_article_tag` (`tag`(30)),
    INDEX `idx_article_category_id` (`category_id`),
    INDEX `idx_article_author_id` (`author_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 将status默认值改为1，表示已发布
ALTER TABLE `articles` 
ALTER COLUMN `status` SET DEFAULT 1;



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

CREATE TABLE IF NOT EXISTS `comments` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `article_id` BIGINT(20) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `content` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `review_status` TINYINT NOT NULL DEFAULT 0,
    `like_count` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_comment_article_id` (`article_id`),
    INDEX `idx_comment_user_id` (`user_id`),
    FULLTEXT INDEX `idx_comment_content` (`content`(50)),
    CONSTRAINT `fk_comment_article` FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `comment_likes` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `article_id` BIGINT(20) NOT NULL,
    `comment_id` BIGINT(20) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_comment_like_comment_id` (`comment_id`),
    INDEX `idx_comment_like_user_id` (`user_id`),
    CONSTRAINT `fk_comment_like_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_comment_like_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;


CREATE TABLE IF NOT EXISTS `tags` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `category_id` BIGINT(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_tag_name` (`name`(30))
    -- UNIQUE INDEX `idx_tag_name` (`name`)  Using both causes issues, and prefix lengths not supported on unique indexes
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- type Category struct {
-- 	ID          int64      `gorm:"column:id;primaryKey" json:"id"`                      // 分类ID
-- 	Name        string     `gorm:"column:name;index" json:"name"`                       // 分类名称
-- 	Description string     `gorm:"column:description" json:"description"`               // 分类描述
-- 	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`  // 创建时间
-- 	UpdatedAt   time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`  // 更新时间
-- 	DeleteAt    *time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"` // 软删除时间
-- 	Articles    []Article  `gorm:"foreignKey:CategoryId" json:"articles"`               // 关联的文章
-- 	Tags        []Tag      `gorm:"foreignKey:CategoryId" json:"tags"`                   // 关联的标签
-- }

CREATE TABLE IF NOT EXISTS `categories` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类ID',
    `name` VARCHAR(255) NOT NULL COMMENT '分类名称',
    `description` VARCHAR(255) DEFAULT NULL COMMENT '分类描述',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    INDEX `idx_category_name` (`name`) COMMENT '分类名称索引',
    INDEX `idx_category_deleted_at` (`deleted_at`) COMMENT '软删除索引'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='分类表';

