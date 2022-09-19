create table enclave(

)

create table transaction(
    block_hash varchar(66) not null,
    transaction_hash varchar(66) not null,
    batch integer not null,
    status varchar(66) not null
    -- UNIQUE(block_hash, transaction_hash, batch)
)

