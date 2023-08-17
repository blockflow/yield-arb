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

// LTokenMetaData contains all meta data concerning the LToken contract.
var LTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AcceptAdminPendingAdminCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAddAmount\",\"type\":\"uint256\"}],\"name\":\"AddReservesFactorFreshCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowCashNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"BorrowComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"LiquidateAccrueBorrowInterestFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"LiquidateAccrueCollateralInterestFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCloseAmountIsUintMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCloseAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCollateralFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"LiquidateComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateLiquidatorIsBorrower\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"LiquidateRepayBorrowFreshFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"LiquidateSeizeComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateSeizeLiquidatorIsBorrower\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"MintComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"RedeemComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedeemFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedeemTransferOutNotPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReduceReservesAdminCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReduceReservesCashNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReduceReservesCashValidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReduceReservesFreshCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReduceReservesGuardianCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"RepayBorrowComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RepayBorrowFreshnessCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetComptrollerOwnerCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetInterestRateModelFreshCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetInterestRateModelOwnerCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetPendingAdminOwnerCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetReserveFactorAdminCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetReserveFactorBoundsCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetReserveFactorFreshCheck\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorCode\",\"type\":\"uint256\"}],\"name\":\"TransferComptrollerRejection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferTooMuch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldReserveGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReserveGuardian\",\"type\":\"address\"}],\"name\":\"NewReserveGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"_becomeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"compLikeDelegatee\",\"type\":\"address\"}],\"name\":\"_delegateCompLikeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_resignImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReserveGuardian\",\"type\":\"address\"}],\"name\":\"_setReserveGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrowee\",\"type\":\"address\"}],\"name\":\"borrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCTokenInterface\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolSeizeShareMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"redeemee\",\"type\":\"address\"}],\"name\":\"redeemBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveGuardian\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEIP20NonStandardInterface\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use LTokenMetaData.ABI instead.
var LTokenABI = LTokenMetaData.ABI

// LToken is an auto generated Go binding around an Ethereum contract.
type LToken struct {
	LTokenCaller     // Read-only binding to the contract
	LTokenTransactor // Write-only binding to the contract
	LTokenFilterer   // Log filterer for contract events
}

// LTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LTokenSession struct {
	Contract     *LToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LTokenCallerSession struct {
	Contract *LTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LTokenTransactorSession struct {
	Contract     *LTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LTokenRaw struct {
	Contract *LToken // Generic contract binding to access the raw methods on
}

// LTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LTokenCallerRaw struct {
	Contract *LTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LTokenTransactorRaw struct {
	Contract *LTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLToken creates a new instance of LToken, bound to a specific deployed contract.
func NewLToken(address common.Address, backend bind.ContractBackend) (*LToken, error) {
	contract, err := bindLToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LToken{LTokenCaller: LTokenCaller{contract: contract}, LTokenTransactor: LTokenTransactor{contract: contract}, LTokenFilterer: LTokenFilterer{contract: contract}}, nil
}

// NewLTokenCaller creates a new read-only instance of LToken, bound to a specific deployed contract.
func NewLTokenCaller(address common.Address, caller bind.ContractCaller) (*LTokenCaller, error) {
	contract, err := bindLToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LTokenCaller{contract: contract}, nil
}

// NewLTokenTransactor creates a new write-only instance of LToken, bound to a specific deployed contract.
func NewLTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LTokenTransactor, error) {
	contract, err := bindLToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LTokenTransactor{contract: contract}, nil
}

// NewLTokenFilterer creates a new log filterer instance of LToken, bound to a specific deployed contract.
func NewLTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LTokenFilterer, error) {
	contract, err := bindLToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LTokenFilterer{contract: contract}, nil
}

// bindLToken binds a generic wrapper to an already deployed contract.
func bindLToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LToken *LTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LToken.Contract.LTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LToken *LTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.Contract.LTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LToken *LTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LToken.Contract.LTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LToken *LTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LToken *LTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LToken *LTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LToken.Contract.contract.Transact(opts, method, params...)
}

// NOERROR is a free data retrieval call binding the contract method 0x69ab3250.
//
// Solidity: function NO_ERROR() view returns(uint256)
func (_LToken *LTokenCaller) NOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NOERROR is a free data retrieval call binding the contract method 0x69ab3250.
//
// Solidity: function NO_ERROR() view returns(uint256)
func (_LToken *LTokenSession) NOERROR() (*big.Int, error) {
	return _LToken.Contract.NOERROR(&_LToken.CallOpts)
}

