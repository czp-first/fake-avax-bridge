# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 22:24:00
"""

from itertools import combinations
import json
import os
import unittest
from uuid import uuid4

from cryptography.fernet import Fernet

from crypto.adapter import get_crypto_obj
from routes import store_info
from settings import enclave_settings
from shamir import combine


class TestStoreInfo(unittest.TestCase):

    def test_local(self):
        wardens_info = [
            {
                "identification": str(uuid4()),
                "key": Fernet.generate_key().decode("utf-8")
            }
            for _ in range(3)
        ]
        wardens_gift = store_info(
            share_version=0,
            threshold=2,
            wardens_info=wardens_info
        )

        for warden_info in wardens_info:
            with open(f"{enclave_settings.warden_path}/{warden_info['identification']}.warden", 'r') as f:
                assert warden_info == json.load(f)

        shares = [
            get_crypto_obj(i["identification"]).decrypt(i["share"])
            for i in wardens_gift
        ]
        mnemonics = set()
        for item in combinations(shares, 2):
            mnemonics.add(combine(",".join(item)))
        assert len(mnemonics) == 1

    def tearDown(self) -> None:
        wardens = os.listdir(enclave_settings.warden_path)
        for warden in wardens:
            os.remove(f"{enclave_settings.warden_path}/{warden}")
