# -*- coding: UTF-8 -*-
"""
@Summary : 测试处理上桥交易
@Author  : Rey
@Time    : 2022-07-12 22:24:00
@Run     : python -m unittest tests/routes/test_process_onboard_txn.py
"""

import json
import sqlite3
import unittest
from uuid import uuid4

from cryptography.fernet import Fernet

from routes import (
    EnclaveTxnStatus,
    process_onboard_txn,
)


class TestProcessOnboardTxn(unittest.TestCase):
    """test routes.process_onboard_txn"""
    def setUp(self) -> None:
        self.conn = sqlite3.connect(":memory:")
        cursor = self.conn.cursor()
        with open("db/schema.sql") as f:
            cursor.executescript(f.read())
        crypto_way = "local"
        self.wardens = [
            {
                "identification": str(uuid4()),
                "credential": json.dumps(dict(key=Fernet.generate_key().decode("utf-8"))),
                "url": f"url{i}"
            }
            for i in range(1, 5)
        ]
        configs = [
            ("crypto_way", crypto_way),
            ("share_version", 0),
            ("threshold", 3)
        ]
        cursor.executemany("INSERT INTO config(key, value)VALUES(?, ?)", configs)
        cursor.executemany(
            "INSERT INTO warden(identification, credential, url)VALUES(?, ?, ?)",
            [(i["identification"], i["credential"], i["url"]) for i in self.wardens]
        )
        cursor.close()

    def test_new_txn(self):
        """new transaction in enclave"""
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification = self.wardens[0]["identification"]
        resp = process_onboard_txn(txn, identification, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Wait.value)

        cursor = self.conn.cursor()
        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], identification)
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Wait.value)

    def test_duplicate_warden_txn(self):
        """warden向enclave发送重复交易"""
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

        resp = process_onboard_txn(txn, identification, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Wait.value)

        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], identification)
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Wait.value)

    def test_append_warden(self):
        """多一个warden共识交易, 但是还未到达共识的阈值"""
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification0 = self.wardens[0]["identification"]

        cursor = self.conn.cursor()
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], identification0, EnclaveTxnStatus.Wait.value),
        )

        identification1 = self.wardens[1]["identification"]
        resp = process_onboard_txn(txn, identification1, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Wait.value)

        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], ",".join([identification0, identification1]))
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Wait.value)

    def test_ready(self):
        """多一个warden共识交易, 并到达共识的阈值"""
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification0 = self.wardens[0]["identification"]
        identification1 = self.wardens[1]["identification"]

        cursor = self.conn.cursor()
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], ",".join([identification0, identification1]), EnclaveTxnStatus.Wait.value),
        )

        identification3 = self.wardens[3]["identification"]
        resp = process_onboard_txn(txn, identification3, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Ready.value)
        self.assertEqual(len(resp["wardens"]), 3)
        self.assertDictEqual(resp["wardens"][0], dict(identification=self.wardens[0]["identification"], url=self.wardens[0]["url"]))
        self.assertDictEqual(resp["wardens"][1], dict(identification=self.wardens[1]["identification"], url=self.wardens[1]["url"]))
        self.assertDictEqual(resp["wardens"][2], dict(identification=self.wardens[3]["identification"], url=self.wardens[3]["url"]))

        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], ",".join([identification0, identification1, identification3]))
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Pending.value)

    def test_pending(self):
        """新的warden, 但是交易状态已经为pending"""
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification0 = self.wardens[0]["identification"]

        cursor = self.conn.cursor()
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], identification0, EnclaveTxnStatus.Pending.value),
        )

        identification1 = self.wardens[1]["identification"]
        resp = process_onboard_txn(txn, identification1, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Pending.value)

        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], identification0)
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Pending.value)

    def test_ago(self):
        """新的warden, 但是交易状态已经为ago"""
        txn = dict(
            block_hash="0xblockhash",
            txn_hash="0xtxnhash",
            batch=1,
        )
        identification0 = self.wardens[0]["identification"]

        cursor = self.conn.cursor()
        cursor.execute(
            """
                INSERT INTO enclave_onboard_txn(block_hash, transaction_hash, batch, wardens, status)
                    VALUES(?, ?, ?, ?, ?)
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"], identification0, EnclaveTxnStatus.Ago.value),
        )

        identification1 = self.wardens[1]["identification"]
        resp = process_onboard_txn(txn, identification1, self.conn)
        self.assertEqual(resp["status"], EnclaveTxnStatus.Ago.value)

        db_txn = cursor.execute(
            """
                SELECT wardens, status
                FROM enclave_onboard_txn
                WHERE block_hash=? AND transaction_hash=? AND batch=?
            """,
            (txn["block_hash"], txn["txn_hash"], txn["batch"]),
        ).fetchone()
        cursor.close()
        self.assertEqual(db_txn[0], identification0)
        self.assertEqual(db_txn[1], EnclaveTxnStatus.Ago.value)

    def tearDown(self) -> None:
        self.conn.close()
