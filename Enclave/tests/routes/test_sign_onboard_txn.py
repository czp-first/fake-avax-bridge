# -*- coding: UTF-8 -*-
"""
@Summary : 测试签名上桥交易
@Author  : Rey
@Time    : 2022-09-29 16:38:41
@Run     : python -m unittest tests/routes/test_sign_onboard_txn.py
"""

import json
import sqlite3
import unittest
from uuid import uuid4

from cryptography.fernet import Fernet
from web3 import Web3

from crypto.adapter import get_crypto_obj
from routes import (
    EnclaveTxnStatus,
    sign_onboard_txn,
)
from wallet import init_wallet


class TestSignOnboardTxn(unittest.TestCase):
    """test routes.sign_onboard_txn"""
    def setUp(self) -> None:
        self.conn = sqlite3.connect(":memory:")
        cursor = self.conn.cursor()
        with open("db/schema.sql") as f:
            cursor.executescript(f.read())
        crypto_way = "local"
        self.threshold = 2
        share_version = 0
        self.wardens = [
            {
                "identification": str(uuid4()),
                "credential": json.dumps(dict(key=Fernet.generate_key().decode("utf-8"))),
                "url": f"url{i}"
            }
            for i in range(1, 4)
        ]
        configs = [
            ("crypto_way", crypto_way),
            ("share_version", share_version),
            ("threshold", self.threshold)
        ]
        cursor.executemany("INSERT INTO config(key, value)VALUES(?, ?)", configs)
        cursor.executemany(
            "INSERT INTO warden(identification, credential, url)VALUES(?, ?, ?)",
            [(i["identification"], i["credential"], i["url"]) for i in self.wardens]
        )
        cursor.close()

        wallet_info = init_wallet(len(self.wardens), self.threshold, 3, 256)
        for index, warden in enumerate(self.wardens):
            warden["encrypt_share"] = get_crypto_obj(warden["identification"], self.conn).encrypt(wallet_info["shares"][index])


    def tearDown(self) -> None:
        self.conn.close()

    def test_success_eip1559(self):
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification = self.wardens[0]["identification"]

        cursor = self.conn.cursor()
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], identification, EnclaveTxnStatus.Wait.value),
        )
        wardens_shares = [
            dict(identification=self.wardens[0]["identification"], encrypt_share=self.wardens[0]["encrypt_share"]),
            dict(identification=self.wardens[1]["identification"], encrypt_share=self.wardens[1]["encrypt_share"]),
        ]
        nonce = 100
        gas_price = Web3.toWei(8, "gwei")
        resp = sign_onboard_txn(
            is_eip1559=True,
            warden_shares=wardens_shares,
            chain_id=256,
            contract_addr="0x71F8BA01598Ee26bf1706a1D49D7af5A2899E53d",
            amount=10000000000000000,
            gas_price=gas_price,
            account_addr="0x44fe5BD0e041aB1E42579812dA1D36234577Cf74",
            nonce=nonce,
            fee=Web3.toWei(6, "gwei"),
            origin_txn="0x7b6785e74dad9ac03b7fcb6d4b20fe2f2178d7eb90a563bcfa1c613d2785f31b",
            origin_block_hash="0xa37a7b5b95a98db4ea0b27972f2c2c02bda99c2e035d7478b67f9c456b43a5c3",
            origin_batch=1,
            db_conn=self.conn,
        )
        self.assertIsNotNone(resp["txn"])
        self.assertEqual(resp["nonce"], nonce)
        self.assertEqual(resp["gas_price"], gas_price)
        self.assertTrue(resp["is_eip1559"])
        self.assertListEqual(resp["urls"], [i["url"] for i in self.wardens])
