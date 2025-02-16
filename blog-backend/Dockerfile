# 使用官方 Go 语言镜像
FROM golang AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 将整个项目复制到容器中
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog-backend ./cmd/main.go

# 使用轻量级的镜像来运行应用程序
FROM alpine:latest

# 设置工作目录
WORKDIR /app  # 工作目录与项目目录一致

# 从构建阶段复制生成的二进制文件和配置文件
COPY --from=builder /app/blog-backend .
COPY --from=builder /app/conf/config.toml ./conf/config.toml  

# 暴露应用运行的端口
EXPOSE 8080

# 运行应用程序
CMD ["./blog-backend"]
