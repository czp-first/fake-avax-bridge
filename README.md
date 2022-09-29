# fake-avax-bridge



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











## 初始化桥配置
```shell
cd FakeCloudService
docker-compose up --build

cd Enclave
docker-compose up --build

cd EnclaveProxy
docker-compose up --build
docker run -it --env-file dev.env  -p 50055:50055 enclaveproxy_enclave-proxy ./EnclaveProxy

cd WardenServer
docker-compose up --build
docker run -it --env-file ../envs/warden1.env -p 50051:50051  wardenserver_warden-server ./WardenServer
docker run -it --env-file ../envs/warden2.env -p 50052:50052  wardenserver_warden-server ./WardenServer
docker run -it --env-file ../envs/warden3.env -p 50053:50053  wardenserver_warden-server ./WardenServer

cd WardenConsumer
docker-compose up --build
docker run -it --env-file ../envs/warden1.env wardenconsumer_warden-consumer ./WardenConsumer
docker run -it --env-file ../envs/warden2.env wardenconsumer_warden-consumer ./WardenConsumer
docker run -it --env-file ../envs/warden3.env wardenconsumer_warden-consumer ./WardenConsumer

cd WardenProducer
docker-compose up --build
docker run -it --env-file ../envs/warden1.env wardenproducer_warden-producer ./WardenProducer
docker run -it --env-file ../envs/warden2.env wardenproducer_warden-producer ./WardenProducer
docker run -it --env-file ../envs/warden3.env wardenproducer_warden-producer ./WardenProducer

cd InitBridge
docker-compose up --build
docker run -it --env-file dev.env initbridge_init-bridge ./InitBridge

cd WardenEtl
docker-compose up --build
docker run -it --env-file ../envs/warden1.env  wardenetl_warden-etl ./WardenEtl
docker run -it --env-file ../envs/warden2.env  wardenetl_warden-etl ./WardenEtl
docker run -it --env-file ../envs/warden3.env  wardenetl_warden-etl ./WardenEtl
```

