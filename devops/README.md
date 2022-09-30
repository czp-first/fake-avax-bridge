<h1 style='text-align:center'>docker本地部署</h1>

## 基本环境



### postgres

```sh
docker run -d --name dev-postgres -e POSTGRES_PASSWORD=Pass2022! -v ${HOME}/postgres-data/:/var/lib/postgresql/data -p 5432:5432 postgres
```



### [pulsar](https://pulsar.apache.org/)

```sh
docker run -it -p 6650:6650  -p 8080:8080 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:2.10.1 bin/pulsar standalone
```



## 数据库

### 命令行连接数据库

```sh
export PGPASSWORD=Pass2022!
psql -h localhost -p 5432  -U postgres
```



### 基本命令

```sh
\l: 查看所有数据库
\c dbname: 选择数据库
create database dbname; 创建数据库
```



### 执行sql脚本

```sh
psql -h localhost -p 5432 -d db1 -U userA -f /pathA/xxx.sql: execute sql document
```





## 启动桥
```shell
# 以三个warden为例

cd devops

# 1 创建warden数据库表
# warden1
psql -h localhost -p 5432 -d warden1 -U postgres -f ../BridgeUtils/sqldb/schema.sql
# warden2
psql -h localhost -p 5432 -d warden2 -U postgres -f ../BridgeUtils/sqldb/schema.sql
# warden3
psql -h localhost -p 5432 -d warden3 -U postgres -f ../BridgeUtils/sqldb/schema.sql

# 2 构建镜像
docker-compose up --build

# 3 启动假的云服务
docker run -it -p 5001:80 devops_fake-cloud-service-api  uvicorn main:app --host 0.0.0.0 --port 80 --workers 3

# 4 启动enclave服务
docker run -it --env-file enclave.env -p 8000:8000 devops_enclave python main.py

# 5 启动enclave代理服务
docker run -it --env-file enclave-proxy.env  -p 50055:50055 devops_enclave-proxy ./EnclaveProxy

# 6 启动warden-server服务
docker run -it --env-file warden1.env -p 50051:50051  devops_warden-server ./WardenServer
docker run -it --env-file warden2.env -p 50052:50052  devops_warden-server ./WardenServer
docker run -it --env-file warden3.env -p 50053:50053  devops_warden-server ./WardenServer

# 7 启动warden-consumer服务
docker run -it --env-file warden1.env devops_warden-consumer ./WardenConsumer
docker run -it --env-file warden2.env devops_warden-consumer ./WardenConsumer
docker run -it --env-file warden3.env devops_warden-consumer ./WardenConsumer

# 8 初始化warden-settings配置
docker run -it --env-file warden1.env devops_warden-producer ./WardenProducer
docker run -it --env-file warden2.env devops_warden-producer ./WardenProducer
docker run -it --env-file warden3.env devops_warden-producer ./WardenProducer

# 9 初始化桥的配置
docker run -it --env-file init-bridge.env devops_init-bridge ./InitBridge

# 10 启动warden-etl服务
docker run -it --env-file warden1.env  devops_warden-etl ./WardenEtl
docker run -it --env-file warden2.env  devops_warden-etl ./WardenEtl
docker run -it --env-file warden3.env  devops_warden-etl ./WardenEtl
```

