# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-09-25 22:19:48
@Run     : python -m unittest tests/crypto/test_adapter.py
"""

import json
import sqlite3
import unittest

from crypto.adapter import get_crypto_obj


class TestGetCryptoObj(unittest.TestCase):

    def setUp(self) -> None:
        self.conn = sqlite3.connect(":memory:")
        cursor = self.conn.cursor()
        with open("db/schema.sql") as f:
            cursor.executescript(f.read())
        self.identification = "warden1"
        self.credential = json.dumps(dict(key="vHjfA2zvVSAY1FPUZRqnoWyJC4zeRJKU_3aBCUqrm8g="))
        cursor.execute("INSERT INTO warden(identification, credential)VALUES(?, ?)", (self.identification, self.credential))
        cursor.close()

    def test_local(self):
        cursor = self.conn.cursor()
        cursor.execute("INSERT INTO config(key, value)VALUES(?, ?)", ("crypto_way", "local"))
        cursor.close()

        obj = get_crypto_obj(self.identification, self.conn)
        plaintext = "immsg"
        ciphertext = obj.encrypt(plaintext)

        result = obj.decrypt(ciphertext)
        assert plaintext == result

    @unittest.skip("todo")
    def test_kms(self):
        ...

    def test_unknow_type(self):
        cursor = self.conn.cursor()
        cursor.execute("INSERT INTO config(key, value)VALUES(?, ?)", ("crypto_way", "ali"))
        cursor.close()
        self.assertRaises(TypeError, get_crypto_obj, self.identification, self.conn)

    def tearDown(self) -> None:
        self.conn.close()
