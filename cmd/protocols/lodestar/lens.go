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

// CompoundLensAccountLimits is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensAccountLimits struct {
	Markets   []common.Address
	Liquidity *big.Int
	Shortfall *big.Int
}

// CompoundLensCTokenBalances is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenBalances struct {
	CToken               common.Address
	BalanceOf            *big.Int
	BorrowBalanceCurrent *big.Int
	BalanceOfUnderlying  *big.Int
	TokenBalance         *big.Int
	TokenAllowance       *big.Int
}

// CompoundLensCTokenMetadata is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenMetadata struct {
	CToken                   common.Address
	ExchangeRateCurrent      *big.Int
	SupplyRatePerBlock       *big.Int
	BorrowRatePerBlock       *big.Int
	ReserveFactorMantissa    *big.Int
	TotalBorrows             *big.Int
	TotalReserves            *big.Int
	TotalSupply              *big.Int
	TotalCash                *big.Int
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	UnderlyingAssetAddress   common.Address
	CTokenDecimals           *big.Int
	UnderlyingDecimals       *big.Int
	CompSupplySpeed          *big.Int
	CompBorrowSpeed          *big.Int
	BorrowCap                *big.Int
}

// CompoundLensCTokenUnderlyingPrice is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCTokenUnderlyingPrice struct {
	CToken          common.Address
	UnderlyingPrice *big.Int
}

// CompoundLensCompBalanceMetadata is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompBalanceMetadata struct {
	Balance  *big.Int
	Votes    *big.Int
	Delegate common.Address
}

// CompoundLensCompBalanceMetadataExt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompBalanceMetadataExt struct {
	Balance   *big.Int
	Votes     *big.Int
	Delegate  common.Address
	Allocated *big.Int
}

// CompoundLensCompVotes is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensCompVotes struct {
	BlockNumber *big.Int
	Votes       *big.Int
}

// CompoundLensGovBravoProposal is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovBravoProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Eta          *big.Int
	Targets      []common.Address
	Values       []*big.Int
	Signatures   []string
	Calldatas    [][]byte
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	AbstainVotes *big.Int
	Canceled     bool
	Executed     bool
}

// CompoundLensGovBravoReceipt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovBravoReceipt struct {
	ProposalId *big.Int
	HasVoted   bool
	Support    uint8
	Votes      *big.Int
}

// CompoundLensGovProposal is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Eta          *big.Int
	Targets      []common.Address
	Values       []*big.Int
	Signatures   []string
	Calldatas    [][]byte
	StartBlock   *big.Int
	EndBlock     *big.Int
	ForVotes     *big.Int
	AgainstVotes *big.Int
	Canceled     bool
	Executed     bool
}

// CompoundLensGovReceipt is an auto generated low-level Go binding around an user-defined struct.
type CompoundLensGovReceipt struct {
	ProposalId *big.Int
	HasVoted   bool
	Support    bool
	Votes      *big.Int
}

