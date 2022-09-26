# fake-avax-bridge

## 初始化桥配置
```shell
cd FakeCloudService
docker-compose up --build

cd ../Enclave
docker-compose up --build

cd ../EnclaveProxy
docker-compose up --build
docker run -it --env-file dev.env  -p 50055:50055 enclaveproxy_enclave-proxy ./EnclaveProxy

cd ../WardenServer
docker-compose up --build
docker run -it --env-file wardens/warden1/dev.env -p 50051:50051  wardenserver_warden-server ./WardenServer
docker run -it --env-file wardens/warden2/dev.env -p 50052:50052  wardenserver_warden-server ./WardenServer
docker run -it --env-file wardens/warden3/dev.env -p 50053:50053  wardenserver_warden-server ./WardenServer

cd ../WardenConsumer
docker-compose up --build
docker run -it --env-file warden1.env wardenconsumer_warden-consumer ./WardenConsumer
docker run -it --env-file warden2.env wardenconsumer_warden-consumer ./WardenConsumer
docker run -it --env-file warden3.env wardenconsumer_warden-consumer ./WardenConsumer

cd ../InitBridge
docker-compose up --build
docker run -it --env-file dev.env initbridge_init-bridge ./InitBridge
```

