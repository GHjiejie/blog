
CREATE TABLE IF NOT EXISTS `casbin_rules` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) comment '策略类型',
    `v0` varchar(100) comment '名称',
    `v1` varchar(100) comment '路由路径',
    `v2` varchar(100) comment 'method GET POST ...',
    `v3` varchar(100),
    `v4` varchar(100),
    `v5` varchar(100),
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rules` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
     
        ('p', 'admin-policy', '/v1/*', '*');       
        


INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
        ('p', 'files-manage-all', '/v1/files/*', '*'),   
        ('p', 'articles-manage-all','/v1/articles/*', '*'), 
        ('p', 'web-all','/v1/web/*','*');     


INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
  
        ('p','users-register','/v1/users/register','POST'),
        ('p','users-login','/v1/users/login','POST'),
        ('p','users-logout','/v1/users/logout','POST'),
        ('p','users-list','/v1/users/*','GET'),
        ('p','users-delete', '/v1/users/*', 'DELETE'),
        ('p','users-password-reset', '/v1/users/change_password', 'PUT'),
        ('p','users-password-update', '/v1/users/update_password', 'POST'),
        ('p','users-update', '/v1/users/*', 'PUT'),
        ('p','users-get', '/v1/users/*', 'GET'),

    
        ('p','files-upload','/v1/files/upload','POST'),
        ('p','files-list','/v1/files/getFileList','GET'),
        ('p','files-delete','/v1/files/deleteFile/*','DELETE'),
        ('p','files-get', '/v1/files/queryFileById/*', 'GET'),
   
        ('p','articles-create','/v1/articles/publishArticle','POST'),
        ('p','articles-list','/v1/articles/getArticleList','GET'),
        ('p','articles-delete','/v1/articles/deleteArticle/*','DELETE'),
        ('p','articles-update','/v1/articles/updateArticle/*','PUT'),
        ('p','articles-get','/v1/articles/getArticleDetail/*','GET'),  
        ('p','article-review','/v1/articles/reviewArticle','POST'),
        ('p','article-query', '/v1/articles/queryArticle', 'GET'),
        ('p','article-like', '/v1/articles/likeArticle', 'POST'),
        ('p','article-dislike', '/v1/articles/cancelLikeArticle', 'POST'),
  
        ('p','comment-create','/v1/articles/publishArticleComment','POST'),
        ('p','comment-list','/v1/articles/getArticleCommentList','GET'),
        ('p','comment-delete','/v1/articles/deleteArticleComment/*','DELETE'),
        ('p','comment-like','/v1/articles/likeArticleComment','POST'),
 
        ('p', 'articles-web-list', '/v1/web/articles/getPublishedArticleList', 'POST');
       


INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 

        ('g', 'users-basic-policy', 'users-register', ''),
        ('g', 'users-basic-policy', 'users-login', ''),
        ('g', 'users-basic-policy', 'users-logout', ''),
        ('g', 'users-basic-policy', 'users-password-update', ''),
        ('g', 'users-basic-policy', 'users-update', ''),
        ('g', 'users-basic-policy', 'users-get', ''),
   
        ('g', 'engine-all', 'files-manage-all', ''), 
        ('g', 'engine-all', 'articles-manage-all', ''), 
        ('g', 'engine-all', 'web-all', ''), 
  
        ('g', 'user-normal-policy', 'user-basic-policy', ''), 
        ('g', 'user-normal-policy', 'files-upload', ''),
        ('g', 'user-normal-policy', 'files-list', ''),
        ('g', 'user-normal-policy', 'files-delete', ''),
        ('g', 'user-normal-policy', 'files-get', ''),
        ('g', 'user-normal-policy', 'articles-create', ''),
        ('g', 'user-normal-policy', 'articles-list', ''),
        ('g', 'user-normal-policy', 'articles-delete', ''),
        ('g', 'user-normal-policy', 'articles-update', ''),
        ('g', 'user-normal-policy', 'articles-get', ''),
        ('g', 'user-normal-policy', 'article-query', ''),
        ('g', 'user-normal-policy', 'comment-create', ''),
        ('g', 'user-normal-policy', 'comment-list', ''),
        ('g', 'user-normal-policy', 'comment-delete', ''),
        ('g', 'user-normal-policy', 'comment-like', ''),
        ('g', 'user-normal-policy', 'comment-query', ''),
        ('g', 'user-normal-policy', 'articles-web-list', ''),
        ('g', 'user-normal-policy', 'article-like', ''),
        ('g', 'user-normal-policy', 'article-dislike',''),

    
        ('g', 'user-admin-policy', 'admin-policy', ''),
        ('g', 'admin', 'user-admin-policy', '');
        




     

         

        
