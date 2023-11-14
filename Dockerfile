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

# 设置环境变量，注意这里使用 ARG 定义默认值，后续会在构建时进行替换
ENV BASE_SERVER_NAMESPACE=386a677f-cc4f-40f3-b596-ee991acf2a68
ENV BASE_SERVER_ADDRESS=82.156.171.8

# 容器启动时执行的命令
CMD ["./flash-pass-server"]
