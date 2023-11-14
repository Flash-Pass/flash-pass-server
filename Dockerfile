# 第一阶段：构建阶段
FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

# 执行 go mod tidy，确保依赖关系最新
RUN go mod tidy

# 在这里你可以执行一些构建命令，如编译代码、运行测试等
RUN CGO_ENABLED=0 go build -o flash-pass-server cmd/server/main.go

# 第二阶段：运行阶段
FROM alpine:latest

WORKDIR /app

# 将构建阶段生成的二进制文件拷贝到运行阶段
COPY --from=builder /app/flash-pass-server .

# 容器启动时执行的命令
CMD ["./flash-pass-server"]
