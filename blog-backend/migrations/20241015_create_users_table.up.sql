-- 20231015_create_users_table.sql

CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `username` varchar(128) NOT NULL comment '用户登录名',
    `password` varchar(128) NOT NULL comment '用户密码',
    `role` tinyint DEFAULT '1' comment '角色',
    `token` varchar(1024) comment '用户token',
    `created_at` datetime NOT NULL DEFAULT NOW() comment '创建时间',
    `updated_at` datetime NOT NULL DEFAULT NOW() ON UPDATE NOW() comment '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_users_on_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- 其他索引、约束等可以在这里添加

-- 创建文件表
CREATE TABLE IF NOT EXISTS `files` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(128) NOT NULL comment '文件名',
    `path` varchar(1024) NOT NULL comment '文件路径',
    `size` int NOT NULL comment '文件大小',
    `type` varchar(128) NOT NULL comment '文件类型',
    `created_at` datetime NOT NULL DEFAULT NOW() comment '创建时间',
    `updated_at` datetime NOT NULL DEFAULT NOW() ON UPDATE NOW() comment '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_files_on_path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

