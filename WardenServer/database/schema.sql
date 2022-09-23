CREATE TABLE onboard_txn (
    id SERIAL PRIMARY KEY,
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



CREATE TABLE onboard (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    onboard_txn_hash character varying(66),
    nonce integer,
    status character varying(66) NOT NULL,
    batch integer NOT NULL
);


CREATE TABLE offboard_txn (
    id SERIAL PRIMARY KEY,
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


CREATE TABLE offboard (
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    block_hash character varying(66) NOT NULL,
    txn_hash character varying(66) NOT NULL,
    offboard_txn_hash character varying(66),
    nonce integer,
    status character varying(66) NOT NULL,
    batch integer NOT NULL
);