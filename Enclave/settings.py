import os

from pydantic import BaseSettings


BASE_DIR = os.path.dirname(os.path.abspath(__file__))


class EnclaveSettings(BaseSettings):

    work_dir: str
    shamir_path: str
    warden_path: str
    fee_address: str
    # crypto_way: str

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

