FROM mysql

# 直接复制 SQL 文件到 /docker-entrypoint-initdb.d 目录
COPY ../blog-backend/migrations/*.sql /docker-entrypoint-initdb.d/
COPY ../blog-backend-article-manage/migrations/*.sql /docker-entrypoint-initdb.d/
COPY ../blog-backend-file-manage/migrations/*.sql /docker-entrypoint-initdb.d/