// LensMetaData contains all meta data concerning the Lens contract.
var LensMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"}],\"name\":\"lodeMetaData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cTokenBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenBalances\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cTokenBalancesAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalanceCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAllowance\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenBalances[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"cTokenMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compSupplySpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compBorrowSpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"cTokenMetadataAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRateCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAssetAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cTokenDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compSupplySpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compBorrowSpeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenMetadata[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"cTokenUnderlyingPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenUnderlyingPrice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"cTokenUnderlyingPriceAll\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CTokenUnderlyingPrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLimits\",\"outputs\":[{\"components\":[{\"internalType\":\"contractCToken[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.AccountLimits\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCompBalanceMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structCompoundLens.CompBalanceMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerLensInterface\",\"name\":\"comptroller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCompBalanceMetadataExt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CompBalanceMetadataExt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractComp\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"blockNumbers\",\"type\":\"uint32[]\"}],\"name\":\"getCompVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structCompoundLens.CompVotes[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGovernorBravoInterface\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovBravoProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structCompoundLens.GovBravoProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGovernorBravoInterface\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovBravoReceipts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structCompoundLens.GovBravoReceipt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structCompoundLens.GovProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractGovernorAlpha\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"proposalIds\",\"type\":\"uint256[]\"}],\"name\":\"getGovReceipts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"support\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structCompoundLens.GovReceipt[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LensABI is the input ABI used to generate the binding from.
// Deprecated: Use LensMetaData.ABI instead.
var LensABI = LensMetaData.ABI

// Lens is an auto generated Go binding around an Ethereum contract.
type Lens struct {
	LensCaller     // Read-only binding to the contract
	LensTransactor // Write-only binding to the contract
	LensFilterer   // Log filterer for contract events
}

// LensCaller is an auto generated read-only Go binding around an Ethereum contract.
type LensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LensSession struct {
	Contract     *Lens             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LensCallerSession struct {
	Contract *LensCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LensTransactorSession struct {
	Contract     *LensTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LensRaw is an auto generated low-level Go binding around an Ethereum contract.
type LensRaw struct {
	Contract *Lens // Generic contract binding to access the raw methods on
}

// LensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LensCallerRaw struct {
	Contract *LensCaller // Generic read-only contract binding to access the raw methods on
}

// LensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LensTransactorRaw struct {
	Contract *LensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLens creates a new instance of Lens, bound to a specific deployed contract.
func NewLens(address common.Address, backend bind.ContractBackend) (*Lens, error) {
	contract, err := bindLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lens{LensCaller: LensCaller{contract: contract}, LensTransactor: LensTransactor{contract: contract}, LensFilterer: LensFilterer{contract: contract}}, nil
}

// NewLensCaller creates a new read-only instance of Lens, bound to a specific deployed contract.
func NewLensCaller(address common.Address, caller bind.ContractCaller) (*LensCaller, error) {
	contract, err := bindLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LensCaller{contract: contract}, nil
}

// NewLensTransactor creates a new write-only instance of Lens, bound to a specific deployed contract.
func NewLensTransactor(address common.Address, transactor bind.ContractTransactor) (*LensTransactor, error) {
	contract, err := bindLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LensTransactor{contract: contract}, nil
}

// NewLensFilterer creates a new log filterer instance of Lens, bound to a specific deployed contract.
func NewLensFilterer(address common.Address, filterer bind.ContractFilterer) (*LensFilterer, error) {
	contract, err := bindLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LensFilterer{contract: contract}, nil
}

// bindLens binds a generic wrapper to an already deployed contract.
func bindLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lens *LensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lens.Contract.LensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lens *LensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lens.Contract.LensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lens *LensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lens.Contract.LensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lens *LensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lens *LensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lens *LensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lens.Contract.contract.Transact(opts, method, params...)
}

// CTokenMetadataAll is a free data retrieval call binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensCaller) CTokenMetadataAll(opts *bind.CallOpts, cTokens []common.Address) ([]CompoundLensCTokenMetadata, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "cTokenMetadataAll", cTokens)

	if err != nil {
		return *new([]CompoundLensCTokenMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensCTokenMetadata)).(*[]CompoundLensCTokenMetadata)

	return out0, err

}

// CTokenMetadataAll is a free data retrieval call binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensSession) CTokenMetadataAll(cTokens []common.Address) ([]CompoundLensCTokenMetadata, error) {
	return _Lens.Contract.CTokenMetadataAll(&_Lens.CallOpts, cTokens)
}

// CTokenMetadataAll is a free data retrieval call binding the contract method 0x4b70d84b.
//
// Solidity: function cTokenMetadataAll(address[] cTokens) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensCallerSession) CTokenMetadataAll(cTokens []common.Address) ([]CompoundLensCTokenMetadata, error) {
	return _Lens.Contract.CTokenMetadataAll(&_Lens.CallOpts, cTokens)
}

// CTokenUnderlyingPrice is a free data retrieval call binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) view returns((address,uint256))
func (_Lens *LensCaller) CTokenUnderlyingPrice(opts *bind.CallOpts, cToken common.Address) (CompoundLensCTokenUnderlyingPrice, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "cTokenUnderlyingPrice", cToken)

	if err != nil {
		return *new(CompoundLensCTokenUnderlyingPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(CompoundLensCTokenUnderlyingPrice)).(*CompoundLensCTokenUnderlyingPrice)

	return out0, err

}

// CTokenUnderlyingPrice is a free data retrieval call binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) view returns((address,uint256))
func (_Lens *LensSession) CTokenUnderlyingPrice(cToken common.Address) (CompoundLensCTokenUnderlyingPrice, error) {
	return _Lens.Contract.CTokenUnderlyingPrice(&_Lens.CallOpts, cToken)
}

// CTokenUnderlyingPrice is a free data retrieval call binding the contract method 0xc5ae5934.
//
// Solidity: function cTokenUnderlyingPrice(address cToken) view returns((address,uint256))
func (_Lens *LensCallerSession) CTokenUnderlyingPrice(cToken common.Address) (CompoundLensCTokenUnderlyingPrice, error) {
	return _Lens.Contract.CTokenUnderlyingPrice(&_Lens.CallOpts, cToken)
}

// CTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) view returns((address,uint256)[])
func (_Lens *LensCaller) CTokenUnderlyingPriceAll(opts *bind.CallOpts, cTokens []common.Address) ([]CompoundLensCTokenUnderlyingPrice, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "cTokenUnderlyingPriceAll", cTokens)

	if err != nil {
		return *new([]CompoundLensCTokenUnderlyingPrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensCTokenUnderlyingPrice)).(*[]CompoundLensCTokenUnderlyingPrice)

	return out0, err

}

// CTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) view returns((address,uint256)[])
func (_Lens *LensSession) CTokenUnderlyingPriceAll(cTokens []common.Address) ([]CompoundLensCTokenUnderlyingPrice, error) {
	return _Lens.Contract.CTokenUnderlyingPriceAll(&_Lens.CallOpts, cTokens)
}

// CTokenUnderlyingPriceAll is a free data retrieval call binding the contract method 0x2b2d5ed6.
//
// Solidity: function cTokenUnderlyingPriceAll(address[] cTokens) view returns((address,uint256)[])
func (_Lens *LensCallerSession) CTokenUnderlyingPriceAll(cTokens []common.Address) ([]CompoundLensCTokenUnderlyingPrice, error) {
	return _Lens.Contract.CTokenUnderlyingPriceAll(&_Lens.CallOpts, cTokens)
}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Lens *LensCaller) GetAccountLimits(opts *bind.CallOpts, comptroller common.Address, account common.Address) (CompoundLensAccountLimits, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getAccountLimits", comptroller, account)

	if err != nil {
		return *new(CompoundLensAccountLimits), err
	}

	out0 := *abi.ConvertType(out[0], new(CompoundLensAccountLimits)).(*CompoundLensAccountLimits)

	return out0, err

}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Lens *LensSession) GetAccountLimits(comptroller common.Address, account common.Address) (CompoundLensAccountLimits, error) {
	return _Lens.Contract.GetAccountLimits(&_Lens.CallOpts, comptroller, account)
}

// GetAccountLimits is a free data retrieval call binding the contract method 0x7dd8f6d9.
//
// Solidity: function getAccountLimits(address comptroller, address account) view returns((address[],uint256,uint256))
func (_Lens *LensCallerSession) GetAccountLimits(comptroller common.Address, account common.Address) (CompoundLensAccountLimits, error) {
	return _Lens.Contract.GetAccountLimits(&_Lens.CallOpts, comptroller, account)
}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_Lens *LensCaller) GetCompBalanceMetadata(opts *bind.CallOpts, comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getCompBalanceMetadata", comp, account)

	if err != nil {
		return *new(CompoundLensCompBalanceMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(CompoundLensCompBalanceMetadata)).(*CompoundLensCompBalanceMetadata)

	return out0, err

}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_Lens *LensSession) GetCompBalanceMetadata(comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	return _Lens.Contract.GetCompBalanceMetadata(&_Lens.CallOpts, comp, account)
}

// GetCompBalanceMetadata is a free data retrieval call binding the contract method 0x416405d7.
//
// Solidity: function getCompBalanceMetadata(address comp, address account) view returns((uint256,uint256,address))
func (_Lens *LensCallerSession) GetCompBalanceMetadata(comp common.Address, account common.Address) (CompoundLensCompBalanceMetadata, error) {
	return _Lens.Contract.GetCompBalanceMetadata(&_Lens.CallOpts, comp, account)
}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Lens *LensCaller) GetCompVotes(opts *bind.CallOpts, comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getCompVotes", comp, account, blockNumbers)

	if err != nil {
		return *new([]CompoundLensCompVotes), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensCompVotes)).(*[]CompoundLensCompVotes)

	return out0, err

}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Lens *LensSession) GetCompVotes(comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	return _Lens.Contract.GetCompVotes(&_Lens.CallOpts, comp, account, blockNumbers)
}

// GetCompVotes is a free data retrieval call binding the contract method 0x59564219.
//
// Solidity: function getCompVotes(address comp, address account, uint32[] blockNumbers) view returns((uint256,uint256)[])
func (_Lens *LensCallerSession) GetCompVotes(comp common.Address, account common.Address, blockNumbers []uint32) ([]CompoundLensCompVotes, error) {
	return _Lens.Contract.GetCompVotes(&_Lens.CallOpts, comp, account, blockNumbers)
}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensCaller) GetGovBravoProposals(opts *bind.CallOpts, governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getGovBravoProposals", governor, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovBravoProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovBravoProposal)).(*[]CompoundLensGovBravoProposal)

	return out0, err

}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensSession) GetGovBravoProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	return _Lens.Contract.GetGovBravoProposals(&_Lens.CallOpts, governor, proposalIds)
}

