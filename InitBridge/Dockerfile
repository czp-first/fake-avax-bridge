FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/czp-first/fake-avax-bridge/InitBridge
COPY WardenServer/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenServer
COPY InitBridge/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/InitBridge
RUN go build .
