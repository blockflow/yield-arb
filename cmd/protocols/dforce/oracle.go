// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dforce

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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poster\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPoster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPoster\",\"type\":\"address\"}],\"name\":\"NewPoster\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"priceModel\",\"type\":\"address\"}],\"name\":\"SetAssetPriceModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newState\",\"type\":\"bool\"}],\"name\":\"SetPaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"_disableAssetPriceModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"}],\"name\":\"_disableAssetStatusOracleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"_executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"}],\"name\":\"_executeTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signature\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"_setAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceModel\",\"type\":\"address\"}],\"name\":\"_setAssetPriceModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_priceModels\",\"type\":\"address[]\"}],\"name\":\"_setAssetPriceModelBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"}],\"name\":\"_setAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_requestedState\",\"type\":\"bool\"}],\"name\":\"_setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"_setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPoster\",\"type\":\"address\"}],\"name\":\"_setPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getAssetPriceStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getUnderlyingPriceAndStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poster\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"priceModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_postSwing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_postBuffer\",\"type\":\"uint256\"}],\"name\":\"readyToUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_requestedPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_postSwings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_postBuffers\",\"type\":\"uint256[]\"}],\"name\":\"readyToUpdates\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_requestedPrices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oracle *OracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oracle *OracleSession) Paused() (bool, error) {
	return _Oracle.Contract.Paused(&_Oracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Oracle *OracleCallerSession) Paused() (bool, error) {
	return _Oracle.Contract.Paused(&_Oracle.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Oracle *OracleCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Oracle *OracleSession) PendingOwner() (common.Address, error) {
	return _Oracle.Contract.PendingOwner(&_Oracle.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Oracle *OracleCallerSession) PendingOwner() (common.Address, error) {
	return _Oracle.Contract.PendingOwner(&_Oracle.CallOpts)
}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_Oracle *OracleCaller) Poster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "poster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_Oracle *OracleSession) Poster() (common.Address, error) {
	return _Oracle.Contract.Poster(&_Oracle.CallOpts)
}

// Poster is a free data retrieval call binding the contract method 0x80959721.
//
// Solidity: function poster() view returns(address)
func (_Oracle *OracleCallerSession) Poster() (common.Address, error) {
	return _Oracle.Contract.Poster(&_Oracle.CallOpts)
}

