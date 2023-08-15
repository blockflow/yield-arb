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

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"BorrowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BorrowedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BorrowedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyCapacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowCapacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distributionFactor\",\"type\":\"uint256\"}],\"name\":\"MarketAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"MintPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBorrowCapacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCapacity\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCapacity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBorrowFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewBorrowFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"NewPendingOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRewardDistributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newRewardDistributor\",\"type\":\"address\"}],\"name\":\"NewRewardDistributor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSupplyCapacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSupplyCapacity\",\"type\":\"uint256\"}],\"name\":\"NewSupplyCapacity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"iToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"RedeemPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"SeizePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"TransferPaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acceptOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_supplyCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_distributionFactor\",\"type\":\"uint256\"}],\"name\":\"_addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setAllBorrowPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setAllMintPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setAllRedeemPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newBorrowCapacity\",\"type\":\"uint256\"}],\"name\":\"_setBorrowCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newBorrowFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setBorrowFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"_setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setRedeemPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardDistributor\",\"type\":\"address\"}],\"name\":\"_setRewardDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newSupplyCapacity\",\"type\":\"uint256\"}],\"name\":\"_setSupplyCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"_setiTokenPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_borrowedAmount\",\"type\":\"uint256\"}],\"name\":\"afterBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"afterFlashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repaidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seizedAmount\",\"type\":\"uint256\"}],\"name\":\"afterLiquidateBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintedAmount\",\"type\":\"uint256\"}],\"name\":\"afterMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_redeemedUnderlying\",\"type\":\"uint256\"}],\"name\":\"afterRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"afterRepayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seizedAmount\",\"type\":\"uint256\"}],\"name\":\"afterSeize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"afterTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_borrowAmount\",\"type\":\"uint256\"}],\"name\":\"beforeBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"beforeFlashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"beforeLiquidateBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"beforeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemAmount\",\"type\":\"uint256\"}],\"name\":\"beforeRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_repayAmount\",\"type\":\"uint256\"}],\"name\":\"beforeRepayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seizeAmount\",\"type\":\"uint256\"}],\"name\":\"beforeSeize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"beforeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"calcAccountEquity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_market\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"enterMarketFromiToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_iTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_results\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_iTokens\",\"type\":\"address[]\"}],\"name\":\"exitMarkets\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_results\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlliTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_alliTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getBorrowedAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_borrowedAssets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getEnteredMarkets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_accountCollaterals\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"}],\"name\":\"hasBorrowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"}],\"name\":\"hasEnteredMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iToken\",\"type\":\"address\"}],\"name\":\"hasiToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isController\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_iTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_seizedTokenCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCapacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCapacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mintPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"redeemPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// CalcAccountEquity is a free data retrieval call binding the contract method 0xa09cf78b.
//
// Solidity: function calcAccountEquity(address _account) view returns(uint256, uint256, uint256, uint256)
func (_Controller *ControllerCaller) CalcAccountEquity(opts *bind.CallOpts, _account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "calcAccountEquity", _account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// CalcAccountEquity is a free data retrieval call binding the contract method 0xa09cf78b.
//
// Solidity: function calcAccountEquity(address _account) view returns(uint256, uint256, uint256, uint256)
func (_Controller *ControllerSession) CalcAccountEquity(_account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Controller.Contract.CalcAccountEquity(&_Controller.CallOpts, _account)
}

// CalcAccountEquity is a free data retrieval call binding the contract method 0xa09cf78b.
//
// Solidity: function calcAccountEquity(address _account) view returns(uint256, uint256, uint256, uint256)
func (_Controller *ControllerCallerSession) CalcAccountEquity(_account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Controller.Contract.CalcAccountEquity(&_Controller.CallOpts, _account)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Controller *ControllerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Controller *ControllerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Controller.Contract.CloseFactorMantissa(&_Controller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Controller *ControllerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Controller.Contract.CloseFactorMantissa(&_Controller.CallOpts)
}

// GetAlliTokens is a free data retrieval call binding the contract method 0x60a8a931.
//
// Solidity: function getAlliTokens() view returns(address[] _alliTokens)
func (_Controller *ControllerCaller) GetAlliTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getAlliTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAlliTokens is a free data retrieval call binding the contract method 0x60a8a931.
//
// Solidity: function getAlliTokens() view returns(address[] _alliTokens)
func (_Controller *ControllerSession) GetAlliTokens() ([]common.Address, error) {
	return _Controller.Contract.GetAlliTokens(&_Controller.CallOpts)
}

// GetAlliTokens is a free data retrieval call binding the contract method 0x60a8a931.
//
// Solidity: function getAlliTokens() view returns(address[] _alliTokens)
func (_Controller *ControllerCallerSession) GetAlliTokens() ([]common.Address, error) {
	return _Controller.Contract.GetAlliTokens(&_Controller.CallOpts)
}

// GetBorrowedAssets is a free data retrieval call binding the contract method 0x075aff36.
//
// Solidity: function getBorrowedAssets(address _account) view returns(address[] _borrowedAssets)
func (_Controller *ControllerCaller) GetBorrowedAssets(opts *bind.CallOpts, _account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getBorrowedAssets", _account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBorrowedAssets is a free data retrieval call binding the contract method 0x075aff36.
//
// Solidity: function getBorrowedAssets(address _account) view returns(address[] _borrowedAssets)
func (_Controller *ControllerSession) GetBorrowedAssets(_account common.Address) ([]common.Address, error) {
	return _Controller.Contract.GetBorrowedAssets(&_Controller.CallOpts, _account)
}

// GetBorrowedAssets is a free data retrieval call binding the contract method 0x075aff36.
//
// Solidity: function getBorrowedAssets(address _account) view returns(address[] _borrowedAssets)
func (_Controller *ControllerCallerSession) GetBorrowedAssets(_account common.Address) ([]common.Address, error) {
	return _Controller.Contract.GetBorrowedAssets(&_Controller.CallOpts, _account)
}

// GetEnteredMarkets is a free data retrieval call binding the contract method 0x8ccb720b.
//
// Solidity: function getEnteredMarkets(address _account) view returns(address[] _accountCollaterals)
func (_Controller *ControllerCaller) GetEnteredMarkets(opts *bind.CallOpts, _account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getEnteredMarkets", _account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEnteredMarkets is a free data retrieval call binding the contract method 0x8ccb720b.
//
// Solidity: function getEnteredMarkets(address _account) view returns(address[] _accountCollaterals)
func (_Controller *ControllerSession) GetEnteredMarkets(_account common.Address) ([]common.Address, error) {
	return _Controller.Contract.GetEnteredMarkets(&_Controller.CallOpts, _account)
}

// GetEnteredMarkets is a free data retrieval call binding the contract method 0x8ccb720b.
//
// Solidity: function getEnteredMarkets(address _account) view returns(address[] _accountCollaterals)
func (_Controller *ControllerCallerSession) GetEnteredMarkets(_account common.Address) ([]common.Address, error) {
	return _Controller.Contract.GetEnteredMarkets(&_Controller.CallOpts, _account)
}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerCaller) HasBorrowed(opts *bind.CallOpts, _account common.Address, _iToken common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "hasBorrowed", _account, _iToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerSession) HasBorrowed(_account common.Address, _iToken common.Address) (bool, error) {
	return _Controller.Contract.HasBorrowed(&_Controller.CallOpts, _account, _iToken)
}

// HasBorrowed is a free data retrieval call binding the contract method 0x8283e3e2.
//
// Solidity: function hasBorrowed(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerCallerSession) HasBorrowed(_account common.Address, _iToken common.Address) (bool, error) {
	return _Controller.Contract.HasBorrowed(&_Controller.CallOpts, _account, _iToken)
}

// HasEnteredMarket is a free data retrieval call binding the contract method 0xd4d48473.
//
// Solidity: function hasEnteredMarket(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerCaller) HasEnteredMarket(opts *bind.CallOpts, _account common.Address, _iToken common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "hasEnteredMarket", _account, _iToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEnteredMarket is a free data retrieval call binding the contract method 0xd4d48473.
//
// Solidity: function hasEnteredMarket(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerSession) HasEnteredMarket(_account common.Address, _iToken common.Address) (bool, error) {
	return _Controller.Contract.HasEnteredMarket(&_Controller.CallOpts, _account, _iToken)
}

// HasEnteredMarket is a free data retrieval call binding the contract method 0xd4d48473.
//
// Solidity: function hasEnteredMarket(address _account, address _iToken) view returns(bool)
func (_Controller *ControllerCallerSession) HasEnteredMarket(_account common.Address, _iToken common.Address) (bool, error) {
	return _Controller.Contract.HasEnteredMarket(&_Controller.CallOpts, _account, _iToken)
}

// HasiToken is a free data retrieval call binding the contract method 0x4428e862.
//
// Solidity: function hasiToken(address _iToken) view returns(bool)
func (_Controller *ControllerCaller) HasiToken(opts *bind.CallOpts, _iToken common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "hasiToken", _iToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasiToken is a free data retrieval call binding the contract method 0x4428e862.
//
// Solidity: function hasiToken(address _iToken) view returns(bool)
func (_Controller *ControllerSession) HasiToken(_iToken common.Address) (bool, error) {
	return _Controller.Contract.HasiToken(&_Controller.CallOpts, _iToken)
}

// HasiToken is a free data retrieval call binding the contract method 0x4428e862.
//
// Solidity: function hasiToken(address _iToken) view returns(bool)
func (_Controller *ControllerCallerSession) HasiToken(_iToken common.Address) (bool, error) {
	return _Controller.Contract.HasiToken(&_Controller.CallOpts, _iToken)
}

// IsController is a free data retrieval call binding the contract method 0x4e1647fb.
//
// Solidity: function isController() view returns(bool)
func (_Controller *ControllerCaller) IsController(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "isController")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsController is a free data retrieval call binding the contract method 0x4e1647fb.
//
// Solidity: function isController() view returns(bool)
func (_Controller *ControllerSession) IsController() (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts)
}

// IsController is a free data retrieval call binding the contract method 0x4e1647fb.
//
// Solidity: function isController() view returns(bool)
func (_Controller *ControllerCallerSession) IsController() (bool, error) {
	return _Controller.Contract.IsController(&_Controller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address _iTokenBorrowed, address _iTokenCollateral, uint256 _actualRepayAmount) view returns(uint256 _seizedTokenCollateral)
func (_Controller *ControllerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, _iTokenBorrowed common.Address, _iTokenCollateral common.Address, _actualRepayAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", _iTokenBorrowed, _iTokenCollateral, _actualRepayAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address _iTokenBorrowed, address _iTokenCollateral, uint256 _actualRepayAmount) view returns(uint256 _seizedTokenCollateral)
func (_Controller *ControllerSession) LiquidateCalculateSeizeTokens(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _actualRepayAmount *big.Int) (*big.Int, error) {
	return _Controller.Contract.LiquidateCalculateSeizeTokens(&_Controller.CallOpts, _iTokenBorrowed, _iTokenCollateral, _actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address _iTokenBorrowed, address _iTokenCollateral, uint256 _actualRepayAmount) view returns(uint256 _seizedTokenCollateral)
func (_Controller *ControllerCallerSession) LiquidateCalculateSeizeTokens(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _actualRepayAmount *big.Int) (*big.Int, error) {
	return _Controller.Contract.LiquidateCalculateSeizeTokens(&_Controller.CallOpts, _iTokenBorrowed, _iTokenCollateral, _actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Controller *ControllerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Controller *ControllerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Controller.Contract.LiquidationIncentiveMantissa(&_Controller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Controller *ControllerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Controller.Contract.LiquidationIncentiveMantissa(&_Controller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(uint256 collateralFactorMantissa, uint256 borrowFactorMantissa, uint256 borrowCapacity, uint256 supplyCapacity, bool mintPaused, bool redeemPaused, bool borrowPaused)
func (_Controller *ControllerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	CollateralFactorMantissa *big.Int
	BorrowFactorMantissa     *big.Int
	BorrowCapacity           *big.Int
	SupplyCapacity           *big.Int
	MintPaused               bool
	RedeemPaused             bool
	BorrowPaused             bool
}, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		CollateralFactorMantissa *big.Int
		BorrowFactorMantissa     *big.Int
		BorrowCapacity           *big.Int
		SupplyCapacity           *big.Int
		MintPaused               bool
		RedeemPaused             bool
		BorrowPaused             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BorrowFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BorrowCapacity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SupplyCapacity = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MintPaused = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.RedeemPaused = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowPaused = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(uint256 collateralFactorMantissa, uint256 borrowFactorMantissa, uint256 borrowCapacity, uint256 supplyCapacity, bool mintPaused, bool redeemPaused, bool borrowPaused)
func (_Controller *ControllerSession) Markets(arg0 common.Address) (struct {
	CollateralFactorMantissa *big.Int
	BorrowFactorMantissa     *big.Int
	BorrowCapacity           *big.Int
	SupplyCapacity           *big.Int
	MintPaused               bool
	RedeemPaused             bool
	BorrowPaused             bool
}, error) {
	return _Controller.Contract.Markets(&_Controller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(uint256 collateralFactorMantissa, uint256 borrowFactorMantissa, uint256 borrowCapacity, uint256 supplyCapacity, bool mintPaused, bool redeemPaused, bool borrowPaused)
func (_Controller *ControllerCallerSession) Markets(arg0 common.Address) (struct {
	CollateralFactorMantissa *big.Int
	BorrowFactorMantissa     *big.Int
	BorrowCapacity           *big.Int
	SupplyCapacity           *big.Int
	MintPaused               bool
	RedeemPaused             bool
	BorrowPaused             bool
}, error) {
	return _Controller.Contract.Markets(&_Controller.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Controller *ControllerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Controller *ControllerSession) PauseGuardian() (common.Address, error) {
	return _Controller.Contract.PauseGuardian(&_Controller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Controller *ControllerCallerSession) PauseGuardian() (common.Address, error) {
	return _Controller.Contract.PauseGuardian(&_Controller.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Controller *ControllerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Controller *ControllerSession) PendingOwner() (common.Address, error) {
	return _Controller.Contract.PendingOwner(&_Controller.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Controller *ControllerCallerSession) PendingOwner() (common.Address, error) {
	return _Controller.Contract.PendingOwner(&_Controller.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Controller *ControllerCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Controller *ControllerSession) PriceOracle() (common.Address, error) {
	return _Controller.Contract.PriceOracle(&_Controller.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Controller *ControllerCallerSession) PriceOracle() (common.Address, error) {
	return _Controller.Contract.PriceOracle(&_Controller.CallOpts)
}

// RewardDistributor is a free data retrieval call binding the contract method 0xacc2166a.
//
// Solidity: function rewardDistributor() view returns(address)
func (_Controller *ControllerCaller) RewardDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "rewardDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardDistributor is a free data retrieval call binding the contract method 0xacc2166a.
//
// Solidity: function rewardDistributor() view returns(address)
func (_Controller *ControllerSession) RewardDistributor() (common.Address, error) {
	return _Controller.Contract.RewardDistributor(&_Controller.CallOpts)
}

// RewardDistributor is a free data retrieval call binding the contract method 0xacc2166a.
//
// Solidity: function rewardDistributor() view returns(address)
func (_Controller *ControllerCallerSession) RewardDistributor() (common.Address, error) {
	return _Controller.Contract.RewardDistributor(&_Controller.CallOpts)
}

// SeizePaused is a free data retrieval call binding the contract method 0xab1777d2.
//
// Solidity: function seizePaused() view returns(bool)
func (_Controller *ControllerCaller) SeizePaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "seizePaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizePaused is a free data retrieval call binding the contract method 0xab1777d2.
//
// Solidity: function seizePaused() view returns(bool)
func (_Controller *ControllerSession) SeizePaused() (bool, error) {
	return _Controller.Contract.SeizePaused(&_Controller.CallOpts)
}

// SeizePaused is a free data retrieval call binding the contract method 0xab1777d2.
//
// Solidity: function seizePaused() view returns(bool)
func (_Controller *ControllerCallerSession) SeizePaused() (bool, error) {
	return _Controller.Contract.SeizePaused(&_Controller.CallOpts)
}

// TransferPaused is a free data retrieval call binding the contract method 0xfb2cb34e.
//
// Solidity: function transferPaused() view returns(bool)
func (_Controller *ControllerCaller) TransferPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "transferPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferPaused is a free data retrieval call binding the contract method 0xfb2cb34e.
//
// Solidity: function transferPaused() view returns(bool)
func (_Controller *ControllerSession) TransferPaused() (bool, error) {
	return _Controller.Contract.TransferPaused(&_Controller.CallOpts)
}

// TransferPaused is a free data retrieval call binding the contract method 0xfb2cb34e.
//
// Solidity: function transferPaused() view returns(bool)
func (_Controller *ControllerCallerSession) TransferPaused() (bool, error) {
	return _Controller.Contract.TransferPaused(&_Controller.CallOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Controller *ControllerTransactor) AcceptOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_acceptOwner")
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Controller *ControllerSession) AcceptOwner() (*types.Transaction, error) {
	return _Controller.Contract.AcceptOwner(&_Controller.TransactOpts)
}

// AcceptOwner is a paid mutator transaction binding the contract method 0xfc4d33f9.
//
// Solidity: function _acceptOwner() returns()
func (_Controller *ControllerTransactorSession) AcceptOwner() (*types.Transaction, error) {
	return _Controller.Contract.AcceptOwner(&_Controller.TransactOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0xd79287d7.
//
// Solidity: function _addMarket(address _iToken, uint256 _collateralFactor, uint256 _borrowFactor, uint256 _supplyCapacity, uint256 _borrowCapacity, uint256 _distributionFactor) returns()
func (_Controller *ControllerTransactor) AddMarket(opts *bind.TransactOpts, _iToken common.Address, _collateralFactor *big.Int, _borrowFactor *big.Int, _supplyCapacity *big.Int, _borrowCapacity *big.Int, _distributionFactor *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_addMarket", _iToken, _collateralFactor, _borrowFactor, _supplyCapacity, _borrowCapacity, _distributionFactor)
}

// AddMarket is a paid mutator transaction binding the contract method 0xd79287d7.
//
// Solidity: function _addMarket(address _iToken, uint256 _collateralFactor, uint256 _borrowFactor, uint256 _supplyCapacity, uint256 _borrowCapacity, uint256 _distributionFactor) returns()
func (_Controller *ControllerSession) AddMarket(_iToken common.Address, _collateralFactor *big.Int, _borrowFactor *big.Int, _supplyCapacity *big.Int, _borrowCapacity *big.Int, _distributionFactor *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddMarket(&_Controller.TransactOpts, _iToken, _collateralFactor, _borrowFactor, _supplyCapacity, _borrowCapacity, _distributionFactor)
}

// AddMarket is a paid mutator transaction binding the contract method 0xd79287d7.
//
// Solidity: function _addMarket(address _iToken, uint256 _collateralFactor, uint256 _borrowFactor, uint256 _supplyCapacity, uint256 _borrowCapacity, uint256 _distributionFactor) returns()
func (_Controller *ControllerTransactorSession) AddMarket(_iToken common.Address, _collateralFactor *big.Int, _borrowFactor *big.Int, _supplyCapacity *big.Int, _borrowCapacity *big.Int, _distributionFactor *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AddMarket(&_Controller.TransactOpts, _iToken, _collateralFactor, _borrowFactor, _supplyCapacity, _borrowCapacity, _distributionFactor)
}

// SetAllBorrowPaused is a paid mutator transaction binding the contract method 0xe715d96a.
//
// Solidity: function _setAllBorrowPaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetAllBorrowPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setAllBorrowPaused", _paused)
}

// SetAllBorrowPaused is a paid mutator transaction binding the contract method 0xe715d96a.
//
// Solidity: function _setAllBorrowPaused(bool _paused) returns()
func (_Controller *ControllerSession) SetAllBorrowPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllBorrowPaused(&_Controller.TransactOpts, _paused)
}

// SetAllBorrowPaused is a paid mutator transaction binding the contract method 0xe715d96a.
//
// Solidity: function _setAllBorrowPaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetAllBorrowPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllBorrowPaused(&_Controller.TransactOpts, _paused)
}

// SetAllMintPaused is a paid mutator transaction binding the contract method 0xdbb32d68.
//
// Solidity: function _setAllMintPaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetAllMintPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setAllMintPaused", _paused)
}

// SetAllMintPaused is a paid mutator transaction binding the contract method 0xdbb32d68.
//
// Solidity: function _setAllMintPaused(bool _paused) returns()
func (_Controller *ControllerSession) SetAllMintPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllMintPaused(&_Controller.TransactOpts, _paused)
}

// SetAllMintPaused is a paid mutator transaction binding the contract method 0xdbb32d68.
//
// Solidity: function _setAllMintPaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetAllMintPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllMintPaused(&_Controller.TransactOpts, _paused)
}

// SetAllRedeemPaused is a paid mutator transaction binding the contract method 0xe29f7f5e.
//
// Solidity: function _setAllRedeemPaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetAllRedeemPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setAllRedeemPaused", _paused)
}

// SetAllRedeemPaused is a paid mutator transaction binding the contract method 0xe29f7f5e.
//
// Solidity: function _setAllRedeemPaused(bool _paused) returns()
func (_Controller *ControllerSession) SetAllRedeemPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllRedeemPaused(&_Controller.TransactOpts, _paused)
}

// SetAllRedeemPaused is a paid mutator transaction binding the contract method 0xe29f7f5e.
//
// Solidity: function _setAllRedeemPaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetAllRedeemPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetAllRedeemPaused(&_Controller.TransactOpts, _paused)
}

// SetBorrowCapacity is a paid mutator transaction binding the contract method 0xbd8296c2.
//
// Solidity: function _setBorrowCapacity(address _iToken, uint256 _newBorrowCapacity) returns()
func (_Controller *ControllerTransactor) SetBorrowCapacity(opts *bind.TransactOpts, _iToken common.Address, _newBorrowCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setBorrowCapacity", _iToken, _newBorrowCapacity)
}

// SetBorrowCapacity is a paid mutator transaction binding the contract method 0xbd8296c2.
//
// Solidity: function _setBorrowCapacity(address _iToken, uint256 _newBorrowCapacity) returns()
func (_Controller *ControllerSession) SetBorrowCapacity(_iToken common.Address, _newBorrowCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowCapacity(&_Controller.TransactOpts, _iToken, _newBorrowCapacity)
}

// SetBorrowCapacity is a paid mutator transaction binding the contract method 0xbd8296c2.
//
// Solidity: function _setBorrowCapacity(address _iToken, uint256 _newBorrowCapacity) returns()
func (_Controller *ControllerTransactorSession) SetBorrowCapacity(_iToken common.Address, _newBorrowCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowCapacity(&_Controller.TransactOpts, _iToken, _newBorrowCapacity)
}

// SetBorrowFactor is a paid mutator transaction binding the contract method 0x9d8cc5c4.
//
// Solidity: function _setBorrowFactor(address _iToken, uint256 _newBorrowFactorMantissa) returns()
func (_Controller *ControllerTransactor) SetBorrowFactor(opts *bind.TransactOpts, _iToken common.Address, _newBorrowFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setBorrowFactor", _iToken, _newBorrowFactorMantissa)
}

// SetBorrowFactor is a paid mutator transaction binding the contract method 0x9d8cc5c4.
//
// Solidity: function _setBorrowFactor(address _iToken, uint256 _newBorrowFactorMantissa) returns()
func (_Controller *ControllerSession) SetBorrowFactor(_iToken common.Address, _newBorrowFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowFactor(&_Controller.TransactOpts, _iToken, _newBorrowFactorMantissa)
}

// SetBorrowFactor is a paid mutator transaction binding the contract method 0x9d8cc5c4.
//
// Solidity: function _setBorrowFactor(address _iToken, uint256 _newBorrowFactorMantissa) returns()
func (_Controller *ControllerTransactorSession) SetBorrowFactor(_iToken common.Address, _newBorrowFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowFactor(&_Controller.TransactOpts, _iToken, _newBorrowFactorMantissa)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactor) SetBorrowPaused(opts *bind.TransactOpts, _iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setBorrowPaused", _iToken, _paused)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerSession) SetBorrowPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetBorrowPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetBorrowPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 _newCloseFactorMantissa) returns()
func (_Controller *ControllerTransactor) SetCloseFactor(opts *bind.TransactOpts, _newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setCloseFactor", _newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 _newCloseFactorMantissa) returns()
func (_Controller *ControllerSession) SetCloseFactor(_newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetCloseFactor(&_Controller.TransactOpts, _newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 _newCloseFactorMantissa) returns()
func (_Controller *ControllerTransactorSession) SetCloseFactor(_newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetCloseFactor(&_Controller.TransactOpts, _newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address _iToken, uint256 _newCollateralFactorMantissa) returns()
func (_Controller *ControllerTransactor) SetCollateralFactor(opts *bind.TransactOpts, _iToken common.Address, _newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setCollateralFactor", _iToken, _newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address _iToken, uint256 _newCollateralFactorMantissa) returns()
func (_Controller *ControllerSession) SetCollateralFactor(_iToken common.Address, _newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetCollateralFactor(&_Controller.TransactOpts, _iToken, _newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address _iToken, uint256 _newCollateralFactorMantissa) returns()
func (_Controller *ControllerTransactorSession) SetCollateralFactor(_iToken common.Address, _newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetCollateralFactor(&_Controller.TransactOpts, _iToken, _newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 _newLiquidationIncentiveMantissa) returns()
func (_Controller *ControllerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, _newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setLiquidationIncentive", _newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 _newLiquidationIncentiveMantissa) returns()
func (_Controller *ControllerSession) SetLiquidationIncentive(_newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetLiquidationIncentive(&_Controller.TransactOpts, _newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 _newLiquidationIncentiveMantissa) returns()
func (_Controller *ControllerTransactorSession) SetLiquidationIncentive(_newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetLiquidationIncentive(&_Controller.TransactOpts, _newLiquidationIncentiveMantissa)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactor) SetMintPaused(opts *bind.TransactOpts, _iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setMintPaused", _iToken, _paused)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerSession) SetMintPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetMintPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetMintPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetMintPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address _newPauseGuardian) returns()
func (_Controller *ControllerTransactor) SetPauseGuardian(opts *bind.TransactOpts, _newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setPauseGuardian", _newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address _newPauseGuardian) returns()
func (_Controller *ControllerSession) SetPauseGuardian(_newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPauseGuardian(&_Controller.TransactOpts, _newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address _newPauseGuardian) returns()
func (_Controller *ControllerTransactorSession) SetPauseGuardian(_newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPauseGuardian(&_Controller.TransactOpts, _newPauseGuardian)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Controller *ControllerTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Controller *ControllerSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPendingOwner(&_Controller.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0x6e96dfd7.
//
// Solidity: function _setPendingOwner(address newPendingOwner) returns()
func (_Controller *ControllerTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPendingOwner(&_Controller.TransactOpts, newPendingOwner)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address _newOracle) returns()
func (_Controller *ControllerTransactor) SetPriceOracle(opts *bind.TransactOpts, _newOracle common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setPriceOracle", _newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address _newOracle) returns()
func (_Controller *ControllerSession) SetPriceOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPriceOracle(&_Controller.TransactOpts, _newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address _newOracle) returns()
func (_Controller *ControllerTransactorSession) SetPriceOracle(_newOracle common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetPriceOracle(&_Controller.TransactOpts, _newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetProtocolPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setProtocolPaused", _paused)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool _paused) returns()
func (_Controller *ControllerSession) SetProtocolPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetProtocolPaused(&_Controller.TransactOpts, _paused)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetProtocolPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetProtocolPaused(&_Controller.TransactOpts, _paused)
}

// SetRedeemPaused is a paid mutator transaction binding the contract method 0x7160a870.
//
// Solidity: function _setRedeemPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactor) SetRedeemPaused(opts *bind.TransactOpts, _iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setRedeemPaused", _iToken, _paused)
}

// SetRedeemPaused is a paid mutator transaction binding the contract method 0x7160a870.
//
// Solidity: function _setRedeemPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerSession) SetRedeemPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetRedeemPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetRedeemPaused is a paid mutator transaction binding the contract method 0x7160a870.
//
// Solidity: function _setRedeemPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetRedeemPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetRedeemPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x5c254d11.
//
// Solidity: function _setRewardDistributor(address _newRewardDistributor) returns()
func (_Controller *ControllerTransactor) SetRewardDistributor(opts *bind.TransactOpts, _newRewardDistributor common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setRewardDistributor", _newRewardDistributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x5c254d11.
//
// Solidity: function _setRewardDistributor(address _newRewardDistributor) returns()
func (_Controller *ControllerSession) SetRewardDistributor(_newRewardDistributor common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetRewardDistributor(&_Controller.TransactOpts, _newRewardDistributor)
}

// SetRewardDistributor is a paid mutator transaction binding the contract method 0x5c254d11.
//
// Solidity: function _setRewardDistributor(address _newRewardDistributor) returns()
func (_Controller *ControllerTransactorSession) SetRewardDistributor(_newRewardDistributor common.Address) (*types.Transaction, error) {
	return _Controller.Contract.SetRewardDistributor(&_Controller.TransactOpts, _newRewardDistributor)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetSeizePaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setSeizePaused", _paused)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool _paused) returns()
func (_Controller *ControllerSession) SetSeizePaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetSeizePaused(&_Controller.TransactOpts, _paused)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetSeizePaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetSeizePaused(&_Controller.TransactOpts, _paused)
}

// SetSupplyCapacity is a paid mutator transaction binding the contract method 0xd5b44360.
//
// Solidity: function _setSupplyCapacity(address _iToken, uint256 _newSupplyCapacity) returns()
func (_Controller *ControllerTransactor) SetSupplyCapacity(opts *bind.TransactOpts, _iToken common.Address, _newSupplyCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setSupplyCapacity", _iToken, _newSupplyCapacity)
}

// SetSupplyCapacity is a paid mutator transaction binding the contract method 0xd5b44360.
//
// Solidity: function _setSupplyCapacity(address _iToken, uint256 _newSupplyCapacity) returns()
func (_Controller *ControllerSession) SetSupplyCapacity(_iToken common.Address, _newSupplyCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetSupplyCapacity(&_Controller.TransactOpts, _iToken, _newSupplyCapacity)
}

// SetSupplyCapacity is a paid mutator transaction binding the contract method 0xd5b44360.
//
// Solidity: function _setSupplyCapacity(address _iToken, uint256 _newSupplyCapacity) returns()
func (_Controller *ControllerTransactorSession) SetSupplyCapacity(_iToken common.Address, _newSupplyCapacity *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.SetSupplyCapacity(&_Controller.TransactOpts, _iToken, _newSupplyCapacity)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool _paused) returns()
func (_Controller *ControllerTransactor) SetTransferPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setTransferPaused", _paused)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool _paused) returns()
func (_Controller *ControllerSession) SetTransferPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetTransferPaused(&_Controller.TransactOpts, _paused)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetTransferPaused(_paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetTransferPaused(&_Controller.TransactOpts, _paused)
}

// SetiTokenPaused is a paid mutator transaction binding the contract method 0xf1575f08.
//
// Solidity: function _setiTokenPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactor) SetiTokenPaused(opts *bind.TransactOpts, _iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "_setiTokenPaused", _iToken, _paused)
}

// SetiTokenPaused is a paid mutator transaction binding the contract method 0xf1575f08.
//
// Solidity: function _setiTokenPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerSession) SetiTokenPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetiTokenPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// SetiTokenPaused is a paid mutator transaction binding the contract method 0xf1575f08.
//
// Solidity: function _setiTokenPaused(address _iToken, bool _paused) returns()
func (_Controller *ControllerTransactorSession) SetiTokenPaused(_iToken common.Address, _paused bool) (*types.Transaction, error) {
	return _Controller.Contract.SetiTokenPaused(&_Controller.TransactOpts, _iToken, _paused)
}

// AfterBorrow is a paid mutator transaction binding the contract method 0xca498ad6.
//
// Solidity: function afterBorrow(address _iToken, address _borrower, uint256 _borrowedAmount) returns()
func (_Controller *ControllerTransactor) AfterBorrow(opts *bind.TransactOpts, _iToken common.Address, _borrower common.Address, _borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterBorrow", _iToken, _borrower, _borrowedAmount)
}

// AfterBorrow is a paid mutator transaction binding the contract method 0xca498ad6.
//
// Solidity: function afterBorrow(address _iToken, address _borrower, uint256 _borrowedAmount) returns()
func (_Controller *ControllerSession) AfterBorrow(_iToken common.Address, _borrower common.Address, _borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterBorrow(&_Controller.TransactOpts, _iToken, _borrower, _borrowedAmount)
}

// AfterBorrow is a paid mutator transaction binding the contract method 0xca498ad6.
//
// Solidity: function afterBorrow(address _iToken, address _borrower, uint256 _borrowedAmount) returns()
func (_Controller *ControllerTransactorSession) AfterBorrow(_iToken common.Address, _borrower common.Address, _borrowedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterBorrow(&_Controller.TransactOpts, _iToken, _borrower, _borrowedAmount)
}

// AfterFlashloan is a paid mutator transaction binding the contract method 0xefced385.
//
// Solidity: function afterFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactor) AfterFlashloan(opts *bind.TransactOpts, _iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterFlashloan", _iToken, _to, _amount)
}

// AfterFlashloan is a paid mutator transaction binding the contract method 0xefced385.
//
// Solidity: function afterFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerSession) AfterFlashloan(_iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterFlashloan(&_Controller.TransactOpts, _iToken, _to, _amount)
}

// AfterFlashloan is a paid mutator transaction binding the contract method 0xefced385.
//
// Solidity: function afterFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) AfterFlashloan(_iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterFlashloan(&_Controller.TransactOpts, _iToken, _to, _amount)
}

// AfterLiquidateBorrow is a paid mutator transaction binding the contract method 0x5f7bc006.
//
// Solidity: function afterLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repaidAmount, uint256 _seizedAmount) returns()
func (_Controller *ControllerTransactor) AfterLiquidateBorrow(opts *bind.TransactOpts, _iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repaidAmount *big.Int, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterLiquidateBorrow", _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repaidAmount, _seizedAmount)
}

// AfterLiquidateBorrow is a paid mutator transaction binding the contract method 0x5f7bc006.
//
// Solidity: function afterLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repaidAmount, uint256 _seizedAmount) returns()
func (_Controller *ControllerSession) AfterLiquidateBorrow(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repaidAmount *big.Int, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterLiquidateBorrow(&_Controller.TransactOpts, _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repaidAmount, _seizedAmount)
}

// AfterLiquidateBorrow is a paid mutator transaction binding the contract method 0x5f7bc006.
//
// Solidity: function afterLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repaidAmount, uint256 _seizedAmount) returns()
func (_Controller *ControllerTransactorSession) AfterLiquidateBorrow(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repaidAmount *big.Int, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterLiquidateBorrow(&_Controller.TransactOpts, _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repaidAmount, _seizedAmount)
}

// AfterMint is a paid mutator transaction binding the contract method 0xde65f41b.
//
// Solidity: function afterMint(address _iToken, address _minter, uint256 _mintAmount, uint256 _mintedAmount) returns()
func (_Controller *ControllerTransactor) AfterMint(opts *bind.TransactOpts, _iToken common.Address, _minter common.Address, _mintAmount *big.Int, _mintedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterMint", _iToken, _minter, _mintAmount, _mintedAmount)
}

// AfterMint is a paid mutator transaction binding the contract method 0xde65f41b.
//
// Solidity: function afterMint(address _iToken, address _minter, uint256 _mintAmount, uint256 _mintedAmount) returns()
func (_Controller *ControllerSession) AfterMint(_iToken common.Address, _minter common.Address, _mintAmount *big.Int, _mintedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterMint(&_Controller.TransactOpts, _iToken, _minter, _mintAmount, _mintedAmount)
}

// AfterMint is a paid mutator transaction binding the contract method 0xde65f41b.
//
// Solidity: function afterMint(address _iToken, address _minter, uint256 _mintAmount, uint256 _mintedAmount) returns()
func (_Controller *ControllerTransactorSession) AfterMint(_iToken common.Address, _minter common.Address, _mintAmount *big.Int, _mintedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterMint(&_Controller.TransactOpts, _iToken, _minter, _mintAmount, _mintedAmount)
}

// AfterRedeem is a paid mutator transaction binding the contract method 0xa2ddeb85.
//
// Solidity: function afterRedeem(address _iToken, address _redeemer, uint256 _redeemAmount, uint256 _redeemedUnderlying) returns()
func (_Controller *ControllerTransactor) AfterRedeem(opts *bind.TransactOpts, _iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int, _redeemedUnderlying *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterRedeem", _iToken, _redeemer, _redeemAmount, _redeemedUnderlying)
}

// AfterRedeem is a paid mutator transaction binding the contract method 0xa2ddeb85.
//
// Solidity: function afterRedeem(address _iToken, address _redeemer, uint256 _redeemAmount, uint256 _redeemedUnderlying) returns()
func (_Controller *ControllerSession) AfterRedeem(_iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int, _redeemedUnderlying *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterRedeem(&_Controller.TransactOpts, _iToken, _redeemer, _redeemAmount, _redeemedUnderlying)
}

// AfterRedeem is a paid mutator transaction binding the contract method 0xa2ddeb85.
//
// Solidity: function afterRedeem(address _iToken, address _redeemer, uint256 _redeemAmount, uint256 _redeemedUnderlying) returns()
func (_Controller *ControllerTransactorSession) AfterRedeem(_iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int, _redeemedUnderlying *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterRedeem(&_Controller.TransactOpts, _iToken, _redeemer, _redeemAmount, _redeemedUnderlying)
}

// AfterRepayBorrow is a paid mutator transaction binding the contract method 0xf079420a.
//
// Solidity: function afterRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactor) AfterRepayBorrow(opts *bind.TransactOpts, _iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterRepayBorrow", _iToken, _payer, _borrower, _repayAmount)
}

// AfterRepayBorrow is a paid mutator transaction binding the contract method 0xf079420a.
//
// Solidity: function afterRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerSession) AfterRepayBorrow(_iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterRepayBorrow(&_Controller.TransactOpts, _iToken, _payer, _borrower, _repayAmount)
}

// AfterRepayBorrow is a paid mutator transaction binding the contract method 0xf079420a.
//
// Solidity: function afterRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactorSession) AfterRepayBorrow(_iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterRepayBorrow(&_Controller.TransactOpts, _iToken, _payer, _borrower, _repayAmount)
}

// AfterSeize is a paid mutator transaction binding the contract method 0x0d7f67a2.
//
// Solidity: function afterSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizedAmount) returns()
func (_Controller *ControllerTransactor) AfterSeize(opts *bind.TransactOpts, _iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterSeize", _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizedAmount)
}

// AfterSeize is a paid mutator transaction binding the contract method 0x0d7f67a2.
//
// Solidity: function afterSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizedAmount) returns()
func (_Controller *ControllerSession) AfterSeize(_iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterSeize(&_Controller.TransactOpts, _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizedAmount)
}

// AfterSeize is a paid mutator transaction binding the contract method 0x0d7f67a2.
//
// Solidity: function afterSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizedAmount) returns()
func (_Controller *ControllerTransactorSession) AfterSeize(_iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizedAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterSeize(&_Controller.TransactOpts, _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizedAmount)
}

// AfterTransfer is a paid mutator transaction binding the contract method 0xe33af48a.
//
// Solidity: function afterTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactor) AfterTransfer(opts *bind.TransactOpts, _iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "afterTransfer", _iToken, _from, _to, _amount)
}

// AfterTransfer is a paid mutator transaction binding the contract method 0xe33af48a.
//
// Solidity: function afterTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerSession) AfterTransfer(_iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterTransfer(&_Controller.TransactOpts, _iToken, _from, _to, _amount)
}

// AfterTransfer is a paid mutator transaction binding the contract method 0xe33af48a.
//
// Solidity: function afterTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) AfterTransfer(_iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.AfterTransfer(&_Controller.TransactOpts, _iToken, _from, _to, _amount)
}

// BeforeBorrow is a paid mutator transaction binding the contract method 0x2b9553ce.
//
// Solidity: function beforeBorrow(address _iToken, address _borrower, uint256 _borrowAmount) returns()
func (_Controller *ControllerTransactor) BeforeBorrow(opts *bind.TransactOpts, _iToken common.Address, _borrower common.Address, _borrowAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeBorrow", _iToken, _borrower, _borrowAmount)
}

// BeforeBorrow is a paid mutator transaction binding the contract method 0x2b9553ce.
//
// Solidity: function beforeBorrow(address _iToken, address _borrower, uint256 _borrowAmount) returns()
func (_Controller *ControllerSession) BeforeBorrow(_iToken common.Address, _borrower common.Address, _borrowAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeBorrow(&_Controller.TransactOpts, _iToken, _borrower, _borrowAmount)
}

// BeforeBorrow is a paid mutator transaction binding the contract method 0x2b9553ce.
//
// Solidity: function beforeBorrow(address _iToken, address _borrower, uint256 _borrowAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeBorrow(_iToken common.Address, _borrower common.Address, _borrowAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeBorrow(&_Controller.TransactOpts, _iToken, _borrower, _borrowAmount)
}

// BeforeFlashloan is a paid mutator transaction binding the contract method 0xa8f4e9f6.
//
// Solidity: function beforeFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactor) BeforeFlashloan(opts *bind.TransactOpts, _iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeFlashloan", _iToken, _to, _amount)
}

// BeforeFlashloan is a paid mutator transaction binding the contract method 0xa8f4e9f6.
//
// Solidity: function beforeFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerSession) BeforeFlashloan(_iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeFlashloan(&_Controller.TransactOpts, _iToken, _to, _amount)
}

// BeforeFlashloan is a paid mutator transaction binding the contract method 0xa8f4e9f6.
//
// Solidity: function beforeFlashloan(address _iToken, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) BeforeFlashloan(_iToken common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeFlashloan(&_Controller.TransactOpts, _iToken, _to, _amount)
}

// BeforeLiquidateBorrow is a paid mutator transaction binding the contract method 0x061682de.
//
// Solidity: function beforeLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactor) BeforeLiquidateBorrow(opts *bind.TransactOpts, _iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeLiquidateBorrow", _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repayAmount)
}

// BeforeLiquidateBorrow is a paid mutator transaction binding the contract method 0x061682de.
//
// Solidity: function beforeLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerSession) BeforeLiquidateBorrow(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeLiquidateBorrow(&_Controller.TransactOpts, _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repayAmount)
}

// BeforeLiquidateBorrow is a paid mutator transaction binding the contract method 0x061682de.
//
// Solidity: function beforeLiquidateBorrow(address _iTokenBorrowed, address _iTokenCollateral, address _liquidator, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeLiquidateBorrow(_iTokenBorrowed common.Address, _iTokenCollateral common.Address, _liquidator common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeLiquidateBorrow(&_Controller.TransactOpts, _iTokenBorrowed, _iTokenCollateral, _liquidator, _borrower, _repayAmount)
}

// BeforeMint is a paid mutator transaction binding the contract method 0x86b84187.
//
// Solidity: function beforeMint(address _iToken, address _minter, uint256 _mintAmount) returns()
func (_Controller *ControllerTransactor) BeforeMint(opts *bind.TransactOpts, _iToken common.Address, _minter common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeMint", _iToken, _minter, _mintAmount)
}

// BeforeMint is a paid mutator transaction binding the contract method 0x86b84187.
//
// Solidity: function beforeMint(address _iToken, address _minter, uint256 _mintAmount) returns()
func (_Controller *ControllerSession) BeforeMint(_iToken common.Address, _minter common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeMint(&_Controller.TransactOpts, _iToken, _minter, _mintAmount)
}

// BeforeMint is a paid mutator transaction binding the contract method 0x86b84187.
//
// Solidity: function beforeMint(address _iToken, address _minter, uint256 _mintAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeMint(_iToken common.Address, _minter common.Address, _mintAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeMint(&_Controller.TransactOpts, _iToken, _minter, _mintAmount)
}

// BeforeRedeem is a paid mutator transaction binding the contract method 0xaf505ad9.
//
// Solidity: function beforeRedeem(address _iToken, address _redeemer, uint256 _redeemAmount) returns()
func (_Controller *ControllerTransactor) BeforeRedeem(opts *bind.TransactOpts, _iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeRedeem", _iToken, _redeemer, _redeemAmount)
}

// BeforeRedeem is a paid mutator transaction binding the contract method 0xaf505ad9.
//
// Solidity: function beforeRedeem(address _iToken, address _redeemer, uint256 _redeemAmount) returns()
func (_Controller *ControllerSession) BeforeRedeem(_iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeRedeem(&_Controller.TransactOpts, _iToken, _redeemer, _redeemAmount)
}

// BeforeRedeem is a paid mutator transaction binding the contract method 0xaf505ad9.
//
// Solidity: function beforeRedeem(address _iToken, address _redeemer, uint256 _redeemAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeRedeem(_iToken common.Address, _redeemer common.Address, _redeemAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeRedeem(&_Controller.TransactOpts, _iToken, _redeemer, _redeemAmount)
}

// BeforeRepayBorrow is a paid mutator transaction binding the contract method 0x1637eefd.
//
// Solidity: function beforeRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactor) BeforeRepayBorrow(opts *bind.TransactOpts, _iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeRepayBorrow", _iToken, _payer, _borrower, _repayAmount)
}

// BeforeRepayBorrow is a paid mutator transaction binding the contract method 0x1637eefd.
//
// Solidity: function beforeRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerSession) BeforeRepayBorrow(_iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeRepayBorrow(&_Controller.TransactOpts, _iToken, _payer, _borrower, _repayAmount)
}

// BeforeRepayBorrow is a paid mutator transaction binding the contract method 0x1637eefd.
//
// Solidity: function beforeRepayBorrow(address _iToken, address _payer, address _borrower, uint256 _repayAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeRepayBorrow(_iToken common.Address, _payer common.Address, _borrower common.Address, _repayAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeRepayBorrow(&_Controller.TransactOpts, _iToken, _payer, _borrower, _repayAmount)
}

// BeforeSeize is a paid mutator transaction binding the contract method 0xcb6db19c.
//
// Solidity: function beforeSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizeAmount) returns()
func (_Controller *ControllerTransactor) BeforeSeize(opts *bind.TransactOpts, _iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizeAmount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeSeize", _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizeAmount)
}

// BeforeSeize is a paid mutator transaction binding the contract method 0xcb6db19c.
//
// Solidity: function beforeSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizeAmount) returns()
func (_Controller *ControllerSession) BeforeSeize(_iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizeAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeSeize(&_Controller.TransactOpts, _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizeAmount)
}

// BeforeSeize is a paid mutator transaction binding the contract method 0xcb6db19c.
//
// Solidity: function beforeSeize(address _iTokenCollateral, address _iTokenBorrowed, address _liquidator, address _borrower, uint256 _seizeAmount) returns()
func (_Controller *ControllerTransactorSession) BeforeSeize(_iTokenCollateral common.Address, _iTokenBorrowed common.Address, _liquidator common.Address, _borrower common.Address, _seizeAmount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeSeize(&_Controller.TransactOpts, _iTokenCollateral, _iTokenBorrowed, _liquidator, _borrower, _seizeAmount)
}

// BeforeTransfer is a paid mutator transaction binding the contract method 0x72b4304a.
//
// Solidity: function beforeTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactor) BeforeTransfer(opts *bind.TransactOpts, _iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "beforeTransfer", _iToken, _from, _to, _amount)
}

// BeforeTransfer is a paid mutator transaction binding the contract method 0x72b4304a.
//
// Solidity: function beforeTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerSession) BeforeTransfer(_iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeTransfer(&_Controller.TransactOpts, _iToken, _from, _to, _amount)
}

// BeforeTransfer is a paid mutator transaction binding the contract method 0x72b4304a.
//
// Solidity: function beforeTransfer(address _iToken, address _from, address _to, uint256 _amount) returns()
func (_Controller *ControllerTransactorSession) BeforeTransfer(_iToken common.Address, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BeforeTransfer(&_Controller.TransactOpts, _iToken, _from, _to, _amount)
}

// EnterMarketFromiToken is a paid mutator transaction binding the contract method 0xd8f5a572.
//
// Solidity: function enterMarketFromiToken(address _market, address _account) returns()
func (_Controller *ControllerTransactor) EnterMarketFromiToken(opts *bind.TransactOpts, _market common.Address, _account common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "enterMarketFromiToken", _market, _account)
}

