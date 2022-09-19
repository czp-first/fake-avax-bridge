# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:30:27
"""

from crypto.base import Crypto
from crypto.kms import KmsCrypto
from crypto.local import LocalCrypto
from settings import enclave_settings


def get_crypto_obj(identification: str) -> Crypto:
    if enclave_settings.crypto_way == "kms":
        return KmsCrypto(identification)
    elif enclave_settings.crypto_way == "local":
        return LocalCrypto(identification)
    raise TypeError(f'Unknown crypto way: {enclave_settings.crypto_way}')