// GetGovBravoProposals is a free data retrieval call binding the contract method 0x75d80e90.
//
// Solidity: function getGovBravoProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensCallerSession) GetGovBravoProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoProposal, error) {
	return _Lens.Contract.GetGovBravoProposals(&_Lens.CallOpts, governor, proposalIds)
}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_Lens *LensCaller) GetGovBravoReceipts(opts *bind.CallOpts, governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getGovBravoReceipts", governor, voter, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovBravoReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovBravoReceipt)).(*[]CompoundLensGovBravoReceipt)

	return out0, err

}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_Lens *LensSession) GetGovBravoReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	return _Lens.Contract.GetGovBravoReceipts(&_Lens.CallOpts, governor, voter, proposalIds)
}

// GetGovBravoReceipts is a free data retrieval call binding the contract method 0x43c811cc.
//
// Solidity: function getGovBravoReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,uint8,uint96)[])
func (_Lens *LensCallerSession) GetGovBravoReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovBravoReceipt, error) {
	return _Lens.Contract.GetGovBravoReceipts(&_Lens.CallOpts, governor, voter, proposalIds)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensCaller) GetGovProposals(opts *bind.CallOpts, governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getGovProposals", governor, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovProposal)).(*[]CompoundLensGovProposal)

	return out0, err

}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	return _Lens.Contract.GetGovProposals(&_Lens.CallOpts, governor, proposalIds)
}

// GetGovProposals is a free data retrieval call binding the contract method 0x96994869.
//
// Solidity: function getGovProposals(address governor, uint256[] proposalIds) view returns((uint256,address,uint256,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,bool,bool)[])
func (_Lens *LensCallerSession) GetGovProposals(governor common.Address, proposalIds []*big.Int) ([]CompoundLensGovProposal, error) {
	return _Lens.Contract.GetGovProposals(&_Lens.CallOpts, governor, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Lens *LensCaller) GetGovReceipts(opts *bind.CallOpts, governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "getGovReceipts", governor, voter, proposalIds)

	if err != nil {
		return *new([]CompoundLensGovReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new([]CompoundLensGovReceipt)).(*[]CompoundLensGovReceipt)

	return out0, err

}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Lens *LensSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	return _Lens.Contract.GetGovReceipts(&_Lens.CallOpts, governor, voter, proposalIds)
}

// GetGovReceipts is a free data retrieval call binding the contract method 0x995ed99f.
//
// Solidity: function getGovReceipts(address governor, address voter, uint256[] proposalIds) view returns((uint256,bool,bool,uint96)[])
func (_Lens *LensCallerSession) GetGovReceipts(governor common.Address, voter common.Address, proposalIds []*big.Int) ([]CompoundLensGovReceipt, error) {
	return _Lens.Contract.GetGovReceipts(&_Lens.CallOpts, governor, voter, proposalIds)
}

// NullAddress is a free data retrieval call binding the contract method 0x551edf90.
//
// Solidity: function nullAddress() view returns(address)
func (_Lens *LensCaller) NullAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lens.contract.Call(opts, &out, "nullAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NullAddress is a free data retrieval call binding the contract method 0x551edf90.
//
// Solidity: function nullAddress() view returns(address)
func (_Lens *LensSession) NullAddress() (common.Address, error) {
	return _Lens.Contract.NullAddress(&_Lens.CallOpts)
}

// NullAddress is a free data retrieval call binding the contract method 0x551edf90.
//
// Solidity: function nullAddress() view returns(address)
func (_Lens *LensCallerSession) NullAddress() (common.Address, error) {
	return _Lens.Contract.NullAddress(&_Lens.CallOpts)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensTransactor) CTokenBalances(opts *bind.TransactOpts, cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.contract.Transact(opts, "cTokenBalances", cToken, account)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensSession) CTokenBalances(cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenBalances(&_Lens.TransactOpts, cToken, account)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0xbdf950c9.
//
// Solidity: function cTokenBalances(address cToken, address account) returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensTransactorSession) CTokenBalances(cToken common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenBalances(&_Lens.TransactOpts, cToken, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensTransactor) CTokenBalancesAll(opts *bind.TransactOpts, cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.contract.Transact(opts, "cTokenBalancesAll", cTokens, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensSession) CTokenBalancesAll(cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenBalancesAll(&_Lens.TransactOpts, cTokens, account)
}

// CTokenBalancesAll is a paid mutator transaction binding the contract method 0x0972bf8b.
//
// Solidity: function cTokenBalancesAll(address[] cTokens, address account) returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_Lens *LensTransactorSession) CTokenBalancesAll(cTokens []common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenBalancesAll(&_Lens.TransactOpts, cTokens, account)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensTransactor) CTokenMetadata(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Lens.contract.Transact(opts, "cTokenMetadata", cToken)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensSession) CTokenMetadata(cToken common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenMetadata(&_Lens.TransactOpts, cToken)
}

