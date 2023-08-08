package protocols

import (
	"context"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain         string
	cl            *ethclient.Client
	chainId       *big.Int
	cometAddress  common.Address
	cometContract *bind.BoundContract
	baseToken     string // symbol
}

type CompV3AssetInfo struct {
	Offset                    uint8
	Asset                     common.Address
	PriceFeed                 common.Address
	Scale                     uint64
	BorrowCollateralFactor    uint64
	LiquidateCollateralFactor uint64
	LiquidationFactor         uint64
	SupplyCap                 *big.Int
}

const CompoundV3Name = "CompoundV3"
const compv3CometName = "comet"

var compv3CometAddresses = map[string]string{
	"arbitrum":        "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	"arbitrum_goerli": "0x1d573274E19174260c5aCE3f2251598959d24456",
	"polygon":         "0xF25212E676D1F7F89Cd72fFEe66158f541246445",
	"polygon_mumbai":  "0xF09F0369aB0a875254fB565E52226c88f10Bc839",
}

func NewCompoundV3Protocol() *CompoundV3 {
	return &CompoundV3{}
}

// Returns the chains supported by the protocol
func (c *CompoundV3) GetChains() ([]string, error) {
	return []string{"arbitrum", "arbitrum_goerli", "polygon", "polygon_mumbai"}, nil
}

// Connects to the protocol on the specified chain
func (c *CompoundV3) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Fetch chainId
	chainId, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
		return err
	}

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// Instantiate comet
	cometPath := filepath.Join(dir, "abis", CompoundV3Name, compv3CometName+".json")
	cometRawABI, err := os.Open(cometPath)
	if err != nil {
		log.Printf("Failed to open %v.json: %v", compv3CometName, err)
		return err
	}
	defer cometRawABI.Close()
	cometParsedABI, err := abi.JSON(cometRawABI)
	if err != nil {
		log.Printf("Failed to parse %v.json: %v", compv3CometName, err)
		return err
	}
	cometAddress := common.HexToAddress(compv3CometAddresses[chain])
	cometContract := bind.NewBoundContract(cometAddress, cometParsedABI, cl, cl, cl)

	c.chain = chain
	c.cl = cl
	c.chainId = chainId
	c.cometAddress = cometAddress
	c.cometContract = cometContract

	// Get base token symbol
	baseToken, err := c.getBaseToken()
	if err != nil {
		log.Printf("Failed to get base token: %v", err)
		return err
	}
	baseTokenSymbol, err := utils.ConvertAddressesToSymbols(chain, []string{baseToken})
	if err != nil {
		log.Printf("Failed to convert addresses to symbols: %v", err)
		return err
	}

	c.baseToken = baseTokenSymbol[0]

	log.Printf("Connected to %v (chainid: %v)", c.chain, c.chainId)
	return nil
}

// Returns the hex address of the base token.
func (c *CompoundV3) getBaseToken() (string, error) {
	baseTokenResult := []interface{}{new(common.Address)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &baseTokenResult, "baseToken"); err != nil {
		log.Printf("Failed to get base token: %v", err)
		return "", err
	}
	return baseTokenResult[0].(*common.Address).Hex(), nil
}

func (c *CompoundV3) getAssetInfo(offset uint8) (*CompV3AssetInfo, error) {
	type AssetInfoOutput struct {
		Output CompV3AssetInfo
	}
	results := []interface{}{new(AssetInfoOutput)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &results, "getAssetInfo", offset); err != nil {
		log.Printf("Failed to get asset info for %v: %v", offset, err)
		return nil, err
	}
	return &results[0].(*AssetInfoOutput).Output, nil
}

func (c *CompoundV3) getAssetInfoByAddress(address string) (*CompV3AssetInfo, error) {
	type AssetInfoOutput struct {
		Output CompV3AssetInfo
	}
	results := []interface{}{new(AssetInfoOutput)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &results, "getAssetInfoByAddress", common.HexToAddress(address)); err != nil {
		log.Printf("Failed to get asset info for %v: %v", address, err)
		return nil, err
	}
	return &results[0].(*AssetInfoOutput).Output, nil
}

func (c *CompoundV3) getAllAssetInfos() ([]*CompV3AssetInfo, error) {
	numAssetsResult := []interface{}{new(uint8)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &numAssetsResult, "numAssets"); err != nil {
		log.Printf("Failed to get num assets: %v", err)
		return nil, err
	}
	numAssets := *numAssetsResult[0].(*uint8)

	result := make([]*CompV3AssetInfo, numAssets)
	for i := uint8(0); i < numAssets; i++ {
		if assetInfo, err := c.getAssetInfo(uint8(i)); err != nil {
			log.Printf("Failed to get asset info for %v: %v", i, err)
			return nil, err
		} else {
			result[i] = assetInfo
		}
	}

	return result, nil
}

// Returns the tokens supported by the protocol on the specified chain.
// Tokens are represented as their symbols.
func (c *CompoundV3) GetLendingTokens() ([]string, error) {
	// Get all assets
	assetInfos, err := c.getAllAssetInfos()
	if err != nil {
		log.Printf("Failed to get all asset infos: %v", err)
		return nil, err
	}

	// Extract addresses
	assets := make([]string, len(assetInfos))
	for i, assetInfo := range assetInfos {
		assets[i] = assetInfo.Asset.Hex()
	}

	// Convert to symbols
	symbols, err := utils.ConvertAddressesToSymbols(c.chain, assets)
	if err != nil {
		log.Printf("Failed to convert addresses to symbols: %v", err)
		return nil, err
	}

	return append(symbols, c.baseToken), nil
}

