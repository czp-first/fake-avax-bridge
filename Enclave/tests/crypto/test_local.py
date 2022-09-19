# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:03:09
"""

import unittest

from crypto.local import LocalCrypto


class TestLocalCrypto(unittest.TestCase):

    @unittest.skip("skip tmp")
    def test_1(self):
        obj = LocalCrypto()
        plaintext = "immsg"
        ciphertext = obj.encrypt(plaintext)

        result = obj.decrypt(ciphertext)
        assert plaintext == result
