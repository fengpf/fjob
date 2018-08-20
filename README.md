# fjob

docker build -t fengpf/fjob:1.0 .

docker run -d --rm --name fjob fengpf/fjob:1.0

docker push fengpf/fjob:1.0


查看network的名称，执行 

docker network list