FROM mysql:latest 

# Copy the database schema to the container
COPY ./blog-backend/migrations/*.sql /docker-entrypoint-initdb.d/
COPY ./blog-backend-article-manage/migrations/*.sql /docker-entrypoint-initdb.d/
COPY ./blog-backend-file-manage/migrations/*.sql /docker-entrypoint-initdb.d/


ENV MYSQL_ROOT_PASSWORD=12345
ENV MYSQL_DATABASE=blog

# Expose the port
EXPOSE 3306

# Start the MySQL server
CMD ["mysqld"]