// NOERROR is a free data retrieval call binding the contract method 0x69ab3250.
//
// Solidity: function NO_ERROR() view returns(uint256)
func (_LToken *LTokenCallerSession) NOERROR() (*big.Int, error) {
	return _LToken.Contract.NOERROR(&_LToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_LToken *LTokenCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_LToken *LTokenSession) AccrualBlockNumber() (*big.Int, error) {
	return _LToken.Contract.AccrualBlockNumber(&_LToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_LToken *LTokenCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _LToken.Contract.AccrualBlockNumber(&_LToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LToken *LTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LToken *LTokenSession) Admin() (common.Address, error) {
	return _LToken.Contract.Admin(&_LToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_LToken *LTokenCallerSession) Admin() (common.Address, error) {
	return _LToken.Contract.Admin(&_LToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LToken *LTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LToken *LTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LToken.Contract.Allowance(&_LToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LToken *LTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LToken.Contract.Allowance(&_LToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LToken *LTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LToken *LTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LToken.Contract.BalanceOf(&_LToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LToken *LTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LToken.Contract.BalanceOf(&_LToken.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_LToken *LTokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_LToken *LTokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _LToken.Contract.BorrowBalanceStored(&_LToken.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_LToken *LTokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _LToken.Contract.BorrowBalanceStored(&_LToken.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LToken *LTokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LToken *LTokenSession) BorrowIndex() (*big.Int, error) {
	return _LToken.Contract.BorrowIndex(&_LToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_LToken *LTokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _LToken.Contract.BorrowIndex(&_LToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_LToken *LTokenCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_LToken *LTokenSession) BorrowRatePerBlock() (*big.Int, error) {
	return _LToken.Contract.BorrowRatePerBlock(&_LToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_LToken *LTokenCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _LToken.Contract.BorrowRatePerBlock(&_LToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_LToken *LTokenCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_LToken *LTokenSession) Comptroller() (common.Address, error) {
	return _LToken.Contract.Comptroller(&_LToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_LToken *LTokenCallerSession) Comptroller() (common.Address, error) {
	return _LToken.Contract.Comptroller(&_LToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LToken *LTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LToken *LTokenSession) Decimals() (uint8, error) {
	return _LToken.Contract.Decimals(&_LToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LToken *LTokenCallerSession) Decimals() (uint8, error) {
	return _LToken.Contract.Decimals(&_LToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_LToken *LTokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_LToken *LTokenSession) ExchangeRateStored() (*big.Int, error) {
	return _LToken.Contract.ExchangeRateStored(&_LToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_LToken *LTokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _LToken.Contract.ExchangeRateStored(&_LToken.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_LToken *LTokenCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_LToken *LTokenSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _LToken.Contract.GetAccountSnapshot(&_LToken.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_LToken *LTokenCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _LToken.Contract.GetAccountSnapshot(&_LToken.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_LToken *LTokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_LToken *LTokenSession) GetCash() (*big.Int, error) {
	return _LToken.Contract.GetCash(&_LToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_LToken *LTokenCallerSession) GetCash() (*big.Int, error) {
	return _LToken.Contract.GetCash(&_LToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LToken *LTokenCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LToken *LTokenSession) Implementation() (common.Address, error) {
	return _LToken.Contract.Implementation(&_LToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_LToken *LTokenCallerSession) Implementation() (common.Address, error) {
	return _LToken.Contract.Implementation(&_LToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_LToken *LTokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_LToken *LTokenSession) InterestRateModel() (common.Address, error) {
	return _LToken.Contract.InterestRateModel(&_LToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_LToken *LTokenCallerSession) InterestRateModel() (common.Address, error) {
	return _LToken.Contract.InterestRateModel(&_LToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_LToken *LTokenCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_LToken *LTokenSession) IsCToken() (bool, error) {
	return _LToken.Contract.IsCToken(&_LToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_LToken *LTokenCallerSession) IsCToken() (bool, error) {
	return _LToken.Contract.IsCToken(&_LToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LToken *LTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LToken *LTokenSession) Name() (string, error) {
	return _LToken.Contract.Name(&_LToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LToken *LTokenCallerSession) Name() (string, error) {
	return _LToken.Contract.Name(&_LToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_LToken *LTokenCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_LToken *LTokenSession) PendingAdmin() (common.Address, error) {
	return _LToken.Contract.PendingAdmin(&_LToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_LToken *LTokenCallerSession) PendingAdmin() (common.Address, error) {
	return _LToken.Contract.PendingAdmin(&_LToken.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_LToken *LTokenCaller) ProtocolSeizeShareMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "protocolSeizeShareMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_LToken *LTokenSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _LToken.Contract.ProtocolSeizeShareMantissa(&_LToken.CallOpts)
}

// ProtocolSeizeShareMantissa is a free data retrieval call binding the contract method 0x6752e702.
//
// Solidity: function protocolSeizeShareMantissa() view returns(uint256)
func (_LToken *LTokenCallerSession) ProtocolSeizeShareMantissa() (*big.Int, error) {
	return _LToken.Contract.ProtocolSeizeShareMantissa(&_LToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_LToken *LTokenCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_LToken *LTokenSession) ReserveFactorMantissa() (*big.Int, error) {
	return _LToken.Contract.ReserveFactorMantissa(&_LToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_LToken *LTokenCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _LToken.Contract.ReserveFactorMantissa(&_LToken.CallOpts)
}

// ReserveGuardian is a free data retrieval call binding the contract method 0x0d983cc6.
//
// Solidity: function reserveGuardian() view returns(address)
func (_LToken *LTokenCaller) ReserveGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "reserveGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveGuardian is a free data retrieval call binding the contract method 0x0d983cc6.
//
// Solidity: function reserveGuardian() view returns(address)
func (_LToken *LTokenSession) ReserveGuardian() (common.Address, error) {
	return _LToken.Contract.ReserveGuardian(&_LToken.CallOpts)
}

// ReserveGuardian is a free data retrieval call binding the contract method 0x0d983cc6.
//
// Solidity: function reserveGuardian() view returns(address)
func (_LToken *LTokenCallerSession) ReserveGuardian() (common.Address, error) {
	return _LToken.Contract.ReserveGuardian(&_LToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_LToken *LTokenCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_LToken *LTokenSession) SupplyRatePerBlock() (*big.Int, error) {
	return _LToken.Contract.SupplyRatePerBlock(&_LToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_LToken *LTokenCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _LToken.Contract.SupplyRatePerBlock(&_LToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LToken *LTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LToken *LTokenSession) Symbol() (string, error) {
	return _LToken.Contract.Symbol(&_LToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LToken *LTokenCallerSession) Symbol() (string, error) {
	return _LToken.Contract.Symbol(&_LToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_LToken *LTokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_LToken *LTokenSession) TotalBorrows() (*big.Int, error) {
	return _LToken.Contract.TotalBorrows(&_LToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_LToken *LTokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _LToken.Contract.TotalBorrows(&_LToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_LToken *LTokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_LToken *LTokenSession) TotalReserves() (*big.Int, error) {
	return _LToken.Contract.TotalReserves(&_LToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_LToken *LTokenCallerSession) TotalReserves() (*big.Int, error) {
	return _LToken.Contract.TotalReserves(&_LToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LToken *LTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LToken *LTokenSession) TotalSupply() (*big.Int, error) {
	return _LToken.Contract.TotalSupply(&_LToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LToken *LTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LToken.Contract.TotalSupply(&_LToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_LToken *LTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_LToken *LTokenSession) Underlying() (common.Address, error) {
	return _LToken.Contract.Underlying(&_LToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_LToken *LTokenCallerSession) Underlying() (common.Address, error) {
	return _LToken.Contract.Underlying(&_LToken.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_LToken *LTokenTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_LToken *LTokenSession) AcceptAdmin() (*types.Transaction, error) {
	return _LToken.Contract.AcceptAdmin(&_LToken.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_LToken *LTokenTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _LToken.Contract.AcceptAdmin(&_LToken.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_LToken *LTokenTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_LToken *LTokenSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.AddReserves(&_LToken.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.AddReserves(&_LToken.TransactOpts, addAmount)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_LToken *LTokenTransactor) BecomeImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_becomeImplementation", data)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_LToken *LTokenSession) BecomeImplementation(data []byte) (*types.Transaction, error) {
	return _LToken.Contract.BecomeImplementation(&_LToken.TransactOpts, data)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_LToken *LTokenTransactorSession) BecomeImplementation(data []byte) (*types.Transaction, error) {
	return _LToken.Contract.BecomeImplementation(&_LToken.TransactOpts, data)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_LToken *LTokenTransactor) DelegateCompLikeTo(opts *bind.TransactOpts, compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_delegateCompLikeTo", compLikeDelegatee)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_LToken *LTokenSession) DelegateCompLikeTo(compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.DelegateCompLikeTo(&_LToken.TransactOpts, compLikeDelegatee)
}

// DelegateCompLikeTo is a paid mutator transaction binding the contract method 0x7f1e06be.
//
// Solidity: function _delegateCompLikeTo(address compLikeDelegatee) returns()
func (_LToken *LTokenTransactorSession) DelegateCompLikeTo(compLikeDelegatee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.DelegateCompLikeTo(&_LToken.TransactOpts, compLikeDelegatee)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_LToken *LTokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_LToken *LTokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.ReduceReserves(&_LToken.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.ReduceReserves(&_LToken.TransactOpts, reduceAmount)
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_LToken *LTokenTransactor) ResignImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_resignImplementation")
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_LToken *LTokenSession) ResignImplementation() (*types.Transaction, error) {
	return _LToken.Contract.ResignImplementation(&_LToken.TransactOpts)
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_LToken *LTokenTransactorSession) ResignImplementation() (*types.Transaction, error) {
	return _LToken.Contract.ResignImplementation(&_LToken.TransactOpts)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_LToken *LTokenTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_LToken *LTokenSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetComptroller(&_LToken.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_LToken *LTokenTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetComptroller(&_LToken.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_LToken *LTokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_LToken *LTokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetInterestRateModel(&_LToken.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_LToken *LTokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetInterestRateModel(&_LToken.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_LToken *LTokenTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_LToken *LTokenSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetPendingAdmin(&_LToken.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_LToken *LTokenTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetPendingAdmin(&_LToken.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_LToken *LTokenTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_LToken *LTokenSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.SetReserveFactor(&_LToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_LToken *LTokenTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.SetReserveFactor(&_LToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveGuardian is a paid mutator transaction binding the contract method 0xc39b127f.
//
// Solidity: function _setReserveGuardian(address newReserveGuardian) returns(uint256)
func (_LToken *LTokenTransactor) SetReserveGuardian(opts *bind.TransactOpts, newReserveGuardian common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "_setReserveGuardian", newReserveGuardian)
}

// SetReserveGuardian is a paid mutator transaction binding the contract method 0xc39b127f.
//
// Solidity: function _setReserveGuardian(address newReserveGuardian) returns(uint256)
func (_LToken *LTokenSession) SetReserveGuardian(newReserveGuardian common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetReserveGuardian(&_LToken.TransactOpts, newReserveGuardian)
}

// SetReserveGuardian is a paid mutator transaction binding the contract method 0xc39b127f.
//
// Solidity: function _setReserveGuardian(address newReserveGuardian) returns(uint256)
func (_LToken *LTokenTransactorSession) SetReserveGuardian(newReserveGuardian common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SetReserveGuardian(&_LToken.TransactOpts, newReserveGuardian)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_LToken *LTokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_LToken *LTokenSession) AccrueInterest() (*types.Transaction, error) {
	return _LToken.Contract.AccrueInterest(&_LToken.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_LToken *LTokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _LToken.Contract.AccrueInterest(&_LToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LToken *LTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LToken *LTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Approve(&_LToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LToken *LTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Approve(&_LToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_LToken *LTokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_LToken *LTokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BalanceOfUnderlying(&_LToken.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_LToken *LTokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BalanceOfUnderlying(&_LToken.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_LToken *LTokenTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_LToken *LTokenSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Borrow(&_LToken.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Borrow(&_LToken.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_LToken *LTokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_LToken *LTokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BorrowBalanceCurrent(&_LToken.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_LToken *LTokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BorrowBalanceCurrent(&_LToken.TransactOpts, account)
}

// BorrowBehalf is a paid mutator transaction binding the contract method 0x60a7cf20.
//
// Solidity: function borrowBehalf(uint256 borrowAmount, address borrowee) returns(uint256)
func (_LToken *LTokenTransactor) BorrowBehalf(opts *bind.TransactOpts, borrowAmount *big.Int, borrowee common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "borrowBehalf", borrowAmount, borrowee)
}

// BorrowBehalf is a paid mutator transaction binding the contract method 0x60a7cf20.
//
// Solidity: function borrowBehalf(uint256 borrowAmount, address borrowee) returns(uint256)
func (_LToken *LTokenSession) BorrowBehalf(borrowAmount *big.Int, borrowee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BorrowBehalf(&_LToken.TransactOpts, borrowAmount, borrowee)
}

// BorrowBehalf is a paid mutator transaction binding the contract method 0x60a7cf20.
//
// Solidity: function borrowBehalf(uint256 borrowAmount, address borrowee) returns(uint256)
func (_LToken *LTokenTransactorSession) BorrowBehalf(borrowAmount *big.Int, borrowee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.BorrowBehalf(&_LToken.TransactOpts, borrowAmount, borrowee)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_LToken *LTokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_LToken *LTokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _LToken.Contract.ExchangeRateCurrent(&_LToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_LToken *LTokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _LToken.Contract.ExchangeRateCurrent(&_LToken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenTransactor) Initialize(opts *bind.TransactOpts, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "initialize", underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenSession) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.Contract.Initialize(&_LToken.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenTransactorSession) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.Contract.Initialize(&_LToken.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenTransactor) Initialize0(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "initialize0", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenSession) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.Contract.Initialize0(&_LToken.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_LToken *LTokenTransactorSession) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _LToken.Contract.Initialize0(&_LToken.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_LToken *LTokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_LToken *LTokenSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _LToken.Contract.LiquidateBorrow(&_LToken.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_LToken *LTokenTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _LToken.Contract.LiquidateBorrow(&_LToken.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_LToken *LTokenTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_LToken *LTokenSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Mint(&_LToken.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Mint(&_LToken.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_LToken *LTokenTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_LToken *LTokenSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Redeem(&_LToken.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_LToken *LTokenTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Redeem(&_LToken.TransactOpts, redeemTokens)
}

// RedeemBehalf is a paid mutator transaction binding the contract method 0xde3c77eb.
//
// Solidity: function redeemBehalf(uint256 redeemTokens, address redeemee) returns(uint256)
func (_LToken *LTokenTransactor) RedeemBehalf(opts *bind.TransactOpts, redeemTokens *big.Int, redeemee common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "redeemBehalf", redeemTokens, redeemee)
}

// RedeemBehalf is a paid mutator transaction binding the contract method 0xde3c77eb.
//
// Solidity: function redeemBehalf(uint256 redeemTokens, address redeemee) returns(uint256)
func (_LToken *LTokenSession) RedeemBehalf(redeemTokens *big.Int, redeemee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.RedeemBehalf(&_LToken.TransactOpts, redeemTokens, redeemee)
}

// RedeemBehalf is a paid mutator transaction binding the contract method 0xde3c77eb.
//
// Solidity: function redeemBehalf(uint256 redeemTokens, address redeemee) returns(uint256)
func (_LToken *LTokenTransactorSession) RedeemBehalf(redeemTokens *big.Int, redeemee common.Address) (*types.Transaction, error) {
	return _LToken.Contract.RedeemBehalf(&_LToken.TransactOpts, redeemTokens, redeemee)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_LToken *LTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_LToken *LTokenSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RedeemUnderlying(&_LToken.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RedeemUnderlying(&_LToken.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_LToken *LTokenTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_LToken *LTokenSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RepayBorrow(&_LToken.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RepayBorrow(&_LToken.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_LToken *LTokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_LToken *LTokenSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RepayBorrowBehalf(&_LToken.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_LToken *LTokenTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.RepayBorrowBehalf(&_LToken.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_LToken *LTokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_LToken *LTokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Seize(&_LToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_LToken *LTokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Seize(&_LToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_LToken *LTokenTransactor) SweepToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "sweepToken", token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_LToken *LTokenSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SweepToken(&_LToken.TransactOpts, token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address token) returns()
func (_LToken *LTokenTransactorSession) SweepToken(token common.Address) (*types.Transaction, error) {
	return _LToken.Contract.SweepToken(&_LToken.TransactOpts, token)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_LToken *LTokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_LToken *LTokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _LToken.Contract.TotalBorrowsCurrent(&_LToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_LToken *LTokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _LToken.Contract.TotalBorrowsCurrent(&_LToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_LToken *LTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_LToken *LTokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Transfer(&_LToken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_LToken *LTokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.Transfer(&_LToken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_LToken *LTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_LToken *LTokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.TransferFrom(&_LToken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_LToken *LTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LToken.Contract.TransferFrom(&_LToken.TransactOpts, src, dst, amount)
}

// LTokenAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the LToken contract.
type LTokenAccrueInterestIterator struct {
	Event *LTokenAccrueInterest // Event containing the contract specifics and raw log

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
func (it *LTokenAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenAccrueInterest)
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
		it.Event = new(LTokenAccrueInterest)
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
func (it *LTokenAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenAccrueInterest represents a AccrueInterest event raised by the LToken contract.
type LTokenAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_LToken *LTokenFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*LTokenAccrueInterestIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &LTokenAccrueInterestIterator{contract: _LToken.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_LToken *LTokenFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *LTokenAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenAccrueInterest)
				if err := _LToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_LToken *LTokenFilterer) ParseAccrueInterest(log types.Log) (*LTokenAccrueInterest, error) {
	event := new(LTokenAccrueInterest)
	if err := _LToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LToken contract.
type LTokenApprovalIterator struct {
	Event *LTokenApproval // Event containing the contract specifics and raw log

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
func (it *LTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenApproval)
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
		it.Event = new(LTokenApproval)
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
func (it *LTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenApproval represents a Approval event raised by the LToken contract.
type LTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_LToken *LTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LTokenApprovalIterator{contract: _LToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_LToken *LTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenApproval)
				if err := _LToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_LToken *LTokenFilterer) ParseApproval(log types.Log) (*LTokenApproval, error) {
	event := new(LTokenApproval)
	if err := _LToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LToken contract.
type LTokenBorrowIterator struct {
	Event *LTokenBorrow // Event containing the contract specifics and raw log

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
func (it *LTokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenBorrow)
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
		it.Event = new(LTokenBorrow)
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
func (it *LTokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenBorrow represents a Borrow event raised by the LToken contract.
type LTokenBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	BorrowIndex    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) FilterBorrow(opts *bind.FilterOpts) (*LTokenBorrowIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &LTokenBorrowIterator{contract: _LToken.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x2dd79f4fccfd18c360ce7f9132f3621bf05eee18f995224badb32d17f172df73.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LTokenBorrow) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenBorrow)
				if err := _LToken.contract.UnpackLog(event, "Borrow", log); err != nil {
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
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) ParseBorrow(log types.Log) (*LTokenBorrow, error) {
	event := new(LTokenBorrow)
	if err := _LToken.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the LToken contract.
type LTokenLiquidateBorrowIterator struct {
	Event *LTokenLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *LTokenLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenLiquidateBorrow)
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
		it.Event = new(LTokenLiquidateBorrow)
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
func (it *LTokenLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenLiquidateBorrow represents a LiquidateBorrow event raised by the LToken contract.
type LTokenLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_LToken *LTokenFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*LTokenLiquidateBorrowIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &LTokenLiquidateBorrowIterator{contract: _LToken.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_LToken *LTokenFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *LTokenLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenLiquidateBorrow)
				if err := _LToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_LToken *LTokenFilterer) ParseLiquidateBorrow(log types.Log) (*LTokenLiquidateBorrow, error) {
	event := new(LTokenLiquidateBorrow)
	if err := _LToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the LToken contract.
type LTokenMintIterator struct {
	Event *LTokenMint // Event containing the contract specifics and raw log

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
func (it *LTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenMint)
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
		it.Event = new(LTokenMint)
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
func (it *LTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenMint represents a Mint event raised by the LToken contract.
type LTokenMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_LToken *LTokenFilterer) FilterMint(opts *bind.FilterOpts) (*LTokenMintIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &LTokenMintIterator{contract: _LToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_LToken *LTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LTokenMint) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenMint)
				if err := _LToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_LToken *LTokenFilterer) ParseMint(log types.Log) (*LTokenMint, error) {
	event := new(LTokenMint)
	if err := _LToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the LToken contract.
type LTokenNewAdminIterator struct {
	Event *LTokenNewAdmin // Event containing the contract specifics and raw log

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
func (it *LTokenNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewAdmin)
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
		it.Event = new(LTokenNewAdmin)
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
func (it *LTokenNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewAdmin represents a NewAdmin event raised by the LToken contract.
type LTokenNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_LToken *LTokenFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*LTokenNewAdminIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &LTokenNewAdminIterator{contract: _LToken.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_LToken *LTokenFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *LTokenNewAdmin) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewAdmin)
				if err := _LToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_LToken *LTokenFilterer) ParseNewAdmin(log types.Log) (*LTokenNewAdmin, error) {
	event := new(LTokenNewAdmin)
	if err := _LToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the LToken contract.
type LTokenNewComptrollerIterator struct {
	Event *LTokenNewComptroller // Event containing the contract specifics and raw log

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
func (it *LTokenNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewComptroller)
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
		it.Event = new(LTokenNewComptroller)
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
func (it *LTokenNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewComptroller represents a NewComptroller event raised by the LToken contract.
type LTokenNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_LToken *LTokenFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*LTokenNewComptrollerIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &LTokenNewComptrollerIterator{contract: _LToken.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_LToken *LTokenFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *LTokenNewComptroller) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewComptroller)
				if err := _LToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_LToken *LTokenFilterer) ParseNewComptroller(log types.Log) (*LTokenNewComptroller, error) {
	event := new(LTokenNewComptroller)
	if err := _LToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the LToken contract.
type LTokenNewMarketInterestRateModelIterator struct {
	Event *LTokenNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *LTokenNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewMarketInterestRateModel)
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
		it.Event = new(LTokenNewMarketInterestRateModel)
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
func (it *LTokenNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the LToken contract.
type LTokenNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_LToken *LTokenFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*LTokenNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &LTokenNewMarketInterestRateModelIterator{contract: _LToken.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_LToken *LTokenFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *LTokenNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewMarketInterestRateModel)
				if err := _LToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_LToken *LTokenFilterer) ParseNewMarketInterestRateModel(log types.Log) (*LTokenNewMarketInterestRateModel, error) {
	event := new(LTokenNewMarketInterestRateModel)
	if err := _LToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the LToken contract.
type LTokenNewPendingAdminIterator struct {
	Event *LTokenNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *LTokenNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewPendingAdmin)
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
		it.Event = new(LTokenNewPendingAdmin)
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
func (it *LTokenNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewPendingAdmin represents a NewPendingAdmin event raised by the LToken contract.
type LTokenNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_LToken *LTokenFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*LTokenNewPendingAdminIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &LTokenNewPendingAdminIterator{contract: _LToken.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_LToken *LTokenFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *LTokenNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewPendingAdmin)
				if err := _LToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_LToken *LTokenFilterer) ParseNewPendingAdmin(log types.Log) (*LTokenNewPendingAdmin, error) {
	event := new(LTokenNewPendingAdmin)
	if err := _LToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the LToken contract.
type LTokenNewReserveFactorIterator struct {
	Event *LTokenNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *LTokenNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewReserveFactor)
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
		it.Event = new(LTokenNewReserveFactor)
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
func (it *LTokenNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewReserveFactor represents a NewReserveFactor event raised by the LToken contract.
type LTokenNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_LToken *LTokenFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*LTokenNewReserveFactorIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &LTokenNewReserveFactorIterator{contract: _LToken.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_LToken *LTokenFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *LTokenNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewReserveFactor)
				if err := _LToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_LToken *LTokenFilterer) ParseNewReserveFactor(log types.Log) (*LTokenNewReserveFactor, error) {
	event := new(LTokenNewReserveFactor)
	if err := _LToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenNewReserveGuardianIterator is returned from FilterNewReserveGuardian and is used to iterate over the raw logs and unpacked data for NewReserveGuardian events raised by the LToken contract.
type LTokenNewReserveGuardianIterator struct {
	Event *LTokenNewReserveGuardian // Event containing the contract specifics and raw log

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
func (it *LTokenNewReserveGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenNewReserveGuardian)
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
		it.Event = new(LTokenNewReserveGuardian)
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
func (it *LTokenNewReserveGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenNewReserveGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenNewReserveGuardian represents a NewReserveGuardian event raised by the LToken contract.
type LTokenNewReserveGuardian struct {
	OldReserveGuardian common.Address
	NewReserveGuardian common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewReserveGuardian is a free log retrieval operation binding the contract event 0xe3f5272426aa3e70e583474d1155b19c7d9a5394147c40191dfaf32b792bc978.
//
// Solidity: event NewReserveGuardian(address oldReserveGuardian, address newReserveGuardian)
func (_LToken *LTokenFilterer) FilterNewReserveGuardian(opts *bind.FilterOpts) (*LTokenNewReserveGuardianIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "NewReserveGuardian")
	if err != nil {
		return nil, err
	}
	return &LTokenNewReserveGuardianIterator{contract: _LToken.contract, event: "NewReserveGuardian", logs: logs, sub: sub}, nil
}

// WatchNewReserveGuardian is a free log subscription operation binding the contract event 0xe3f5272426aa3e70e583474d1155b19c7d9a5394147c40191dfaf32b792bc978.
//
// Solidity: event NewReserveGuardian(address oldReserveGuardian, address newReserveGuardian)
func (_LToken *LTokenFilterer) WatchNewReserveGuardian(opts *bind.WatchOpts, sink chan<- *LTokenNewReserveGuardian) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "NewReserveGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenNewReserveGuardian)
				if err := _LToken.contract.UnpackLog(event, "NewReserveGuardian", log); err != nil {
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

// ParseNewReserveGuardian is a log parse operation binding the contract event 0xe3f5272426aa3e70e583474d1155b19c7d9a5394147c40191dfaf32b792bc978.
//
// Solidity: event NewReserveGuardian(address oldReserveGuardian, address newReserveGuardian)
func (_LToken *LTokenFilterer) ParseNewReserveGuardian(log types.Log) (*LTokenNewReserveGuardian, error) {
	event := new(LTokenNewReserveGuardian)
	if err := _LToken.contract.UnpackLog(event, "NewReserveGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the LToken contract.
type LTokenRedeemIterator struct {
	Event *LTokenRedeem // Event containing the contract specifics and raw log

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
func (it *LTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenRedeem)
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
		it.Event = new(LTokenRedeem)
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
func (it *LTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenRedeem represents a Redeem event raised by the LToken contract.
type LTokenRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_LToken *LTokenFilterer) FilterRedeem(opts *bind.FilterOpts) (*LTokenRedeemIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &LTokenRedeemIterator{contract: _LToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_LToken *LTokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *LTokenRedeem) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenRedeem)
				if err := _LToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_LToken *LTokenFilterer) ParseRedeem(log types.Log) (*LTokenRedeem, error) {
	event := new(LTokenRedeem)
	if err := _LToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the LToken contract.
type LTokenRepayBorrowIterator struct {
	Event *LTokenRepayBorrow // Event containing the contract specifics and raw log

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
func (it *LTokenRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenRepayBorrow)
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
		it.Event = new(LTokenRepayBorrow)
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
func (it *LTokenRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenRepayBorrow represents a RepayBorrow event raised by the LToken contract.
type LTokenRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	BorrowIndex    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x6fadbf7329d21f278e724fa0d4511001a158f2a97ee35c5bc4cf8b64417399ef.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*LTokenRepayBorrowIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &LTokenRepayBorrowIterator{contract: _LToken.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x6fadbf7329d21f278e724fa0d4511001a158f2a97ee35c5bc4cf8b64417399ef.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *LTokenRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenRepayBorrow)
				if err := _LToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows, uint256 borrowIndex)
func (_LToken *LTokenFilterer) ParseRepayBorrow(log types.Log) (*LTokenRepayBorrow, error) {
	event := new(LTokenRepayBorrow)
	if err := _LToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the LToken contract.
type LTokenReservesAddedIterator struct {
	Event *LTokenReservesAdded // Event containing the contract specifics and raw log

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
func (it *LTokenReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenReservesAdded)
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
		it.Event = new(LTokenReservesAdded)
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
func (it *LTokenReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenReservesAdded represents a ReservesAdded event raised by the LToken contract.
type LTokenReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*LTokenReservesAddedIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &LTokenReservesAddedIterator{contract: _LToken.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *LTokenReservesAdded) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenReservesAdded)
				if err := _LToken.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) ParseReservesAdded(log types.Log) (*LTokenReservesAdded, error) {
	event := new(LTokenReservesAdded)
	if err := _LToken.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the LToken contract.
type LTokenReservesReducedIterator struct {
	Event *LTokenReservesReduced // Event containing the contract specifics and raw log

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
func (it *LTokenReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenReservesReduced)
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
		it.Event = new(LTokenReservesReduced)
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
func (it *LTokenReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenReservesReduced represents a ReservesReduced event raised by the LToken contract.
type LTokenReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*LTokenReservesReducedIterator, error) {

	logs, sub, err := _LToken.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &LTokenReservesReducedIterator{contract: _LToken.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *LTokenReservesReduced) (event.Subscription, error) {

	logs, sub, err := _LToken.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenReservesReduced)
				if err := _LToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_LToken *LTokenFilterer) ParseReservesReduced(log types.Log) (*LTokenReservesReduced, error) {
	event := new(LTokenReservesReduced)
	if err := _LToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LToken contract.
type LTokenTransferIterator struct {
	Event *LTokenTransfer // Event containing the contract specifics and raw log

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
func (it *LTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LTokenTransfer)
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
		it.Event = new(LTokenTransfer)
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
func (it *LTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LTokenTransfer represents a Transfer event raised by the LToken contract.
type LTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_LToken *LTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LTokenTransferIterator{contract: _LToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_LToken *LTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LTokenTransfer)
				if err := _LToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_LToken *LTokenFilterer) ParseTransfer(log types.Log) (*LTokenTransfer, error) {
	event := new(LTokenTransfer)
	if err := _LToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
