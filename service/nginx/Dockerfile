
FROM nginx:latest
MAINTAINER mountainfpf <mountainfpf@gmail.com>
RUN apt-get update && apt-get install -y curl unzip
COPY default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80