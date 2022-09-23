FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenConsumer
COPY WardenConsumer/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenConsumer
RUN go build .
