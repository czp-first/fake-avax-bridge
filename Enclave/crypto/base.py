# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 17:44:28
"""

from abc import ABCMeta, abstractmethod
import json

from settings import enclave_settings


class Crypto(metaclass=ABCMeta):

    def __init__(self, identification: str) -> None:
        with open(f'{enclave_settings.warden_path}/{identification}.warden', 'r') as f:
            self.warden_info = json.load(f)

    @abstractmethod
    def encrypt(self, plaintext: str):
        ...

    @abstractmethod
    def decrypt(self, ciphertext: str):
        ...
