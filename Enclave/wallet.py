from decimal import Decimal

from web3 import Web3
from web3.auto import w3
from hdwallet import BIP44HDWallet
from hdwallet.cryptocurrencies import EthereumMainnet
from hdwallet.derivations import BIP44Derivation
from hdwallet.utils import generate_mnemonic

from abi import DXCHAIN_TOKEN_ABI, ERC20_TOKEN_ABI
from settings import enclave_settings
import shamir


def get_wallet(mnemonic: str):
    passphrase = None
    bip44_hdwallet = BIP44HDWallet(cryptocurrency=EthereumMainnet)
    bip44_hdwallet.from_mnemonic(mnemonic=mnemonic, language='english', passphrase=passphrase)
    bip44_hdwallet.clean_derivation()
    return bip44_hdwallet


def init_wallet(parts: int, threshold: int, onboard_chain_id: int, offboard_chain_id: int):
    """
    :param parts: the number of shares
    :param threshold: the number of shares to recover
    """
    mnemonic = generate_mnemonic()

    bip44_hdwallet = get_wallet(mnemonic)

    onboard_bip44_derivation = BIP44Derivation(cryptocurrency=EthereumMainnet, account=onboard_chain_id, change=False, address=0)
    bip44_hdwallet.from_path(path=onboard_bip44_derivation)
    onboard_account_address = bip44_hdwallet.address()
    bip44_hdwallet.clean_derivation()

    offboard_bip44_derivation = BIP44Derivation(cryptocurrency=EthereumMainnet, account=offboard_chain_id, change=False, address=0)
    bip44_hdwallet.from_path(path=offboard_bip44_derivation)
    offboard_account_address = bip44_hdwallet.address()
    bip44_hdwallet.clean_derivation()

    return dict(
        shares=shamir.split(parts, threshold, mnemonic),
        onboard_account_address=onboard_account_address,
        offboard_account_address=offboard_account_address,
    )


def sign_onboard_transaction(
    is_eip1559, mnemonic, chain_id, contract_addr, amount, gas_price, account_addr, nonce, origin_txn, fee
):
    """
    :param amount: unit(wei)
    :param gas_price: unit(wei)
    """

    bip44_hdwallet = get_wallet(mnemonic)

    bip44_derivation = BIP44Derivation(cryptocurrency=EthereumMainnet, account=chain_id, change=False, address=0)
    bip44_hdwallet.from_path(path=bip44_derivation)
    # sender_addr = bip44_hdwallet.address()
    private_key = f'0x{bip44_hdwallet.private_key()}'
    bip44_hdwallet.clean_derivation()

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
        origin_txn
    ).buildTransaction(transaction_params)

    sign_transaction = w3.eth.account.sign_transaction(
        transaction_dict=transaction,
        private_key=private_key
    )

    return Web3.toHex(sign_transaction.rawTransaction)


def sign_offboard_transaction(is_eip1559, mnemonic, chain_id, contract_addr, amount, gas_price, account_addr, nonce):

    passphrase = None
    bip44_hdwallet = BIP44HDWallet(cryptocurrency=EthereumMainnet)
    bip44_hdwallet.from_mnemonic(mnemonic=mnemonic, language='english', passphrase=passphrase)
    bip44_hdwallet.clean_derivation()

    bip44_derivation = BIP44Derivation(
        cryptocurrency=EthereumMainnet, account=chain_id, change=False, address=0
    )
    bip44_hdwallet.from_path(path=bip44_derivation)
    # sender_addr = bip44_hdwallet.address()
    private_key = f'0x{bip44_hdwallet.private_key()}'
    bip44_hdwallet.clean_derivation()

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
    print(mnemonic)
    passphrase = None
    bip44_hdwallet = BIP44HDWallet(cryptocurrency=EthereumMainnet)
    bip44_hdwallet.from_mnemonic(mnemonic=mnemonic, language='english', passphrase=passphrase)
    bip44_hdwallet.clean_derivation()
    return mnemonic

