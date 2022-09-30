# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-09-30 11:11:54
"""

from pydantic import BaseSettings


class EnclaveSettings(BaseSettings):

    shamir_path: str
    fee_address: str

    default_gas: int
    max_fee_per_gas: int
    db_path: str = "db/enclave.db"

    class Config:
        env_file = '.env'
        case_sensitive = False
        env_file_encoding = 'utf-8'


enclave_settings = EnclaveSettings()


if __name__ == '__main__':
    print(enclave_settings)
