# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:35:49
"""

from enum import Enum

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


def process_onboard_txn(txn, identification, db_conn):
    cursor = db_conn.cursor()
    cursor.execute(
        """
            SELECT status, wardens
            FROM enclave_onboard_txn
            WHERE block_hash=? AND transaction_hash=? AND batch=?
        """,
        (txn["block_hash"], txn["txn_hash"], txn["batch"])
    )
    row = cursor.fetchone()
    if not row:
        logger.info("onboard: new onboard transaction")
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], identification, EnclaveTxnStatus.Wait.value)
        )
        db_conn.commit()
        return dict(status=EnclaveTxnStatus.Wait.value)
    else:
        wardens = row[1].split(",")
        if identification in wardens:
            logger.info("onboard: duplicate")
            return dict(status=row[0])

        if row[0] != EnclaveTxnStatus.Wait.value:
            logger.info("onboard: status {}", row[0])
            return dict(status=row[0])

        cursor.execute("SELECT value FROM config WHERE key=?", ("threshold",))
        threshold = int(cursor.fetchone()[0])
        logger.info("onboard: current wardens {}", wardens)
        wardens.append(identification)
        if len(wardens) >= threshold:
            logger.info("onboard: ready")
            cursor.execute(
                """
                    UPDATE enclave_onboard_txn
                    SET wardens=?, status=?
                    WHERE block_hash=? AND transaction_hash=? AND batch=?
                """,
                (",".join(wardens), EnclaveTxnStatus.Pending.value, txn["block_hash"], txn["txn_hash"], txn["batch"])
            )
            cursor.execute(
                """
                    SELECT identification, url
                    FROM warden
                    WHERE identification IN ({seq})
                    ORDER BY id
                """.format(seq=",".join(["?"]*threshold)),
                wardens[:threshold],
            )
            wardens_info = [
                dict(
                    identification=i[0],
                    url=i[1],
                )
                for i in cursor.fetchall()
            ]
            cursor.close()
            db_conn.commit()
            return dict(
                status=EnclaveTxnStatus.Ready.value,
                wardens=wardens_info,
            )
        else:
            logger.info("onboard: go on waiting")
            cursor.execute(
                """
                    UPDATE enclave_onboard_txn
                    SET wardens=?
                    WHERE block_hash=? AND transaction_hash=? AND batch=?
                """,
                (",".join(wardens), txn["block_hash"], txn["txn_hash"], txn["batch"])
            )
            cursor.close()
            db_conn.commit()
            return dict(status=EnclaveTxnStatus.Wait.value)


# TODO(Rey): refactor like proceess_onboard_txn
def process_offboard_txn(txn, identification):

    unique_txn = f'{txn["block_hash"]}{txn["txn_hash"]}{txn["batch"]}'
    if unique_txn in pending_offboard_txns:
        return {'status': 'pending'}
    if unique_txn in ago_offboard_txns:
        return {'status': 'ago'}

    if unique_txn in comfirming_offboard_txns:
        if identification in comfirming_offboard_txns[unique_txn]:
            return {'status': 'wait'}
        comfirming_offboard_txns[unique_txn].append(identification)
        if len(comfirming_offboard_txns[unique_txn]) >= threshold:
            pending_offboard_txns.append(unique_txn)
            return {
                'status': 'ready',
                'wardens': comfirming_offboard_txns[unique_txn][:threshold]
            }
        return {'status': 'wait'}

    comfirming_offboard_txns[unique_txn] = [identification]
    return {'status': 'wait'}


def sign_onboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce, fee, origin_txn, origin_block_hash, origin_batch, db_conn):
    decrypt_shares = []
    for warden_share in warden_shares:
        decrypt_share = get_crypto_obj(warden_share["identification"], db_conn).decrypt(warden_share["encrypt_share"])
        decrypt_shares.append(decrypt_share)
    mnemonic = wallet.recover_wallet(tuple(decrypt_shares))

    params = dict(
        is_eip1559=is_eip1559,
        mnemonic=mnemonic,
        chain_id=chain_id,
        contract_addr=contract_addr,
        amount=amount,
        gas_price=gas_price,
        account_addr=account_addr,
        nonce=nonce,
        origin_txn=origin_txn,
        fee=fee,
    )

    cursor = db_conn.cursor()
    cursor.execute(
        """
            UPDATE enclave_onboard_txn
            SET status=?
            WHERE block_hash=? AND transaction_hash=? AND batch=?
        """,
        (EnclaveTxnStatus.Ago.value, origin_block_hash, origin_txn, origin_batch)
    )
    cursor.execute(
        """
            SELECT url
            FROM warden
            ORDER BY id
        """,
    )
    urls = [i[0] for i in cursor.fetchall()]
    cursor.close()
    db_conn.commit()

    return {
        # web3 需要0x开头, go-ethereum不需要
        'txn': wallet.sign_onboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
        "urls": urls,
    }


# TODO(Rey): 
def sign_offboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce):
    decrypt_shares = []
    for warden_share in warden_shares:
        decrypt_share = get_crypto_obj(warden_share["identification"]).decrypt(warden_share["encrypt_share"])
        decrypt_shares.append(decrypt_share)
    mnemonic = wallet.recover_wallet(tuple(decrypt_shares))

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

    return {
        'txn': wallet.sign_offboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
    }


ROUTES_MAP = {
    'storeCredential': store_credential,
    # 'recove': recover,

    'onboardTxn': process_onboard_txn,
    'signOnboardTxn': sign_onboard_txn,

    'offboardTxn': process_offboard_txn,
    'signOffboardTxn': sign_offboard_txn,
}
