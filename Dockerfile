# FROM 基于 golang:1.16-alpine
FROM golang:1.18-alpine AS builder

# ENV 设置环境变量
ENV GOPATH=/golang/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
ENV PROJECTNAME=gin-hong-api

# RUN 创建目录
RUN mkdir -p $GOPATH/src/
RUN mkdir -p $GOPATH/bin/
RUN mkdir -p $GOPATH/pkg/

# COPY 源路径 目标路径
COPY . $GOPATH/src/$PROJECTNAME

RUN cd $GOPATH/src/$PROJECTNAME && go install .
RUN cd $GOPATH/src/$PROJECTNAME/cmd/queue && go install .
#
# FROM 基于 alpine:latest
FROM alpine:latest

# RUN 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /golang/go/bin /opt


RUN apt-get install -y supervisor

COPY --from=builder /golang/go/src/supervisor.ini /etc/supervisor/conf.d/supervisor.ini

RUN supervisord -c /etc/supervisor/supervisord.conf

# EXPOSE 设置端口映射
EXPOSE 8800/tcp

# WORKDIR 设置工作目录
WORKDIR /opt/bin

# CMD 设置启动命令
CMD ["./$PROJECTNAME", "-env", "docker"]

# 设置命令行执行命令
CMD ["./queue", "-env", "docker"]