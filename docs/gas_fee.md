![gas-fee](https://github.com/czp-first/fake-avax-bridge/blob/master/docs/imgs/gas-fee.png)



- Gas Price：baseFee + maxPriority
- Gas Limit：gas
- Gas Fees
  - Base：<font color='red'>???链上自己计算，无法指定，只能读取</font>
  - Max：实际交易的最大的gas price
  - Max Priority：给矿工的小费的价格的最大值



# web3.py



## maxFee > baseFee + maxPriority

```python
txn = dict(
    nonce=w3.eth.get_transaction_count(account),
    maxFeePerGas=Web3.toWei(1508, 'gwei'),
    maxPriorityFeePerGas=8000000000,
    gas=52000,
    to='',
    value=Web3.toWei(0.01, 'ether'),
    data=b'',
    type=2,  # (optional) the type is now implicitly set based on appropriate transaction params
    chainId=w3.eth.chain_id,
)
```



![image-20220716150054362](/Users/rey/Library/Application Support/typora-user-images/image-20220716150054362.png)



## maxFee < baseFee + maxPriority

```python
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
```



![image-20220716150249081](/Users/rey/Library/Application Support/typora-user-images/image-20220716150249081.png)





# geth



```go
type DynamicFeeTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList AccessList

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}
```

