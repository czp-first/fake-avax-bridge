# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-16 14:18:09
"""

import time

from web3 import Web3, HTTPProvider
from web3.exceptions import TransactionNotFound

w3 = Web3(HTTPProvider("https://ropsten.infura.io/v3/"))

account = ""
private_key = ""

print(w3.eth.chain_id)
txn = dict(
    nonce=w3.eth.get_transaction_count(account),
    gasPrice=Web3.toWei(10, "gwei"),
    gas=52000,
    to='',
    value=Web3.toWei(0.01, 'ether'),
    data=b'',
    chainId=w3.eth.chain_id,
)

signed_tx = w3.eth.account.sign_transaction(txn, private_key)

#send transaction
tx_hash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)

#get transaction hash
print(w3.toHex(tx_hash))

transaction_hash = Web3.toHex(tx_hash)

while True:
    try:
        receipt = w3.eth.get_transaction_receipt(transaction_hash)
        break
    except TransactionNotFound:
        print("not yet")
        time.sleep(5)
        continue
