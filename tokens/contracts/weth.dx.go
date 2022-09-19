// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weth

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WethMetaData contains all meta data concerning the Weth contract.
var WethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"AddSupportedChainId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyIncrement\",\"type\":\"uint256\"}],\"name\":\"AddSwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBridgeRoleAddress\",\"type\":\"address\"}],\"name\":\"MigrateBridgeRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originTxId\",\"type\":\"bytes32\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyDecrement\",\"type\":\"uint256\"}],\"name\":\"RemoveSwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"addSupportedChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supplyIncrement\",\"type\":\"uint256\"}],\"name\":\"addSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBridgeRoleAddress\",\"type\":\"address\"}],\"name\":\"migrateBridgeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"originTxId\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supplyDecrement\",\"type\":\"uint256\"}],\"name\":\"removeSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"swapSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WethABI is the input ABI used to generate the binding from.
// Deprecated: Use WethMetaData.ABI instead.
var WethABI = WethMetaData.ABI

// Weth is an auto generated Go binding around an Ethereum contract.
type Weth struct {
	WethCaller     // Read-only binding to the contract
	WethTransactor // Write-only binding to the contract
	WethFilterer   // Log filterer for contract events
}

// WethCaller is an auto generated read-only Go binding around an Ethereum contract.
type WethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WethSession struct {
	Contract     *Weth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WethCallerSession struct {
	Contract *WethCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WethTransactorSession struct {
	Contract     *WethTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WethRaw is an auto generated low-level Go binding around an Ethereum contract.
type WethRaw struct {
	Contract *Weth // Generic contract binding to access the raw methods on
}

// WethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WethCallerRaw struct {
	Contract *WethCaller // Generic read-only contract binding to access the raw methods on
}

// WethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WethTransactorRaw struct {
	Contract *WethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeth creates a new instance of Weth, bound to a specific deployed contract.
func NewWeth(address common.Address, backend bind.ContractBackend) (*Weth, error) {
	contract, err := bindWeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Weth{WethCaller: WethCaller{contract: contract}, WethTransactor: WethTransactor{contract: contract}, WethFilterer: WethFilterer{contract: contract}}, nil
}

// NewWethCaller creates a new read-only instance of Weth, bound to a specific deployed contract.
func NewWethCaller(address common.Address, caller bind.ContractCaller) (*WethCaller, error) {
	contract, err := bindWeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WethCaller{contract: contract}, nil
}

// NewWethTransactor creates a new write-only instance of Weth, bound to a specific deployed contract.
func NewWethTransactor(address common.Address, transactor bind.ContractTransactor) (*WethTransactor, error) {
	contract, err := bindWeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WethTransactor{contract: contract}, nil
}

// NewWethFilterer creates a new log filterer instance of Weth, bound to a specific deployed contract.
func NewWethFilterer(address common.Address, filterer bind.ContractFilterer) (*WethFilterer, error) {
	contract, err := bindWeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WethFilterer{contract: contract}, nil
}

// bindWeth binds a generic wrapper to an already deployed contract.
func bindWeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Weth *WethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Weth.Contract.WethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Weth *WethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Weth.Contract.WethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Weth *WethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Weth.Contract.WethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Weth *WethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Weth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Weth *WethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Weth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Weth *WethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Weth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Weth *WethCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Weth *WethSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Weth.Contract.Allowance(&_Weth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Weth *WethCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Weth.Contract.Allowance(&_Weth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Weth *WethCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Weth *WethSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Weth.Contract.BalanceOf(&_Weth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Weth *WethCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Weth.Contract.BalanceOf(&_Weth.CallOpts, account)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_Weth *WethCaller) ChainIds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "chainIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_Weth *WethSession) ChainIds(arg0 *big.Int) (bool, error) {
	return _Weth.Contract.ChainIds(&_Weth.CallOpts, arg0)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_Weth *WethCallerSession) ChainIds(arg0 *big.Int) (bool, error) {
	return _Weth.Contract.ChainIds(&_Weth.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Weth *WethCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Weth *WethSession) Decimals() (uint8, error) {
	return _Weth.Contract.Decimals(&_Weth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Weth *WethCallerSession) Decimals() (uint8, error) {
	return _Weth.Contract.Decimals(&_Weth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Weth *WethCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Weth *WethSession) Name() (string, error) {
	return _Weth.Contract.Name(&_Weth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Weth *WethCallerSession) Name() (string, error) {
	return _Weth.Contract.Name(&_Weth.CallOpts)
}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_Weth *WethCaller) SwapSupply(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "swapSupply", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_Weth *WethSession) SwapSupply(token common.Address) (*big.Int, error) {
	return _Weth.Contract.SwapSupply(&_Weth.CallOpts, token)
}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_Weth *WethCallerSession) SwapSupply(token common.Address) (*big.Int, error) {
	return _Weth.Contract.SwapSupply(&_Weth.CallOpts, token)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Weth *WethCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Weth *WethSession) Symbol() (string, error) {
	return _Weth.Contract.Symbol(&_Weth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Weth *WethCallerSession) Symbol() (string, error) {
	return _Weth.Contract.Symbol(&_Weth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Weth *WethCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Weth *WethSession) TotalSupply() (*big.Int, error) {
	return _Weth.Contract.TotalSupply(&_Weth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Weth *WethCallerSession) TotalSupply() (*big.Int, error) {
	return _Weth.Contract.TotalSupply(&_Weth.CallOpts)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_Weth *WethTransactor) AddSupportedChainId(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "addSupportedChainId", chainId)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_Weth *WethSession) AddSupportedChainId(chainId *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.AddSupportedChainId(&_Weth.TransactOpts, chainId)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_Weth *WethTransactorSession) AddSupportedChainId(chainId *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.AddSupportedChainId(&_Weth.TransactOpts, chainId)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_Weth *WethTransactor) AddSwapToken(opts *bind.TransactOpts, contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "addSwapToken", contractAddress, supplyIncrement)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_Weth *WethSession) AddSwapToken(contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.AddSwapToken(&_Weth.TransactOpts, contractAddress, supplyIncrement)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_Weth *WethTransactorSession) AddSwapToken(contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.AddSwapToken(&_Weth.TransactOpts, contractAddress, supplyIncrement)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Weth *WethTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Weth *WethSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Approve(&_Weth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Weth *WethTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Approve(&_Weth.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Weth *WethTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Weth *WethSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Burn(&_Weth.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Weth *WethTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Burn(&_Weth.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Weth *WethTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Weth *WethSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.BurnFrom(&_Weth.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Weth *WethTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.BurnFrom(&_Weth.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Weth *WethTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Weth *WethSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.DecreaseAllowance(&_Weth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Weth *WethTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.DecreaseAllowance(&_Weth.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Weth *WethTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Weth *WethSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.IncreaseAllowance(&_Weth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Weth *WethTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.IncreaseAllowance(&_Weth.TransactOpts, spender, addedValue)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_Weth *WethTransactor) MigrateBridgeRole(opts *bind.TransactOpts, newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "migrateBridgeRole", newBridgeRoleAddress)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_Weth *WethSession) MigrateBridgeRole(newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _Weth.Contract.MigrateBridgeRole(&_Weth.TransactOpts, newBridgeRoleAddress)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_Weth *WethTransactorSession) MigrateBridgeRole(newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _Weth.Contract.MigrateBridgeRole(&_Weth.TransactOpts, newBridgeRoleAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_Weth *WethTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "mint", to, amount, feeAddress, feeAmount, originTxId)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_Weth *WethSession) Mint(to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _Weth.Contract.Mint(&_Weth.TransactOpts, to, amount, feeAddress, feeAmount, originTxId)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_Weth *WethTransactorSession) Mint(to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _Weth.Contract.Mint(&_Weth.TransactOpts, to, amount, feeAddress, feeAmount, originTxId)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_Weth *WethTransactor) RemoveSwapToken(opts *bind.TransactOpts, contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "removeSwapToken", contractAddress, supplyDecrement)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_Weth *WethSession) RemoveSwapToken(contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.RemoveSwapToken(&_Weth.TransactOpts, contractAddress, supplyDecrement)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_Weth *WethTransactorSession) RemoveSwapToken(contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.RemoveSwapToken(&_Weth.TransactOpts, contractAddress, supplyDecrement)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Weth *WethTransactor) Swap(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "swap", token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Weth *WethSession) Swap(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Swap(&_Weth.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Weth *WethTransactorSession) Swap(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Swap(&_Weth.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Weth *WethTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Weth *WethSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Transfer(&_Weth.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Weth *WethTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Transfer(&_Weth.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Weth *WethTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Weth *WethSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.TransferFrom(&_Weth.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Weth *WethTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.TransferFrom(&_Weth.TransactOpts, sender, recipient, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_Weth *WethTransactor) Unwrap(opts *bind.TransactOpts, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Weth.contract.Transact(opts, "unwrap", amount, chainId)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_Weth *WethSession) Unwrap(amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Unwrap(&_Weth.TransactOpts, amount, chainId)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_Weth *WethTransactorSession) Unwrap(amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _Weth.Contract.Unwrap(&_Weth.TransactOpts, amount, chainId)
}

// WethAddSupportedChainIdIterator is returned from FilterAddSupportedChainId and is used to iterate over the raw logs and unpacked data for AddSupportedChainId events raised by the Weth contract.
type WethAddSupportedChainIdIterator struct {
	Event *WethAddSupportedChainId // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethAddSupportedChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddSupportedChainId)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethAddSupportedChainId)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethAddSupportedChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddSupportedChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddSupportedChainId represents a AddSupportedChainId event raised by the Weth contract.
type WethAddSupportedChainId struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddSupportedChainId is a free log retrieval operation binding the contract event 0x677e2d9a4ed9201aa86725fef875137fc53876e6b68036b974404762682bd122.
//
// Solidity: event AddSupportedChainId(uint256 chainId)
func (_Weth *WethFilterer) FilterAddSupportedChainId(opts *bind.FilterOpts) (*WethAddSupportedChainIdIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "AddSupportedChainId")
	if err != nil {
		return nil, err
	}
	return &WethAddSupportedChainIdIterator{contract: _Weth.contract, event: "AddSupportedChainId", logs: logs, sub: sub}, nil
}

// WatchAddSupportedChainId is a free log subscription operation binding the contract event 0x677e2d9a4ed9201aa86725fef875137fc53876e6b68036b974404762682bd122.
//
// Solidity: event AddSupportedChainId(uint256 chainId)
func (_Weth *WethFilterer) WatchAddSupportedChainId(opts *bind.WatchOpts, sink chan<- *WethAddSupportedChainId) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "AddSupportedChainId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddSupportedChainId)
				if err := _Weth.contract.UnpackLog(event, "AddSupportedChainId", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddSupportedChainId is a log parse operation binding the contract event 0x677e2d9a4ed9201aa86725fef875137fc53876e6b68036b974404762682bd122.
//
// Solidity: event AddSupportedChainId(uint256 chainId)
func (_Weth *WethFilterer) ParseAddSupportedChainId(log types.Log) (*WethAddSupportedChainId, error) {
	event := new(WethAddSupportedChainId)
	if err := _Weth.contract.UnpackLog(event, "AddSupportedChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethAddSwapTokenIterator is returned from FilterAddSwapToken and is used to iterate over the raw logs and unpacked data for AddSwapToken events raised by the Weth contract.
type WethAddSwapTokenIterator struct {
	Event *WethAddSwapToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethAddSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddSwapToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethAddSwapToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethAddSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddSwapToken represents a AddSwapToken event raised by the Weth contract.
type WethAddSwapToken struct {
	ContractAddress common.Address
	SupplyIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddSwapToken is a free log retrieval operation binding the contract event 0x3e4fdfb0f47da284fe8b5b3a7e5d10b211e323c9a0c144c421ae1d211873f853.
//
// Solidity: event AddSwapToken(address contractAddress, uint256 supplyIncrement)
func (_Weth *WethFilterer) FilterAddSwapToken(opts *bind.FilterOpts) (*WethAddSwapTokenIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "AddSwapToken")
	if err != nil {
		return nil, err
	}
	return &WethAddSwapTokenIterator{contract: _Weth.contract, event: "AddSwapToken", logs: logs, sub: sub}, nil
}

// WatchAddSwapToken is a free log subscription operation binding the contract event 0x3e4fdfb0f47da284fe8b5b3a7e5d10b211e323c9a0c144c421ae1d211873f853.
//
// Solidity: event AddSwapToken(address contractAddress, uint256 supplyIncrement)
func (_Weth *WethFilterer) WatchAddSwapToken(opts *bind.WatchOpts, sink chan<- *WethAddSwapToken) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "AddSwapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddSwapToken)
				if err := _Weth.contract.UnpackLog(event, "AddSwapToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddSwapToken is a log parse operation binding the contract event 0x3e4fdfb0f47da284fe8b5b3a7e5d10b211e323c9a0c144c421ae1d211873f853.
//
// Solidity: event AddSwapToken(address contractAddress, uint256 supplyIncrement)
func (_Weth *WethFilterer) ParseAddSwapToken(log types.Log) (*WethAddSwapToken, error) {
	event := new(WethAddSwapToken)
	if err := _Weth.contract.UnpackLog(event, "AddSwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Weth contract.
type WethApprovalIterator struct {
	Event *WethApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethApproval represents a Approval event raised by the Weth contract.
type WethApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Weth *WethFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WethApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Weth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WethApprovalIterator{contract: _Weth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Weth *WethFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WethApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Weth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethApproval)
				if err := _Weth.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Weth *WethFilterer) ParseApproval(log types.Log) (*WethApproval, error) {
	event := new(WethApproval)
	if err := _Weth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethMigrateBridgeRoleIterator is returned from FilterMigrateBridgeRole and is used to iterate over the raw logs and unpacked data for MigrateBridgeRole events raised by the Weth contract.
type WethMigrateBridgeRoleIterator struct {
	Event *WethMigrateBridgeRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethMigrateBridgeRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethMigrateBridgeRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethMigrateBridgeRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethMigrateBridgeRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethMigrateBridgeRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethMigrateBridgeRole represents a MigrateBridgeRole event raised by the Weth contract.
type WethMigrateBridgeRole struct {
	NewBridgeRoleAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMigrateBridgeRole is a free log retrieval operation binding the contract event 0x871b00a4e20f8436702d0174eb87d84d7cd1dd5c34d4bb1b4e75438b3398d512.
//
// Solidity: event MigrateBridgeRole(address newBridgeRoleAddress)
func (_Weth *WethFilterer) FilterMigrateBridgeRole(opts *bind.FilterOpts) (*WethMigrateBridgeRoleIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "MigrateBridgeRole")
	if err != nil {
		return nil, err
	}
	return &WethMigrateBridgeRoleIterator{contract: _Weth.contract, event: "MigrateBridgeRole", logs: logs, sub: sub}, nil
}

// WatchMigrateBridgeRole is a free log subscription operation binding the contract event 0x871b00a4e20f8436702d0174eb87d84d7cd1dd5c34d4bb1b4e75438b3398d512.
//
// Solidity: event MigrateBridgeRole(address newBridgeRoleAddress)
func (_Weth *WethFilterer) WatchMigrateBridgeRole(opts *bind.WatchOpts, sink chan<- *WethMigrateBridgeRole) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "MigrateBridgeRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethMigrateBridgeRole)
				if err := _Weth.contract.UnpackLog(event, "MigrateBridgeRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMigrateBridgeRole is a log parse operation binding the contract event 0x871b00a4e20f8436702d0174eb87d84d7cd1dd5c34d4bb1b4e75438b3398d512.
//
// Solidity: event MigrateBridgeRole(address newBridgeRoleAddress)
func (_Weth *WethFilterer) ParseMigrateBridgeRole(log types.Log) (*WethMigrateBridgeRole, error) {
	event := new(WethMigrateBridgeRole)
	if err := _Weth.contract.UnpackLog(event, "MigrateBridgeRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Weth contract.
type WethMintIterator struct {
	Event *WethMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethMint represents a Mint event raised by the Weth contract.
type WethMint struct {
	To         common.Address
	Amount     *big.Int
	FeeAddress common.Address
	FeeAmount  *big.Int
	OriginTxId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x918d77674bb88eaf75afb307c9723ea6037706de68d6fc07dd0c6cba423a5250.
//
// Solidity: event Mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId)
func (_Weth *WethFilterer) FilterMint(opts *bind.FilterOpts) (*WethMintIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &WethMintIterator{contract: _Weth.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x918d77674bb88eaf75afb307c9723ea6037706de68d6fc07dd0c6cba423a5250.
//
// Solidity: event Mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId)
func (_Weth *WethFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *WethMint) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethMint)
				if err := _Weth.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x918d77674bb88eaf75afb307c9723ea6037706de68d6fc07dd0c6cba423a5250.
//
// Solidity: event Mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId)
func (_Weth *WethFilterer) ParseMint(log types.Log) (*WethMint, error) {
	event := new(WethMint)
	if err := _Weth.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethRemoveSwapTokenIterator is returned from FilterRemoveSwapToken and is used to iterate over the raw logs and unpacked data for RemoveSwapToken events raised by the Weth contract.
type WethRemoveSwapTokenIterator struct {
	Event *WethRemoveSwapToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethRemoveSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethRemoveSwapToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethRemoveSwapToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethRemoveSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethRemoveSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethRemoveSwapToken represents a RemoveSwapToken event raised by the Weth contract.
type WethRemoveSwapToken struct {
	ContractAddress common.Address
	SupplyDecrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemoveSwapToken is a free log retrieval operation binding the contract event 0xd3b4025ff115b79bf2ec5a73c9c784ba8aa9f8f6ba9186b255895c1a9f9042a3.
//
// Solidity: event RemoveSwapToken(address contractAddress, uint256 supplyDecrement)
func (_Weth *WethFilterer) FilterRemoveSwapToken(opts *bind.FilterOpts) (*WethRemoveSwapTokenIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "RemoveSwapToken")
	if err != nil {
		return nil, err
	}
	return &WethRemoveSwapTokenIterator{contract: _Weth.contract, event: "RemoveSwapToken", logs: logs, sub: sub}, nil
}

// WatchRemoveSwapToken is a free log subscription operation binding the contract event 0xd3b4025ff115b79bf2ec5a73c9c784ba8aa9f8f6ba9186b255895c1a9f9042a3.
//
// Solidity: event RemoveSwapToken(address contractAddress, uint256 supplyDecrement)
func (_Weth *WethFilterer) WatchRemoveSwapToken(opts *bind.WatchOpts, sink chan<- *WethRemoveSwapToken) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "RemoveSwapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethRemoveSwapToken)
				if err := _Weth.contract.UnpackLog(event, "RemoveSwapToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveSwapToken is a log parse operation binding the contract event 0xd3b4025ff115b79bf2ec5a73c9c784ba8aa9f8f6ba9186b255895c1a9f9042a3.
//
// Solidity: event RemoveSwapToken(address contractAddress, uint256 supplyDecrement)
func (_Weth *WethFilterer) ParseRemoveSwapToken(log types.Log) (*WethRemoveSwapToken, error) {
	event := new(WethRemoveSwapToken)
	if err := _Weth.contract.UnpackLog(event, "RemoveSwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Weth contract.
type WethSwapIterator struct {
	Event *WethSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethSwap represents a Swap event raised by the Weth contract.
type WethSwap struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x562c219552544ec4c9d7a8eb850f80ea152973e315372bf4999fe7c953ea004f.
//
// Solidity: event Swap(address token, uint256 amount)
func (_Weth *WethFilterer) FilterSwap(opts *bind.FilterOpts) (*WethSwapIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &WethSwapIterator{contract: _Weth.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x562c219552544ec4c9d7a8eb850f80ea152973e315372bf4999fe7c953ea004f.
//
// Solidity: event Swap(address token, uint256 amount)
func (_Weth *WethFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *WethSwap) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethSwap)
				if err := _Weth.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0x562c219552544ec4c9d7a8eb850f80ea152973e315372bf4999fe7c953ea004f.
//
// Solidity: event Swap(address token, uint256 amount)
func (_Weth *WethFilterer) ParseSwap(log types.Log) (*WethSwap, error) {
	event := new(WethSwap)
	if err := _Weth.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Weth contract.
type WethTransferIterator struct {
	Event *WethTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethTransfer represents a Transfer event raised by the Weth contract.
type WethTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Weth *WethFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WethTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Weth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WethTransferIterator{contract: _Weth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Weth *WethFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WethTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Weth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethTransfer)
				if err := _Weth.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Weth *WethFilterer) ParseTransfer(log types.Log) (*WethTransfer, error) {
	event := new(WethTransfer)
	if err := _Weth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethUnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the Weth contract.
type WethUnwrapIterator struct {
	Event *WethUnwrap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WethUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethUnwrap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WethUnwrap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WethUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethUnwrap represents a Unwrap event raised by the Weth contract.
type WethUnwrap struct {
	Amount  *big.Int
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0x37a06799a3500428a773d00284aa706101f5ad94dae9ec37e1c3773aa54c3304.
//
// Solidity: event Unwrap(uint256 amount, uint256 chainId)
func (_Weth *WethFilterer) FilterUnwrap(opts *bind.FilterOpts) (*WethUnwrapIterator, error) {

	logs, sub, err := _Weth.contract.FilterLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return &WethUnwrapIterator{contract: _Weth.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0x37a06799a3500428a773d00284aa706101f5ad94dae9ec37e1c3773aa54c3304.
//
// Solidity: event Unwrap(uint256 amount, uint256 chainId)
func (_Weth *WethFilterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *WethUnwrap) (event.Subscription, error) {

	logs, sub, err := _Weth.contract.WatchLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethUnwrap)
				if err := _Weth.contract.UnpackLog(event, "Unwrap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnwrap is a log parse operation binding the contract event 0x37a06799a3500428a773d00284aa706101f5ad94dae9ec37e1c3773aa54c3304.
//
// Solidity: event Unwrap(uint256 amount, uint256 chainId)
func (_Weth *WethFilterer) ParseUnwrap(log types.Log) (*WethUnwrap, error) {
	event := new(WethUnwrap)
	if err := _Weth.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
