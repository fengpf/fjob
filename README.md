# fjob

### 编译go二进制
go build -o fjob main.go

### 打包镜像tag
docker build -t mountainfeng/fjob:1.1 .

### 单步调试查看log
docker run -it --name fjob mountainfeng/fjob:1.1

### 容器退出删掉
docker run -d --rm --name fjob mountainfeng/fjob:1.1

### 将当前镜像发送到远端 hub.docker.com
docker push mountainfeng/fjob:1.1


查看network的名称，执行 

docker network list