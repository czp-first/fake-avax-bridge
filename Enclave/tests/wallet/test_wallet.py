# -*- coding: UTF-8 -*-
"""
@Summary : 测试 Wallet类
@Author  : Rey
@Time    : 2022-10-05 15:53:55
@Run     : python -m unittest tests/wallet/test_wallet.py
"""

import unittest

from wallet import Wallet


class TestWallet(unittest.TestCase):
    """test wallet.Wallet"""

    def test_1(self):
        mnemonic = "arch video pool whisper margin clip burger art code arm owner ugly ten trick cabbage cattle winter another negative middle argue harbor solution raw"
        wallet = Wallet(mnemonic=mnemonic)
        self.assertEqual(wallet.mnemonic, mnemonic)

        account1 = wallet.get_account(account=3)
        self.assertEqual(account1[0], "0xa91cd90E3d58E4EFCc9D1BF86C6767256354578B")
        account256 = wallet.get_account(account=256)
        self.assertEqual(account256[0], "0x6541765e494B07bd44fcE6D08b1AE8e0ac805F32")