// PriceModel is a free data retrieval call binding the contract method 0x932e295f.
//
// Solidity: function priceModel(address _asset) view returns(address)
func (_Oracle *OracleCaller) PriceModel(opts *bind.CallOpts, _asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "priceModel", _asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceModel is a free data retrieval call binding the contract method 0x932e295f.
//
// Solidity: function priceModel(address _asset) view returns(address)
func (_Oracle *OracleSession) PriceModel(_asset common.Address) (common.Address, error) {
	return _Oracle.Contract.PriceModel(&_Oracle.CallOpts, _asset)
}

// PriceModel is a free data retrieval call binding the contract method 0x932e295f.
//
// Solidity: function priceModel(address _asset) view returns(address)
func (_Oracle *OracleCallerSession) PriceModel(_asset common.Address) (common.Address, error) {
	return _Oracle.Contract.PriceModel(&_Oracle.CallOpts, _asset)
}

// ReadyToUpdate is a free data retrieval call binding the contract method 0x16484cd1.
//
// Solidity: function readyToUpdate(address _asset, uint256 _requestedPrice, uint256 _postSwing, uint256 _postBuffer) view returns(bool)
func (_Oracle *OracleCaller) ReadyToUpdate(opts *bind.CallOpts, _asset common.Address, _requestedPrice *big.Int, _postSwing *big.Int, _postBuffer *big.Int) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "readyToUpdate", _asset, _requestedPrice, _postSwing, _postBuffer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReadyToUpdate is a free data retrieval call binding the contract method 0x16484cd1.
//
// Solidity: function readyToUpdate(address _asset, uint256 _requestedPrice, uint256 _postSwing, uint256 _postBuffer) view returns(bool)
func (_Oracle *OracleSession) ReadyToUpdate(_asset common.Address, _requestedPrice *big.Int, _postSwing *big.Int, _postBuffer *big.Int) (bool, error) {
	return _Oracle.Contract.ReadyToUpdate(&_Oracle.CallOpts, _asset, _requestedPrice, _postSwing, _postBuffer)
}

// ReadyToUpdate is a free data retrieval call binding the contract method 0x16484cd1.
//
// Solidity: function readyToUpdate(address _asset, uint256 _requestedPrice, uint256 _postSwing, uint256 _postBuffer) view returns(bool)
func (_Oracle *OracleCallerSession) ReadyToUpdate(_asset common.Address, _requestedPrice *big.Int, _postSwing *big.Int, _postBuffer *big.Int) (bool, error) {
	return _Oracle.Contract.ReadyToUpdate(&_Oracle.CallOpts, _asset, _requestedPrice, _postSwing, _postBuffer)
}

// ReadyToUpdates is a free data retrieval call binding the contract method 0x70ccd4aa.
//
// Solidity: function readyToUpdates(address[] _assets, uint256[] _requestedPrices, uint256[] _postSwings, uint256[] _postBuffers) view returns(bool[])
func (_Oracle *OracleCaller) ReadyToUpdates(opts *bind.CallOpts, _assets []common.Address, _requestedPrices []*big.Int, _postSwings []*big.Int, _postBuffers []*big.Int) ([]bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "readyToUpdates", _assets, _requestedPrices, _postSwings, _postBuffers)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ReadyToUpdates is a free data retrieval call binding the contract method 0x70ccd4aa.
//
// Solidity: function readyToUpdates(address[] _assets, uint256[] _requestedPrices, uint256[] _postSwings, uint256[] _postBuffers) view returns(bool[])
func (_Oracle *OracleSession) ReadyToUpdates(_assets []common.Address, _requestedPrices []*big.Int, _postSwings []*big.Int, _postBuffers []*big.Int) ([]bool, error) {
	return _Oracle.Contract.ReadyToUpdates(&_Oracle.CallOpts, _assets, _requestedPrices, _postSwings, _postBuffers)
}

// ReadyToUpdates is a free data retrieval call binding the contract method 0x70ccd4aa.
//
// Solidity: function readyToUpdates(address[] _assets, uint256[] _requestedPrices, uint256[] _postSwings, uint256[] _postBuffers) view returns(bool[])
func (_Oracle *OracleCallerSession) ReadyToUpdates(_assets []common.Address, _requestedPrices []*big.Int, _postSwings []*big.Int, _postBuffers []*big.Int) ([]bool, error) {
	return _Oracle.Contract.ReadyToUpdates(&_Oracle.CallOpts, _assets, _requestedPrices, _postSwings, _postBuffers)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Oracle *OracleTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Oracle *OracleSession) AcceptOwner() (*types.Transaction, error) {
	return _Oracle.Contract.AcceptOwner(&_Oracle.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Oracle *OracleTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Oracle.Contract.AcceptOwner(&_Oracle.TransactOpts)
}

// DisableAssetPriceModel is a paid mutator transaction binding the contract method 0xb3fa9183.
//
// Solidity: function _disableAssetPriceModel(address _asset) returns()
func (_Oracle *OracleTransactor) DisableAssetPriceModel(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_disableAssetPriceModel", _asset)
}

// DisableAssetPriceModel is a paid mutator transaction binding the contract method 0xb3fa9183.
//
// Solidity: function _disableAssetPriceModel(address _asset) returns()
func (_Oracle *OracleSession) DisableAssetPriceModel(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DisableAssetPriceModel(&_Oracle.TransactOpts, _asset)
}

// DisableAssetPriceModel is a paid mutator transaction binding the contract method 0xb3fa9183.
//
// Solidity: function _disableAssetPriceModel(address _asset) returns()
func (_Oracle *OracleTransactorSession) DisableAssetPriceModel(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DisableAssetPriceModel(&_Oracle.TransactOpts, _asset)
}

// DisableAssetStatusOracleBatch is a paid mutator transaction binding the contract method 0xe709194d.
//
// Solidity: function _disableAssetStatusOracleBatch(address[] _assets) returns()
func (_Oracle *OracleTransactor) DisableAssetStatusOracleBatch(opts *bind.TransactOpts, _assets []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_disableAssetStatusOracleBatch", _assets)
}

// DisableAssetStatusOracleBatch is a paid mutator transaction binding the contract method 0xe709194d.
//
// Solidity: function _disableAssetStatusOracleBatch(address[] _assets) returns()
func (_Oracle *OracleSession) DisableAssetStatusOracleBatch(_assets []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DisableAssetStatusOracleBatch(&_Oracle.TransactOpts, _assets)
}

// DisableAssetStatusOracleBatch is a paid mutator transaction binding the contract method 0xe709194d.
//
// Solidity: function _disableAssetStatusOracleBatch(address[] _assets) returns()
func (_Oracle *OracleTransactorSession) DisableAssetStatusOracleBatch(_assets []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DisableAssetStatusOracleBatch(&_Oracle.TransactOpts, _assets)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x45f66d61.
//
// Solidity: function _executeTransaction(address _target, string _signature, bytes _data) returns()
func (_Oracle *OracleTransactor) ExecuteTransaction(opts *bind.TransactOpts, _target common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_executeTransaction", _target, _signature, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x45f66d61.
//
// Solidity: function _executeTransaction(address _target, string _signature, bytes _data) returns()
func (_Oracle *OracleSession) ExecuteTransaction(_target common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.ExecuteTransaction(&_Oracle.TransactOpts, _target, _signature, _data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x45f66d61.
//
// Solidity: function _executeTransaction(address _target, string _signature, bytes _data) returns()
func (_Oracle *OracleTransactorSession) ExecuteTransaction(_target common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.ExecuteTransaction(&_Oracle.TransactOpts, _target, _signature, _data)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0xf7b252bc.
//
// Solidity: function _executeTransactions(address[] _targets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleTransactor) ExecuteTransactions(opts *bind.TransactOpts, _targets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_executeTransactions", _targets, _signatures, _calldatas)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0xf7b252bc.
//
// Solidity: function _executeTransactions(address[] _targets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleSession) ExecuteTransactions(_targets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.Contract.ExecuteTransactions(&_Oracle.TransactOpts, _targets, _signatures, _calldatas)
}

// ExecuteTransactions is a paid mutator transaction binding the contract method 0xf7b252bc.
//
// Solidity: function _executeTransactions(address[] _targets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleTransactorSession) ExecuteTransactions(_targets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.Contract.ExecuteTransactions(&_Oracle.TransactOpts, _targets, _signatures, _calldatas)
}

// SetAsset is a paid mutator transaction binding the contract method 0x432593bb.
//
// Solidity: function _setAsset(address _asset, string _signature, bytes _data) returns()
func (_Oracle *OracleTransactor) SetAsset(opts *bind.TransactOpts, _asset common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setAsset", _asset, _signature, _data)
}

// SetAsset is a paid mutator transaction binding the contract method 0x432593bb.
//
// Solidity: function _setAsset(address _asset, string _signature, bytes _data) returns()
func (_Oracle *OracleSession) SetAsset(_asset common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetAsset(&_Oracle.TransactOpts, _asset, _signature, _data)
}

// SetAsset is a paid mutator transaction binding the contract method 0x432593bb.
//
// Solidity: function _setAsset(address _asset, string _signature, bytes _data) returns()
func (_Oracle *OracleTransactorSession) SetAsset(_asset common.Address, _signature string, _data []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetAsset(&_Oracle.TransactOpts, _asset, _signature, _data)
}

// SetAssetPriceModel is a paid mutator transaction binding the contract method 0x7c3d5221.
//
// Solidity: function _setAssetPriceModel(address _asset, address _priceModel) returns()
func (_Oracle *OracleTransactor) SetAssetPriceModel(opts *bind.TransactOpts, _asset common.Address, _priceModel common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setAssetPriceModel", _asset, _priceModel)
}

// SetAssetPriceModel is a paid mutator transaction binding the contract method 0x7c3d5221.
//
// Solidity: function _setAssetPriceModel(address _asset, address _priceModel) returns()
func (_Oracle *OracleSession) SetAssetPriceModel(_asset common.Address, _priceModel common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetPriceModel(&_Oracle.TransactOpts, _asset, _priceModel)
}

// SetAssetPriceModel is a paid mutator transaction binding the contract method 0x7c3d5221.
//
// Solidity: function _setAssetPriceModel(address _asset, address _priceModel) returns()
func (_Oracle *OracleTransactorSession) SetAssetPriceModel(_asset common.Address, _priceModel common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetPriceModel(&_Oracle.TransactOpts, _asset, _priceModel)
}

// SetAssetPriceModelBatch is a paid mutator transaction binding the contract method 0xb2681608.
//
// Solidity: function _setAssetPriceModelBatch(address[] _assets, address[] _priceModels) returns()
func (_Oracle *OracleTransactor) SetAssetPriceModelBatch(opts *bind.TransactOpts, _assets []common.Address, _priceModels []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setAssetPriceModelBatch", _assets, _priceModels)
}

// SetAssetPriceModelBatch is a paid mutator transaction binding the contract method 0xb2681608.
//
// Solidity: function _setAssetPriceModelBatch(address[] _assets, address[] _priceModels) returns()
func (_Oracle *OracleSession) SetAssetPriceModelBatch(_assets []common.Address, _priceModels []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetPriceModelBatch(&_Oracle.TransactOpts, _assets, _priceModels)
}

// SetAssetPriceModelBatch is a paid mutator transaction binding the contract method 0xb2681608.
//
// Solidity: function _setAssetPriceModelBatch(address[] _assets, address[] _priceModels) returns()
func (_Oracle *OracleTransactorSession) SetAssetPriceModelBatch(_assets []common.Address, _priceModels []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssetPriceModelBatch(&_Oracle.TransactOpts, _assets, _priceModels)
}

// SetAssets is a paid mutator transaction binding the contract method 0xce03752b.
//
// Solidity: function _setAssets(address[] _assets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleTransactor) SetAssets(opts *bind.TransactOpts, _assets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setAssets", _assets, _signatures, _calldatas)
}

// SetAssets is a paid mutator transaction binding the contract method 0xce03752b.
//
// Solidity: function _setAssets(address[] _assets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleSession) SetAssets(_assets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssets(&_Oracle.TransactOpts, _assets, _signatures, _calldatas)
}

// SetAssets is a paid mutator transaction binding the contract method 0xce03752b.
//
// Solidity: function _setAssets(address[] _assets, string[] _signatures, bytes[] _calldatas) returns()
func (_Oracle *OracleTransactorSession) SetAssets(_assets []common.Address, _signatures []string, _calldatas [][]byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetAssets(&_Oracle.TransactOpts, _assets, _signatures, _calldatas)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool _requestedState) returns()
func (_Oracle *OracleTransactor) SetPaused(opts *bind.TransactOpts, _requestedState bool) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setPaused", _requestedState)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool _requestedState) returns()
func (_Oracle *OracleSession) SetPaused(_requestedState bool) (*types.Transaction, error) {
	return _Oracle.Contract.SetPaused(&_Oracle.TransactOpts, _requestedState)
}

// SetPaused is a paid mutator transaction binding the contract method 0x26617c28.
//
// Solidity: function _setPaused(bool _requestedState) returns()
func (_Oracle *OracleTransactorSession) SetPaused(_requestedState bool) (*types.Transaction, error) {
	return _Oracle.Contract.SetPaused(&_Oracle.TransactOpts, _requestedState)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Oracle *OracleTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Oracle *OracleSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetPendingOwner(&_Oracle.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Oracle *OracleTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetPendingOwner(&_Oracle.TransactOpts, newPendingOwner)
}

// SetPoster is a paid mutator transaction binding the contract method 0x5d496818.
//
// Solidity: function _setPoster(address _newPoster) returns()
func (_Oracle *OracleTransactor) SetPoster(opts *bind.TransactOpts, _newPoster common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "_setPoster", _newPoster)
}

// SetPoster is a paid mutator transaction binding the contract method 0x5d496818.
//
// Solidity: function _setPoster(address _newPoster) returns()
func (_Oracle *OracleSession) SetPoster(_newPoster common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetPoster(&_Oracle.TransactOpts, _newPoster)
}

// SetPoster is a paid mutator transaction binding the contract method 0x5d496818.
//
// Solidity: function _setPoster(address _newPoster) returns()
func (_Oracle *OracleTransactorSession) SetPoster(_newPoster common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetPoster(&_Oracle.TransactOpts, _newPoster)
}

// GetAssetPriceStatus is a paid mutator transaction binding the contract method 0x0a2b53bf.
//
// Solidity: function getAssetPriceStatus(address _asset) returns(bool)
func (_Oracle *OracleTransactor) GetAssetPriceStatus(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "getAssetPriceStatus", _asset)
}

// GetAssetPriceStatus is a paid mutator transaction binding the contract method 0x0a2b53bf.
//
// Solidity: function getAssetPriceStatus(address _asset) returns(bool)
func (_Oracle *OracleSession) GetAssetPriceStatus(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetAssetPriceStatus(&_Oracle.TransactOpts, _asset)
}

// GetAssetPriceStatus is a paid mutator transaction binding the contract method 0x0a2b53bf.
//
// Solidity: function getAssetPriceStatus(address _asset) returns(bool)
func (_Oracle *OracleTransactorSession) GetAssetPriceStatus(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetAssetPriceStatus(&_Oracle.TransactOpts, _asset)
}

// GetUnderlyingPrice is a paid mutator transaction binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address _asset) returns(uint256 _price)
func (_Oracle *OracleTransactor) GetUnderlyingPrice(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "getUnderlyingPrice", _asset)
}

// GetUnderlyingPrice is a paid mutator transaction binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address _asset) returns(uint256 _price)
func (_Oracle *OracleSession) GetUnderlyingPrice(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.TransactOpts, _asset)
}

