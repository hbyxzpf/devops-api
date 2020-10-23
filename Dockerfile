FROM golang:alpine AS builder
MAINTAINER "hbyx.zpf@gmail.com"

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn
# 移动到工作目录：/build
WORKDIR /build
# 将代码复制到容器中
COPY . .
RUN  go build -o devops-api .
WORKDIR /dist
RUN cp /build/devops-api .
CMD ./devops-api init
#ENTRYPOINT ["./devops-api","init"]
ENTRYPOINT ["./devops-api","server"]
