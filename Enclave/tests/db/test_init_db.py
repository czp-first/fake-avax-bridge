# -*- coding: UTF-8 -*-
"""
@Summary : 测试初始化数据库
@Author  : Rey
@Time    : 2022-09-30 21:10:48
@Run     : python -m unittest tests/db/test_init_db.py
"""

import unittest

from db.init_db import init_db


class TestInitDb(unittest.TestCase):
    """test db.init_db.init_db"""
    def test_1(self):
        db_path = ":memory:"
        conn = init_db(db_path)
        conn.execute("SELECT * FROM enclave_onboard_txn")
        conn.execute("SELECT * FROM enclave_offboard_txn")
        conn.execute("SELECT * FROM warden")
        conn.execute("SELECT * FROM config")
        conn.close()
