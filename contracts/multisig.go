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
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_requiredConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitTransaction\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredConfirmations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsCaller) GetTransaction(opts *bind.CallOpts, _txIndex *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTransaction", _txIndex)

	outstruct := new(struct {
		To            common.Address
		Value         *big.Int
		Data          []byte
		Executed      bool
		Confirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Confirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsSession) GetTransaction(_txIndex *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	return _Contracts.Contract.GetTransaction(&_Contracts.CallOpts, _txIndex)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsCallerSession) GetTransaction(_txIndex *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	return _Contracts.Contract.GetTransaction(&_Contracts.CallOpts, _txIndex)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Contracts *ContractsCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Contracts *ContractsSession) GetTransactionCount() (*big.Int, error) {
	return _Contracts.Contract.GetTransactionCount(&_Contracts.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetTransactionCount() (*big.Int, error) {
	return _Contracts.Contract.GetTransactionCount(&_Contracts.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsCaller) IsConfirmed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isConfirmed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.IsConfirmed(&_Contracts.CallOpts, arg0, arg1)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsCallerSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.IsConfirmed(&_Contracts.CallOpts, arg0, arg1)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Contracts *ContractsCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Contracts *ContractsSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.IsOwner(&_Contracts.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.IsOwner(&_Contracts.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contracts *ContractsSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.Owners(&_Contracts.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.Owners(&_Contracts.CallOpts, arg0)
}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Contracts *ContractsCaller) RequiredConfirmations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "requiredConfirmations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Contracts *ContractsSession) RequiredConfirmations() (*big.Int, error) {
	return _Contracts.Contract.RequiredConfirmations(&_Contracts.CallOpts)
}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Contracts *ContractsCallerSession) RequiredConfirmations() (*big.Int, error) {
	return _Contracts.Contract.RequiredConfirmations(&_Contracts.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		To            common.Address
		Value         *big.Int
		Data          []byte
		Executed      bool
		Confirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Confirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsSession) Transactions(arg0 *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	return _Contracts.Contract.Transactions(&_Contracts.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 confirmations)
func (_Contracts *ContractsCallerSession) Transactions(arg0 *big.Int) (struct {
	To            common.Address
	Value         *big.Int
	Data          []byte
	Executed      bool
	Confirmations *big.Int
}, error) {
	return _Contracts.Contract.Transactions(&_Contracts.CallOpts, arg0)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsTransactor) ConfirmTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "confirmTransaction", _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfirmTransaction(&_Contracts.TransactOpts, _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsTransactorSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfirmTransaction(&_Contracts.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsTransactor) ExecuteTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "executeTransaction", _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ExecuteTransaction(&_Contracts.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_Contracts *ContractsTransactorSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ExecuteTransaction(&_Contracts.TransactOpts, _txIndex)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Contracts *ContractsTransactor) SubmitTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitTransaction", _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Contracts *ContractsSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitTransaction(&_Contracts.TransactOpts, _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_Contracts *ContractsTransactorSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitTransaction(&_Contracts.TransactOpts, _to, _value, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the Contracts contract.
type ContractsConfirmTransactionIterator struct {
	Event *ContractsConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *ContractsConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsConfirmTransaction)
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
		it.Event = new(ContractsConfirmTransaction)
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
func (it *ContractsConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsConfirmTransaction represents a ConfirmTransaction event raised by the Contracts contract.
type ContractsConfirmTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*ContractsConfirmTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsConfirmTransactionIterator{contract: _Contracts.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *ContractsConfirmTransaction, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsConfirmTransaction)
				if err := _Contracts.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ParseConfirmTransaction is a log parse operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) ParseConfirmTransaction(log types.Log) (*ContractsConfirmTransaction, error) {
	event := new(ContractsConfirmTransaction)
	if err := _Contracts.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contracts contract.
type ContractsDepositIterator struct {
	Event *ContractsDeposit // Event containing the contract specifics and raw log

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
func (it *ContractsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDeposit)
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
		it.Event = new(ContractsDeposit)
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
func (it *ContractsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDeposit represents a Deposit event raised by the Contracts contract.
type ContractsDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ContractsDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDepositIterator{contract: _Contracts.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractsDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDeposit)
				if err := _Contracts.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Contracts *ContractsFilterer) ParseDeposit(log types.Log) (*ContractsDeposit, error) {
	event := new(ContractsDeposit)
	if err := _Contracts.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsExecuteTransactionIterator is returned from FilterExecuteTransaction and is used to iterate over the raw logs and unpacked data for ExecuteTransaction events raised by the Contracts contract.
type ContractsExecuteTransactionIterator struct {
	Event *ContractsExecuteTransaction // Event containing the contract specifics and raw log

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
func (it *ContractsExecuteTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsExecuteTransaction)
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
		it.Event = new(ContractsExecuteTransaction)
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
func (it *ContractsExecuteTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsExecuteTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsExecuteTransaction represents a ExecuteTransaction event raised by the Contracts contract.
type ContractsExecuteTransaction struct {
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransaction is a free log retrieval operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) FilterExecuteTransaction(opts *bind.FilterOpts, txIndex []*big.Int) (*ContractsExecuteTransactionIterator, error) {

	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ExecuteTransaction", txIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsExecuteTransactionIterator{contract: _Contracts.contract, event: "ExecuteTransaction", logs: logs, sub: sub}, nil
}

// WatchExecuteTransaction is a free log subscription operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) WatchExecuteTransaction(opts *bind.WatchOpts, sink chan<- *ContractsExecuteTransaction, txIndex []*big.Int) (event.Subscription, error) {

	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ExecuteTransaction", txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsExecuteTransaction)
				if err := _Contracts.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
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

// ParseExecuteTransaction is a log parse operation binding the contract event 0xae30dc3f11bb6b178aafe5e7fc568fb6d87200068a944a8015c0db1b4533dbb8.
//
// Solidity: event ExecuteTransaction(uint256 indexed txIndex)
func (_Contracts *ContractsFilterer) ParseExecuteTransaction(log types.Log) (*ContractsExecuteTransaction, error) {
	event := new(ContractsExecuteTransaction)
	if err := _Contracts.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSubmitTransactionIterator is returned from FilterSubmitTransaction and is used to iterate over the raw logs and unpacked data for SubmitTransaction events raised by the Contracts contract.
type ContractsSubmitTransactionIterator struct {
	Event *ContractsSubmitTransaction // Event containing the contract specifics and raw log

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
func (it *ContractsSubmitTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSubmitTransaction)
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
		it.Event = new(ContractsSubmitTransaction)
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
func (it *ContractsSubmitTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSubmitTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSubmitTransaction represents a SubmitTransaction event raised by the Contracts contract.
type ContractsSubmitTransaction struct {
	TxIndex *big.Int
	To      common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitTransaction is a free log retrieval operation binding the contract event 0x00c2937519b07b47ea25f52e915dd99023f9b3e1aaeef041ddecbac3e9400bbe.
//
// Solidity: event SubmitTransaction(uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) FilterSubmitTransaction(opts *bind.FilterOpts, txIndex []*big.Int, to []common.Address) (*ContractsSubmitTransactionIterator, error) {

	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SubmitTransaction", txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSubmitTransactionIterator{contract: _Contracts.contract, event: "SubmitTransaction", logs: logs, sub: sub}, nil
}

// WatchSubmitTransaction is a free log subscription operation binding the contract event 0x00c2937519b07b47ea25f52e915dd99023f9b3e1aaeef041ddecbac3e9400bbe.
//
// Solidity: event SubmitTransaction(uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) WatchSubmitTransaction(opts *bind.WatchOpts, sink chan<- *ContractsSubmitTransaction, txIndex []*big.Int, to []common.Address) (event.Subscription, error) {

	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SubmitTransaction", txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSubmitTransaction)
				if err := _Contracts.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
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

// ParseSubmitTransaction is a log parse operation binding the contract event 0x00c2937519b07b47ea25f52e915dd99023f9b3e1aaeef041ddecbac3e9400bbe.
//
// Solidity: event SubmitTransaction(uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_Contracts *ContractsFilterer) ParseSubmitTransaction(log types.Log) (*ContractsSubmitTransaction, error) {
	event := new(ContractsSubmitTransaction)
	if err := _Contracts.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
