# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:03:09
@Run     : python -m unittest tests/crypto/test_local.py
"""

import json
import sqlite3
import unittest

from crypto.local import LocalCrypto


class TestLocalCrypto(unittest.TestCase):

    def setUp(self) -> None:
        self.conn = sqlite3.connect(":memory:")
        cursor = self.conn.cursor()
        with open("db/schema.sql") as f:
            cursor.executescript(f.read())
        self.identification = "warden1"
        self.credential = json.dumps(dict(key="vHjfA2zvVSAY1FPUZRqnoWyJC4zeRJKU_3aBCUqrm8g="))
        cursor.execute("INSERT INTO warden(identification, credential)VALUES(?, ?)", (self.identification, self.credential))
        cursor.close()

    def test_1(self):
        obj = LocalCrypto(self.identification, self.conn)
        plaintext = "immsg"
        ciphertext = obj.encrypt(plaintext)

        result = obj.decrypt(ciphertext)
        assert plaintext == result

    def tearDown(self) -> None:
        self.conn.close()
