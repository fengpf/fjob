# work

### 打包镜像tag
docker build -t mountainfeng/work:1.0 .

### 单步调试查看log
docker run -it --name work mountainfeng/work:1.0

### 将当前镜像发送到远端 hub.docker.com
docker push mountainfeng/work:1.1

docker-compose build && docker-compose up -d

### 查看网络
docker network list