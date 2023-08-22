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

// ITokenMetaData contains all meta data concerning the IToken contract.
var ITokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountInterestIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loaner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flashloanFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Flashloan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIControllerInterface\",\"name\":\"oldController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIControllerInterface\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashloanFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashloanFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProtocolFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFeeRatio\",\"type\":\"uint256\"}],\"name\":\"NewFlashloanFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashloanFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashloanFeeRatio\",\"type\":\"uint256\"}],\"name\":\"NewFlashloanFeeRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIInterestRateModelInterface\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIInterestRateModelInterface\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProtocolFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFeeRatio\",\"type\":\"uint256\"}],\"name\":\"NewProtocolFeeRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveRatio\",\"type\":\"uint256\"}],\"name\":\"NewReserveRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemiTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemUnderlyingAmount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountInterestIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"}],\"name\":\"UpdateInterest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIControllerInterface\",\"name\":\"_newController\",\"type\":\"address\"}],\"name\":\"_setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIInterestRateModelInterface\",\"name\":\"_newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFlashloanFeeRatio\",\"type\":\"uint256\"}],\"name\":\"_setNewFlashloanFeeRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newProtocolFeeRatio\",\"type\":\"uint256\"}],\"name\":\"_setNewProtocolFeeRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newReserveRatio\",\"type\":\"uint256\"}],\"name\":\"_setNewReserveRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"_setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"_withdrawReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"borrowSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashloanFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"contractIControllerInterface\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"contractIInterestRateModelInterface\",\"name\":\"_interestRateModel\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractIInterestRateModelInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isiToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_assetCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAndEnterMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintForSelfAndEnterMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemiToken\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemUnderlying\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMetaData.ABI instead.
var ITokenABI = ITokenMetaData.ABI

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IToken *ITokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IToken *ITokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IToken.Contract.DOMAINSEPARATOR(&_IToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IToken *ITokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IToken.Contract.DOMAINSEPARATOR(&_IToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_IToken *ITokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_IToken *ITokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IToken.Contract.PERMITTYPEHASH(&_IToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_IToken *ITokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _IToken.Contract.PERMITTYPEHASH(&_IToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IToken *ITokenCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IToken *ITokenSession) AccrualBlockNumber() (*big.Int, error) {
	return _IToken.Contract.AccrualBlockNumber(&_IToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_IToken *ITokenCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _IToken.Contract.AccrualBlockNumber(&_IToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IToken *ITokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IToken *ITokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_IToken *ITokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IToken *ITokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IToken *ITokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IToken *ITokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, arg0)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address _account) view returns(uint256)
func (_IToken *ITokenCaller) BorrowBalanceStored(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "borrowBalanceStored", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address _account) view returns(uint256)
func (_IToken *ITokenSession) BorrowBalanceStored(_account common.Address) (*big.Int, error) {
	return _IToken.Contract.BorrowBalanceStored(&_IToken.CallOpts, _account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address _account) view returns(uint256)
func (_IToken *ITokenCallerSession) BorrowBalanceStored(_account common.Address) (*big.Int, error) {
	return _IToken.Contract.BorrowBalanceStored(&_IToken.CallOpts, _account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IToken *ITokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IToken *ITokenSession) BorrowIndex() (*big.Int, error) {
	return _IToken.Contract.BorrowIndex(&_IToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_IToken *ITokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _IToken.Contract.BorrowIndex(&_IToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IToken *ITokenCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IToken *ITokenSession) BorrowRatePerBlock() (*big.Int, error) {
	return _IToken.Contract.BorrowRatePerBlock(&_IToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_IToken *ITokenCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _IToken.Contract.BorrowRatePerBlock(&_IToken.CallOpts)
}

// BorrowSnapshot is a free data retrieval call binding the contract method 0x37d33618.
//
// Solidity: function borrowSnapshot(address _account) view returns(uint256, uint256)
func (_IToken *ITokenCaller) BorrowSnapshot(opts *bind.CallOpts, _account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "borrowSnapshot", _account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BorrowSnapshot is a free data retrieval call binding the contract method 0x37d33618.
//
// Solidity: function borrowSnapshot(address _account) view returns(uint256, uint256)
func (_IToken *ITokenSession) BorrowSnapshot(_account common.Address) (*big.Int, *big.Int, error) {
	return _IToken.Contract.BorrowSnapshot(&_IToken.CallOpts, _account)
}

// BorrowSnapshot is a free data retrieval call binding the contract method 0x37d33618.
//
// Solidity: function borrowSnapshot(address _account) view returns(uint256, uint256)
func (_IToken *ITokenCallerSession) BorrowSnapshot(_account common.Address) (*big.Int, *big.Int, error) {
	return _IToken.Contract.BorrowSnapshot(&_IToken.CallOpts, _account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_IToken *ITokenCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_IToken *ITokenSession) Controller() (common.Address, error) {
	return _IToken.Contract.Controller(&_IToken.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_IToken *ITokenCallerSession) Controller() (common.Address, error) {
	return _IToken.Contract.Controller(&_IToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IToken *ITokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IToken *ITokenSession) Decimals() (uint8, error) {
	return _IToken.Contract.Decimals(&_IToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IToken *ITokenCallerSession) Decimals() (uint8, error) {
	return _IToken.Contract.Decimals(&_IToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IToken *ITokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IToken *ITokenSession) ExchangeRateStored() (*big.Int, error) {
	return _IToken.Contract.ExchangeRateStored(&_IToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_IToken *ITokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _IToken.Contract.ExchangeRateStored(&_IToken.CallOpts)
}

// FlashloanFeeRatio is a free data retrieval call binding the contract method 0x197523f8.
//
// Solidity: function flashloanFeeRatio() view returns(uint256)
func (_IToken *ITokenCaller) FlashloanFeeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "flashloanFeeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashloanFeeRatio is a free data retrieval call binding the contract method 0x197523f8.
//
// Solidity: function flashloanFeeRatio() view returns(uint256)
func (_IToken *ITokenSession) FlashloanFeeRatio() (*big.Int, error) {
	return _IToken.Contract.FlashloanFeeRatio(&_IToken.CallOpts)
}

// FlashloanFeeRatio is a free data retrieval call binding the contract method 0x197523f8.
//
// Solidity: function flashloanFeeRatio() view returns(uint256)
func (_IToken *ITokenCallerSession) FlashloanFeeRatio() (*big.Int, error) {
	return _IToken.Contract.FlashloanFeeRatio(&_IToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IToken *ITokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IToken *ITokenSession) GetCash() (*big.Int, error) {
	return _IToken.Contract.GetCash(&_IToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_IToken *ITokenCallerSession) GetCash() (*big.Int, error) {
	return _IToken.Contract.GetCash(&_IToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IToken *ITokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IToken *ITokenSession) InterestRateModel() (common.Address, error) {
	return _IToken.Contract.InterestRateModel(&_IToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_IToken *ITokenCallerSession) InterestRateModel() (common.Address, error) {
	return _IToken.Contract.InterestRateModel(&_IToken.CallOpts)
}

// IsSupported is a free data retrieval call binding the contract method 0x74427937.
//
// Solidity: function isSupported() view returns(bool)
func (_IToken *ITokenCaller) IsSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "isSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupported is a free data retrieval call binding the contract method 0x74427937.
//
// Solidity: function isSupported() view returns(bool)
func (_IToken *ITokenSession) IsSupported() (bool, error) {
	return _IToken.Contract.IsSupported(&_IToken.CallOpts)
}

// IsSupported is a free data retrieval call binding the contract method 0x74427937.
//
// Solidity: function isSupported() view returns(bool)
func (_IToken *ITokenCallerSession) IsSupported() (bool, error) {
	return _IToken.Contract.IsSupported(&_IToken.CallOpts)
}

// IsiToken is a free data retrieval call binding the contract method 0x621fd507.
//
// Solidity: function isiToken() pure returns(bool)
func (_IToken *ITokenCaller) IsiToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "isiToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsiToken is a free data retrieval call binding the contract method 0x621fd507.
//
// Solidity: function isiToken() pure returns(bool)
func (_IToken *ITokenSession) IsiToken() (bool, error) {
	return _IToken.Contract.IsiToken(&_IToken.CallOpts)
}

// IsiToken is a free data retrieval call binding the contract method 0x621fd507.
//
// Solidity: function isiToken() pure returns(bool)
func (_IToken *ITokenCallerSession) IsiToken() (bool, error) {
	return _IToken.Contract.IsiToken(&_IToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IToken *ITokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IToken *ITokenSession) Name() (string, error) {
	return _IToken.Contract.Name(&_IToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IToken *ITokenCallerSession) Name() (string, error) {
	return _IToken.Contract.Name(&_IToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IToken *ITokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IToken *ITokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _IToken.Contract.Nonces(&_IToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_IToken *ITokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _IToken.Contract.Nonces(&_IToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IToken *ITokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IToken *ITokenSession) Owner() (common.Address, error) {
	return _IToken.Contract.Owner(&_IToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IToken *ITokenCallerSession) Owner() (common.Address, error) {
	return _IToken.Contract.Owner(&_IToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IToken *ITokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IToken *ITokenSession) PendingOwner() (common.Address, error) {
	return _IToken.Contract.PendingOwner(&_IToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IToken *ITokenCallerSession) PendingOwner() (common.Address, error) {
	return _IToken.Contract.PendingOwner(&_IToken.CallOpts)
}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint256)
func (_IToken *ITokenCaller) ProtocolFeeRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "protocolFeeRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint256)
func (_IToken *ITokenSession) ProtocolFeeRatio() (*big.Int, error) {
	return _IToken.Contract.ProtocolFeeRatio(&_IToken.CallOpts)
}

// ProtocolFeeRatio is a free data retrieval call binding the contract method 0x7a27d9f6.
//
// Solidity: function protocolFeeRatio() view returns(uint256)
func (_IToken *ITokenCallerSession) ProtocolFeeRatio() (*big.Int, error) {
	return _IToken.Contract.ProtocolFeeRatio(&_IToken.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_IToken *ITokenCaller) ReserveRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "reserveRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_IToken *ITokenSession) ReserveRatio() (*big.Int, error) {
	return _IToken.Contract.ReserveRatio(&_IToken.CallOpts)
}

// ReserveRatio is a free data retrieval call binding the contract method 0x0c7d5cd8.
//
// Solidity: function reserveRatio() view returns(uint256)
func (_IToken *ITokenCallerSession) ReserveRatio() (*big.Int, error) {
	return _IToken.Contract.ReserveRatio(&_IToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IToken *ITokenCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IToken *ITokenSession) SupplyRatePerBlock() (*big.Int, error) {
	return _IToken.Contract.SupplyRatePerBlock(&_IToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_IToken *ITokenCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _IToken.Contract.SupplyRatePerBlock(&_IToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IToken *ITokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IToken *ITokenSession) Symbol() (string, error) {
	return _IToken.Contract.Symbol(&_IToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IToken *ITokenCallerSession) Symbol() (string, error) {
	return _IToken.Contract.Symbol(&_IToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IToken *ITokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IToken *ITokenSession) TotalBorrows() (*big.Int, error) {
	return _IToken.Contract.TotalBorrows(&_IToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_IToken *ITokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _IToken.Contract.TotalBorrows(&_IToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IToken *ITokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IToken *ITokenSession) TotalReserves() (*big.Int, error) {
	return _IToken.Contract.TotalReserves(&_IToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_IToken *ITokenCallerSession) TotalReserves() (*big.Int, error) {
	return _IToken.Contract.TotalReserves(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_IToken *ITokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_IToken *ITokenSession) Underlying() (common.Address, error) {
	return _IToken.Contract.Underlying(&_IToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_IToken *ITokenCallerSession) Underlying() (common.Address, error) {
	return _IToken.Contract.Underlying(&_IToken.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_IToken *ITokenTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_IToken *ITokenSession) AcceptOwner() (*types.Transaction, error) {
	return _IToken.Contract.AcceptOwner(&_IToken.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_IToken *ITokenTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _IToken.Contract.AcceptOwner(&_IToken.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _newController) returns()
func (_IToken *ITokenTransactor) SetController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setController", _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _newController) returns()
func (_IToken *ITokenSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetController(&_IToken.TransactOpts, _newController)
}

// SetController is a paid mutator transaction binding the contract method 0x83de424e.
//
// Solidity: function _setController(address _newController) returns()
func (_IToken *ITokenTransactorSession) SetController(_newController common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetController(&_IToken.TransactOpts, _newController)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address _newInterestRateModel) returns()
func (_IToken *ITokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, _newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setInterestRateModel", _newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address _newInterestRateModel) returns()
func (_IToken *ITokenSession) SetInterestRateModel(_newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetInterestRateModel(&_IToken.TransactOpts, _newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address _newInterestRateModel) returns()
func (_IToken *ITokenTransactorSession) SetInterestRateModel(_newInterestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetInterestRateModel(&_IToken.TransactOpts, _newInterestRateModel)
}

// SetNewFlashloanFeeRatio is a paid mutator transaction binding the contract method 0x508fe21f.
//
// Solidity: function _setNewFlashloanFeeRatio(uint256 _newFlashloanFeeRatio) returns()
func (_IToken *ITokenTransactor) SetNewFlashloanFeeRatio(opts *bind.TransactOpts, _newFlashloanFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setNewFlashloanFeeRatio", _newFlashloanFeeRatio)
}

// SetNewFlashloanFeeRatio is a paid mutator transaction binding the contract method 0x508fe21f.
//
// Solidity: function _setNewFlashloanFeeRatio(uint256 _newFlashloanFeeRatio) returns()
func (_IToken *ITokenSession) SetNewFlashloanFeeRatio(_newFlashloanFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewFlashloanFeeRatio(&_IToken.TransactOpts, _newFlashloanFeeRatio)
}

// SetNewFlashloanFeeRatio is a paid mutator transaction binding the contract method 0x508fe21f.
//
// Solidity: function _setNewFlashloanFeeRatio(uint256 _newFlashloanFeeRatio) returns()
func (_IToken *ITokenTransactorSession) SetNewFlashloanFeeRatio(_newFlashloanFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewFlashloanFeeRatio(&_IToken.TransactOpts, _newFlashloanFeeRatio)
}

// SetNewProtocolFeeRatio is a paid mutator transaction binding the contract method 0x269aafee.
//
// Solidity: function _setNewProtocolFeeRatio(uint256 _newProtocolFeeRatio) returns()
func (_IToken *ITokenTransactor) SetNewProtocolFeeRatio(opts *bind.TransactOpts, _newProtocolFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setNewProtocolFeeRatio", _newProtocolFeeRatio)
}

// SetNewProtocolFeeRatio is a paid mutator transaction binding the contract method 0x269aafee.
//
// Solidity: function _setNewProtocolFeeRatio(uint256 _newProtocolFeeRatio) returns()
func (_IToken *ITokenSession) SetNewProtocolFeeRatio(_newProtocolFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewProtocolFeeRatio(&_IToken.TransactOpts, _newProtocolFeeRatio)
}

// SetNewProtocolFeeRatio is a paid mutator transaction binding the contract method 0x269aafee.
//
// Solidity: function _setNewProtocolFeeRatio(uint256 _newProtocolFeeRatio) returns()
func (_IToken *ITokenTransactorSession) SetNewProtocolFeeRatio(_newProtocolFeeRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewProtocolFeeRatio(&_IToken.TransactOpts, _newProtocolFeeRatio)
}

// SetNewReserveRatio is a paid mutator transaction binding the contract method 0xc9f30e53.
//
// Solidity: function _setNewReserveRatio(uint256 _newReserveRatio) returns()
func (_IToken *ITokenTransactor) SetNewReserveRatio(opts *bind.TransactOpts, _newReserveRatio *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setNewReserveRatio", _newReserveRatio)
}

// SetNewReserveRatio is a paid mutator transaction binding the contract method 0xc9f30e53.
//
// Solidity: function _setNewReserveRatio(uint256 _newReserveRatio) returns()
func (_IToken *ITokenSession) SetNewReserveRatio(_newReserveRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewReserveRatio(&_IToken.TransactOpts, _newReserveRatio)
}

// SetNewReserveRatio is a paid mutator transaction binding the contract method 0xc9f30e53.
//
// Solidity: function _setNewReserveRatio(uint256 _newReserveRatio) returns()
func (_IToken *ITokenTransactorSession) SetNewReserveRatio(_newReserveRatio *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.SetNewReserveRatio(&_IToken.TransactOpts, _newReserveRatio)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_IToken *ITokenTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_IToken *ITokenSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetPendingOwner(&_IToken.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_IToken *ITokenTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _IToken.Contract.SetPendingOwner(&_IToken.TransactOpts, newPendingOwner)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xb3efd5b4.
//
// Solidity: function _withdrawReserves(uint256 _withdrawAmount) returns()
func (_IToken *ITokenTransactor) WithdrawReserves(opts *bind.TransactOpts, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "_withdrawReserves", _withdrawAmount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xb3efd5b4.
//
// Solidity: function _withdrawReserves(uint256 _withdrawAmount) returns()
func (_IToken *ITokenSession) WithdrawReserves(_withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.WithdrawReserves(&_IToken.TransactOpts, _withdrawAmount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xb3efd5b4.
//
// Solidity: function _withdrawReserves(uint256 _withdrawAmount) returns()
func (_IToken *ITokenTransactorSession) WithdrawReserves(_withdrawAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.WithdrawReserves(&_IToken.TransactOpts, _withdrawAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) returns(uint256)
func (_IToken *ITokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "balanceOfUnderlying", _account)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) returns(uint256)
func (_IToken *ITokenSession) BalanceOfUnderlying(_account common.Address) (*types.Transaction, error) {
	return _IToken.Contract.BalanceOfUnderlying(&_IToken.TransactOpts, _account)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address _account) returns(uint256)
func (_IToken *ITokenTransactorSession) BalanceOfUnderlying(_account common.Address) (*types.Transaction, error) {
	return _IToken.Contract.BalanceOfUnderlying(&_IToken.TransactOpts, _account)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _borrowAmount) returns()
func (_IToken *ITokenTransactor) Borrow(opts *bind.TransactOpts, _borrowAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "borrow", _borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _borrowAmount) returns()
func (_IToken *ITokenSession) Borrow(_borrowAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Borrow(&_IToken.TransactOpts, _borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 _borrowAmount) returns()
func (_IToken *ITokenTransactorSession) Borrow(_borrowAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Borrow(&_IToken.TransactOpts, _borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address _account) returns(uint256)
func (_IToken *ITokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "borrowBalanceCurrent", _account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address _account) returns(uint256)
func (_IToken *ITokenSession) BorrowBalanceCurrent(_account common.Address) (*types.Transaction, error) {
	return _IToken.Contract.BorrowBalanceCurrent(&_IToken.TransactOpts, _account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address _account) returns(uint256)
func (_IToken *ITokenTransactorSession) BorrowBalanceCurrent(_account common.Address) (*types.Transaction, error) {
	return _IToken.Contract.BorrowBalanceCurrent(&_IToken.TransactOpts, _account)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IToken *ITokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IToken *ITokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DecreaseAllowance(&_IToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IToken *ITokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.DecreaseAllowance(&_IToken.TransactOpts, spender, subtractedValue)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IToken *ITokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IToken *ITokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _IToken.Contract.ExchangeRateCurrent(&_IToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_IToken *ITokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _IToken.Contract.ExchangeRateCurrent(&_IToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IToken *ITokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IToken *ITokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.IncreaseAllowance(&_IToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IToken *ITokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.IncreaseAllowance(&_IToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x6cf1dbed.
//
// Solidity: function initialize(address _underlyingToken, string _name, string _symbol, address _controller, address _interestRateModel) returns()
func (_IToken *ITokenTransactor) Initialize(opts *bind.TransactOpts, _underlyingToken common.Address, _name string, _symbol string, _controller common.Address, _interestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "initialize", _underlyingToken, _name, _symbol, _controller, _interestRateModel)
}

// Initialize is a paid mutator transaction binding the contract method 0x6cf1dbed.
//
// Solidity: function initialize(address _underlyingToken, string _name, string _symbol, address _controller, address _interestRateModel) returns()
func (_IToken *ITokenSession) Initialize(_underlyingToken common.Address, _name string, _symbol string, _controller common.Address, _interestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.Contract.Initialize(&_IToken.TransactOpts, _underlyingToken, _name, _symbol, _controller, _interestRateModel)
}

// Initialize is a paid mutator transaction binding the contract method 0x6cf1dbed.
//
// Solidity: function initialize(address _underlyingToken, string _name, string _symbol, address _controller, address _interestRateModel) returns()
func (_IToken *ITokenTransactorSession) Initialize(_underlyingToken common.Address, _name string, _symbol string, _controller common.Address, _interestRateModel common.Address) (*types.Transaction, error) {
	return _IToken.Contract.Initialize(&_IToken.TransactOpts, _underlyingToken, _name, _symbol, _controller, _interestRateModel)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address _borrower, uint256 _repayAmount, address _assetCollateral) returns()
func (_IToken *ITokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, _borrower common.Address, _repayAmount *big.Int, _assetCollateral common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "liquidateBorrow", _borrower, _repayAmount, _assetCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address _borrower, uint256 _repayAmount, address _assetCollateral) returns()
func (_IToken *ITokenSession) LiquidateBorrow(_borrower common.Address, _repayAmount *big.Int, _assetCollateral common.Address) (*types.Transaction, error) {
	return _IToken.Contract.LiquidateBorrow(&_IToken.TransactOpts, _borrower, _repayAmount, _assetCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address _borrower, uint256 _repayAmount, address _assetCollateral) returns()
func (_IToken *ITokenTransactorSession) LiquidateBorrow(_borrower common.Address, _repayAmount *big.Int, _assetCollateral common.Address) (*types.Transaction, error) {
	return _IToken.Contract.LiquidateBorrow(&_IToken.TransactOpts, _borrower, _repayAmount, _assetCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenTransactor) Mint(opts *bind.TransactOpts, _recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "mint", _recipient, _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenSession) Mint(_recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Mint(&_IToken.TransactOpts, _recipient, _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenTransactorSession) Mint(_recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Mint(&_IToken.TransactOpts, _recipient, _mintAmount)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _recipient) payable returns()
func (_IToken *ITokenTransactor) Mint0(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "mint0", _recipient)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _recipient) payable returns()
func (_IToken *ITokenSession) Mint0(_recipient common.Address) (*types.Transaction, error) {
	return _IToken.Contract.Mint0(&_IToken.TransactOpts, _recipient)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _recipient) payable returns()
func (_IToken *ITokenTransactorSession) Mint0(_recipient common.Address) (*types.Transaction, error) {
	return _IToken.Contract.Mint0(&_IToken.TransactOpts, _recipient)
}

// MintAndEnterMarket is a paid mutator transaction binding the contract method 0xe22b755f.
//
// Solidity: function mintAndEnterMarket(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenTransactor) MintAndEnterMarket(opts *bind.TransactOpts, _recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "mintAndEnterMarket", _recipient, _mintAmount)
}

// MintAndEnterMarket is a paid mutator transaction binding the contract method 0xe22b755f.
//
// Solidity: function mintAndEnterMarket(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenSession) MintAndEnterMarket(_recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.MintAndEnterMarket(&_IToken.TransactOpts, _recipient, _mintAmount)
}

// MintAndEnterMarket is a paid mutator transaction binding the contract method 0xe22b755f.
//
// Solidity: function mintAndEnterMarket(address _recipient, uint256 _mintAmount) returns()
func (_IToken *ITokenTransactorSession) MintAndEnterMarket(_recipient common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.MintAndEnterMarket(&_IToken.TransactOpts, _recipient, _mintAmount)
}

// MintForSelfAndEnterMarket is a paid mutator transaction binding the contract method 0xc0bd21fd.
//
// Solidity: function mintForSelfAndEnterMarket(uint256 _mintAmount) returns()
func (_IToken *ITokenTransactor) MintForSelfAndEnterMarket(opts *bind.TransactOpts, _mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "mintForSelfAndEnterMarket", _mintAmount)
}

// MintForSelfAndEnterMarket is a paid mutator transaction binding the contract method 0xc0bd21fd.
//
// Solidity: function mintForSelfAndEnterMarket(uint256 _mintAmount) returns()
func (_IToken *ITokenSession) MintForSelfAndEnterMarket(_mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.MintForSelfAndEnterMarket(&_IToken.TransactOpts, _mintAmount)
}

// MintForSelfAndEnterMarket is a paid mutator transaction binding the contract method 0xc0bd21fd.
//
// Solidity: function mintForSelfAndEnterMarket(uint256 _mintAmount) returns()
func (_IToken *ITokenTransactorSession) MintForSelfAndEnterMarket(_mintAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.MintForSelfAndEnterMarket(&_IToken.TransactOpts, _mintAmount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IToken *ITokenTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IToken *ITokenSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IToken.Contract.Permit(&_IToken.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IToken *ITokenTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IToken.Contract.Permit(&_IToken.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _from, uint256 _redeemiToken) returns()
func (_IToken *ITokenTransactor) Redeem(opts *bind.TransactOpts, _from common.Address, _redeemiToken *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "redeem", _from, _redeemiToken)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _from, uint256 _redeemiToken) returns()
func (_IToken *ITokenSession) Redeem(_from common.Address, _redeemiToken *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Redeem(&_IToken.TransactOpts, _from, _redeemiToken)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _from, uint256 _redeemiToken) returns()
func (_IToken *ITokenTransactorSession) Redeem(_from common.Address, _redeemiToken *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Redeem(&_IToken.TransactOpts, _from, _redeemiToken)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address _from, uint256 _redeemUnderlying) returns()
func (_IToken *ITokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, _from common.Address, _redeemUnderlying *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "redeemUnderlying", _from, _redeemUnderlying)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address _from, uint256 _redeemUnderlying) returns()
func (_IToken *ITokenSession) RedeemUnderlying(_from common.Address, _redeemUnderlying *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RedeemUnderlying(&_IToken.TransactOpts, _from, _redeemUnderlying)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address _from, uint256 _redeemUnderlying) returns()
func (_IToken *ITokenTransactorSession) RedeemUnderlying(_from common.Address, _redeemUnderlying *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RedeemUnderlying(&_IToken.TransactOpts, _from, _redeemUnderlying)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 _repayAmount) returns()
func (_IToken *ITokenTransactor) RepayBorrow(opts *bind.TransactOpts, _repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "repayBorrow", _repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 _repayAmount) returns()
func (_IToken *ITokenSession) RepayBorrow(_repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RepayBorrow(&_IToken.TransactOpts, _repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 _repayAmount) returns()
func (_IToken *ITokenTransactorSession) RepayBorrow(_repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RepayBorrow(&_IToken.TransactOpts, _repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address _borrower, uint256 _repayAmount) returns()
func (_IToken *ITokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "repayBorrowBehalf", _borrower, _repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address _borrower, uint256 _repayAmount) returns()
func (_IToken *ITokenSession) RepayBorrowBehalf(_borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RepayBorrowBehalf(&_IToken.TransactOpts, _borrower, _repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address _borrower, uint256 _repayAmount) returns()
func (_IToken *ITokenTransactorSession) RepayBorrowBehalf(_borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.RepayBorrowBehalf(&_IToken.TransactOpts, _borrower, _repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address _liquidator, address _borrower, uint256 _seizeTokens) returns()
func (_IToken *ITokenTransactor) Seize(opts *bind.TransactOpts, _liquidator common.Address, _borrower common.Address, _seizeTokens *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "seize", _liquidator, _borrower, _seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address _liquidator, address _borrower, uint256 _seizeTokens) returns()
func (_IToken *ITokenSession) Seize(_liquidator common.Address, _borrower common.Address, _seizeTokens *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Seize(&_IToken.TransactOpts, _liquidator, _borrower, _seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address _liquidator, address _borrower, uint256 _seizeTokens) returns()
func (_IToken *ITokenTransactorSession) Seize(_liquidator common.Address, _borrower common.Address, _seizeTokens *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Seize(&_IToken.TransactOpts, _liquidator, _borrower, _seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IToken *ITokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IToken *ITokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _IToken.Contract.TotalBorrowsCurrent(&_IToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_IToken *ITokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _IToken.Contract.TotalBorrowsCurrent(&_IToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_IToken *ITokenTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, _sender, _recipient, _amount)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns(bool)
func (_IToken *ITokenTransactor) UpdateInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "updateInterest")
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns(bool)
func (_IToken *ITokenSession) UpdateInterest() (*types.Transaction, error) {
	return _IToken.Contract.UpdateInterest(&_IToken.TransactOpts)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns(bool)
func (_IToken *ITokenTransactorSession) UpdateInterest() (*types.Transaction, error) {
	return _IToken.Contract.UpdateInterest(&_IToken.TransactOpts)
}

// ITokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IToken contract.
type ITokenApprovalIterator struct {
	Event *ITokenApproval // Event containing the contract specifics and raw log

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
func (it *ITokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenApproval)
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
		it.Event = new(ITokenApproval)
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
func (it *ITokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenApproval represents a Approval event raised by the IToken contract.
type ITokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITokenApprovalIterator{contract: _IToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenApproval)
				if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IToken *ITokenFilterer) ParseApproval(log types.Log) (*ITokenApproval, error) {
	event := new(ITokenApproval)
	if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the IToken contract.
type ITokenBorrowIterator struct {
	Event *ITokenBorrow // Event containing the contract specifics and raw log

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
func (it *ITokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenBorrow)
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
		it.Event = new(ITokenBorrow)
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
func (it *ITokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenBorrow represents a Borrow event raised by the IToken contract.
type ITokenBorrow struct {
	Borrower             common.Address
	BorrowAmount         *big.Int
	AccountBorrows       *big.Int
	AccountInterestIndex *big.Int
	TotalBorrows         *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) FilterBorrow(opts *bind.FilterOpts) (*ITokenBorrowIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &ITokenBorrowIterator{contract: _IToken.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *ITokenBorrow) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenBorrow)
				if err := _IToken.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) ParseBorrow(log types.Log) (*ITokenBorrow, error) {
	event := new(ITokenBorrow)
	if err := _IToken.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenFlashloanIterator is returned from FilterFlashloan and is used to iterate over the raw logs and unpacked data for Flashloan events raised by the IToken contract.
type ITokenFlashloanIterator struct {
	Event *ITokenFlashloan // Event containing the contract specifics and raw log

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
func (it *ITokenFlashloanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenFlashloan)
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
		it.Event = new(ITokenFlashloan)
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
func (it *ITokenFlashloanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenFlashloanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenFlashloan represents a Flashloan event raised by the IToken contract.
type ITokenFlashloan struct {
	Loaner       common.Address
	LoanAmount   *big.Int
	FlashloanFee *big.Int
	ProtocolFee  *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlashloan is a free log retrieval operation binding the contract event 0xb4fbc5a274ff98f626cbcba96eb3a67e7ec30bed7c28160d43685a7e025308e1.
//
// Solidity: event Flashloan(address loaner, uint256 loanAmount, uint256 flashloanFee, uint256 protocolFee, uint256 timestamp)
func (_IToken *ITokenFilterer) FilterFlashloan(opts *bind.FilterOpts) (*ITokenFlashloanIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Flashloan")
	if err != nil {
		return nil, err
	}
	return &ITokenFlashloanIterator{contract: _IToken.contract, event: "Flashloan", logs: logs, sub: sub}, nil
}

// WatchFlashloan is a free log subscription operation binding the contract event 0xb4fbc5a274ff98f626cbcba96eb3a67e7ec30bed7c28160d43685a7e025308e1.
//
// Solidity: event Flashloan(address loaner, uint256 loanAmount, uint256 flashloanFee, uint256 protocolFee, uint256 timestamp)
func (_IToken *ITokenFilterer) WatchFlashloan(opts *bind.WatchOpts, sink chan<- *ITokenFlashloan) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Flashloan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenFlashloan)
				if err := _IToken.contract.UnpackLog(event, "Flashloan", log); err != nil {
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

// ParseFlashloan is a log parse operation binding the contract event 0xb4fbc5a274ff98f626cbcba96eb3a67e7ec30bed7c28160d43685a7e025308e1.
//
// Solidity: event Flashloan(address loaner, uint256 loanAmount, uint256 flashloanFee, uint256 protocolFee, uint256 timestamp)
func (_IToken *ITokenFilterer) ParseFlashloan(log types.Log) (*ITokenFlashloan, error) {
	event := new(ITokenFlashloan)
	if err := _IToken.contract.UnpackLog(event, "Flashloan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the IToken contract.
type ITokenLiquidateBorrowIterator struct {
	Event *ITokenLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *ITokenLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenLiquidateBorrow)
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
		it.Event = new(ITokenLiquidateBorrow)
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
func (it *ITokenLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenLiquidateBorrow represents a LiquidateBorrow event raised by the IToken contract.
type ITokenLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	ITokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address iTokenCollateral, uint256 seizeTokens)
func (_IToken *ITokenFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*ITokenLiquidateBorrowIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &ITokenLiquidateBorrowIterator{contract: _IToken.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address iTokenCollateral, uint256 seizeTokens)
func (_IToken *ITokenFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *ITokenLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenLiquidateBorrow)
				if err := _IToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address iTokenCollateral, uint256 seizeTokens)
func (_IToken *ITokenFilterer) ParseLiquidateBorrow(log types.Log) (*ITokenLiquidateBorrow, error) {
	event := new(ITokenLiquidateBorrow)
	if err := _IToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IToken contract.
type ITokenMintIterator struct {
	Event *ITokenMint // Event containing the contract specifics and raw log

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
func (it *ITokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenMint)
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
		it.Event = new(ITokenMint)
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
func (it *ITokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenMint represents a Mint event raised by the IToken contract.
type ITokenMint struct {
	Sender     common.Address
	Recipient  common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address sender, address recipient, uint256 mintAmount, uint256 mintTokens)
func (_IToken *ITokenFilterer) FilterMint(opts *bind.FilterOpts) (*ITokenMintIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &ITokenMintIterator{contract: _IToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address sender, address recipient, uint256 mintAmount, uint256 mintTokens)
func (_IToken *ITokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ITokenMint) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenMint)
				if err := _IToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address sender, address recipient, uint256 mintAmount, uint256 mintTokens)
func (_IToken *ITokenFilterer) ParseMint(log types.Log) (*ITokenMint, error) {
	event := new(ITokenMint)
	if err := _IToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the IToken contract.
type ITokenNewControllerIterator struct {
	Event *ITokenNewController // Event containing the contract specifics and raw log

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
func (it *ITokenNewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewController)
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
		it.Event = new(ITokenNewController)
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
func (it *ITokenNewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewController represents a NewController event raised by the IToken contract.
type ITokenNewController struct {
	OldController common.Address
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xf9b6a28700579d5c8fab50f0ac2dcaa52109b85c369c4f511fcc873330ab6cbb.
//
// Solidity: event NewController(address oldController, address newController)
func (_IToken *ITokenFilterer) FilterNewController(opts *bind.FilterOpts) (*ITokenNewControllerIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewController")
	if err != nil {
		return nil, err
	}
	return &ITokenNewControllerIterator{contract: _IToken.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xf9b6a28700579d5c8fab50f0ac2dcaa52109b85c369c4f511fcc873330ab6cbb.
//
// Solidity: event NewController(address oldController, address newController)
func (_IToken *ITokenFilterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *ITokenNewController) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewController)
				if err := _IToken.contract.UnpackLog(event, "NewController", log); err != nil {
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

// ParseNewController is a log parse operation binding the contract event 0xf9b6a28700579d5c8fab50f0ac2dcaa52109b85c369c4f511fcc873330ab6cbb.
//
// Solidity: event NewController(address oldController, address newController)
func (_IToken *ITokenFilterer) ParseNewController(log types.Log) (*ITokenNewController, error) {
	event := new(ITokenNewController)
	if err := _IToken.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewFlashloanFeeIterator is returned from FilterNewFlashloanFee and is used to iterate over the raw logs and unpacked data for NewFlashloanFee events raised by the IToken contract.
type ITokenNewFlashloanFeeIterator struct {
	Event *ITokenNewFlashloanFee // Event containing the contract specifics and raw log

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
func (it *ITokenNewFlashloanFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewFlashloanFee)
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
		it.Event = new(ITokenNewFlashloanFee)
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
func (it *ITokenNewFlashloanFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewFlashloanFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewFlashloanFee represents a NewFlashloanFee event raised by the IToken contract.
type ITokenNewFlashloanFee struct {
	OldFlashloanFeeRatio *big.Int
	NewFlashloanFeeRatio *big.Int
	OldProtocolFeeRatio  *big.Int
	NewProtocolFeeRatio  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewFlashloanFee is a free log retrieval operation binding the contract event 0xe5b7591d834474f1454af47b636cd49301f5d38c7e2f1569be292f51cabb2cd8.
//
// Solidity: event NewFlashloanFee(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio, uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) FilterNewFlashloanFee(opts *bind.FilterOpts) (*ITokenNewFlashloanFeeIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewFlashloanFee")
	if err != nil {
		return nil, err
	}
	return &ITokenNewFlashloanFeeIterator{contract: _IToken.contract, event: "NewFlashloanFee", logs: logs, sub: sub}, nil
}

// WatchNewFlashloanFee is a free log subscription operation binding the contract event 0xe5b7591d834474f1454af47b636cd49301f5d38c7e2f1569be292f51cabb2cd8.
//
// Solidity: event NewFlashloanFee(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio, uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) WatchNewFlashloanFee(opts *bind.WatchOpts, sink chan<- *ITokenNewFlashloanFee) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewFlashloanFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewFlashloanFee)
				if err := _IToken.contract.UnpackLog(event, "NewFlashloanFee", log); err != nil {
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

// ParseNewFlashloanFee is a log parse operation binding the contract event 0xe5b7591d834474f1454af47b636cd49301f5d38c7e2f1569be292f51cabb2cd8.
//
// Solidity: event NewFlashloanFee(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio, uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) ParseNewFlashloanFee(log types.Log) (*ITokenNewFlashloanFee, error) {
	event := new(ITokenNewFlashloanFee)
	if err := _IToken.contract.UnpackLog(event, "NewFlashloanFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewFlashloanFeeRatioIterator is returned from FilterNewFlashloanFeeRatio and is used to iterate over the raw logs and unpacked data for NewFlashloanFeeRatio events raised by the IToken contract.
type ITokenNewFlashloanFeeRatioIterator struct {
	Event *ITokenNewFlashloanFeeRatio // Event containing the contract specifics and raw log

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
func (it *ITokenNewFlashloanFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewFlashloanFeeRatio)
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
		it.Event = new(ITokenNewFlashloanFeeRatio)
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
func (it *ITokenNewFlashloanFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewFlashloanFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewFlashloanFeeRatio represents a NewFlashloanFeeRatio event raised by the IToken contract.
type ITokenNewFlashloanFeeRatio struct {
	OldFlashloanFeeRatio *big.Int
	NewFlashloanFeeRatio *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewFlashloanFeeRatio is a free log retrieval operation binding the contract event 0xcc8d29bd7a6ccb34210e5349873398367afa955e6b745621430e8152677d7c75.
//
// Solidity: event NewFlashloanFeeRatio(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio)
func (_IToken *ITokenFilterer) FilterNewFlashloanFeeRatio(opts *bind.FilterOpts) (*ITokenNewFlashloanFeeRatioIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewFlashloanFeeRatio")
	if err != nil {
		return nil, err
	}
	return &ITokenNewFlashloanFeeRatioIterator{contract: _IToken.contract, event: "NewFlashloanFeeRatio", logs: logs, sub: sub}, nil
}

// WatchNewFlashloanFeeRatio is a free log subscription operation binding the contract event 0xcc8d29bd7a6ccb34210e5349873398367afa955e6b745621430e8152677d7c75.
//
// Solidity: event NewFlashloanFeeRatio(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio)
func (_IToken *ITokenFilterer) WatchNewFlashloanFeeRatio(opts *bind.WatchOpts, sink chan<- *ITokenNewFlashloanFeeRatio) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewFlashloanFeeRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewFlashloanFeeRatio)
				if err := _IToken.contract.UnpackLog(event, "NewFlashloanFeeRatio", log); err != nil {
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

// ParseNewFlashloanFeeRatio is a log parse operation binding the contract event 0xcc8d29bd7a6ccb34210e5349873398367afa955e6b745621430e8152677d7c75.
//
// Solidity: event NewFlashloanFeeRatio(uint256 oldFlashloanFeeRatio, uint256 newFlashloanFeeRatio)
func (_IToken *ITokenFilterer) ParseNewFlashloanFeeRatio(log types.Log) (*ITokenNewFlashloanFeeRatio, error) {
	event := new(ITokenNewFlashloanFeeRatio)
	if err := _IToken.contract.UnpackLog(event, "NewFlashloanFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewInterestRateModelIterator is returned from FilterNewInterestRateModel and is used to iterate over the raw logs and unpacked data for NewInterestRateModel events raised by the IToken contract.
type ITokenNewInterestRateModelIterator struct {
	Event *ITokenNewInterestRateModel // Event containing the contract specifics and raw log

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
func (it *ITokenNewInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewInterestRateModel)
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
		it.Event = new(ITokenNewInterestRateModel)
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
func (it *ITokenNewInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewInterestRateModel represents a NewInterestRateModel event raised by the IToken contract.
type ITokenNewInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewInterestRateModel is a free log retrieval operation binding the contract event 0xeb5cc99f497dc2d7106563bb080e06c9b09e3d81a38623ac4d0839575658d1fa.
//
// Solidity: event NewInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_IToken *ITokenFilterer) FilterNewInterestRateModel(opts *bind.FilterOpts) (*ITokenNewInterestRateModelIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &ITokenNewInterestRateModelIterator{contract: _IToken.contract, event: "NewInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewInterestRateModel is a free log subscription operation binding the contract event 0xeb5cc99f497dc2d7106563bb080e06c9b09e3d81a38623ac4d0839575658d1fa.
//
// Solidity: event NewInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_IToken *ITokenFilterer) WatchNewInterestRateModel(opts *bind.WatchOpts, sink chan<- *ITokenNewInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewInterestRateModel)
				if err := _IToken.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
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

// ParseNewInterestRateModel is a log parse operation binding the contract event 0xeb5cc99f497dc2d7106563bb080e06c9b09e3d81a38623ac4d0839575658d1fa.
//
// Solidity: event NewInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_IToken *ITokenFilterer) ParseNewInterestRateModel(log types.Log) (*ITokenNewInterestRateModel, error) {
	event := new(ITokenNewInterestRateModel)
	if err := _IToken.contract.UnpackLog(event, "NewInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the IToken contract.
type ITokenNewOwnerIterator struct {
	Event *ITokenNewOwner // Event containing the contract specifics and raw log

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
func (it *ITokenNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewOwner)
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
		it.Event = new(ITokenNewOwner)
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
func (it *ITokenNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewOwner represents a NewOwner event raised by the IToken contract.
type ITokenNewOwner struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_IToken *ITokenFilterer) FilterNewOwner(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ITokenNewOwnerIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ITokenNewOwnerIterator{contract: _IToken.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_IToken *ITokenFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ITokenNewOwner, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewOwner)
				if err := _IToken.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_IToken *ITokenFilterer) ParseNewOwner(log types.Log) (*ITokenNewOwner, error) {
	event := new(ITokenNewOwner)
	if err := _IToken.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the IToken contract.
type ITokenNewPendingOwnerIterator struct {
	Event *ITokenNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *ITokenNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewPendingOwner)
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
		it.Event = new(ITokenNewPendingOwner)
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
func (it *ITokenNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewPendingOwner represents a NewPendingOwner event raised by the IToken contract.
type ITokenNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_IToken *ITokenFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*ITokenNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ITokenNewPendingOwnerIterator{contract: _IToken.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_IToken *ITokenFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *ITokenNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewPendingOwner)
				if err := _IToken.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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
func (_IToken *ITokenFilterer) ParseNewPendingOwner(log types.Log) (*ITokenNewPendingOwner, error) {
	event := new(ITokenNewPendingOwner)
	if err := _IToken.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewProtocolFeeRatioIterator is returned from FilterNewProtocolFeeRatio and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRatio events raised by the IToken contract.
type ITokenNewProtocolFeeRatioIterator struct {
	Event *ITokenNewProtocolFeeRatio // Event containing the contract specifics and raw log

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
func (it *ITokenNewProtocolFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewProtocolFeeRatio)
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
		it.Event = new(ITokenNewProtocolFeeRatio)
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
func (it *ITokenNewProtocolFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewProtocolFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewProtocolFeeRatio represents a NewProtocolFeeRatio event raised by the IToken contract.
type ITokenNewProtocolFeeRatio struct {
	OldProtocolFeeRatio *big.Int
	NewProtocolFeeRatio *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRatio is a free log retrieval operation binding the contract event 0x1f6917c8223f142d2549d7531b9897b963f67c4cd1d266e9771080a608ebe188.
//
// Solidity: event NewProtocolFeeRatio(uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) FilterNewProtocolFeeRatio(opts *bind.FilterOpts) (*ITokenNewProtocolFeeRatioIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return &ITokenNewProtocolFeeRatioIterator{contract: _IToken.contract, event: "NewProtocolFeeRatio", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRatio is a free log subscription operation binding the contract event 0x1f6917c8223f142d2549d7531b9897b963f67c4cd1d266e9771080a608ebe188.
//
// Solidity: event NewProtocolFeeRatio(uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) WatchNewProtocolFeeRatio(opts *bind.WatchOpts, sink chan<- *ITokenNewProtocolFeeRatio) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewProtocolFeeRatio)
				if err := _IToken.contract.UnpackLog(event, "NewProtocolFeeRatio", log); err != nil {
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

// ParseNewProtocolFeeRatio is a log parse operation binding the contract event 0x1f6917c8223f142d2549d7531b9897b963f67c4cd1d266e9771080a608ebe188.
//
// Solidity: event NewProtocolFeeRatio(uint256 oldProtocolFeeRatio, uint256 newProtocolFeeRatio)
func (_IToken *ITokenFilterer) ParseNewProtocolFeeRatio(log types.Log) (*ITokenNewProtocolFeeRatio, error) {
	event := new(ITokenNewProtocolFeeRatio)
	if err := _IToken.contract.UnpackLog(event, "NewProtocolFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenNewReserveRatioIterator is returned from FilterNewReserveRatio and is used to iterate over the raw logs and unpacked data for NewReserveRatio events raised by the IToken contract.
type ITokenNewReserveRatioIterator struct {
	Event *ITokenNewReserveRatio // Event containing the contract specifics and raw log

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
func (it *ITokenNewReserveRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenNewReserveRatio)
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
		it.Event = new(ITokenNewReserveRatio)
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
func (it *ITokenNewReserveRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenNewReserveRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenNewReserveRatio represents a NewReserveRatio event raised by the IToken contract.
type ITokenNewReserveRatio struct {
	OldReserveRatio *big.Int
	NewReserveRatio *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewReserveRatio is a free log retrieval operation binding the contract event 0x619970c7b73e9bc2b93e0fc379d50dbf7eced564f397fba395b2e7efc0b4894d.
//
// Solidity: event NewReserveRatio(uint256 oldReserveRatio, uint256 newReserveRatio)
func (_IToken *ITokenFilterer) FilterNewReserveRatio(opts *bind.FilterOpts) (*ITokenNewReserveRatioIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "NewReserveRatio")
	if err != nil {
		return nil, err
	}
	return &ITokenNewReserveRatioIterator{contract: _IToken.contract, event: "NewReserveRatio", logs: logs, sub: sub}, nil
}

// WatchNewReserveRatio is a free log subscription operation binding the contract event 0x619970c7b73e9bc2b93e0fc379d50dbf7eced564f397fba395b2e7efc0b4894d.
//
// Solidity: event NewReserveRatio(uint256 oldReserveRatio, uint256 newReserveRatio)
func (_IToken *ITokenFilterer) WatchNewReserveRatio(opts *bind.WatchOpts, sink chan<- *ITokenNewReserveRatio) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "NewReserveRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenNewReserveRatio)
				if err := _IToken.contract.UnpackLog(event, "NewReserveRatio", log); err != nil {
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

// ParseNewReserveRatio is a log parse operation binding the contract event 0x619970c7b73e9bc2b93e0fc379d50dbf7eced564f397fba395b2e7efc0b4894d.
//
// Solidity: event NewReserveRatio(uint256 oldReserveRatio, uint256 newReserveRatio)
func (_IToken *ITokenFilterer) ParseNewReserveRatio(log types.Log) (*ITokenNewReserveRatio, error) {
	event := new(ITokenNewReserveRatio)
	if err := _IToken.contract.UnpackLog(event, "NewReserveRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the IToken contract.
type ITokenRedeemIterator struct {
	Event *ITokenRedeem // Event containing the contract specifics and raw log

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
func (it *ITokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenRedeem)
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
		it.Event = new(ITokenRedeem)
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
func (it *ITokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenRedeem represents a Redeem event raised by the IToken contract.
type ITokenRedeem struct {
	From                   common.Address
	Recipient              common.Address
	RedeemiTokenAmount     *big.Int
	RedeemUnderlyingAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x3f693fff038bb8a046aa76d9516190ac7444f7d69cf952c4cbdc086fdef2d6fc.
//
// Solidity: event Redeem(address from, address recipient, uint256 redeemiTokenAmount, uint256 redeemUnderlyingAmount)
func (_IToken *ITokenFilterer) FilterRedeem(opts *bind.FilterOpts) (*ITokenRedeemIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &ITokenRedeemIterator{contract: _IToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x3f693fff038bb8a046aa76d9516190ac7444f7d69cf952c4cbdc086fdef2d6fc.
//
// Solidity: event Redeem(address from, address recipient, uint256 redeemiTokenAmount, uint256 redeemUnderlyingAmount)
func (_IToken *ITokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ITokenRedeem) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenRedeem)
				if err := _IToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x3f693fff038bb8a046aa76d9516190ac7444f7d69cf952c4cbdc086fdef2d6fc.
//
// Solidity: event Redeem(address from, address recipient, uint256 redeemiTokenAmount, uint256 redeemUnderlyingAmount)
func (_IToken *ITokenFilterer) ParseRedeem(log types.Log) (*ITokenRedeem, error) {
	event := new(ITokenRedeem)
	if err := _IToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the IToken contract.
type ITokenRepayBorrowIterator struct {
	Event *ITokenRepayBorrow // Event containing the contract specifics and raw log

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
func (it *ITokenRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenRepayBorrow)
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
		it.Event = new(ITokenRepayBorrow)
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
func (it *ITokenRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenRepayBorrow represents a RepayBorrow event raised by the IToken contract.
type ITokenRepayBorrow struct {
	Payer                common.Address
	Borrower             common.Address
	RepayAmount          *big.Int
	AccountBorrows       *big.Int
	AccountInterestIndex *big.Int
	TotalBorrows         *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x6fadbf7329d21f278e724fa0d4511001a158f2a97ee35c5bc4cf8b64417399ef.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*ITokenRepayBorrowIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &ITokenRepayBorrowIterator{contract: _IToken.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x6fadbf7329d21f278e724fa0d4511001a158f2a97ee35c5bc4cf8b64417399ef.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *ITokenRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenRepayBorrow)
				if err := _IToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x6fadbf7329d21f278e724fa0d4511001a158f2a97ee35c5bc4cf8b64417399ef.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 accountInterestIndex, uint256 totalBorrows)
func (_IToken *ITokenFilterer) ParseRepayBorrow(log types.Log) (*ITokenRepayBorrow, error) {
	event := new(ITokenRepayBorrow)
	if err := _IToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenReservesWithdrawnIterator is returned from FilterReservesWithdrawn and is used to iterate over the raw logs and unpacked data for ReservesWithdrawn events raised by the IToken contract.
type ITokenReservesWithdrawnIterator struct {
	Event *ITokenReservesWithdrawn // Event containing the contract specifics and raw log

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
func (it *ITokenReservesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenReservesWithdrawn)
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
		it.Event = new(ITokenReservesWithdrawn)
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
func (it *ITokenReservesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenReservesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenReservesWithdrawn represents a ReservesWithdrawn event raised by the IToken contract.
type ITokenReservesWithdrawn struct {
	Admin            common.Address
	Amount           *big.Int
	NewTotalReserves *big.Int
	OldTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesWithdrawn is a free log retrieval operation binding the contract event 0x2e8843ddc3123732d720f1cb17a6e81d71d5bb90a346f13498b87c5639d47440.
//
// Solidity: event ReservesWithdrawn(address admin, uint256 amount, uint256 newTotalReserves, uint256 oldTotalReserves)
func (_IToken *ITokenFilterer) FilterReservesWithdrawn(opts *bind.FilterOpts) (*ITokenReservesWithdrawnIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "ReservesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &ITokenReservesWithdrawnIterator{contract: _IToken.contract, event: "ReservesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchReservesWithdrawn is a free log subscription operation binding the contract event 0x2e8843ddc3123732d720f1cb17a6e81d71d5bb90a346f13498b87c5639d47440.
//
// Solidity: event ReservesWithdrawn(address admin, uint256 amount, uint256 newTotalReserves, uint256 oldTotalReserves)
func (_IToken *ITokenFilterer) WatchReservesWithdrawn(opts *bind.WatchOpts, sink chan<- *ITokenReservesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "ReservesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenReservesWithdrawn)
				if err := _IToken.contract.UnpackLog(event, "ReservesWithdrawn", log); err != nil {
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

// ParseReservesWithdrawn is a log parse operation binding the contract event 0x2e8843ddc3123732d720f1cb17a6e81d71d5bb90a346f13498b87c5639d47440.
//
// Solidity: event ReservesWithdrawn(address admin, uint256 amount, uint256 newTotalReserves, uint256 oldTotalReserves)
func (_IToken *ITokenFilterer) ParseReservesWithdrawn(log types.Log) (*ITokenReservesWithdrawn, error) {
	event := new(ITokenReservesWithdrawn)
	if err := _IToken.contract.UnpackLog(event, "ReservesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IToken contract.
type ITokenTransferIterator struct {
	Event *ITokenTransfer // Event containing the contract specifics and raw log

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
func (it *ITokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenTransfer)
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
		it.Event = new(ITokenTransfer)
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
func (it *ITokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenTransfer represents a Transfer event raised by the IToken contract.
type ITokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferIterator{contract: _IToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenTransfer)
				if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IToken *ITokenFilterer) ParseTransfer(log types.Log) (*ITokenTransfer, error) {
	event := new(ITokenTransfer)
	if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenUpdateInterestIterator is returned from FilterUpdateInterest and is used to iterate over the raw logs and unpacked data for UpdateInterest events raised by the IToken contract.
type ITokenUpdateInterestIterator struct {
	Event *ITokenUpdateInterest // Event containing the contract specifics and raw log

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
func (it *ITokenUpdateInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenUpdateInterest)
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
		it.Event = new(ITokenUpdateInterest)
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
func (it *ITokenUpdateInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenUpdateInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenUpdateInterest represents a UpdateInterest event raised by the IToken contract.
type ITokenUpdateInterest struct {
	CurrentBlockNumber  *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	Cash                *big.Int
	TotalBorrows        *big.Int
	TotalReserves       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateInterest is a free log retrieval operation binding the contract event 0x59693255bedc2974b761b077cd2fdb47b3bde759f64b247f599c6941525655e1.
//
// Solidity: event UpdateInterest(uint256 currentBlockNumber, uint256 interestAccumulated, uint256 borrowIndex, uint256 cash, uint256 totalBorrows, uint256 totalReserves)
func (_IToken *ITokenFilterer) FilterUpdateInterest(opts *bind.FilterOpts) (*ITokenUpdateInterestIterator, error) {

	logs, sub, err := _IToken.contract.FilterLogs(opts, "UpdateInterest")
	if err != nil {
		return nil, err
	}
	return &ITokenUpdateInterestIterator{contract: _IToken.contract, event: "UpdateInterest", logs: logs, sub: sub}, nil
}

// WatchUpdateInterest is a free log subscription operation binding the contract event 0x59693255bedc2974b761b077cd2fdb47b3bde759f64b247f599c6941525655e1.
//
// Solidity: event UpdateInterest(uint256 currentBlockNumber, uint256 interestAccumulated, uint256 borrowIndex, uint256 cash, uint256 totalBorrows, uint256 totalReserves)
func (_IToken *ITokenFilterer) WatchUpdateInterest(opts *bind.WatchOpts, sink chan<- *ITokenUpdateInterest) (event.Subscription, error) {

	logs, sub, err := _IToken.contract.WatchLogs(opts, "UpdateInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenUpdateInterest)
				if err := _IToken.contract.UnpackLog(event, "UpdateInterest", log); err != nil {
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

// ParseUpdateInterest is a log parse operation binding the contract event 0x59693255bedc2974b761b077cd2fdb47b3bde759f64b247f599c6941525655e1.
//
// Solidity: event UpdateInterest(uint256 currentBlockNumber, uint256 interestAccumulated, uint256 borrowIndex, uint256 cash, uint256 totalBorrows, uint256 totalReserves)
func (_IToken *ITokenFilterer) ParseUpdateInterest(log types.Log) (*ITokenUpdateInterest, error) {
	event := new(ITokenUpdateInterest)
	if err := _IToken.contract.UnpackLog(event, "UpdateInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
