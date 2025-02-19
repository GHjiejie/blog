# 使用 Node.js 作为基础镜像
FROM node AS builder

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install 

# 复制项目文件
COPY . .

# 构建项目 (确保你的项目有 build 命令, 例如 `npm run build`)
RUN npm run build

# 使用 Nginx 作为生产环境服务器 (也可选择其他的, 例如 Caddy)
FROM nginx:latest

# 将构建好的静态文件复制到 Nginx 的默认目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 暴露 9000 端口
EXPOSE 9000

# 可选: 如果你的 Vue 应用使用了 history 模式, 需要配置 Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 启动 Nginx (默认 CMD 已经存在, 通常不需要修改)
CMD ["nginx", "-g", "daemon off;"]

