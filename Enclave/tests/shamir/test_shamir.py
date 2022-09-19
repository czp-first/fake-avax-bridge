# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 22:57:58
"""

from itertools import combinations
import unittest

from hdwallet.utils import generate_mnemonic

from shamir import split, combine


class TestShamir(unittest.TestCase):
    def test1(self):
        mnemonic = generate_mnemonic()
        shares = split(3, 2, mnemonic)
        for item in combinations(shares, 2):
             assert mnemonic == combine(",".join(item))
