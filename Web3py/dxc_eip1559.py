# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-16 14:18:09
"""

from web3 import Web3, HTTPProvider

w3 = Web3(HTTPProvider("https://testnet-http.dxchain.com"))

account = ""
private_key = ""

print(w3.eth.chain_id)
txn = dict(
    nonce=w3.eth.get_transaction_count(account),
    maxFeePerGas=Web3.toWei(8, 'gwei'),
    maxPriorityFeePerGas=8000000000,
    gas=52000,
    to='',
    value=Web3.toWei(0.01, 'ether'),
    data=b'',
    type=2,  # (optional) the type is now implicitly set based on appropriate transaction params
    chainId=w3.eth.chain_id,
)

signed_tx = w3.eth.account.sign_transaction(txn, private_key)
print(Web3.toHex(signed_tx.rawTransaction))
#send transaction
# tx_hash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)

#get transaction hash
# print(w3.toHex(tx_hash))