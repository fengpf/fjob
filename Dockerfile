#源镜像
FROM alpine:3.8
# FROM alpine:3.7
#作者
MAINTAINER mountainfpf "mountainfpf@gmail.com"

ENV TZ "Asia/Shanghai"

#设置工作目录
# WORKDIR $GOPATH/src/github.com/fengpf/job
WORKDIR /go
#将服务器的go工程代码加入到docker容器中
# ADD . $GOPATH/src/github.com/fengpf/job
ADD ./fjob ./
#go构建可执行文件
# CMD go version && go build -v
# RUN go build .
#暴露端口
EXPOSE 6064
#最终运行docker的命令
ENTRYPOINT  ["./fjob"]