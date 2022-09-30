CREATE TABLE IF NOT EXISTS enclave_onboard_txn
-- 上桥交易
(
    id INTEGER PRIMARY KEY,  -- 主键ID
    block_hash CHARACTER(66) NOT NULL,  -- 区块哈希
    transaction_hash CHARACTER(66) NOT NULL,  -- 交易哈希
    batch INTEGER NOT NULL,  -- 跨链交易的批次
    status CHARACTER(16) NOT NULL,  -- 上桥交易的状态
    wardens TEXT NOT NULL  -- 达成共识的wardens的标识(逗号分割)
);

CREATE UNIQUE INDEX uni_on_blk_txn_batch ON enclave_onboard_txn(block_hash, transaction_hash, batch);


CREATE TABLE IF NOT EXISTS enclave_offboard_txn
-- 下桥交易
(
    id INTEGER PRIMARY KEY,  -- 主键ID
    block_hash CHARACTER(66) NOT NULL,  -- 区块哈希
    transaction_hash CHARACTER(66) NOT NULL,  -- 交易哈希
    batch INTEGER NOT NULL,  -- 跨链交易的批次
    status CHARACTER(16) NOT NULL,  -- 下桥交易的状态
    wardens TEXT NOT NULL  -- 达成共识的wardens的标识(逗号分割)
);

CREATE UNIQUE INDEX uni_off_blk_txn_batch ON enclave_offboard_txn(block_hash, transaction_hash, batch);


CREATE TABLE IF NOT EXISTS warden
-- warden信息
(
    id INTEGER PRIMARY KEY,  -- 主键ID
    identification VARCHAR(64) UNIQUE NOT NULL,  -- 唯一标识
    credential TEXT NOT NULL,  -- 身份信息
    url VARCHAR(64) NOT NULL  -- warden server的url
);


CREATE TABLE IF NOT EXISTS config
-- 全局配置
(
    id INTEGER PRIMARY KEY,  -- 主键ID
    key VARCHAR(64) UNIQUE NOT NULL,  -- 键
    value VARCHAR(64) NOT NULL -- 值
);

