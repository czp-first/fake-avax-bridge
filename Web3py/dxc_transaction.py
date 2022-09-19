# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-16 14:18:09
"""

from web3 import Web3, HTTPProvider

w3 = Web3(HTTPProvider("https://testnet-http.dxchain.com"))

transaction = w3.eth.get_transaction("")
print(transaction)

print(w3.eth.fee_history(4, 2207288, None))