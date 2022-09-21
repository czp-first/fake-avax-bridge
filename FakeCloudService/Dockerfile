FROM python:3.8

ENV PROJECT_DIR=fake-cloud-service

RUN mkdir /$PROJECT_DIR
WORKDIR /$PROJECT_DIR

COPY requirements.txt ./

# 更换国内源
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple/
# 安装python包
RUN pip install --upgrade pip && pip install -r /$PROJECT_DIR/requirements.txt --no-dependencies

COPY . .
WORKDIR /$PROJECT_DIR
