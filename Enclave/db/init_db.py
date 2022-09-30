# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-09-26 11:29:08
"""

import sqlite3

from settings import enclave_settings


def init_db():
    conn = sqlite3.connect(enclave_settings.db_path)
    cursor = conn.cursor()

    with open("db/schema.sql") as f:
        cursor.executescript(f.read())

    cursor.close()
    conn.close()


if __name__ == '__main__':
    init_db()