// GetUnderlyingPrice is a paid mutator transaction binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address _asset) returns(uint256 _price)
func (_Oracle *OracleTransactorSession) GetUnderlyingPrice(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetUnderlyingPrice(&_Oracle.TransactOpts, _asset)
}

// GetUnderlyingPriceAndStatus is a paid mutator transaction binding the contract method 0xed5cf2dc.
//
// Solidity: function getUnderlyingPriceAndStatus(address _asset) returns(uint256 _price, bool _status)
func (_Oracle *OracleTransactor) GetUnderlyingPriceAndStatus(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "getUnderlyingPriceAndStatus", _asset)
}

// GetUnderlyingPriceAndStatus is a paid mutator transaction binding the contract method 0xed5cf2dc.
//
// Solidity: function getUnderlyingPriceAndStatus(address _asset) returns(uint256 _price, bool _status)
func (_Oracle *OracleSession) GetUnderlyingPriceAndStatus(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetUnderlyingPriceAndStatus(&_Oracle.TransactOpts, _asset)
}

// GetUnderlyingPriceAndStatus is a paid mutator transaction binding the contract method 0xed5cf2dc.
//
// Solidity: function getUnderlyingPriceAndStatus(address _asset) returns(uint256 _price, bool _status)
func (_Oracle *OracleTransactorSession) GetUnderlyingPriceAndStatus(_asset common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.GetUnderlyingPriceAndStatus(&_Oracle.TransactOpts, _asset)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _poster) returns()
func (_Oracle *OracleTransactor) Initialize(opts *bind.TransactOpts, _poster common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "initialize", _poster)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _poster) returns()
func (_Oracle *OracleSession) Initialize(_poster common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, _poster)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _poster) returns()
func (_Oracle *OracleTransactorSession) Initialize(_poster common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts, _poster)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _asset, uint256 _requestedPrice) returns(bool)
func (_Oracle *OracleTransactor) SetPrice(opts *bind.TransactOpts, _asset common.Address, _requestedPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setPrice", _asset, _requestedPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _asset, uint256 _requestedPrice) returns(bool)
func (_Oracle *OracleSession) SetPrice(_asset common.Address, _requestedPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrice(&_Oracle.TransactOpts, _asset, _requestedPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _asset, uint256 _requestedPrice) returns(bool)
func (_Oracle *OracleTransactorSession) SetPrice(_asset common.Address, _requestedPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrice(&_Oracle.TransactOpts, _asset, _requestedPrice)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] _assets, uint256[] _requestedPrices) returns(bool[])
func (_Oracle *OracleTransactor) SetPrices(opts *bind.TransactOpts, _assets []common.Address, _requestedPrices []*big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setPrices", _assets, _requestedPrices)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] _assets, uint256[] _requestedPrices) returns(bool[])
func (_Oracle *OracleSession) SetPrices(_assets []common.Address, _requestedPrices []*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrices(&_Oracle.TransactOpts, _assets, _requestedPrices)
}

// SetPrices is a paid mutator transaction binding the contract method 0x4352fa9f.
//
// Solidity: function setPrices(address[] _assets, uint256[] _requestedPrices) returns(bool[])
func (_Oracle *OracleTransactorSession) SetPrices(_assets []common.Address, _requestedPrices []*big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetPrices(&_Oracle.TransactOpts, _assets, _requestedPrices)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Oracle *OracleTransactorSession) Receive() (*types.Transaction, error) {
	return _Oracle.Contract.Receive(&_Oracle.TransactOpts)
}

// OracleNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Oracle contract.
type OracleNewOwnerIterator struct {
	Event *OracleNewOwner // Event containing the contract specifics and raw log

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
func (it *OracleNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewOwner)
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
		it.Event = new(OracleNewOwner)
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
func (it *OracleNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewOwner represents a NewOwner event raised by the Oracle contract.
type OracleNewOwner struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) FilterNewOwner(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OracleNewOwnerIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleNewOwnerIterator{contract: _Oracle.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *OracleNewOwner, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewOwner)
				if err := _Oracle.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_Oracle *OracleFilterer) ParseNewOwner(log types.Log) (*OracleNewOwner, error) {
	event := new(OracleNewOwner)
	if err := _Oracle.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the Oracle contract.
type OracleNewPendingOwnerIterator struct {
	Event *OracleNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *OracleNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewPendingOwner)
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
		it.Event = new(OracleNewPendingOwner)
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
func (it *OracleNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewPendingOwner represents a NewPendingOwner event raised by the Oracle contract.
type OracleNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Oracle *OracleFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*OracleNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OracleNewPendingOwnerIterator{contract: _Oracle.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Oracle *OracleFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *OracleNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewPendingOwner)
				if err := _Oracle.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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

// ParseNewPendingOwner is a log parse operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Oracle *OracleFilterer) ParseNewPendingOwner(log types.Log) (*OracleNewPendingOwner, error) {
	event := new(OracleNewPendingOwner)
	if err := _Oracle.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleNewPosterIterator is returned from FilterNewPoster and is used to iterate over the raw logs and unpacked data for NewPoster events raised by the Oracle contract.
type OracleNewPosterIterator struct {
	Event *OracleNewPoster // Event containing the contract specifics and raw log

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
func (it *OracleNewPosterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewPoster)
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
		it.Event = new(OracleNewPoster)
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
func (it *OracleNewPosterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewPosterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewPoster represents a NewPoster event raised by the Oracle contract.
type OracleNewPoster struct {
	OldPoster common.Address
	NewPoster common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPoster is a free log retrieval operation binding the contract event 0xd8209502cb6dab3843acd5ce51ae027a6f3488c57c2d4aedac2d94501dfcfc8d.
//
// Solidity: event NewPoster(address oldPoster, address newPoster)
func (_Oracle *OracleFilterer) FilterNewPoster(opts *bind.FilterOpts) (*OracleNewPosterIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewPoster")
	if err != nil {
		return nil, err
	}
	return &OracleNewPosterIterator{contract: _Oracle.contract, event: "NewPoster", logs: logs, sub: sub}, nil
}

// WatchNewPoster is a free log subscription operation binding the contract event 0xd8209502cb6dab3843acd5ce51ae027a6f3488c57c2d4aedac2d94501dfcfc8d.
//
// Solidity: event NewPoster(address oldPoster, address newPoster)
func (_Oracle *OracleFilterer) WatchNewPoster(opts *bind.WatchOpts, sink chan<- *OracleNewPoster) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewPoster")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewPoster)
				if err := _Oracle.contract.UnpackLog(event, "NewPoster", log); err != nil {
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

// ParseNewPoster is a log parse operation binding the contract event 0xd8209502cb6dab3843acd5ce51ae027a6f3488c57c2d4aedac2d94501dfcfc8d.
//
// Solidity: event NewPoster(address oldPoster, address newPoster)
func (_Oracle *OracleFilterer) ParseNewPoster(log types.Log) (*OracleNewPoster, error) {
	event := new(OracleNewPoster)
	if err := _Oracle.contract.UnpackLog(event, "NewPoster", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleSetAssetPriceModelIterator is returned from FilterSetAssetPriceModel and is used to iterate over the raw logs and unpacked data for SetAssetPriceModel events raised by the Oracle contract.
type OracleSetAssetPriceModelIterator struct {
	Event *OracleSetAssetPriceModel // Event containing the contract specifics and raw log

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
func (it *OracleSetAssetPriceModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSetAssetPriceModel)
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
		it.Event = new(OracleSetAssetPriceModel)
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
func (it *OracleSetAssetPriceModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSetAssetPriceModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSetAssetPriceModel represents a SetAssetPriceModel event raised by the Oracle contract.
type OracleSetAssetPriceModel struct {
	Asset      common.Address
	PriceModel common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetAssetPriceModel is a free log retrieval operation binding the contract event 0xac57e3171e85d835e3f74f03852ac5327a0eab7e582f3cbe138edbf744772562.
//
// Solidity: event SetAssetPriceModel(address asset, address priceModel)
func (_Oracle *OracleFilterer) FilterSetAssetPriceModel(opts *bind.FilterOpts) (*OracleSetAssetPriceModelIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "SetAssetPriceModel")
	if err != nil {
		return nil, err
	}
	return &OracleSetAssetPriceModelIterator{contract: _Oracle.contract, event: "SetAssetPriceModel", logs: logs, sub: sub}, nil
}

// WatchSetAssetPriceModel is a free log subscription operation binding the contract event 0xac57e3171e85d835e3f74f03852ac5327a0eab7e582f3cbe138edbf744772562.
//
// Solidity: event SetAssetPriceModel(address asset, address priceModel)
func (_Oracle *OracleFilterer) WatchSetAssetPriceModel(opts *bind.WatchOpts, sink chan<- *OracleSetAssetPriceModel) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "SetAssetPriceModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSetAssetPriceModel)
				if err := _Oracle.contract.UnpackLog(event, "SetAssetPriceModel", log); err != nil {
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

// ParseSetAssetPriceModel is a log parse operation binding the contract event 0xac57e3171e85d835e3f74f03852ac5327a0eab7e582f3cbe138edbf744772562.
//
// Solidity: event SetAssetPriceModel(address asset, address priceModel)
func (_Oracle *OracleFilterer) ParseSetAssetPriceModel(log types.Log) (*OracleSetAssetPriceModel, error) {
	event := new(OracleSetAssetPriceModel)
	if err := _Oracle.contract.UnpackLog(event, "SetAssetPriceModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleSetPausedIterator is returned from FilterSetPaused and is used to iterate over the raw logs and unpacked data for SetPaused events raised by the Oracle contract.
type OracleSetPausedIterator struct {
	Event *OracleSetPaused // Event containing the contract specifics and raw log

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
func (it *OracleSetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleSetPaused)
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
		it.Event = new(OracleSetPaused)
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
func (it *OracleSetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleSetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleSetPaused represents a SetPaused event raised by the Oracle contract.
type OracleSetPaused struct {
	NewState bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPaused is a free log retrieval operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_Oracle *OracleFilterer) FilterSetPaused(opts *bind.FilterOpts) (*OracleSetPausedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "SetPaused")
	if err != nil {
		return nil, err
	}
	return &OracleSetPausedIterator{contract: _Oracle.contract, event: "SetPaused", logs: logs, sub: sub}, nil
}

// WatchSetPaused is a free log subscription operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_Oracle *OracleFilterer) WatchSetPaused(opts *bind.WatchOpts, sink chan<- *OracleSetPaused) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "SetPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleSetPaused)
				if err := _Oracle.contract.UnpackLog(event, "SetPaused", log); err != nil {
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

// ParseSetPaused is a log parse operation binding the contract event 0x3c70af01296aef045b2f5c9d3c30b05d4428fd257145b9c7fcd76418e65b5980.
//
// Solidity: event SetPaused(bool newState)
func (_Oracle *OracleFilterer) ParseSetPaused(log types.Log) (*OracleSetPaused, error) {
	event := new(OracleSetPaused)
	if err := _Oracle.contract.UnpackLog(event, "SetPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
