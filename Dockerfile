FROM golang:latest

WORKDIR $GOPATH/src/go-docker-test
COPY . $GOPATH/src/go-docker-test

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn


RUN go build .

EXPOSE 8011
ENTRYPOINT ["./go-docker-test"]