// EnterMarketFromiToken is a paid mutator transaction binding the contract method 0xd8f5a572.
//
// Solidity: function enterMarketFromiToken(address _market, address _account) returns()
func (_Controller *ControllerSession) EnterMarketFromiToken(_market common.Address, _account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnterMarketFromiToken(&_Controller.TransactOpts, _market, _account)
}

// EnterMarketFromiToken is a paid mutator transaction binding the contract method 0xd8f5a572.
//
// Solidity: function enterMarketFromiToken(address _market, address _account) returns()
func (_Controller *ControllerTransactorSession) EnterMarketFromiToken(_market common.Address, _account common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnterMarketFromiToken(&_Controller.TransactOpts, _market, _account)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerTransactor) EnterMarkets(opts *bind.TransactOpts, _iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "enterMarkets", _iTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerSession) EnterMarkets(_iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnterMarkets(&_Controller.TransactOpts, _iTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerTransactorSession) EnterMarkets(_iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.EnterMarkets(&_Controller.TransactOpts, _iTokens)
}

// ExitMarkets is a paid mutator transaction binding the contract method 0xd0af4138.
//
// Solidity: function exitMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerTransactor) ExitMarkets(opts *bind.TransactOpts, _iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "exitMarkets", _iTokens)
}

// ExitMarkets is a paid mutator transaction binding the contract method 0xd0af4138.
//
// Solidity: function exitMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerSession) ExitMarkets(_iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ExitMarkets(&_Controller.TransactOpts, _iTokens)
}

