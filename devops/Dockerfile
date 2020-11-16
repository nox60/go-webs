FROM golang:1.14-alpine

ENV GO111MODULE on

ENV GOPROXY https://goproxy.baidu.com/

ENV GOCACHE "/root/.cache/go-build"

ADD testdb /opt/gopath/src/
#
ADD testdb/verdor/* /opt/gopath/src
#
WORKDIR /opt/gopath/src

RUN go mod init testdb

WORKDIR /opt/gopath/src/testdb

RUN go build

#WORKDIR /opt/gopath/src
#
#RUN go mod vendor