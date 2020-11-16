FROM golang:1.14-alpine

ENV GO111MODULE on

ENV GOPROXY https://goproxy.cn/

ENV GOCACHE "/root/.cache/go-build"

RUN mkdir /opt/logs/

ADD testdb /opt/gopath/src/testdb

WORKDIR /opt/gopath/src/testdb/

RUN go mod tidy

WORKDIR /opt/gopath/src/testdb/

RUN go build

ENTRYPOINT /opt/gopath/src/testdb/testdb