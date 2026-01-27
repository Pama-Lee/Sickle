# 构建阶段 - 使用 Debian 基础镜像避免 musl libc 兼容性问题
FROM golang:1.21-bookworm AS builder

# 安装必要的构建工具
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建二进制文件
RUN CGO_ENABLED=1 GOOS=linux go build -o sickle .

# 运行阶段 - 使用轻量级 Debian
FROM debian:bookworm-slim

# 安装运行时依赖
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

# 设置时区
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/sickle .

# 创建必要的目录
RUN mkdir -p /app/webhooks /app/logs /app/request /app/config

# 暴露端口
EXPOSE 8080

# 设置数据卷
VOLUME ["/app/webhooks", "/app/logs", "/app/config"]

# 运行程序
CMD ["./sickle"]
