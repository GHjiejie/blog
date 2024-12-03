@echo off
@REM cd C:\jie-github-projects\blog\blog-backend\cmd
@REM go run main.go
start cmd /k "cd C:\jie-github-projects\blog\blog-backend\cmd && go run main.go"
start cmd /k "cd C:\jie-github-projects\blog\blog-backend-article-manage\cmd && go run main.go"
start cmd /k "cd C:\jie-github-projects\blog\blog-backend-file-manage\cmd && go run main.go"


@REM start cmd /k "cd C:\jie-github-projects\blog\blog-frontend && npm run dev"
pause

