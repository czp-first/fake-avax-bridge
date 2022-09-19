# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-12 18:35:49
"""

import json
import os

from crypto.adapter import get_crypto_obj
from settings import enclave_settings
import wallet
from txns import (
    pending_txns,
    ago_onboard_txns,
    comfirm_info,
    threshold,
    ago_offboard_txns,
    pending_offboard_txns,
    comfirming_offboard_txns,
)


def store_credential(share_version: int, threshold: int, wardens: list, onboard_chain_id: int, offboard_chain_id: int):
    # TODO(Rey): return account address
    for warden in wardens:
        with open(f'{enclave_settings.warden_path}/{warden["identification"]}.warden', 'w') as f:
            f.write(json.dumps(warden))

    wallet_info = wallet.init_wallet(len(wardens), threshold, onboard_chain_id, offboard_chain_id)
    if share_version == 0:
        wardens_path = os.listdir(enclave_settings.warden_path)
        encrypt_shares = [
            {
                "identification": item.split(".")[0],
                "share": get_crypto_obj(item.split(".")[0]).encrypt(wallet["shares"][index])
            }
            for index, item in enumerate(wardens_path)
        ]
        return dict(
            encrypt_shares=encrypt_shares,
            onboard_account_address=wallet_info["onboard_account_address"],
            offboard_account_address=wallet_info["offboard_account_address"],
        )
    # TODO: shareVersion>0
    else:
        ...


def recover(wardens_info):
    shares = (
        get_crypto_obj(warden["identification"]).decrypt(warden["share"])
        for warden in wardens_info
    )
    return wallet.recover_wallet(shares)


def process_onboard_txn(txn, identification):

    # print(comfirm_info)
    # print(ago_onboard_txns)
    # print(pending_txns)

    unique_txn = f'{txn["block_hash"]}{txn["txn_hash"]}{txn["batch"]}'

    # txn:{txn_hash: '', block_hash: ''}
    if unique_txn in pending_txns:
        return {'status': 'pending'}
    if unique_txn in ago_onboard_txns:
        return {'status': 'ago'}

    if unique_txn in comfirm_info:
        if identification in comfirm_info[unique_txn]:
            return {'status': 'wait'}
        comfirm_info[unique_txn].append(identification)
        if len(comfirm_info[unique_txn]) >= threshold:
            pending_txns.append(unique_txn)
            wardens = comfirm_info.pop(unique_txn)
            return {
                'status': 'ready',
                'wardens': wardens[:threshold],
            }
        return {'status': 'wait'}

    comfirm_info[unique_txn] = [identification]

    return {'status': 'wait'}


def process_offboard_txn(txn, identification):
    # print(comfirming_offboard_txns)
    # print(ago_offboard_txns)
    # print(pending_offboard_txns)

    unique_txn = f'{txn["block_hash"]}{txn["txn_hash"]}{txn["batch"]}'
    if unique_txn in pending_offboard_txns:
        return {'status': 'pending'}
    if unique_txn in ago_offboard_txns:
        return {'status': 'ago'}

    if unique_txn in comfirming_offboard_txns:
        if identification in comfirming_offboard_txns[unique_txn]:
            return {'status': 'wait'}
        comfirming_offboard_txns[unique_txn].append(identification)
        if len(comfirming_offboard_txns[unique_txn]) >= threshold:
            pending_offboard_txns.append(unique_txn)
            return {
                'status': 'ready',
                'wardens': comfirming_offboard_txns[unique_txn][:threshold]
            }
        return {'status': 'wait'}

    comfirming_offboard_txns[unique_txn] = [identification]
    return {'status': 'wait'}


def sign_onboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce, origin_txn, fee):
    decrypt_shares = []
    for warden_share in warden_shares:
        decrypt_share = get_crypto_obj(warden_share["identification"]).decrypt(warden_share["encrypt_share"])
        decrypt_shares.append(decrypt_share)
    mnemonic = wallet.recover_wallet(tuple(decrypt_shares))

    params = dict(
        is_eip1559=is_eip1559,
        mnemonic=mnemonic,
        chain_id=chain_id,
        contract_addr=contract_addr,
        amount=amount,
        gas_price=gas_price,
        account_addr=account_addr,
        nonce=nonce,
        origin_txn=origin_txn,
        fee=fee,
    )

    return {
        # web3 需要0x开头, go-ethereum不需要
        'txn': wallet.sign_onboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
    }


def sign_offboard_txn(is_eip1559, warden_shares, chain_id, contract_addr, amount, gas_price, account_addr, nonce):
    decrypt_shares = []
    for warden_share in warden_shares:
        decrypt_share = get_crypto_obj(warden_share["identification"]).decrypt(warden_share["encrypt_share"])
        decrypt_shares.append(decrypt_share)
    mnemonic = wallet.recover_wallet(tuple(decrypt_shares))

    params = dict(
        is_eip1559=is_eip1559,
        mnemonic=mnemonic,
        chain_id=chain_id,
        contract_addr=contract_addr,
        amount=amount,
        gas_price=gas_price,
        account_addr=account_addr,
        nonce=nonce
    )

    return {
        'txn': wallet.sign_offboard_transaction(**params)[2:],
        'nonce': nonce,
        'gas_price': gas_price,
        "is_eip1559": is_eip1559,
    }


ROUTES_MAP = {
    'storeCredential': store_credential,
    'recove': recover,

    'onboardTxn': process_onboard_txn,
    'signOnboardTxn': sign_onboard_txn,

    'offboardTxn': process_offboard_txn,
    'signOffboardTxn': sign_offboard_txn,
}