// CTokenMetadata is a paid mutator transaction binding the contract method 0x158eca8b.
//
// Solidity: function cTokenMetadata(address cToken) returns((address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256,address,uint256,uint256,uint256,uint256,uint256))
func (_Lens *LensTransactorSession) CTokenMetadata(cToken common.Address) (*types.Transaction, error) {
	return _Lens.Contract.CTokenMetadata(&_Lens.TransactOpts, cToken)
}

// GetCompBalanceMetadataExt is a paid mutator transaction binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Lens *LensTransactor) GetCompBalanceMetadataExt(opts *bind.TransactOpts, comp common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.contract.Transact(opts, "getCompBalanceMetadataExt", comp, comptroller, account)
}

// GetCompBalanceMetadataExt is a paid mutator transaction binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Lens *LensSession) GetCompBalanceMetadataExt(comp common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.GetCompBalanceMetadataExt(&_Lens.TransactOpts, comp, comptroller, account)
}

// GetCompBalanceMetadataExt is a paid mutator transaction binding the contract method 0x1ea63741.
//
// Solidity: function getCompBalanceMetadataExt(address comp, address comptroller, address account) returns((uint256,uint256,address,uint256))
func (_Lens *LensTransactorSession) GetCompBalanceMetadataExt(comp common.Address, comptroller common.Address, account common.Address) (*types.Transaction, error) {
	return _Lens.Contract.GetCompBalanceMetadataExt(&_Lens.TransactOpts, comp, comptroller, account)
}

// LensLodeMetaDataIterator is returned from FilterLodeMetaData and is used to iterate over the raw logs and unpacked data for LodeMetaData events raised by the Lens contract.
type LensLodeMetaDataIterator struct {
	Event *LensLodeMetaData // Event containing the contract specifics and raw log

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
func (it *LensLodeMetaDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LensLodeMetaData)
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
		it.Event = new(LensLodeMetaData)
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
func (it *LensLodeMetaDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LensLodeMetaDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LensLodeMetaData represents a LodeMetaData event raised by the Lens contract.
type LensLodeMetaData struct {
	Balance   *big.Int
	Allocated *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLodeMetaData is a free log retrieval operation binding the contract event 0x72193a0da001f50c93b6968353455757d2ff3d7a200d835e38b12829b14a05a4.
//
// Solidity: event lodeMetaData(uint256 balance, uint256 allocated)
func (_Lens *LensFilterer) FilterLodeMetaData(opts *bind.FilterOpts) (*LensLodeMetaDataIterator, error) {

	logs, sub, err := _Lens.contract.FilterLogs(opts, "lodeMetaData")
	if err != nil {
		return nil, err
	}
	return &LensLodeMetaDataIterator{contract: _Lens.contract, event: "lodeMetaData", logs: logs, sub: sub}, nil
}

// WatchLodeMetaData is a free log subscription operation binding the contract event 0x72193a0da001f50c93b6968353455757d2ff3d7a200d835e38b12829b14a05a4.
//
// Solidity: event lodeMetaData(uint256 balance, uint256 allocated)
func (_Lens *LensFilterer) WatchLodeMetaData(opts *bind.WatchOpts, sink chan<- *LensLodeMetaData) (event.Subscription, error) {

	logs, sub, err := _Lens.contract.WatchLogs(opts, "lodeMetaData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LensLodeMetaData)
				if err := _Lens.contract.UnpackLog(event, "lodeMetaData", log); err != nil {
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

// ParseLodeMetaData is a log parse operation binding the contract event 0x72193a0da001f50c93b6968353455757d2ff3d7a200d835e38b12829b14a05a4.
//
// Solidity: event lodeMetaData(uint256 balance, uint256 allocated)
func (_Lens *LensFilterer) ParseLodeMetaData(log types.Log) (*LensLodeMetaData, error) {
	event := new(LensLodeMetaData)
	if err := _Lens.contract.UnpackLog(event, "lodeMetaData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
