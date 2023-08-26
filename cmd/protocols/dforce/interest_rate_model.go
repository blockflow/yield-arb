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

// InterestRateModelMetaData contains all meta data concerning the InterestRateModel contract.
var InterestRateModelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocksPerYear\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocksPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slope_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slope_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InterestRateModelABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateModelMetaData.ABI instead.
var InterestRateModelABI = InterestRateModelMetaData.ABI

// InterestRateModel is an auto generated Go binding around an Ethereum contract.
type InterestRateModel struct {
	InterestRateModelCaller     // Read-only binding to the contract
	InterestRateModelTransactor // Write-only binding to the contract
	InterestRateModelFilterer   // Log filterer for contract events
}

// InterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelSession struct {
	Contract     *InterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelCallerSession struct {
	Contract *InterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelTransactorSession struct {
	Contract     *InterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelRaw struct {
	Contract *InterestRateModel // Generic contract binding to access the raw methods on
}

// InterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelCallerRaw struct {
	Contract *InterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelTransactorRaw struct {
	Contract *InterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModel creates a new instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModel(address common.Address, backend bind.ContractBackend) (*InterestRateModel, error) {
	contract, err := bindInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModel{InterestRateModelCaller: InterestRateModelCaller{contract: contract}, InterestRateModelTransactor: InterestRateModelTransactor{contract: contract}, InterestRateModelFilterer: InterestRateModelFilterer{contract: contract}}, nil
}

// NewInterestRateModelCaller creates a new read-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*InterestRateModelCaller, error) {
	contract, err := bindInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelCaller{contract: contract}, nil
}

// NewInterestRateModelTransactor creates a new write-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelTransactor, error) {
	contract, err := bindInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelTransactor{contract: contract}, nil
}

// NewInterestRateModelFilterer creates a new log filterer instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelFilterer, error) {
	contract, err := bindInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelFilterer{contract: contract}, nil
}

// bindInterestRateModel binds a generic wrapper to an already deployed contract.
func bindInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateModelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.InterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) Base(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) Base() (*big.Int, error) {
	return _InterestRateModel.Contract.Base(&_InterestRateModel.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) Base() (*big.Int, error) {
	return _InterestRateModel.Contract.Base(&_InterestRateModel.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) BlocksPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "blocksPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) BlocksPerYear() (*big.Int, error) {
	return _InterestRateModel.Contract.BlocksPerYear(&_InterestRateModel.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) BlocksPerYear() (*big.Int, error) {
	return _InterestRateModel.Contract.BlocksPerYear(&_InterestRateModel.CallOpts)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 _balance, uint256 _borrows, uint256 _reserves) view returns(uint256 _borrowRate)
func (_InterestRateModel *InterestRateModelCaller) GetBorrowRate(opts *bind.CallOpts, _balance *big.Int, _borrows *big.Int, _reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getBorrowRate", _balance, _borrows, _reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 _balance, uint256 _borrows, uint256 _reserves) view returns(uint256 _borrowRate)
func (_InterestRateModel *InterestRateModelSession) GetBorrowRate(_balance *big.Int, _borrows *big.Int, _reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, _balance, _borrows, _reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 _balance, uint256 _borrows, uint256 _reserves) view returns(uint256 _borrowRate)
func (_InterestRateModel *InterestRateModelCallerSession) GetBorrowRate(_balance *big.Int, _borrows *big.Int, _reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, _balance, _borrows, _reserves)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() pure returns(bool)
func (_InterestRateModel *InterestRateModelCaller) IsInterestRateModel(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "isInterestRateModel")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() pure returns(bool)
func (_InterestRateModel *InterestRateModelSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() pure returns(bool)
func (_InterestRateModel *InterestRateModelCallerSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// Optimal is a free data retrieval call binding the contract method 0x4d24044c.
//
// Solidity: function optimal() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) Optimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "optimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Optimal is a free data retrieval call binding the contract method 0x4d24044c.
//
// Solidity: function optimal() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) Optimal() (*big.Int, error) {
	return _InterestRateModel.Contract.Optimal(&_InterestRateModel.CallOpts)
}

// Optimal is a free data retrieval call binding the contract method 0x4d24044c.
//
// Solidity: function optimal() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) Optimal() (*big.Int, error) {
	return _InterestRateModel.Contract.Optimal(&_InterestRateModel.CallOpts)
}

// Slope1 is a free data retrieval call binding the contract method 0x7dbefcb6.
//
// Solidity: function slope_1() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) Slope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "slope_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slope1 is a free data retrieval call binding the contract method 0x7dbefcb6.
//
// Solidity: function slope_1() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) Slope1() (*big.Int, error) {
	return _InterestRateModel.Contract.Slope1(&_InterestRateModel.CallOpts)
}

// Slope1 is a free data retrieval call binding the contract method 0x7dbefcb6.
//
// Solidity: function slope_1() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) Slope1() (*big.Int, error) {
	return _InterestRateModel.Contract.Slope1(&_InterestRateModel.CallOpts)
}

// Slope2 is a free data retrieval call binding the contract method 0xe7bed802.
//
// Solidity: function slope_2() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) Slope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "slope_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Slope2 is a free data retrieval call binding the contract method 0xe7bed802.
//
// Solidity: function slope_2() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) Slope2() (*big.Int, error) {
	return _InterestRateModel.Contract.Slope2(&_InterestRateModel.CallOpts)
}

// Slope2 is a free data retrieval call binding the contract method 0xe7bed802.
//
// Solidity: function slope_2() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) Slope2() (*big.Int, error) {
	return _InterestRateModel.Contract.Slope2(&_InterestRateModel.CallOpts)
}
