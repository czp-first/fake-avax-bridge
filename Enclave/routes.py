# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:35:49
"""

from enum import Enum
import sqlite3

from loguru import logger

from crypto.adapter import get_crypto_obj
import wallet


class EnclaveTxnStatus(str, Enum):
    Wait = "wait"  # 交易还在共识的过程中
    Ready = "ready"  # 交易已经共识, 可以进行跨链, 但是该状态不会在数据库中存在
    Pending = "pending"  # 交易已经共识
    Ago = "ago"  # 交易已经签名完毕


def store_credential(share_version: int, threshold: int, wardens: list, from_chain_id: int, to_chain_id: int, db_conn):
    wardens_info = []
    crypto_way = set()
    configs = []
    for warden in wardens:
        crypto_way.add(warden["type"])
        wardens_info.append((warden["identification"], warden["credential"], warden["url"]))
    if len(crypto_way) != 1:
        raise ValueError(f"crypto way is not a consensus: {crypto_way}")

    configs.append(("crypto_way", list(crypto_way)[0]))
    configs.append(("share_version", share_version))
    configs.append(("threshold", threshold))
    db_conn.executemany("INSERT INTO config(key, value)VALUES(?, ?)", configs)
    db_conn.executemany("INSERT INTO warden(identification, credential, url)VALUES(?, ?, ?)", wardens_info)
    db_conn.commit()

    wallet_info = wallet.init_wallet(len(wardens), threshold, from_chain_id, to_chain_id)
    logger.info("wallet info: {}", wallet_info)
    if share_version == 0:
        encrypt_shares = [
            {
                "identification": item["identification"],
                "share": get_crypto_obj(item["identification"], db_conn).encrypt(wallet_info["shares"][index])
            }
            for index, item in enumerate(wardens)
        ]
        return dict(
            encrypt_shares=encrypt_shares,
            from_account_address=wallet_info["onboard_account_address"],
            to_account_address=wallet_info["offboard_account_address"],
        )
    # TODO: shareVersion>0
    else:
        ...


# def recover(wardens_info):
#     shares = (
#         get_crypto_obj(warden["identification"]).decrypt(warden["share"])
#         for warden in wardens_info
#     )
#     return wallet.recover_wallet(shares)


class ProcessTxn:
    def __init__(self, is_onboard: bool, txn: dict, identification: str, db_conn: sqlite3.Connection) -> None:
        """_summary_

        :param bool is_onboard: 是否是上桥交易
        :param dict txn: 交易信息
        :param str identification: warden标识
        :param sqlite3.Connection db_conn: 数据库链接
        """
        if is_onboard:
            self.txn_table = "enclave_onboard_txn"
        else:
            self.txn_table = "enclave_offboard_txn"

        self._txn = txn
        self.block_hash = txn["block_hash"]
        self.txn_hash = txn["txn_hash"]
        self.batch = txn["batch"]
        self.identification = identification
        self.db_conn = db_conn
        self.cursor: sqlite3.Cursor = self.db_conn.cursor()

    def get_sql(self, base_sql):
        """添加sql的表名"""
        return base_sql.format(table=self.txn_table)

    def get_enclave_txn(self):
        """获取enclave中对应的交易信息"""
        sql = self.get_sql("""
                SELECT status, wardens
                FROM {table}
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """
        )
        self.cursor.execute(sql, (self.block_hash, self.txn_hash, self.batch))
        row = self.cursor.fetchone()
        if not row:
            return None
        return dict(status=row[0], wardens=row[1].split(","))

    def insert_new_enclave_txn(self):
        """在enclave中插入新的交易"""
        logger.info("{}: new {} transaction", self.txn_table, self.txn_table)
        sql = self.get_sql(
            """
                INSERT INTO {table}(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """
        )
        self.cursor.execute(sql, (self.block_hash, self.txn_hash, self.batch, self.identification, EnclaveTxnStatus.Wait.value))
        return dict(status=EnclaveTxnStatus.Wait.value)

    def process_ready(self, threshold, latest_wardens):
        """处理达到共识的交易"""
        logger.info("{}: ready", self.txn_table)
        sql = self.get_sql(
            """
                UPDATE {table}
                SET wardens=?, status=?
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """
        )
        self.cursor.execute(sql, (",".join(latest_wardens), EnclaveTxnStatus.Pending.value, self.block_hash, self.txn_hash, self.batch))
        self.cursor.execute(
            """
                SELECT identification, url
                FROM warden
                WHERE identification IN ({seq})
                ORDER BY id
            """.format(seq=",".join(["?"]*threshold)),
            latest_wardens[:threshold],
        )
        wardens_info = [
            dict(
                identification=i[0],
                url=i[1],
            )
            for i in self.cursor.fetchall()
        ]
        return dict(
            status=EnclaveTxnStatus.Ready.value,
            wardens=wardens_info
        )

    def process_unready(self, latest_wardens):
        """处理未达到共识的交易"""
        logger.info("{}: go on waiting", self.txn_table)
        sql = self.get_sql(
            """
                UPDATE {table}
                SET wardens=?
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """
        )
        self.cursor.execute(sql, (",".join(latest_wardens), self.block_hash, self.txn_hash, self.batch))
        return dict(status=EnclaveTxnStatus.Wait.value)

    def process_new_warden(self, wardens):
        """处理交易加入新的warden"""
        self.cursor.execute("SELECT value FROM config WHERE key=?", ("threshold",))
        threshold = int(self.cursor.fetchone()[0])
        logger.info("{}: current wardens {}", self.txn_table, wardens)
        wardens.append(self.identification)
        if len(wardens) >= threshold:
            return self.process_ready(threshold, wardens)
        return self.process_unready(wardens)

    def process_old_enclave_txn(self, enclave_txn: dict):
        """处理在enclave已经存在的交易"""
        if self.identification in enclave_txn["wardens"] or enclave_txn["status"] != EnclaveTxnStatus.Wait.value:
            logger.info("{}: status {}", self.txn_table, enclave_txn["status"])
            return dict(status=enclave_txn["status"])

        return self.process_new_warden(enclave_txn["wardens"])

    def process_enclave_txn(self, enclave_txn):
        """处理对应的enclave中的交易"""
        if enclave_txn:
            return self.process_old_enclave_txn(enclave_txn)
        return self.insert_new_enclave_txn()

    def run(self):
        encalve_txn = self.get_enclave_txn()
        result = self.process_enclave_txn(encalve_txn)
        self.cursor.close()
        self.db_conn.commit()
        return result


def process_onboard_txn(txn, identification, db_conn: sqlite3.Connection):
    """处理上桥交易"""
    return ProcessTxn(True, txn, identification, db_conn).run()


def process_offboard_txn(txn, identification, db_conn: sqlite3.Connection):
    """处理下桥交易"""
    return ProcessTxn(False, txn, identification, db_conn).run()


def get_mnemonic(warden_shares, db_conn):
    decrypt_shares = []
    for warden_share in warden_shares:
        decrypt_share = get_crypto_obj(warden_share["identification"], db_conn).decrypt(warden_share["encrypt_share"])
        decrypt_shares.append(decrypt_share)
    return wallet.recover_wallet(tuple(decrypt_shares))


def set_enclave_txn_ago(is_onboard, db_conn, block_hash, txn_hash, batch):
    table = "enclave_onboard_txn" if is_onboard else "enclave_offboard_txn"
    cursor = db_conn.cursor()
    cursor.execute(
        """
            UPDATE {table}
            SET status=?
            WHERE block_hash=? AND transaction_hash=? AND batch=?
        """.format(table=table),
        (EnclaveTxnStatus.Ago.value, block_hash, txn_hash, batch)
    )
    cursor.execute("SELECT url FROM warden ORDER BY id")
    urls = [i[0] for i in cursor.fetchall()]
    cursor.close()
    return urls


def sign_onboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce, fee, origin_txn_hash, origin_block_hash, origin_batch, db_conn):
    """签名上桥交易对应的跨链交易"""
    mnemonic = get_mnemonic(warden_shares, db_conn)

    params = dict(
        is_eip1559=is_eip1559,
        mnemonic=mnemonic,
        chain_id=chain_id,
        contract_addr=contract_addr,
        amount=amount,
        gas_price=gas_price,
        account_addr=account_addr,
        nonce=nonce,
        origin_txn_hash=origin_txn_hash,
        fee=fee,
    )

    urls = set_enclave_txn_ago(True, db_conn, origin_block_hash, origin_txn_hash, origin_batch)
    db_conn.commit()

    return {
        # web3 需要0x开头, go-ethereum不需要
        'txn': wallet.sign_onboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
        "urls": urls,
    }


def sign_offboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce, origin_block_hash, origin_txn_hash, origin_batch, db_conn):
    """签名下桥交易对应的跨链交易"""
    mnemonic = get_mnemonic(warden_shares, db_conn)

    params = dict(
        is_eip1559=is_eip1559,
        mnemonic=mnemonic,
        chain_id=chain_id,
        contract_addr=contract_addr,
        amount=amount,
        gas_price=gas_price,
        account_addr=account_addr,
        nonce=nonce
    )

    urls = set_enclave_txn_ago(False, db_conn, origin_block_hash, origin_txn_hash, origin_batch)
    db_conn.commit()

    return {
        'txn': wallet.sign_offboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
        "urls": urls,
    }


ROUTES_MAP = {
    'storeCredential': store_credential,
    # 'recove': recover,

    'onboardTxn': process_onboard_txn,
    'signOnboardTxn': sign_onboard_txn,

    'offboardTxn': process_offboard_txn,
    'signOffboardTxn': sign_offboard_txn,
}
