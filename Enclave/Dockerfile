FROM golang:latest AS builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY ShamirAlgorithm/ /build
RUN go build -o shamirbin cli/main.go

##################################################

FROM python:3.8

ENV PROJECT_DIR=Enclave

RUN mkdir /$PROJECT_DIR
WORKDIR /$PROJECT_DIR

COPY ./Enclave/requirements.txt ./

# 更换国内源
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple/
# 安装python包
RUN pip install --upgrade pip && pip install -r /$PROJECT_DIR/requirements.txt --no-dependencies

COPY ./Enclave/ .
COPY --from=builder /build/shamirbin .
WORKDIR /$PROJECT_DIR
