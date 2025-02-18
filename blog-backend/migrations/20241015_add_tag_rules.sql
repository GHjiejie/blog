INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
        ('p','articles-tag-list','/v1/articles/getArticleTagList','GET'),
        ('p','articles-tag-create','/v1/articles/addArticleTag','POST'),
        ('p','articles-tag-delete','/v1/articles/deleteArticleTag/*','DELETE');

-- 设置用户组权限
INSERT INTO `casbin_rules`(ptype,v0,v1,v2) 
    VALUES 
        ('g', 'user-normal-policy', 'articles-tag-list',''),
        ('g', 'user-normal-policy', 'articles-tag-create',''),
        ('g', 'user-normal-policy', 'articles-tag-delete','');



    
      
        
       
        