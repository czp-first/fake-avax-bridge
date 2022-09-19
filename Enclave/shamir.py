# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 23:05:25
"""

import os

from settings import enclave_settings


def split(parts, threshold, secret):
    path = os.path.join(os.getcwd(), enclave_settings.shamir_path)
    f = os.popen(
        f'{path} split -p {parts} -t {threshold} -s "{secret}"'
    )
    res = f.read()
    f.close()

    return res.strip().split('\n')


def combine(shares):
    path = os.path.join(os.getcwd(), enclave_settings.shamir_path)
    f = os.popen(
        f'{path} combine -s {shares}'
    )
    res = f.read()
    f.close()
    return res.strip()
