FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenServer
COPY EnclaveProxy/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/EnclaveProxy
COPY WardenServer/ $GOPATH/src/github.com/czp-first/fake-avax-bridge/WardenServer
RUN go build .
