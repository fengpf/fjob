
FROM alpine:3.8
# Godep for vendoring
# RUN go get github.com/beego/bee
# ENV APP_DIR $GOPATH/src/work
# RUN mkdir -p $APP_DIR
# ADD ./work $APP_DIR/

WORKDIR /go
ADD ./work ./

# Set the entrypoint
ENTRYPOINT  ["./work"]
# ENTRYPOINT (cd $APP_DIR && bee run > debug_log.log 2>&1)
#ADD . $APP_DIR
EXPOSE 80