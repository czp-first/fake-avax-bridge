# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:30:27
"""

from crypto.base import Crypto
# from crypto.kms import KmsCrypto
from crypto.local import LocalCrypto


def get_crypto_obj(identification: str, db_conn) -> Crypto:
    cursor = db_conn.cursor()
    cursor.execute("SELECT value FROM config WHERE key=?", ("crypto_way",))
    crypto_way = cursor.fetchone()[0]
    cursor.close()
    if crypto_way == "local":
        return LocalCrypto(identification, db_conn)
    # elif crypto_way == "kms":
        # return KmsCrypto(identification, db_conn)
    raise TypeError(f'Unknown crypto way: {crypto_way}')
