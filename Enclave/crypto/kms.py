# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:14:54
"""

import base64

import boto3

from crypto.base import Crypto


class KmsCrypto(Crypto):

    def encrypt(self, plaintext: str):
        kms = boto3.client(
            service_name='kms',
            region_name=self.warden_info['region'],
            aws_access_key_id=self.warden_info['accessKeyId'],
            aws_secret_access_key=self.warden_info['secretAccessKey']
        )

        stuff = kms.encrypt(KeyId=self.warden_info['accessKeyId'], Plaintext=plaintext)
        binary_encrypted = stuff[u'CiphertextBlob']
        encrypted_password = base64.b64encode(binary_encrypted)
        return encrypted_password.decode()

    def decrypt(self, ciphertext: str):
        kms = boto3.client(
            service_name='kms',
            region_name=self.warden_info['region'],
            aws_access_key_id=self.warden_info['accessKeyId'],
            aws_secret_access_key=self.warden_info['secretAccessKey']
        )
        binary_data = base64.b64decode(ciphertext)
        meta = kms.decrypt(CiphertextBlob=binary_data)
        plaintext = meta[u'Plaintext']
        return plaintext.decode()
