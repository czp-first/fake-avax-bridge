FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenEtl
COPY BridgeUtils/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/BridgeUtils
# COPY WardenServer/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenServer
COPY WardenEtl/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenEtl
RUN go build .
