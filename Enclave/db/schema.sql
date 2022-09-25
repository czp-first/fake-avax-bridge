
create table enclave_onboard_txn(
    id integer PRIMARY KEY,
    block_hash text not null,
    transaction_hash text not null,
    batch int not null,
    status text not null,
    wardens text not null
);

-- create unique index blk_txn_batch on transaction(block_hash, transaction_hash, batch);

create table warden(
    id integer primary key,
    identification varchar(64) unique not null,
    credential text not null
);

create table config(
    id integet PRIMARY KEY,
    key varchar(64) unique not null,
    value varchar(64) not null
);

