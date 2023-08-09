package compoundv3

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"
	t "yield-arb/cmd/protocols/types"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain         string
	cl            *ethclient.Client
	chainId       *big.Int
	cometAddress  common.Address
	cometContract *Comet
	baseToken     string // symbol
}

const CompoundV3Name = "CompoundV3"

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

	// Instantiate comet
	cometAddress := common.HexToAddress(compv3CometAddresses[chain])
	cometContract, err := NewComet(cometAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate comet: %v", err)
		return err
	}

	// Get base token symbol
	baseTokenAddress, err := cometContract.CometCaller.BaseToken(nil)
	if err != nil {
		log.Printf("Failed to fetch base token: %v", err)
		return err
	}

	symbols, err := utils.ConvertAddressesToSymbols(chain, []string{baseTokenAddress.Hex()})
	if err != nil {
		log.Printf("Failed to convert base token: %v", err)
		return err
	}

	c.chain = chain
	c.cl = cl
	c.chainId = chainId
	c.cometAddress = cometAddress
	c.cometContract = cometContract
	c.baseToken = symbols[0]

	log.Printf("Connected to %v (chainid: %v)", c.chain, c.chainId)
	return nil
}

// TODO: Deploy a contract to reduce API calls
func (c *CompoundV3) getAllAssetInfos() ([]*CometCoreAssetInfo, error) {
	numAssets, err := c.cometContract.NumAssets(nil)
	if err != nil {
		log.Printf("Failed to fetch num assets: %v", err)
		return nil, err
	}

	result := make([]*CometCoreAssetInfo, numAssets)
	for i := uint8(0); i < numAssets; i++ {
		if assetInfo, err := c.cometContract.GetAssetInfo(nil, i); err != nil {
			log.Printf("Failed to get asset info for %v: %v", i, err)
			return nil, err
		} else {
			result[i] = &assetInfo
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
func (c *CompoundV3) getBaseAPR(isSupplying bool) (*big.Float, error) {
	// Get utilization
	utilization, err := c.cometContract.GetUtilization(nil)
	if err != nil {
		log.Printf("Failed to get utilization: %v", err)
		return nil, err
	}

	// Get supply rate
	var rateInt uint64
	if isSupplying {
		rateInt, err = c.cometContract.GetSupplyRate(nil, utilization)
	} else {
		rateInt, err = c.cometContract.GetBorrowRate(nil, utilization)
	}
	if err != nil {
		log.Printf("Failed to get rate: %v", err)
	}
	rate := big.NewFloat(float64(rateInt))

	// Calculate APR
	numerator := new(big.Float).Mul(rate, utils.SecPerYear)
	numerator.Mul(numerator, big.NewFloat(100))
	apr := new(big.Float).Quo(numerator, utils.ETHMantissa)

	return apr, nil
}

// Returns the TokenSpecs for the specified tokens
// Tokens are represented as their symbols
// Only base token earns APR. Base token LTV is 0.
func (c *CompoundV3) GetLendingSpecs(symbols []string) ([]*t.TokenSpecs, error) {
	addresses, err := utils.ConvertSymbolsToAddresses(c.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}

	result := make([]*t.TokenSpecs, len(symbols))
	for i, symbol := range symbols {
		if symbol == c.baseToken {
			baseAPR, err := c.getBaseAPR(true)
			if err != nil {
				log.Printf("Failed to get base apr: %v", err)
				return nil, err
			}
			result[i] = &t.TokenSpecs{
				Protocol: CompoundV3Name,
				Chain:    c.chain,
				Token:    symbol,
				LTV:      big.NewFloat(0),
				APY:      baseAPR,
			}
		} else {
			assetInfo, err := c.cometContract.GetAssetInfoByAddress(nil, common.HexToAddress(addresses[i]))
			if err != nil {
				log.Printf("Failed to get asset info by address: %v", err)
				return nil, err
			}
			ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
			ltv.Quo(ltv, big.NewFloat(10000000000000000))
			result[i] = &t.TokenSpecs{
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
func (c *CompoundV3) GetBorrowingSpecs(symbols []string) ([]*t.TokenSpecs, error) {
	for _, symbol := range symbols {
		if symbol == c.baseToken {
			baseAPR, err := c.getBaseAPR(false)
			if err != nil {
				log.Printf("Failed to get base apr: %v", err)
				return nil, err
			}
			return []*t.TokenSpecs{{
				Protocol: CompoundV3Name,
				Chain:    c.chain,
				Token:    symbol,
				LTV:      big.NewFloat(0),
				APY:      baseAPR,
			}}, nil
		}
	}

	return []*t.TokenSpecs{}, nil
}

// Returns the markets for the protocol
func (c *CompoundV3) GetMarkets() (*t.ProtocolMarkets, error) {
	startTime := time.Now()
	log.Println("Fetching markets...")

	// Get all token info
	allAssetInfos, err := c.getAllAssetInfos()
	if err != nil {
		log.Printf("Failed to fetch all asset infos: %v", err)
		return nil, err
	}

	// Fill in LTV and APY for collateral tokens
	supplySpecs := make([]*t.TokenSpecs, len(allAssetInfos)+1)
	for i, assetInfo := range allAssetInfos {
		symbols, err := utils.ConvertAddressesToSymbols(c.chain, []string{assetInfo.Asset.Hex()})
		if err != nil {
			log.Printf("Failed to convert symbol: %v", err)
			return nil, err
		} else if len(symbols) == 0 {
			msg := "token address mapping missing"
			log.Println(msg)
			return nil, errors.New(msg)
		}
		symbol := symbols[0]
		// Has LTV, no APY
		ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
		ltv.Quo(ltv, big.NewFloat(10000000000000000))
		supplySpecs[i] = &t.TokenSpecs{
			Protocol: CompoundV3Name,
			Chain:    c.chain,
			Token:    symbol,
			LTV:      ltv,
			APY:      big.NewFloat(0),
		}
	}

	// Base token, has APY, no LTV
	supplyAPY, err := c.getBaseAPR(true)
	if err != nil {
		log.Printf("Failed to get supply apr: %v", err)
	}
	borrowAPY, err := c.getBaseAPR(false)
	if err != nil {
		log.Printf("Failed to get borrow apr: %v", err)
	}
	supplySpecs[len(allAssetInfos)] = &t.TokenSpecs{
		Protocol: CompoundV3Name,
		Chain:    c.chain,
		Token:    c.baseToken,
		LTV:      big.NewFloat(0),
		APY:      supplyAPY,
	}
	borrowSpecs := []*t.TokenSpecs{{
		Protocol: CompoundV3Name,
		Chain:    c.chain,
		Token:    c.baseToken,
		LTV:      big.NewFloat(0),
		APY:      borrowAPY,
	}}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplySpecs), len(borrowSpecs))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolMarkets{
		Protocol:       CompoundV3Name,
		Chain:          c.chain,
		LendingSpecs:   supplySpecs,
		BorrowingSpecs: borrowSpecs,
	}, nil
}

// // Lends the token to the protocol
func (c *CompoundV3) Supply(from string, token string, amount *big.Int) (*common.Hash, error) {
	return nil, nil
}

// TODO: Handle ETH transactions differently!!!
