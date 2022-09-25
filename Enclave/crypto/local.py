# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 17:48:33
"""

from cryptography.fernet import Fernet

from crypto.base import Crypto


class LocalCrypto(Crypto):

    def encrypt(self, plaintext: str):
        f = Fernet(self.credential["key"].encode("utf-8"))
        return f.encrypt(plaintext.encode("utf-8")).decode("utf-8")

    def decrypt(self, ciphertext: str):
        f = Fernet(self.credential["key"].encode("utf-8"))
        return f.decrypt(ciphertext.encode("utf-8")).decode("utf-8")
