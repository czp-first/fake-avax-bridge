# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 22:18:15
"""

import unittest

from crypto.kms import KmsCrypto


class TestKmsCrypto(unittest.TestCase):

    @unittest.skip("no kms resource")
    def test_1(self):
        obj = KmsCrypto("8186de91-c28f-4741-a69d-62b740bc1929")
        plaintext = "immsg"
        print(obj.encrypt(plaintext))
