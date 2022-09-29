FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenProducer
COPY BridgeUtils/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/BridgeUtils
COPY WardenProducer/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenProducer
RUN go build .
