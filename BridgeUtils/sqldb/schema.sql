create sequence warden_onboard_id_seq;
CREATE TABLE if not EXISTS warden_onboard (
    id bigint not null default nextval('warden_onboard_id_seq') PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    chain_id integer NOT NULL,
    contract character varying(66) NOT NULL,
    account character varying(66) NOT NULL,
    amount numeric(29,0) NOT NULL,
    gas_price numeric(29,0),
    block_number integer NOT NULL,
    txn_index integer NOT NULL,
    onboard_txn_hash character varying(66),
    status character varying(66) NOT NULL,
    onboard_txn_amount numeric(29,0),
    nonce integer,
    batch integer NOT NULL
);

comment on table warden_onboard is 'warden上桥交易';

alter sequence warden_onboard_id_seq owned by warden_onboard.id;


create sequence enclave_onboard_id_seq;
CREATE TABLE if not EXISTS enclave_onboard (
    id bigint not null default nextval('enclave_onboard_id_seq') PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    onboard_txn_hash character varying(66),
    nonce integer,
    status character varying(66) NOT NULL,
    batch integer NOT NULL
);

comment on table enclave_onboard is 'enclave上桥交易';

alter sequence enclave_onboard_id_seq owned by enclave_onboard.id;


create sequence warden_offboard_id_seq;
CREATE TABLE if not EXISTS warden_offboard (
    id bigint not null default nextval('warden_offboard_id_seq') PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    chain_id integer NOT NULL,
    contract character varying(66) NOT NULL,
    account character varying(66) NOT NULL,
    amount numeric(29,0) NOT NULL,
    gas_price numeric(29,0),
    block_number integer NOT NULL,
    txn_index integer NOT NULL,
    offboard_txn_hash character varying(66),
    status character varying(66) NOT NULL,
    offboard_txn_amount numeric(29,0),
    nonce integer,
    batch integer NOT NULL
);

comment on table warden_offboard is 'warden下桥交易';

alter sequence warden_offboard_id_seq owned by warden_offboard.id;


create sequence enclave_offboard_id_seq;
CREATE TABLE if not EXISTS enclave_offboard (
    id bigint not null DEFAULT nextval('enclave_offboard_id_seq') PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    offboard_txn_hash character varying(66),
    nonce integer,
    status character varying(66) NOT NULL,
    batch integer NOT NULL
);

comment on table enclave_offboard is 'enclave下桥交易';

alter sequence enclave_offboard_id_seq owned by enclave_offboard.id;
