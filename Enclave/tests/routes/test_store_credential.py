# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 22:24:00
@Run     : python -m unittest tests/routes/test_store_credential.py
"""

from itertools import combinations
import json
import sqlite3
import unittest
from uuid import uuid4

from cryptography.fernet import Fernet

from crypto.adapter import get_crypto_obj
from routes import store_credential
from shamir import combine


class TestStoreInfo(unittest.TestCase):

    def setUp(self) -> None:
        self.conn = sqlite3.connect(":memory:")
        cursor = self.conn.cursor()
        with open("db/schema.sql") as f:
            cursor.executescript(f.read())
        cursor.close()

    def test_local(self):
        crypto_way = "local"
        wardens = [
            {
                "identification": str(uuid4()),
                "credential": dict(key=Fernet.generate_key().decode("utf-8")),
                "type": crypto_way,
            }
            for _ in range(3)
        ]
        share_version = 0
        wardens_resp = store_credential(
            share_version=share_version,
            threshold=2,
            wardens=wardens,
            from_chain_id=2,
            to_chain_id=256,
            db_conn=self.conn,
        )

        cursor = self.conn.cursor()
        for warden in wardens:
            cursor.execute("SELECT credential FROM warden WHERE identification=?", (warden["identification"],))
            self.assertDictEqual(warden["credential"], json.loads(cursor.fetchone()[0]))

        query_config_sql = "SELECT value FROM config WHERE key=?"
        cursor.execute(query_config_sql, ("crypto_way",))
        self.assertEqual(cursor.fetchone()[0], crypto_way)

        cursor.execute(query_config_sql, ("share_version",))
        self.assertEqual(cursor.fetchone()[0], str(share_version))

        shares = [
            get_crypto_obj(i["identification"], self.conn).decrypt(i["share"])
            for i in wardens_resp["encrypt_shares"]
        ]
        mnemonics = set()
        for item in combinations(shares, 2):
            mnemonics.add(combine(",".join(item)))
        assert len(mnemonics) == 1

    def tearDown(self) -> None:
        self.conn.close()
