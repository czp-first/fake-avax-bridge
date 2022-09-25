# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 17:44:28
"""

from abc import ABCMeta, abstractmethod
import json


class Crypto(metaclass=ABCMeta):

    def __init__(self, identification: str, db_conn) -> None:
        cursor = db_conn.cursor()
        cursor.execute("SELECT credential FROM warden WHERE identification=?", (identification,))
        row = cursor.fetchone()
        if not row:
            raise ValueError("uknown identification")
        self.credential = json.loads(row[0])

    @abstractmethod
    def encrypt(self, plaintext: str):
        ...

    @abstractmethod
    def decrypt(self, ciphertext: str):
        ...