// Returns the tokens supported by the protocol on the specified chain.
// Tokens are represented as their symbols.
// CompV3 only allows borrowing the base token.
func (c *CompoundV3) GetBorrowingTokens() ([]string, error) {
	return []string{c.baseToken}, nil
}

/*
method: getSupplyRate
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Supply Rate = getSupplyRate(Utilization)
Supply APR = Supply Rate / (10 ^ 18) * Seconds Per Year * 100

method: getBorrowRate
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Borrow Rate = getBorrowRate(Utilization)
Borrow APR = Borrow Rate / (10 ^ 18) * Seconds Per Year * 100
*/
func (c *CompoundV3) getBaseAPR(method string) (*big.Float, error) {
	// Get utilization
	type UtilResult struct {
		Utilization *big.Int
	}
	utilResult := []interface{}{new(UtilResult)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &utilResult, "getUtilization"); err != nil {
		log.Printf("Failed to get utilization: %v", err)
		return nil, err
	}
	utilization := utilResult[0].(*UtilResult).Utilization

	// Get supply rate
	rateResult := []interface{}{new(uint64)}
	if err := c.cometContract.Call(&bind.CallOpts{}, &rateResult, method, utilization); err != nil {
		log.Printf("Failed to %v: %v", method, err)
		return nil, err
	}
	rate := big.NewFloat(float64(*rateResult[0].(*uint64)))

	// Calculate APR
	numerator := new(big.Float).Mul(rate, utils.SecPerYear)
	numerator.Mul(numerator, big.NewFloat(100))
	apr := new(big.Float).Quo(numerator, utils.ETHMantissa)

	return apr, nil
}

// Returns the TokenSpecs for the specified tokens
// Tokens are represented as their symbols
// Only base token earns APR. Base token LTV is 0.
func (c *CompoundV3) GetLendingSpecs(symbols []string) ([]*TokenSpecs, error) {
	addresses, err := utils.ConvertSymbolsToAddresses(c.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}

	result := make([]*TokenSpecs, len(symbols))
	for i, symbol := range symbols {
		if symbol == c.baseToken {
			baseAPR, err := c.getBaseAPR("getSupplyRate")
			if err != nil {
				log.Printf("Failed to get base apr: %v", err)
				return nil, err
			}
			result[i] = &TokenSpecs{
				Protocol: CompoundV3Name,
				Chain:    c.chain,
				Token:    symbol,
				LTV:      big.NewFloat(0),
				APY:      baseAPR,
			}
		} else {
			assetInfo, err := c.getAssetInfoByAddress(addresses[i])
			if err != nil {
				log.Printf("Failed to get asset info: %v", err)
				return nil, err
			}
			ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
			ltv.Quo(ltv, big.NewFloat(10000000000000000))
			result[i] = &TokenSpecs{
				Protocol: CompoundV3Name,
				Chain:    c.chain,
				Token:    symbol,
				LTV:      ltv,
				APY:      big.NewFloat(0),
			}
		}
	}

	return result, nil
}

// Returns the TokenSpecs for the specified tokens.
// Tokens are represented as their symbols.
// Only base token can be borrowed. Sets LTV to 0.
func (c *CompoundV3) GetBorrowingSpecs(symbols []string) ([]*TokenSpecs, error) {
	for _, symbol := range symbols {
		if symbol == c.baseToken {
			baseAPR, err := c.getBaseAPR("getBorrowRate")
			if err != nil {
				log.Printf("Failed to get base apr: %v", err)
				return nil, err
			}
			return []*TokenSpecs{{
				Protocol: CompoundV3Name,
				Chain:    c.chain,
				Token:    symbol,
				LTV:      big.NewFloat(0),
				APY:      baseAPR,
			}}, nil
		}
	}

	return []*TokenSpecs{}, nil
}

// Returns the markets for the protocol
func (c *CompoundV3) GetMarkets() (*ProtocolMarkets, error) {
	log.Println("Fetching markets...")
	symbols, err := c.GetLendingTokens()
	if err != nil {
		log.Printf("Failed to fetch tokens: %v", err)
	}
	lendingSpecs, err := c.GetLendingSpecs(symbols)
	if err != nil {
		log.Printf("Failed to fetch lending specs: %v", err)
	}
	borrowingSpecs, err := c.GetBorrowingSpecs(symbols)
	if err != nil {
		log.Printf("Failed to fetch borrowing specs: %v", err)
	}
	return &ProtocolMarkets{
		Protocol:       CompoundV3Name,
		Chain:          c.chain,
		LendingSpecs:   lendingSpecs,
		BorrowingSpecs: borrowingSpecs,
	}, nil
}

// // Lends the token to the protocol
func (c *CompoundV3) Deposit(from string, token string, amount *big.Int) (*common.Hash, error) {
	return nil, nil
}

// TODO: Handle ETH transactions differently!!!
