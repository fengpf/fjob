# fjob

### 进入镜像
docker run -t -i mountainfeng/fjob:1.1 /bin/bash

### 修改已有镜像
docker commit -m "test" -a "fpf test" 0860637e3bbc mountainfeng/fjob:1.2

其中，-m 来指定提交的说明信息，跟我们使用的版本控制工具一样；-a 可以指定更新的用户信息；之后是用来创建镜像的容器的 ID；
最后指定目标镜像的仓库名和 tag 信息。创建成功后会返回这个镜像的 ID 信息。


## 利用 Dockerfile 来创建镜像
使用 docker commit 来扩展一个镜像比较简单，但是不方便在一个团队中分享。
我们可以使用 docker build来创建一个新的镜像。为此，首先需要创建一个 Dockerfile，包含一些如何创建镜像的指令。

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