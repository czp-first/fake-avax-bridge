# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-08-03 17:40:57
"""

from web3 import Web3, HTTPProvider

from abi import (
    ERC20_TOKEN_ABI,
)


# eth -> heco
sender = ""
private_key = ""
receiver = ""
weth_contract_address = ""
dx_contract_address = ""
amount = Web3.toWei(0.1, "ether")


def transfer(contract_address: str):
    w3 = Web3(HTTPProvider("https://ropsten.infura.io/v3/"))

    w3.eth.default_account = sender

    weth_contract = w3.eth.contract(
        address=w3.toChecksumAddress(contract_address),
        abi=ERC20_TOKEN_ABI
    )

    txn = weth_contract.functions.transfer(
        Web3.toChecksumAddress(receiver), amount
    ).buildTransaction({
        "gasPrice": w3.toWei(8, "gwei"),
        "gas": 210000,
        "chainId": w3.eth.chain_id,
        "nonce": w3.eth.get_transaction_count(sender)
    })
    signed_txn = w3.eth.account.signTransaction(
        txn, private_key=private_key
    )
    txn_hash = w3.eth.sendRawTransaction(signed_txn.rawTransaction)

    print(w3.toHex(txn_hash))


if __name__ == '__main__':
    transfer(dx_contract_address)
