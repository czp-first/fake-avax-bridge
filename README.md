<h1 style='text-align:center'>fake-avax-bridge</h1>



模仿avax-bridge的理念，实现的跨链桥。



<font color='red'>目前只实现了上桥功能，下桥功能尚不可用。</font>

# 概要

![summary](https://github.com/czp-first/fake-avax-bridge/blob/master/docs/imgs/summary.jpg)



Warden监听来源链和目标链的合法交易，并将其发送至Enclave，如果Enclave接收到的相同的跨链交易，共识warden的数量已经达到阈值，将发送跨链交易。



# Enclave

由两部分组成

- Enclave Server

- Enclave Proxy



## Enclave Server

模拟可信任执行环境(TEE)，与外界的交互均通过Enclave Proxy完成。

主要的功能包括

- 初始化桥的钱包
- 将桥的钱包的助记词，运用shamir算法，拆分成碎片通过分发给Warden
- 存储Warden发送的上桥交易和下桥交易
- 对跨链交易进行签名



交易数据的存储是通过SQLite3实现的，交易在Enclave的状态有以下几种

| status  | comment                        |
| ------- | ------------------------------ |
| wait    | 交易等待Warden共识             |
| pending | 交易已经共识                   |
| ago     | 交易对应的跨链交易已经签名完毕 |



P.S

- 钱包的管理使用的是[hdwallet](https://github.com/meherett/python-hdwallet)
- Shamir算法的实现[在这里](https://github.com/czp-first/fake-avax-bridge/tree/master/ShamirAlgorithm)
- 请求的处理是阻塞模式的



## Enclave Proxy

主要功能包括

- 是Enclave Server与外界通信的代理。接收Warden发送的上桥和下桥交易，转发给Enclave Server，然后告知Warden，该交易在Enclave中的状态。

- Enclave Server签名的跨链交易会在这里进行上链。



# Warden

由三部分组成

- Warden Etl
- Warden Server
- Warden Consumer



上桥、下桥交易数据在Warden中的传递过程

![交易数据流](https://github.com/czp-first/fake-avax-bridge/blob/master/docs/imgs/warden.jpg)



上桥、下桥交易在Warden数据库中的状态有以下几种

| status  | comment                                            |
| ------- | -------------------------------------------------- |
| init    | Warden Etl新发现的上桥和下桥交易                   |
| wait    | 等待Enclave共识                                    |
| pending | 已经发送相应的跨链交易, 但是相应的跨链交易未被确认 |
| done    | 相应的跨链交易已被确认                             |
| timeout | 相应的跨链交易长时间pending                        |



## Warden Etl

Warden监听来源链和目标链的合法交易，发送至Pulsar消息队列。



## Warden Server

主要功能包括

- 轮询Warden数据库中待处理的上桥交易和下桥交易，发送至Enclave Proxy
- 确认跨链交易的状态
- 处理长时间pending的跨链交易，发送至Pulsar消息队列



## Warden Consumer

消费Pulsar消息队列

每个Warden有三个消息主题

- 上桥交易
- 下桥交易
- 配置文件修改



### 上桥交易

- Warden Etl 监听到的新的上桥交易
- Warden Server 发现的长时间pending的跨链交易
- Enclave Proxy 返回的已上链的跨链交易



### 下桥交易

TODO



### 配置文件修改

每个Warden的公开的配置文件存储，模拟的是 AWS的S3 技术，但是文件存储在并发时进行局部修改，结果不可预测。所以通过消息队列控制并发，进行局部修改。



# Fake-Cloud-Service

模仿AWS云服务，实现AWS的以下功能

- S3
- KMS



# 技术栈

- Python、Golang
- PostgreSQL、SQLite3
- Pulsar
- gRPC