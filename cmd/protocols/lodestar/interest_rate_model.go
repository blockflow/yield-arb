// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lodestar

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRatePerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kink_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRatePerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplierPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kink\",\"type\":\"uint256\"}],\"name\":\"NewInterestParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocksPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jumpMultiplierPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRatePerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kink_\",\"type\":\"uint256\"}],\"name\":\"updateJumpRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"utilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) BaseRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "baseRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) BaseRatePerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.BaseRatePerBlock(&_InterestRateModel.CallOpts)
}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) BaseRatePerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.BaseRatePerBlock(&_InterestRateModel.CallOpts)
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
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) GetBorrowRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getBorrowRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) GetSupplyRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getSupplyRate", cash, borrows, reserves, reserveFactorMantissa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetSupplyRate(&_InterestRateModel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.GetSupplyRate(&_InterestRateModel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
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
// Solidity: function isInterestRateModel() view returns(bool)
func (_InterestRateModel *InterestRateModelSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_InterestRateModel *InterestRateModelCallerSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) JumpMultiplierPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "jumpMultiplierPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) JumpMultiplierPerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.JumpMultiplierPerBlock(&_InterestRateModel.CallOpts)
}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) JumpMultiplierPerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.JumpMultiplierPerBlock(&_InterestRateModel.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) Kink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "kink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) Kink() (*big.Int, error) {
	return _InterestRateModel.Contract.Kink(&_InterestRateModel.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) Kink() (*big.Int, error) {
	return _InterestRateModel.Contract.Kink(&_InterestRateModel.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) MultiplierPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "multiplierPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelSession) MultiplierPerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.MultiplierPerBlock(&_InterestRateModel.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) MultiplierPerBlock() (*big.Int, error) {
	return _InterestRateModel.Contract.MultiplierPerBlock(&_InterestRateModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InterestRateModel *InterestRateModelCallerSession) Owner() (common.Address, error) {
	return _InterestRateModel.Contract.Owner(&_InterestRateModel.CallOpts)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_InterestRateModel *InterestRateModelCaller) UtilizationRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "utilizationRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_InterestRateModel *InterestRateModelSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.UtilizationRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_InterestRateModel *InterestRateModelCallerSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _InterestRateModel.Contract.UtilizationRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_InterestRateModel *InterestRateModelTransactor) UpdateJumpRateModel(opts *bind.TransactOpts, baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.contract.Transact(opts, "updateJumpRateModel", baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_InterestRateModel *InterestRateModelSession) UpdateJumpRateModel(baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.Contract.UpdateJumpRateModel(&_InterestRateModel.TransactOpts, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_InterestRateModel *InterestRateModelTransactorSession) UpdateJumpRateModel(baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _InterestRateModel.Contract.UpdateJumpRateModel(&_InterestRateModel.TransactOpts, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// InterestRateModelNewInterestParamsIterator is returned from FilterNewInterestParams and is used to iterate over the raw logs and unpacked data for NewInterestParams events raised by the InterestRateModel contract.
type InterestRateModelNewInterestParamsIterator struct {
	Event *InterestRateModelNewInterestParams // Event containing the contract specifics and raw log

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
func (it *InterestRateModelNewInterestParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterestRateModelNewInterestParams)
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
		it.Event = new(InterestRateModelNewInterestParams)
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
func (it *InterestRateModelNewInterestParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterestRateModelNewInterestParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterestRateModelNewInterestParams represents a NewInterestParams event raised by the InterestRateModel contract.
type InterestRateModelNewInterestParams struct {
	BaseRatePerBlock       *big.Int
	MultiplierPerBlock     *big.Int
	JumpMultiplierPerBlock *big.Int
	Kink                   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewInterestParams is a free log retrieval operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_InterestRateModel *InterestRateModelFilterer) FilterNewInterestParams(opts *bind.FilterOpts) (*InterestRateModelNewInterestParamsIterator, error) {

	logs, sub, err := _InterestRateModel.contract.FilterLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return &InterestRateModelNewInterestParamsIterator{contract: _InterestRateModel.contract, event: "NewInterestParams", logs: logs, sub: sub}, nil
}

// WatchNewInterestParams is a free log subscription operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_InterestRateModel *InterestRateModelFilterer) WatchNewInterestParams(opts *bind.WatchOpts, sink chan<- *InterestRateModelNewInterestParams) (event.Subscription, error) {

	logs, sub, err := _InterestRateModel.contract.WatchLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterestRateModelNewInterestParams)
				if err := _InterestRateModel.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
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

// ParseNewInterestParams is a log parse operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_InterestRateModel *InterestRateModelFilterer) ParseNewInterestParams(log types.Log) (*InterestRateModelNewInterestParams, error) {
	event := new(InterestRateModelNewInterestParams)
	if err := _InterestRateModel.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