// ExitMarkets is a paid mutator transaction binding the contract method 0xd0af4138.
//
// Solidity: function exitMarkets(address[] _iTokens) returns(bool[] _results)
func (_Controller *ControllerTransactorSession) ExitMarkets(_iTokens []common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ExitMarkets(&_Controller.TransactOpts, _iTokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerSession) Initialize() (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Controller *ControllerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Controller.Contract.Initialize(&_Controller.TransactOpts)
}

// ControllerBorrowPausedIterator is returned from FilterBorrowPaused and is used to iterate over the raw logs and unpacked data for BorrowPaused events raised by the Controller contract.
type ControllerBorrowPausedIterator struct {
	Event *ControllerBorrowPaused // Event containing the contract specifics and raw log

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
func (it *ControllerBorrowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerBorrowPaused)
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
		it.Event = new(ControllerBorrowPaused)
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
func (it *ControllerBorrowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerBorrowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerBorrowPaused represents a BorrowPaused event raised by the Controller contract.
type ControllerBorrowPaused struct {
	IToken common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrowPaused is a free log retrieval operation binding the contract event 0xa501bd5ac3de9924ce0c13576750267130fd835780f1ec6e1ae9fb13ee746503.
//
// Solidity: event BorrowPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) FilterBorrowPaused(opts *bind.FilterOpts) (*ControllerBorrowPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "BorrowPaused")
	if err != nil {
		return nil, err
	}
	return &ControllerBorrowPausedIterator{contract: _Controller.contract, event: "BorrowPaused", logs: logs, sub: sub}, nil
}

// WatchBorrowPaused is a free log subscription operation binding the contract event 0xa501bd5ac3de9924ce0c13576750267130fd835780f1ec6e1ae9fb13ee746503.
//
// Solidity: event BorrowPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) WatchBorrowPaused(opts *bind.WatchOpts, sink chan<- *ControllerBorrowPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "BorrowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerBorrowPaused)
				if err := _Controller.contract.UnpackLog(event, "BorrowPaused", log); err != nil {
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

// ParseBorrowPaused is a log parse operation binding the contract event 0xa501bd5ac3de9924ce0c13576750267130fd835780f1ec6e1ae9fb13ee746503.
//
// Solidity: event BorrowPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) ParseBorrowPaused(log types.Log) (*ControllerBorrowPaused, error) {
	event := new(ControllerBorrowPaused)
	if err := _Controller.contract.UnpackLog(event, "BorrowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerBorrowedAddedIterator is returned from FilterBorrowedAdded and is used to iterate over the raw logs and unpacked data for BorrowedAdded events raised by the Controller contract.
type ControllerBorrowedAddedIterator struct {
	Event *ControllerBorrowedAdded // Event containing the contract specifics and raw log

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
func (it *ControllerBorrowedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerBorrowedAdded)
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
		it.Event = new(ControllerBorrowedAdded)
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
func (it *ControllerBorrowedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerBorrowedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerBorrowedAdded represents a BorrowedAdded event raised by the Controller contract.
type ControllerBorrowedAdded struct {
	IToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrowedAdded is a free log retrieval operation binding the contract event 0x61400a0050615e96f6a4a234e4fc162ab71775345f32f4641cd884601603a455.
//
// Solidity: event BorrowedAdded(address iToken, address account)
func (_Controller *ControllerFilterer) FilterBorrowedAdded(opts *bind.FilterOpts) (*ControllerBorrowedAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "BorrowedAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerBorrowedAddedIterator{contract: _Controller.contract, event: "BorrowedAdded", logs: logs, sub: sub}, nil
}

// WatchBorrowedAdded is a free log subscription operation binding the contract event 0x61400a0050615e96f6a4a234e4fc162ab71775345f32f4641cd884601603a455.
//
// Solidity: event BorrowedAdded(address iToken, address account)
func (_Controller *ControllerFilterer) WatchBorrowedAdded(opts *bind.WatchOpts, sink chan<- *ControllerBorrowedAdded) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "BorrowedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerBorrowedAdded)
				if err := _Controller.contract.UnpackLog(event, "BorrowedAdded", log); err != nil {
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

// ParseBorrowedAdded is a log parse operation binding the contract event 0x61400a0050615e96f6a4a234e4fc162ab71775345f32f4641cd884601603a455.
//
// Solidity: event BorrowedAdded(address iToken, address account)
func (_Controller *ControllerFilterer) ParseBorrowedAdded(log types.Log) (*ControllerBorrowedAdded, error) {
	event := new(ControllerBorrowedAdded)
	if err := _Controller.contract.UnpackLog(event, "BorrowedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerBorrowedRemovedIterator is returned from FilterBorrowedRemoved and is used to iterate over the raw logs and unpacked data for BorrowedRemoved events raised by the Controller contract.
type ControllerBorrowedRemovedIterator struct {
	Event *ControllerBorrowedRemoved // Event containing the contract specifics and raw log

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
func (it *ControllerBorrowedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerBorrowedRemoved)
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
		it.Event = new(ControllerBorrowedRemoved)
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
func (it *ControllerBorrowedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerBorrowedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerBorrowedRemoved represents a BorrowedRemoved event raised by the Controller contract.
type ControllerBorrowedRemoved struct {
	IToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrowedRemoved is a free log retrieval operation binding the contract event 0x6df0340c4ae9990f4b8f3c3e0e5408dcf6ba2c086b92966dfdbbef1f2f052e1a.
//
// Solidity: event BorrowedRemoved(address iToken, address account)
func (_Controller *ControllerFilterer) FilterBorrowedRemoved(opts *bind.FilterOpts) (*ControllerBorrowedRemovedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "BorrowedRemoved")
	if err != nil {
		return nil, err
	}
	return &ControllerBorrowedRemovedIterator{contract: _Controller.contract, event: "BorrowedRemoved", logs: logs, sub: sub}, nil
}

// WatchBorrowedRemoved is a free log subscription operation binding the contract event 0x6df0340c4ae9990f4b8f3c3e0e5408dcf6ba2c086b92966dfdbbef1f2f052e1a.
//
// Solidity: event BorrowedRemoved(address iToken, address account)
func (_Controller *ControllerFilterer) WatchBorrowedRemoved(opts *bind.WatchOpts, sink chan<- *ControllerBorrowedRemoved) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "BorrowedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerBorrowedRemoved)
				if err := _Controller.contract.UnpackLog(event, "BorrowedRemoved", log); err != nil {
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

// ParseBorrowedRemoved is a log parse operation binding the contract event 0x6df0340c4ae9990f4b8f3c3e0e5408dcf6ba2c086b92966dfdbbef1f2f052e1a.
//
// Solidity: event BorrowedRemoved(address iToken, address account)
func (_Controller *ControllerFilterer) ParseBorrowedRemoved(log types.Log) (*ControllerBorrowedRemoved, error) {
	event := new(ControllerBorrowedRemoved)
	if err := _Controller.contract.UnpackLog(event, "BorrowedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMarketAddedIterator is returned from FilterMarketAdded and is used to iterate over the raw logs and unpacked data for MarketAdded events raised by the Controller contract.
type ControllerMarketAddedIterator struct {
	Event *ControllerMarketAdded // Event containing the contract specifics and raw log

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
func (it *ControllerMarketAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMarketAdded)
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
		it.Event = new(ControllerMarketAdded)
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
func (it *ControllerMarketAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMarketAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMarketAdded represents a MarketAdded event raised by the Controller contract.
type ControllerMarketAdded struct {
	IToken             common.Address
	CollateralFactor   *big.Int
	BorrowFactor       *big.Int
	SupplyCapacity     *big.Int
	BorrowCapacity     *big.Int
	DistributionFactor *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketAdded is a free log retrieval operation binding the contract event 0x96f15e56d27d14b82ed736fd9387871d60f721cd59f324699722b4f0d11ee1b1.
//
// Solidity: event MarketAdded(address iToken, uint256 collateralFactor, uint256 borrowFactor, uint256 supplyCapacity, uint256 borrowCapacity, uint256 distributionFactor)
func (_Controller *ControllerFilterer) FilterMarketAdded(opts *bind.FilterOpts) (*ControllerMarketAddedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return &ControllerMarketAddedIterator{contract: _Controller.contract, event: "MarketAdded", logs: logs, sub: sub}, nil
}

// WatchMarketAdded is a free log subscription operation binding the contract event 0x96f15e56d27d14b82ed736fd9387871d60f721cd59f324699722b4f0d11ee1b1.
//
// Solidity: event MarketAdded(address iToken, uint256 collateralFactor, uint256 borrowFactor, uint256 supplyCapacity, uint256 borrowCapacity, uint256 distributionFactor)
func (_Controller *ControllerFilterer) WatchMarketAdded(opts *bind.WatchOpts, sink chan<- *ControllerMarketAdded) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MarketAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMarketAdded)
				if err := _Controller.contract.UnpackLog(event, "MarketAdded", log); err != nil {
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

// ParseMarketAdded is a log parse operation binding the contract event 0x96f15e56d27d14b82ed736fd9387871d60f721cd59f324699722b4f0d11ee1b1.
//
// Solidity: event MarketAdded(address iToken, uint256 collateralFactor, uint256 borrowFactor, uint256 supplyCapacity, uint256 borrowCapacity, uint256 distributionFactor)
func (_Controller *ControllerFilterer) ParseMarketAdded(log types.Log) (*ControllerMarketAdded, error) {
	event := new(ControllerMarketAdded)
	if err := _Controller.contract.UnpackLog(event, "MarketAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Controller contract.
type ControllerMarketEnteredIterator struct {
	Event *ControllerMarketEntered // Event containing the contract specifics and raw log

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
func (it *ControllerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMarketEntered)
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
		it.Event = new(ControllerMarketEntered)
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
func (it *ControllerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMarketEntered represents a MarketEntered event raised by the Controller contract.
type ControllerMarketEntered struct {
	IToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address iToken, address account)
func (_Controller *ControllerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*ControllerMarketEnteredIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ControllerMarketEnteredIterator{contract: _Controller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address iToken, address account)
func (_Controller *ControllerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ControllerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMarketEntered)
				if err := _Controller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address iToken, address account)
func (_Controller *ControllerFilterer) ParseMarketEntered(log types.Log) (*ControllerMarketEntered, error) {
	event := new(ControllerMarketEntered)
	if err := _Controller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Controller contract.
type ControllerMarketExitedIterator struct {
	Event *ControllerMarketExited // Event containing the contract specifics and raw log

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
func (it *ControllerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMarketExited)
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
		it.Event = new(ControllerMarketExited)
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
func (it *ControllerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMarketExited represents a MarketExited event raised by the Controller contract.
type ControllerMarketExited struct {
	IToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address iToken, address account)
func (_Controller *ControllerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*ControllerMarketExitedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ControllerMarketExitedIterator{contract: _Controller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address iToken, address account)
func (_Controller *ControllerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ControllerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMarketExited)
				if err := _Controller.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address iToken, address account)
func (_Controller *ControllerFilterer) ParseMarketExited(log types.Log) (*ControllerMarketExited, error) {
	event := new(ControllerMarketExited)
	if err := _Controller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerMintPausedIterator is returned from FilterMintPaused and is used to iterate over the raw logs and unpacked data for MintPaused events raised by the Controller contract.
type ControllerMintPausedIterator struct {
	Event *ControllerMintPaused // Event containing the contract specifics and raw log

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
func (it *ControllerMintPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerMintPaused)
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
		it.Event = new(ControllerMintPaused)
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
func (it *ControllerMintPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerMintPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerMintPaused represents a MintPaused event raised by the Controller contract.
type ControllerMintPaused struct {
	IToken common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMintPaused is a free log retrieval operation binding the contract event 0x48cff1fdfe035051f817c77855362a92f96f941342d41008f2d596f530c4e96b.
//
// Solidity: event MintPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) FilterMintPaused(opts *bind.FilterOpts) (*ControllerMintPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "MintPaused")
	if err != nil {
		return nil, err
	}
	return &ControllerMintPausedIterator{contract: _Controller.contract, event: "MintPaused", logs: logs, sub: sub}, nil
}

// WatchMintPaused is a free log subscription operation binding the contract event 0x48cff1fdfe035051f817c77855362a92f96f941342d41008f2d596f530c4e96b.
//
// Solidity: event MintPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) WatchMintPaused(opts *bind.WatchOpts, sink chan<- *ControllerMintPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "MintPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerMintPaused)
				if err := _Controller.contract.UnpackLog(event, "MintPaused", log); err != nil {
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

// ParseMintPaused is a log parse operation binding the contract event 0x48cff1fdfe035051f817c77855362a92f96f941342d41008f2d596f530c4e96b.
//
// Solidity: event MintPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) ParseMintPaused(log types.Log) (*ControllerMintPaused, error) {
	event := new(ControllerMintPaused)
	if err := _Controller.contract.UnpackLog(event, "MintPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewBorrowCapacityIterator is returned from FilterNewBorrowCapacity and is used to iterate over the raw logs and unpacked data for NewBorrowCapacity events raised by the Controller contract.
type ControllerNewBorrowCapacityIterator struct {
	Event *ControllerNewBorrowCapacity // Event containing the contract specifics and raw log

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
func (it *ControllerNewBorrowCapacityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewBorrowCapacity)
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
		it.Event = new(ControllerNewBorrowCapacity)
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
func (it *ControllerNewBorrowCapacityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewBorrowCapacityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewBorrowCapacity represents a NewBorrowCapacity event raised by the Controller contract.
type ControllerNewBorrowCapacity struct {
	IToken            common.Address
	OldBorrowCapacity *big.Int
	NewBorrowCapacity *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapacity is a free log retrieval operation binding the contract event 0x6c07964980a680dc12652ccab077a57c65740bf2ebc204a7670438051ee6fa5b.
//
// Solidity: event NewBorrowCapacity(address iToken, uint256 oldBorrowCapacity, uint256 newBorrowCapacity)
func (_Controller *ControllerFilterer) FilterNewBorrowCapacity(opts *bind.FilterOpts) (*ControllerNewBorrowCapacityIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewBorrowCapacity")
	if err != nil {
		return nil, err
	}
	return &ControllerNewBorrowCapacityIterator{contract: _Controller.contract, event: "NewBorrowCapacity", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapacity is a free log subscription operation binding the contract event 0x6c07964980a680dc12652ccab077a57c65740bf2ebc204a7670438051ee6fa5b.
//
// Solidity: event NewBorrowCapacity(address iToken, uint256 oldBorrowCapacity, uint256 newBorrowCapacity)
func (_Controller *ControllerFilterer) WatchNewBorrowCapacity(opts *bind.WatchOpts, sink chan<- *ControllerNewBorrowCapacity) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewBorrowCapacity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewBorrowCapacity)
				if err := _Controller.contract.UnpackLog(event, "NewBorrowCapacity", log); err != nil {
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

// ParseNewBorrowCapacity is a log parse operation binding the contract event 0x6c07964980a680dc12652ccab077a57c65740bf2ebc204a7670438051ee6fa5b.
//
// Solidity: event NewBorrowCapacity(address iToken, uint256 oldBorrowCapacity, uint256 newBorrowCapacity)
func (_Controller *ControllerFilterer) ParseNewBorrowCapacity(log types.Log) (*ControllerNewBorrowCapacity, error) {
	event := new(ControllerNewBorrowCapacity)
	if err := _Controller.contract.UnpackLog(event, "NewBorrowCapacity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewBorrowFactorIterator is returned from FilterNewBorrowFactor and is used to iterate over the raw logs and unpacked data for NewBorrowFactor events raised by the Controller contract.
type ControllerNewBorrowFactorIterator struct {
	Event *ControllerNewBorrowFactor // Event containing the contract specifics and raw log

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
func (it *ControllerNewBorrowFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewBorrowFactor)
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
		it.Event = new(ControllerNewBorrowFactor)
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
func (it *ControllerNewBorrowFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewBorrowFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewBorrowFactor represents a NewBorrowFactor event raised by the Controller contract.
type ControllerNewBorrowFactor struct {
	IToken                  common.Address
	OldBorrowFactorMantissa *big.Int
	NewBorrowFactorMantissa *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowFactor is a free log retrieval operation binding the contract event 0xd69d3c48186e96a0af0545cf26e970b49ed3e29a306a37118a5a787bae6b90be.
//
// Solidity: event NewBorrowFactor(address iToken, uint256 oldBorrowFactorMantissa, uint256 newBorrowFactorMantissa)
func (_Controller *ControllerFilterer) FilterNewBorrowFactor(opts *bind.FilterOpts) (*ControllerNewBorrowFactorIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewBorrowFactor")
	if err != nil {
		return nil, err
	}
	return &ControllerNewBorrowFactorIterator{contract: _Controller.contract, event: "NewBorrowFactor", logs: logs, sub: sub}, nil
}

// WatchNewBorrowFactor is a free log subscription operation binding the contract event 0xd69d3c48186e96a0af0545cf26e970b49ed3e29a306a37118a5a787bae6b90be.
//
// Solidity: event NewBorrowFactor(address iToken, uint256 oldBorrowFactorMantissa, uint256 newBorrowFactorMantissa)
func (_Controller *ControllerFilterer) WatchNewBorrowFactor(opts *bind.WatchOpts, sink chan<- *ControllerNewBorrowFactor) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewBorrowFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewBorrowFactor)
				if err := _Controller.contract.UnpackLog(event, "NewBorrowFactor", log); err != nil {
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

// ParseNewBorrowFactor is a log parse operation binding the contract event 0xd69d3c48186e96a0af0545cf26e970b49ed3e29a306a37118a5a787bae6b90be.
//
// Solidity: event NewBorrowFactor(address iToken, uint256 oldBorrowFactorMantissa, uint256 newBorrowFactorMantissa)
func (_Controller *ControllerFilterer) ParseNewBorrowFactor(log types.Log) (*ControllerNewBorrowFactor, error) {
	event := new(ControllerNewBorrowFactor)
	if err := _Controller.contract.UnpackLog(event, "NewBorrowFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Controller contract.
type ControllerNewCloseFactorIterator struct {
	Event *ControllerNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *ControllerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewCloseFactor)
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
		it.Event = new(ControllerNewCloseFactor)
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
func (it *ControllerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewCloseFactor represents a NewCloseFactor event raised by the Controller contract.
type ControllerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Controller *ControllerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ControllerNewCloseFactorIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ControllerNewCloseFactorIterator{contract: _Controller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Controller *ControllerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ControllerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewCloseFactor)
				if err := _Controller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Controller *ControllerFilterer) ParseNewCloseFactor(log types.Log) (*ControllerNewCloseFactor, error) {
	event := new(ControllerNewCloseFactor)
	if err := _Controller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Controller contract.
type ControllerNewCollateralFactorIterator struct {
	Event *ControllerNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *ControllerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewCollateralFactor)
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
		it.Event = new(ControllerNewCollateralFactor)
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
func (it *ControllerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewCollateralFactor represents a NewCollateralFactor event raised by the Controller contract.
type ControllerNewCollateralFactor struct {
	IToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address iToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Controller *ControllerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ControllerNewCollateralFactorIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ControllerNewCollateralFactorIterator{contract: _Controller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address iToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Controller *ControllerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ControllerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewCollateralFactor)
				if err := _Controller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address iToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Controller *ControllerFilterer) ParseNewCollateralFactor(log types.Log) (*ControllerNewCollateralFactor, error) {
	event := new(ControllerNewCollateralFactor)
	if err := _Controller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Controller contract.
type ControllerNewLiquidationIncentiveIterator struct {
	Event *ControllerNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *ControllerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewLiquidationIncentive)
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
		it.Event = new(ControllerNewLiquidationIncentive)
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
func (it *ControllerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Controller contract.
type ControllerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Controller *ControllerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ControllerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ControllerNewLiquidationIncentiveIterator{contract: _Controller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Controller *ControllerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ControllerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewLiquidationIncentive)
				if err := _Controller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Controller *ControllerFilterer) ParseNewLiquidationIncentive(log types.Log) (*ControllerNewLiquidationIncentive, error) {
	event := new(ControllerNewLiquidationIncentive)
	if err := _Controller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Controller contract.
type ControllerNewOwnerIterator struct {
	Event *ControllerNewOwner // Event containing the contract specifics and raw log

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
func (it *ControllerNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewOwner)
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
		it.Event = new(ControllerNewOwner)
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
func (it *ControllerNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewOwner represents a NewOwner event raised by the Controller contract.
type ControllerNewOwner struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) FilterNewOwner(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerNewOwnerIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerNewOwnerIterator{contract: _Controller.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x70aea8d848e8a90fb7661b227dc522eb6395c3dac71b63cb59edd5c9899b2364.
//
// Solidity: event NewOwner(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ControllerNewOwner, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewOwner", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewOwner)
				if err := _Controller.contract.UnpackLog(event, "NewOwner", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseNewOwner(log types.Log) (*ControllerNewOwner, error) {
	event := new(ControllerNewOwner)
	if err := _Controller.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Controller contract.
type ControllerNewPauseGuardianIterator struct {
	Event *ControllerNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *ControllerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewPauseGuardian)
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
		it.Event = new(ControllerNewPauseGuardian)
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
func (it *ControllerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewPauseGuardian represents a NewPauseGuardian event raised by the Controller contract.
type ControllerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Controller *ControllerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ControllerNewPauseGuardianIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ControllerNewPauseGuardianIterator{contract: _Controller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Controller *ControllerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ControllerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewPauseGuardian)
				if err := _Controller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Controller *ControllerFilterer) ParseNewPauseGuardian(log types.Log) (*ControllerNewPauseGuardian, error) {
	event := new(ControllerNewPauseGuardian)
	if err := _Controller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewPendingOwnerIterator is returned from FilterNewPendingOwner and is used to iterate over the raw logs and unpacked data for NewPendingOwner events raised by the Controller contract.
type ControllerNewPendingOwnerIterator struct {
	Event *ControllerNewPendingOwner // Event containing the contract specifics and raw log

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
func (it *ControllerNewPendingOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewPendingOwner)
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
		it.Event = new(ControllerNewPendingOwner)
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
func (it *ControllerNewPendingOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewPendingOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewPendingOwner represents a NewPendingOwner event raised by the Controller contract.
type ControllerNewPendingOwner struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingOwner is a free log retrieval operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Controller *ControllerFilterer) FilterNewPendingOwner(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*ControllerNewPendingOwnerIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerNewPendingOwnerIterator{contract: _Controller.contract, event: "NewPendingOwner", logs: logs, sub: sub}, nil
}

// WatchNewPendingOwner is a free log subscription operation binding the contract event 0xb3d55174552271a4f1aaf36b72f50381e892171636b3fb5447fe00e995e7a37b.
//
// Solidity: event NewPendingOwner(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Controller *ControllerFilterer) WatchNewPendingOwner(opts *bind.WatchOpts, sink chan<- *ControllerNewPendingOwner, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewPendingOwner", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewPendingOwner)
				if err := _Controller.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseNewPendingOwner(log types.Log) (*ControllerNewPendingOwner, error) {
	event := new(ControllerNewPendingOwner)
	if err := _Controller.contract.UnpackLog(event, "NewPendingOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Controller contract.
type ControllerNewPriceOracleIterator struct {
	Event *ControllerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ControllerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewPriceOracle)
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
		it.Event = new(ControllerNewPriceOracle)
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
func (it *ControllerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewPriceOracle represents a NewPriceOracle event raised by the Controller contract.
type ControllerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Controller *ControllerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ControllerNewPriceOracleIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ControllerNewPriceOracleIterator{contract: _Controller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Controller *ControllerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ControllerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewPriceOracle)
				if err := _Controller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Controller *ControllerFilterer) ParseNewPriceOracle(log types.Log) (*ControllerNewPriceOracle, error) {
	event := new(ControllerNewPriceOracle)
	if err := _Controller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewRewardDistributorIterator is returned from FilterNewRewardDistributor and is used to iterate over the raw logs and unpacked data for NewRewardDistributor events raised by the Controller contract.
type ControllerNewRewardDistributorIterator struct {
	Event *ControllerNewRewardDistributor // Event containing the contract specifics and raw log

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
func (it *ControllerNewRewardDistributorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewRewardDistributor)
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
		it.Event = new(ControllerNewRewardDistributor)
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
func (it *ControllerNewRewardDistributorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewRewardDistributorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewRewardDistributor represents a NewRewardDistributor event raised by the Controller contract.
type ControllerNewRewardDistributor struct {
	OldRewardDistributor common.Address
	NewRewardDistributor common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewRewardDistributor is a free log retrieval operation binding the contract event 0x8ddca872a7a62d68235cff1a03badc845dc3007cfaa6145379f7bf3452ecb9b9.
//
// Solidity: event NewRewardDistributor(address oldRewardDistributor, address _newRewardDistributor)
func (_Controller *ControllerFilterer) FilterNewRewardDistributor(opts *bind.FilterOpts) (*ControllerNewRewardDistributorIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewRewardDistributor")
	if err != nil {
		return nil, err
	}
	return &ControllerNewRewardDistributorIterator{contract: _Controller.contract, event: "NewRewardDistributor", logs: logs, sub: sub}, nil
}

// WatchNewRewardDistributor is a free log subscription operation binding the contract event 0x8ddca872a7a62d68235cff1a03badc845dc3007cfaa6145379f7bf3452ecb9b9.
//
// Solidity: event NewRewardDistributor(address oldRewardDistributor, address _newRewardDistributor)
func (_Controller *ControllerFilterer) WatchNewRewardDistributor(opts *bind.WatchOpts, sink chan<- *ControllerNewRewardDistributor) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewRewardDistributor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewRewardDistributor)
				if err := _Controller.contract.UnpackLog(event, "NewRewardDistributor", log); err != nil {
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

// ParseNewRewardDistributor is a log parse operation binding the contract event 0x8ddca872a7a62d68235cff1a03badc845dc3007cfaa6145379f7bf3452ecb9b9.
//
// Solidity: event NewRewardDistributor(address oldRewardDistributor, address _newRewardDistributor)
func (_Controller *ControllerFilterer) ParseNewRewardDistributor(log types.Log) (*ControllerNewRewardDistributor, error) {
	event := new(ControllerNewRewardDistributor)
	if err := _Controller.contract.UnpackLog(event, "NewRewardDistributor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerNewSupplyCapacityIterator is returned from FilterNewSupplyCapacity and is used to iterate over the raw logs and unpacked data for NewSupplyCapacity events raised by the Controller contract.
type ControllerNewSupplyCapacityIterator struct {
	Event *ControllerNewSupplyCapacity // Event containing the contract specifics and raw log

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
func (it *ControllerNewSupplyCapacityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNewSupplyCapacity)
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
		it.Event = new(ControllerNewSupplyCapacity)
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
func (it *ControllerNewSupplyCapacityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNewSupplyCapacityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNewSupplyCapacity represents a NewSupplyCapacity event raised by the Controller contract.
type ControllerNewSupplyCapacity struct {
	IToken            common.Address
	OldSupplyCapacity *big.Int
	NewSupplyCapacity *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewSupplyCapacity is a free log retrieval operation binding the contract event 0x6ffc41eacf6bc473fab68880680a35214210358ff10d45c53cee339f1547f0bb.
//
// Solidity: event NewSupplyCapacity(address iToken, uint256 oldSupplyCapacity, uint256 newSupplyCapacity)
func (_Controller *ControllerFilterer) FilterNewSupplyCapacity(opts *bind.FilterOpts) (*ControllerNewSupplyCapacityIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NewSupplyCapacity")
	if err != nil {
		return nil, err
	}
	return &ControllerNewSupplyCapacityIterator{contract: _Controller.contract, event: "NewSupplyCapacity", logs: logs, sub: sub}, nil
}

// WatchNewSupplyCapacity is a free log subscription operation binding the contract event 0x6ffc41eacf6bc473fab68880680a35214210358ff10d45c53cee339f1547f0bb.
//
// Solidity: event NewSupplyCapacity(address iToken, uint256 oldSupplyCapacity, uint256 newSupplyCapacity)
func (_Controller *ControllerFilterer) WatchNewSupplyCapacity(opts *bind.WatchOpts, sink chan<- *ControllerNewSupplyCapacity) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NewSupplyCapacity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNewSupplyCapacity)
				if err := _Controller.contract.UnpackLog(event, "NewSupplyCapacity", log); err != nil {
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

// ParseNewSupplyCapacity is a log parse operation binding the contract event 0x6ffc41eacf6bc473fab68880680a35214210358ff10d45c53cee339f1547f0bb.
//
// Solidity: event NewSupplyCapacity(address iToken, uint256 oldSupplyCapacity, uint256 newSupplyCapacity)
func (_Controller *ControllerFilterer) ParseNewSupplyCapacity(log types.Log) (*ControllerNewSupplyCapacity, error) {
	event := new(ControllerNewSupplyCapacity)
	if err := _Controller.contract.UnpackLog(event, "NewSupplyCapacity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerRedeemPausedIterator is returned from FilterRedeemPaused and is used to iterate over the raw logs and unpacked data for RedeemPaused events raised by the Controller contract.
type ControllerRedeemPausedIterator struct {
	Event *ControllerRedeemPaused // Event containing the contract specifics and raw log

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
func (it *ControllerRedeemPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerRedeemPaused)
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
		it.Event = new(ControllerRedeemPaused)
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
func (it *ControllerRedeemPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerRedeemPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerRedeemPaused represents a RedeemPaused event raised by the Controller contract.
type ControllerRedeemPaused struct {
	IToken common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemPaused is a free log retrieval operation binding the contract event 0x3d95d87adcb06de16ac0e322ea97a06e30967fedcc83a6d1d0b9fda60300801a.
//
// Solidity: event RedeemPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) FilterRedeemPaused(opts *bind.FilterOpts) (*ControllerRedeemPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "RedeemPaused")
	if err != nil {
		return nil, err
	}
	return &ControllerRedeemPausedIterator{contract: _Controller.contract, event: "RedeemPaused", logs: logs, sub: sub}, nil
}

// WatchRedeemPaused is a free log subscription operation binding the contract event 0x3d95d87adcb06de16ac0e322ea97a06e30967fedcc83a6d1d0b9fda60300801a.
//
// Solidity: event RedeemPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) WatchRedeemPaused(opts *bind.WatchOpts, sink chan<- *ControllerRedeemPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "RedeemPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerRedeemPaused)
				if err := _Controller.contract.UnpackLog(event, "RedeemPaused", log); err != nil {
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

// ParseRedeemPaused is a log parse operation binding the contract event 0x3d95d87adcb06de16ac0e322ea97a06e30967fedcc83a6d1d0b9fda60300801a.
//
// Solidity: event RedeemPaused(address iToken, bool paused)
func (_Controller *ControllerFilterer) ParseRedeemPaused(log types.Log) (*ControllerRedeemPaused, error) {
	event := new(ControllerRedeemPaused)
	if err := _Controller.contract.UnpackLog(event, "RedeemPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerSeizePausedIterator is returned from FilterSeizePaused and is used to iterate over the raw logs and unpacked data for SeizePaused events raised by the Controller contract.
type ControllerSeizePausedIterator struct {
	Event *ControllerSeizePaused // Event containing the contract specifics and raw log

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
func (it *ControllerSeizePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerSeizePaused)
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
		it.Event = new(ControllerSeizePaused)
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
func (it *ControllerSeizePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerSeizePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerSeizePaused represents a SeizePaused event raised by the Controller contract.
type ControllerSeizePaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSeizePaused is a free log retrieval operation binding the contract event 0xcaf79886b466b41f609cc39fdc375ab5484cc3811570b7faa22fcb2b55516fdf.
//
// Solidity: event SeizePaused(bool paused)
func (_Controller *ControllerFilterer) FilterSeizePaused(opts *bind.FilterOpts) (*ControllerSeizePausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "SeizePaused")
	if err != nil {
		return nil, err
	}
	return &ControllerSeizePausedIterator{contract: _Controller.contract, event: "SeizePaused", logs: logs, sub: sub}, nil
}

// WatchSeizePaused is a free log subscription operation binding the contract event 0xcaf79886b466b41f609cc39fdc375ab5484cc3811570b7faa22fcb2b55516fdf.
//
// Solidity: event SeizePaused(bool paused)
func (_Controller *ControllerFilterer) WatchSeizePaused(opts *bind.WatchOpts, sink chan<- *ControllerSeizePaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "SeizePaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerSeizePaused)
				if err := _Controller.contract.UnpackLog(event, "SeizePaused", log); err != nil {
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

// ParseSeizePaused is a log parse operation binding the contract event 0xcaf79886b466b41f609cc39fdc375ab5484cc3811570b7faa22fcb2b55516fdf.
//
// Solidity: event SeizePaused(bool paused)
func (_Controller *ControllerFilterer) ParseSeizePaused(log types.Log) (*ControllerSeizePaused, error) {
	event := new(ControllerSeizePaused)
	if err := _Controller.contract.UnpackLog(event, "SeizePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerTransferPausedIterator is returned from FilterTransferPaused and is used to iterate over the raw logs and unpacked data for TransferPaused events raised by the Controller contract.
type ControllerTransferPausedIterator struct {
	Event *ControllerTransferPaused // Event containing the contract specifics and raw log

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
func (it *ControllerTransferPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerTransferPaused)
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
		it.Event = new(ControllerTransferPaused)
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
func (it *ControllerTransferPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerTransferPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerTransferPaused represents a TransferPaused event raised by the Controller contract.
type ControllerTransferPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferPaused is a free log retrieval operation binding the contract event 0x7b92c580e5b53e34f85950f4117143b86c4031ca9568a3cd30f4b3cd5bb95dc8.
//
// Solidity: event TransferPaused(bool paused)
func (_Controller *ControllerFilterer) FilterTransferPaused(opts *bind.FilterOpts) (*ControllerTransferPausedIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "TransferPaused")
	if err != nil {
		return nil, err
	}
	return &ControllerTransferPausedIterator{contract: _Controller.contract, event: "TransferPaused", logs: logs, sub: sub}, nil
}

// WatchTransferPaused is a free log subscription operation binding the contract event 0x7b92c580e5b53e34f85950f4117143b86c4031ca9568a3cd30f4b3cd5bb95dc8.
//
// Solidity: event TransferPaused(bool paused)
func (_Controller *ControllerFilterer) WatchTransferPaused(opts *bind.WatchOpts, sink chan<- *ControllerTransferPaused) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "TransferPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerTransferPaused)
				if err := _Controller.contract.UnpackLog(event, "TransferPaused", log); err != nil {
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

// ParseTransferPaused is a log parse operation binding the contract event 0x7b92c580e5b53e34f85950f4117143b86c4031ca9568a3cd30f4b3cd5bb95dc8.
//
// Solidity: event TransferPaused(bool paused)
func (_Controller *ControllerFilterer) ParseTransferPaused(log types.Log) (*ControllerTransferPaused, error) {
	event := new(ControllerTransferPaused)
	if err := _Controller.contract.UnpackLog(event, "TransferPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
