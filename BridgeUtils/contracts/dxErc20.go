// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DxErc20MetaData contains all meta data concerning the DxErc20 contract.
var DxErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"AddSupportedChainId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyIncrement\",\"type\":\"uint256\"}],\"name\":\"AddSwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBridgeRoleAddress\",\"type\":\"address\"}],\"name\":\"MigrateBridgeRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"originTxId\",\"type\":\"bytes32\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyDecrement\",\"type\":\"uint256\"}],\"name\":\"RemoveSwapToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"addSupportedChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supplyIncrement\",\"type\":\"uint256\"}],\"name\":\"addSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBridgeRoleAddress\",\"type\":\"address\"}],\"name\":\"migrateBridgeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"originTxId\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supplyDecrement\",\"type\":\"uint256\"}],\"name\":\"removeSwapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"swapSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DxErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use DxErc20MetaData.ABI instead.
var DxErc20ABI = DxErc20MetaData.ABI

// DxErc20 is an auto generated Go binding around an Ethereum contract.
type DxErc20 struct {
	DxErc20Caller     // Read-only binding to the contract
	DxErc20Transactor // Write-only binding to the contract
	DxErc20Filterer   // Log filterer for contract events
}

// DxErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type DxErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DxErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DxErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DxErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DxErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DxErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DxErc20Session struct {
	Contract     *DxErc20          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DxErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DxErc20CallerSession struct {
	Contract *DxErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DxErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DxErc20TransactorSession struct {
	Contract     *DxErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DxErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type DxErc20Raw struct {
	Contract *DxErc20 // Generic contract binding to access the raw methods on
}

// DxErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DxErc20CallerRaw struct {
	Contract *DxErc20Caller // Generic read-only contract binding to access the raw methods on
}

// DxErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DxErc20TransactorRaw struct {
	Contract *DxErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDxErc20 creates a new instance of DxErc20, bound to a specific deployed contract.
func NewDxErc20(address common.Address, backend bind.ContractBackend) (*DxErc20, error) {
	contract, err := bindDxErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DxErc20{DxErc20Caller: DxErc20Caller{contract: contract}, DxErc20Transactor: DxErc20Transactor{contract: contract}, DxErc20Filterer: DxErc20Filterer{contract: contract}}, nil
}

// NewDxErc20Caller creates a new read-only instance of DxErc20, bound to a specific deployed contract.
func NewDxErc20Caller(address common.Address, caller bind.ContractCaller) (*DxErc20Caller, error) {
	contract, err := bindDxErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DxErc20Caller{contract: contract}, nil
}

// NewDxErc20Transactor creates a new write-only instance of DxErc20, bound to a specific deployed contract.
func NewDxErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*DxErc20Transactor, error) {
	contract, err := bindDxErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DxErc20Transactor{contract: contract}, nil
}

// NewDxErc20Filterer creates a new log filterer instance of DxErc20, bound to a specific deployed contract.
func NewDxErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*DxErc20Filterer, error) {
	contract, err := bindDxErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DxErc20Filterer{contract: contract}, nil
}

