# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-10-05 15:05:11
"""

from decimal import Decimal
from typing import Tuple

from web3 import Web3
from web3.auto import w3
from hdwallet import BIP44HDWallet
from hdwallet.cryptocurrencies import EthereumMainnet
from hdwallet.derivations import BIP44Derivation
from hdwallet.utils import generate_mnemonic

from abi import DXCHAIN_TOKEN_ABI, ERC20_TOKEN_ABI
from settings import enclave_settings
import shamir


class Wallet:
    def __init__(self, mnemonic: str) -> None:
        if not mnemonic:
            self.mnemonic = generate_mnemonic()
        else:
            self.mnemonic = mnemonic
        self._wallet = None

    @property
    def wallet(self) -> BIP44HDWallet:
        if not self._wallet:
            passphrase = None
            bip44_hdwallet = BIP44HDWallet(cryptocurrency=EthereumMainnet)
            bip44_hdwallet.from_mnemonic(mnemonic=self.mnemonic, language='english', passphrase=passphrase)
            bip44_hdwallet.clean_derivation()
            self._wallet = bip44_hdwallet
        return self._wallet

    def get_account(self, account, address=0) -> Tuple[str, str]:
        bip44_derivation = BIP44Derivation(cryptocurrency=EthereumMainnet, account=account, change=False, address=address)
        self.wallet.from_path(path=bip44_derivation)
        address = self.wallet.address()
        private_key = f'0x{self.wallet.private_key()}'
        self.wallet.clean_derivation()
        return address, private_key


def init_wallet(parts: int, threshold: int, onboard_chain_id: int, offboard_chain_id: int):
    """
    :param parts: the number of shares
    :param threshold: the number of shares to recover
    """
    mnemonic = ""
    wallet = Wallet(mnemonic=mnemonic)
    onboard_account_address, _ = wallet.get_account(account=onboard_chain_id)
    offboard_account_address, _ = wallet.get_account(account=offboard_chain_id)

    return dict(
        shares=shamir.split(parts, threshold, wallet.mnemonic),
        onboard_account_address=onboard_account_address,
        offboard_account_address=offboard_account_address,
    )


def sign_onboard_transaction(
    is_eip1559, mnemonic, chain_id, contract_addr, amount, gas_price, account_addr, nonce, origin_txn_hash, fee
):

    _, private_key = Wallet(mnemonic=mnemonic).get_account(account=chain_id)

    dx_contract = w3.eth.contract(
        address=Web3.toChecksumAddress(contract_addr),
        abi=DXCHAIN_TOKEN_ABI
    )

    transaction_params = dict(
        gas=enclave_settings.default_gas,
        chainId=int(chain_id),
        nonce=int(nonce),
    )

    if is_eip1559:
        transaction_params["maxFeePerGas"] = Web3.toWei(enclave_settings.max_fee_per_gas, 'gwei')
        transaction_params["maxPriorityFeePerGas"] = Web3.toWei(Decimal(gas_price), 'wei')
        transaction_params["type"] = 2
    else:
        transaction_params["gasPrice"] = Web3.toWei(Decimal(gas_price), 'wei')

    transaction = dx_contract.functions.mint(
        account_addr,
        Web3.toWei(Decimal(amount), 'wei'),
        enclave_settings.fee_address,
        Web3.toWei(Decimal(fee), 'wei'),
        origin_txn_hash
    ).buildTransaction(transaction_params)

    sign_transaction = w3.eth.account.sign_transaction(
        transaction_dict=transaction,
        private_key=private_key
    )

    return Web3.toHex(sign_transaction.rawTransaction)


def sign_offboard_transaction(is_eip1559, mnemonic, chain_id, contract_addr, amount, gas_price, account_addr, nonce):

    _, private_key = Wallet(mnemonic=mnemonic).get_account(account=chain_id)

    contract = w3.eth.contract(
        address=Web3.toChecksumAddress(contract_addr),
        abi=ERC20_TOKEN_ABI
    )

    transaction_params = dict(
        gas=enclave_settings.default_gas,
        chainId=int(chain_id),
        nonce=int(nonce),
    )

    if is_eip1559:
        transaction_params["maxFeePerGas"] = Web3.toWei(enclave_settings.max_fee_per_gas, 'gwei')
        transaction_params["maxPriorityFeePerGas"] = Web3.toWei(Decimal(gas_price), 'wei')
        transaction_params["type"] = 2
    else:
        transaction_params["gasPrice"] = Web3.toWei(Decimal(gas_price), 'wei')

    txn = contract.functions.transfer(
        account_addr, Web3.toWei(Decimal(amount), 'wei')
    ).buildTransaction(transaction_params)

    sign_txn = w3.eth.account.sign_transaction(
        transaction_dict=txn,
        private_key=private_key
    )
    return Web3.toHex(sign_txn.rawTransaction)


def recover_wallet(shares):
    mnemonic = shamir.combine(','.join(shares))
    passphrase = None
    bip44_hdwallet = BIP44HDWallet(cryptocurrency=EthereumMainnet)
    bip44_hdwallet.from_mnemonic(mnemonic=mnemonic, language='english', passphrase=passphrase)
    bip44_hdwallet.clean_derivation()
    return mnemonic