// bindDxErc20 binds a generic wrapper to an already deployed contract.
func bindDxErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DxErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DxErc20 *DxErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DxErc20.Contract.DxErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DxErc20 *DxErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DxErc20.Contract.DxErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DxErc20 *DxErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DxErc20.Contract.DxErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DxErc20 *DxErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DxErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DxErc20 *DxErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DxErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DxErc20 *DxErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DxErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DxErc20 *DxErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DxErc20 *DxErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DxErc20.Contract.Allowance(&_DxErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DxErc20 *DxErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DxErc20.Contract.Allowance(&_DxErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DxErc20 *DxErc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DxErc20 *DxErc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _DxErc20.Contract.BalanceOf(&_DxErc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DxErc20 *DxErc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DxErc20.Contract.BalanceOf(&_DxErc20.CallOpts, account)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_DxErc20 *DxErc20Caller) ChainIds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "chainIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_DxErc20 *DxErc20Session) ChainIds(arg0 *big.Int) (bool, error) {
	return _DxErc20.Contract.ChainIds(&_DxErc20.CallOpts, arg0)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(bool)
func (_DxErc20 *DxErc20CallerSession) ChainIds(arg0 *big.Int) (bool, error) {
	return _DxErc20.Contract.ChainIds(&_DxErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DxErc20 *DxErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DxErc20 *DxErc20Session) Decimals() (uint8, error) {
	return _DxErc20.Contract.Decimals(&_DxErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DxErc20 *DxErc20CallerSession) Decimals() (uint8, error) {
	return _DxErc20.Contract.Decimals(&_DxErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DxErc20 *DxErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DxErc20 *DxErc20Session) Name() (string, error) {
	return _DxErc20.Contract.Name(&_DxErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DxErc20 *DxErc20CallerSession) Name() (string, error) {
	return _DxErc20.Contract.Name(&_DxErc20.CallOpts)
}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_DxErc20 *DxErc20Caller) SwapSupply(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "swapSupply", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_DxErc20 *DxErc20Session) SwapSupply(token common.Address) (*big.Int, error) {
	return _DxErc20.Contract.SwapSupply(&_DxErc20.CallOpts, token)
}

// SwapSupply is a free data retrieval call binding the contract method 0xab32dbb7.
//
// Solidity: function swapSupply(address token) view returns(uint256)
func (_DxErc20 *DxErc20CallerSession) SwapSupply(token common.Address) (*big.Int, error) {
	return _DxErc20.Contract.SwapSupply(&_DxErc20.CallOpts, token)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DxErc20 *DxErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DxErc20 *DxErc20Session) Symbol() (string, error) {
	return _DxErc20.Contract.Symbol(&_DxErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DxErc20 *DxErc20CallerSession) Symbol() (string, error) {
	return _DxErc20.Contract.Symbol(&_DxErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DxErc20 *DxErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DxErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DxErc20 *DxErc20Session) TotalSupply() (*big.Int, error) {
	return _DxErc20.Contract.TotalSupply(&_DxErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DxErc20 *DxErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _DxErc20.Contract.TotalSupply(&_DxErc20.CallOpts)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_DxErc20 *DxErc20Transactor) AddSupportedChainId(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "addSupportedChainId", chainId)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_DxErc20 *DxErc20Session) AddSupportedChainId(chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.AddSupportedChainId(&_DxErc20.TransactOpts, chainId)
}

// AddSupportedChainId is a paid mutator transaction binding the contract method 0x66de3b36.
//
// Solidity: function addSupportedChainId(uint256 chainId) returns()
func (_DxErc20 *DxErc20TransactorSession) AddSupportedChainId(chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.AddSupportedChainId(&_DxErc20.TransactOpts, chainId)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_DxErc20 *DxErc20Transactor) AddSwapToken(opts *bind.TransactOpts, contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "addSwapToken", contractAddress, supplyIncrement)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_DxErc20 *DxErc20Session) AddSwapToken(contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.AddSwapToken(&_DxErc20.TransactOpts, contractAddress, supplyIncrement)
}

// AddSwapToken is a paid mutator transaction binding the contract method 0xeff03830.
//
// Solidity: function addSwapToken(address contractAddress, uint256 supplyIncrement) returns()
func (_DxErc20 *DxErc20TransactorSession) AddSwapToken(contractAddress common.Address, supplyIncrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.AddSwapToken(&_DxErc20.TransactOpts, contractAddress, supplyIncrement)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Approve(&_DxErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Approve(&_DxErc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DxErc20 *DxErc20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DxErc20 *DxErc20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Burn(&_DxErc20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DxErc20 *DxErc20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Burn(&_DxErc20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_DxErc20 *DxErc20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_DxErc20 *DxErc20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.BurnFrom(&_DxErc20.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_DxErc20 *DxErc20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.BurnFrom(&_DxErc20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DxErc20 *DxErc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DxErc20 *DxErc20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.DecreaseAllowance(&_DxErc20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DxErc20 *DxErc20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.DecreaseAllowance(&_DxErc20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DxErc20 *DxErc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DxErc20 *DxErc20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.IncreaseAllowance(&_DxErc20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DxErc20 *DxErc20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.IncreaseAllowance(&_DxErc20.TransactOpts, spender, addedValue)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_DxErc20 *DxErc20Transactor) MigrateBridgeRole(opts *bind.TransactOpts, newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "migrateBridgeRole", newBridgeRoleAddress)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_DxErc20 *DxErc20Session) MigrateBridgeRole(newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _DxErc20.Contract.MigrateBridgeRole(&_DxErc20.TransactOpts, newBridgeRoleAddress)
}

// MigrateBridgeRole is a paid mutator transaction binding the contract method 0x5d9898d3.
//
// Solidity: function migrateBridgeRole(address newBridgeRoleAddress) returns()
func (_DxErc20 *DxErc20TransactorSession) MigrateBridgeRole(newBridgeRoleAddress common.Address) (*types.Transaction, error) {
	return _DxErc20.Contract.MigrateBridgeRole(&_DxErc20.TransactOpts, newBridgeRoleAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_DxErc20 *DxErc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "mint", to, amount, feeAddress, feeAmount, originTxId)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_DxErc20 *DxErc20Session) Mint(to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _DxErc20.Contract.Mint(&_DxErc20.TransactOpts, to, amount, feeAddress, feeAmount, originTxId)
}

// Mint is a paid mutator transaction binding the contract method 0x67fc19bb.
//
// Solidity: function mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId) returns()
func (_DxErc20 *DxErc20TransactorSession) Mint(to common.Address, amount *big.Int, feeAddress common.Address, feeAmount *big.Int, originTxId [32]byte) (*types.Transaction, error) {
	return _DxErc20.Contract.Mint(&_DxErc20.TransactOpts, to, amount, feeAddress, feeAmount, originTxId)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_DxErc20 *DxErc20Transactor) RemoveSwapToken(opts *bind.TransactOpts, contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "removeSwapToken", contractAddress, supplyDecrement)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_DxErc20 *DxErc20Session) RemoveSwapToken(contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.RemoveSwapToken(&_DxErc20.TransactOpts, contractAddress, supplyDecrement)
}

// RemoveSwapToken is a paid mutator transaction binding the contract method 0x7c38b457.
//
// Solidity: function removeSwapToken(address contractAddress, uint256 supplyDecrement) returns()
func (_DxErc20 *DxErc20TransactorSession) RemoveSwapToken(contractAddress common.Address, supplyDecrement *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.RemoveSwapToken(&_DxErc20.TransactOpts, contractAddress, supplyDecrement)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_DxErc20 *DxErc20Transactor) Swap(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "swap", token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_DxErc20 *DxErc20Session) Swap(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Swap(&_DxErc20.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_DxErc20 *DxErc20TransactorSession) Swap(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Swap(&_DxErc20.TransactOpts, token, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Transfer(&_DxErc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Transfer(&_DxErc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.TransferFrom(&_DxErc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_DxErc20 *DxErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.TransferFrom(&_DxErc20.TransactOpts, sender, recipient, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_DxErc20 *DxErc20Transactor) Unwrap(opts *bind.TransactOpts, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.contract.Transact(opts, "unwrap", amount, chainId)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_DxErc20 *DxErc20Session) Unwrap(amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Unwrap(&_DxErc20.TransactOpts, amount, chainId)
}

// Unwrap is a paid mutator transaction binding the contract method 0x6e286671.
//
// Solidity: function unwrap(uint256 amount, uint256 chainId) returns()
func (_DxErc20 *DxErc20TransactorSession) Unwrap(amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _DxErc20.Contract.Unwrap(&_DxErc20.TransactOpts, amount, chainId)
}

// DxErc20AddSupportedChainIdIterator is returned from FilterAddSupportedChainId and is used to iterate over the raw logs and unpacked data for AddSupportedChainId events raised by the DxErc20 contract.
type DxErc20AddSupportedChainIdIterator struct {
	Event *DxErc20AddSupportedChainId // Event containing the contract specifics and raw log

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
func (it *DxErc20AddSupportedChainIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20AddSupportedChainId)
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
		it.Event = new(DxErc20AddSupportedChainId)
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
func (it *DxErc20AddSupportedChainIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20AddSupportedChainIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20AddSupportedChainId represents a AddSupportedChainId event raised by the DxErc20 contract.
type DxErc20AddSupportedChainId struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddSupportedChainId is a free log retrieval operation binding the contract event 0x677e2d9a4ed9201aa86725fef875137fc53876e6b68036b974404762682bd122.
//
// Solidity: event AddSupportedChainId(uint256 chainId)
func (_DxErc20 *DxErc20Filterer) FilterAddSupportedChainId(opts *bind.FilterOpts) (*DxErc20AddSupportedChainIdIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "AddSupportedChainId")
	if err != nil {
		return nil, err
	}
	return &DxErc20AddSupportedChainIdIterator{contract: _DxErc20.contract, event: "AddSupportedChainId", logs: logs, sub: sub}, nil
}

// WatchAddSupportedChainId is a free log subscription operation binding the contract event 0x677e2d9a4ed9201aa86725fef875137fc53876e6b68036b974404762682bd122.
//
// Solidity: event AddSupportedChainId(uint256 chainId)
func (_DxErc20 *DxErc20Filterer) WatchAddSupportedChainId(opts *bind.WatchOpts, sink chan<- *DxErc20AddSupportedChainId) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "AddSupportedChainId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20AddSupportedChainId)
				if err := _DxErc20.contract.UnpackLog(event, "AddSupportedChainId", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseAddSupportedChainId(log types.Log) (*DxErc20AddSupportedChainId, error) {
	event := new(DxErc20AddSupportedChainId)
	if err := _DxErc20.contract.UnpackLog(event, "AddSupportedChainId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20AddSwapTokenIterator is returned from FilterAddSwapToken and is used to iterate over the raw logs and unpacked data for AddSwapToken events raised by the DxErc20 contract.
type DxErc20AddSwapTokenIterator struct {
	Event *DxErc20AddSwapToken // Event containing the contract specifics and raw log

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
func (it *DxErc20AddSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20AddSwapToken)
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
		it.Event = new(DxErc20AddSwapToken)
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
func (it *DxErc20AddSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20AddSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20AddSwapToken represents a AddSwapToken event raised by the DxErc20 contract.
type DxErc20AddSwapToken struct {
	ContractAddress common.Address
	SupplyIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddSwapToken is a free log retrieval operation binding the contract event 0x3e4fdfb0f47da284fe8b5b3a7e5d10b211e323c9a0c144c421ae1d211873f853.
//
// Solidity: event AddSwapToken(address contractAddress, uint256 supplyIncrement)
func (_DxErc20 *DxErc20Filterer) FilterAddSwapToken(opts *bind.FilterOpts) (*DxErc20AddSwapTokenIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "AddSwapToken")
	if err != nil {
		return nil, err
	}
	return &DxErc20AddSwapTokenIterator{contract: _DxErc20.contract, event: "AddSwapToken", logs: logs, sub: sub}, nil
}

// WatchAddSwapToken is a free log subscription operation binding the contract event 0x3e4fdfb0f47da284fe8b5b3a7e5d10b211e323c9a0c144c421ae1d211873f853.
//
// Solidity: event AddSwapToken(address contractAddress, uint256 supplyIncrement)
func (_DxErc20 *DxErc20Filterer) WatchAddSwapToken(opts *bind.WatchOpts, sink chan<- *DxErc20AddSwapToken) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "AddSwapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20AddSwapToken)
				if err := _DxErc20.contract.UnpackLog(event, "AddSwapToken", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseAddSwapToken(log types.Log) (*DxErc20AddSwapToken, error) {
	event := new(DxErc20AddSwapToken)
	if err := _DxErc20.contract.UnpackLog(event, "AddSwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DxErc20 contract.
type DxErc20ApprovalIterator struct {
	Event *DxErc20Approval // Event containing the contract specifics and raw log

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
func (it *DxErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20Approval)
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
		it.Event = new(DxErc20Approval)
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
func (it *DxErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20Approval represents a Approval event raised by the DxErc20 contract.
type DxErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DxErc20 *DxErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DxErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DxErc20ApprovalIterator{contract: _DxErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DxErc20 *DxErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DxErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20Approval)
				if err := _DxErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseApproval(log types.Log) (*DxErc20Approval, error) {
	event := new(DxErc20Approval)
	if err := _DxErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20MigrateBridgeRoleIterator is returned from FilterMigrateBridgeRole and is used to iterate over the raw logs and unpacked data for MigrateBridgeRole events raised by the DxErc20 contract.
type DxErc20MigrateBridgeRoleIterator struct {
	Event *DxErc20MigrateBridgeRole // Event containing the contract specifics and raw log

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
func (it *DxErc20MigrateBridgeRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20MigrateBridgeRole)
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
		it.Event = new(DxErc20MigrateBridgeRole)
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
func (it *DxErc20MigrateBridgeRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20MigrateBridgeRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20MigrateBridgeRole represents a MigrateBridgeRole event raised by the DxErc20 contract.
type DxErc20MigrateBridgeRole struct {
	NewBridgeRoleAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMigrateBridgeRole is a free log retrieval operation binding the contract event 0x871b00a4e20f8436702d0174eb87d84d7cd1dd5c34d4bb1b4e75438b3398d512.
//
// Solidity: event MigrateBridgeRole(address newBridgeRoleAddress)
func (_DxErc20 *DxErc20Filterer) FilterMigrateBridgeRole(opts *bind.FilterOpts) (*DxErc20MigrateBridgeRoleIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "MigrateBridgeRole")
	if err != nil {
		return nil, err
	}
	return &DxErc20MigrateBridgeRoleIterator{contract: _DxErc20.contract, event: "MigrateBridgeRole", logs: logs, sub: sub}, nil
}

// WatchMigrateBridgeRole is a free log subscription operation binding the contract event 0x871b00a4e20f8436702d0174eb87d84d7cd1dd5c34d4bb1b4e75438b3398d512.
//
// Solidity: event MigrateBridgeRole(address newBridgeRoleAddress)
func (_DxErc20 *DxErc20Filterer) WatchMigrateBridgeRole(opts *bind.WatchOpts, sink chan<- *DxErc20MigrateBridgeRole) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "MigrateBridgeRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20MigrateBridgeRole)
				if err := _DxErc20.contract.UnpackLog(event, "MigrateBridgeRole", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseMigrateBridgeRole(log types.Log) (*DxErc20MigrateBridgeRole, error) {
	event := new(DxErc20MigrateBridgeRole)
	if err := _DxErc20.contract.UnpackLog(event, "MigrateBridgeRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DxErc20 contract.
type DxErc20MintIterator struct {
	Event *DxErc20Mint // Event containing the contract specifics and raw log

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
func (it *DxErc20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20Mint)
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
		it.Event = new(DxErc20Mint)
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
func (it *DxErc20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20Mint represents a Mint event raised by the DxErc20 contract.
type DxErc20Mint struct {
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
func (_DxErc20 *DxErc20Filterer) FilterMint(opts *bind.FilterOpts) (*DxErc20MintIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &DxErc20MintIterator{contract: _DxErc20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x918d77674bb88eaf75afb307c9723ea6037706de68d6fc07dd0c6cba423a5250.
//
// Solidity: event Mint(address to, uint256 amount, address feeAddress, uint256 feeAmount, bytes32 originTxId)
func (_DxErc20 *DxErc20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DxErc20Mint) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20Mint)
				if err := _DxErc20.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseMint(log types.Log) (*DxErc20Mint, error) {
	event := new(DxErc20Mint)
	if err := _DxErc20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20RemoveSwapTokenIterator is returned from FilterRemoveSwapToken and is used to iterate over the raw logs and unpacked data for RemoveSwapToken events raised by the DxErc20 contract.
type DxErc20RemoveSwapTokenIterator struct {
	Event *DxErc20RemoveSwapToken // Event containing the contract specifics and raw log

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
func (it *DxErc20RemoveSwapTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20RemoveSwapToken)
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
		it.Event = new(DxErc20RemoveSwapToken)
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
func (it *DxErc20RemoveSwapTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20RemoveSwapTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20RemoveSwapToken represents a RemoveSwapToken event raised by the DxErc20 contract.
type DxErc20RemoveSwapToken struct {
	ContractAddress common.Address
	SupplyDecrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemoveSwapToken is a free log retrieval operation binding the contract event 0xd3b4025ff115b79bf2ec5a73c9c784ba8aa9f8f6ba9186b255895c1a9f9042a3.
//
// Solidity: event RemoveSwapToken(address contractAddress, uint256 supplyDecrement)
func (_DxErc20 *DxErc20Filterer) FilterRemoveSwapToken(opts *bind.FilterOpts) (*DxErc20RemoveSwapTokenIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "RemoveSwapToken")
	if err != nil {
		return nil, err
	}
	return &DxErc20RemoveSwapTokenIterator{contract: _DxErc20.contract, event: "RemoveSwapToken", logs: logs, sub: sub}, nil
}

// WatchRemoveSwapToken is a free log subscription operation binding the contract event 0xd3b4025ff115b79bf2ec5a73c9c784ba8aa9f8f6ba9186b255895c1a9f9042a3.
//
// Solidity: event RemoveSwapToken(address contractAddress, uint256 supplyDecrement)
func (_DxErc20 *DxErc20Filterer) WatchRemoveSwapToken(opts *bind.WatchOpts, sink chan<- *DxErc20RemoveSwapToken) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "RemoveSwapToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20RemoveSwapToken)
				if err := _DxErc20.contract.UnpackLog(event, "RemoveSwapToken", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseRemoveSwapToken(log types.Log) (*DxErc20RemoveSwapToken, error) {
	event := new(DxErc20RemoveSwapToken)
	if err := _DxErc20.contract.UnpackLog(event, "RemoveSwapToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the DxErc20 contract.
type DxErc20SwapIterator struct {
	Event *DxErc20Swap // Event containing the contract specifics and raw log

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
func (it *DxErc20SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20Swap)
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
		it.Event = new(DxErc20Swap)
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
func (it *DxErc20SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20Swap represents a Swap event raised by the DxErc20 contract.
type DxErc20Swap struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x562c219552544ec4c9d7a8eb850f80ea152973e315372bf4999fe7c953ea004f.
//
// Solidity: event Swap(address token, uint256 amount)
func (_DxErc20 *DxErc20Filterer) FilterSwap(opts *bind.FilterOpts) (*DxErc20SwapIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &DxErc20SwapIterator{contract: _DxErc20.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x562c219552544ec4c9d7a8eb850f80ea152973e315372bf4999fe7c953ea004f.
//
// Solidity: event Swap(address token, uint256 amount)
func (_DxErc20 *DxErc20Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *DxErc20Swap) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20Swap)
				if err := _DxErc20.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseSwap(log types.Log) (*DxErc20Swap, error) {
	event := new(DxErc20Swap)
	if err := _DxErc20.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DxErc20 contract.
type DxErc20TransferIterator struct {
	Event *DxErc20Transfer // Event containing the contract specifics and raw log

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
func (it *DxErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20Transfer)
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
		it.Event = new(DxErc20Transfer)
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
func (it *DxErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20Transfer represents a Transfer event raised by the DxErc20 contract.
type DxErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DxErc20 *DxErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DxErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DxErc20TransferIterator{contract: _DxErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DxErc20 *DxErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DxErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20Transfer)
				if err := _DxErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseTransfer(log types.Log) (*DxErc20Transfer, error) {
	event := new(DxErc20Transfer)
	if err := _DxErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DxErc20UnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the DxErc20 contract.
type DxErc20UnwrapIterator struct {
	Event *DxErc20Unwrap // Event containing the contract specifics and raw log

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
func (it *DxErc20UnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DxErc20Unwrap)
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
		it.Event = new(DxErc20Unwrap)
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
func (it *DxErc20UnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DxErc20UnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DxErc20Unwrap represents a Unwrap event raised by the DxErc20 contract.
type DxErc20Unwrap struct {
	Amount  *big.Int
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0x37a06799a3500428a773d00284aa706101f5ad94dae9ec37e1c3773aa54c3304.
//
// Solidity: event Unwrap(uint256 amount, uint256 chainId)
func (_DxErc20 *DxErc20Filterer) FilterUnwrap(opts *bind.FilterOpts) (*DxErc20UnwrapIterator, error) {

	logs, sub, err := _DxErc20.contract.FilterLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return &DxErc20UnwrapIterator{contract: _DxErc20.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0x37a06799a3500428a773d00284aa706101f5ad94dae9ec37e1c3773aa54c3304.
//
// Solidity: event Unwrap(uint256 amount, uint256 chainId)
func (_DxErc20 *DxErc20Filterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *DxErc20Unwrap) (event.Subscription, error) {

	logs, sub, err := _DxErc20.contract.WatchLogs(opts, "Unwrap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DxErc20Unwrap)
				if err := _DxErc20.contract.UnpackLog(event, "Unwrap", log); err != nil {
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
func (_DxErc20 *DxErc20Filterer) ParseUnwrap(log types.Log) (*DxErc20Unwrap, error) {
	event := new(DxErc20Unwrap)
	if err := _DxErc20.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
