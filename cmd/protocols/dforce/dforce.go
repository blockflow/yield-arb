package dforce

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
	t "yield-arb/cmd/protocols/types"
	txs "yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slices"
)

type DForce struct {
	chain              string
	cl                 *ethclient.Client
	chainID            *big.Int
	controllerAddress  common.Address
	controllerContract *Controller
	oracleAddress      common.Address
}

const DForceName = "dforce"

var controllerAddresses = map[string]string{
	"arbitrum": "0x8E7e9eA9023B81457Ae7E6D2a51b003D421E5408",
}

var ignoreTokens = []common.Address{
	common.HexToAddress("0xe8c85B60Cb3bA32369c699015621813fb2fEA56c"), // TODO: iMUSX, no supply rate?
	common.HexToAddress("0x5BE49B2e04aC55A17c72aC37E3a85D9602322021"), // TODO: iMEUX, no supply rate?
}

// Returns the chains supported by the protocol
func (d *DForce) GetChains() ([]string, error) {
	return []string{
		"ethereum",
		"polygon",
		"avalanche",
		"arbitrum",
		"conflux",
		"kava",
		"optimism",
		"binance",
	}, nil
}

func NewDForceProtocol() *DForce {
	return &DForce{}
}

func (d *DForce) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return fmt.Errorf("failed to connect to the %v client: %v", chain, err)
	}

	// Fetch chainid
	// TODO: Fetch this in multicall with pool address.
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch chainid: %v", err)
	}

	// Instantiate controller
	controllerAddress := common.HexToAddress(controllerAddresses[chain])
	controllerContract, err := NewController(controllerAddress, cl)
	if err != nil {
		return fmt.Errorf("failed to instantiate AddressesProvider: %v", err)
	}

	// Fetch price oracle
	oracleAddress, err := controllerContract.PriceOracle(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch oracle address: %v", err)
	}

	d.chain = chain
	d.cl = cl
	d.chainID = chainid
	d.controllerAddress = controllerAddress
	d.controllerContract = controllerContract
	d.oracleAddress = oracleAddress
	log.Printf("%v connected to %v (chainid: %v, controller: %v)", DForceName, d.chain, d.chainID, controllerAddress)
	return nil
}

// Returns the markets for the protocol
func (d *DForce) GetMarkets() (*t.ProtocolChain, error) {
	log.Printf("Fetching markets for %v on %v", DForceName, d.chain)
	startTime := time.Now()
	// Get all tokens
	allITokens, err := d.controllerContract.GetAlliTokens(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all iTokens: %v", err)
	}

	// Aggregate calldata
	iTokenABI, err := ITokenMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get iToken abi: %v", err)
	}
	controllerABI, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get controller abi: %v", err)
	}
	oracleABI, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get oracle abi: %v", err)
	}
	var calls []txs.Multicall3Call3
	iTokenMethods := []string{"symbol", "decimals", "supplyRatePerBlock", "borrowRatePerBlock", "totalSupply", "getCash"}
	for _, itoken := range allITokens {
		if slices.Contains[common.Address](ignoreTokens, itoken) {
			continue
		}
		// iToken calls
		for _, method := range iTokenMethods {
			data, err := iTokenABI.Pack(method)
			if err != nil {
				return nil, fmt.Errorf("failed to pack %v calldata: %v", method, err)
			}
			calls = append(calls, txs.Multicall3Call3{
				Target:   itoken,
				CallData: data,
			})
		}
		// Controller calls. Fetch markets.
		marketsData, err := controllerABI.Pack("markets", itoken)
		if err != nil {
			return nil, fmt.Errorf("failed to pack markets calldata: %v", err)
		}
		calls = append(calls, txs.Multicall3Call3{
			Target:   d.controllerAddress,
			CallData: marketsData,
		})
		// Price Oracle calls. Get underlying price.
		priceData, err := oracleABI.Pack("getUnderlyingPrice", itoken)
		if err != nil {
			return nil, fmt.Errorf("failed to pack getUnderlyingPrice calldata: %v", err)
		}
		calls = append(calls, txs.Multicall3Call3{
			Target:   d.oracleAddress,
			CallData: priceData,
		})
	}

	// Perform multicall
	responses, err := txs.HandleMulticall(d.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}

	// Unpack
	numTokens := len(allITokens) - len(ignoreTokens) // Exclude ignore token
	supplyMarkets := make([]*t.MarketInfo, numTokens)
	borrowMarkets := make([]*t.MarketInfo, numTokens)
	factor := len(iTokenMethods) + 2
	for i := 0; i < numTokens; i++ {
		j := i * factor
		// Unpack itoken calls
		var symbol string
		var decimals uint8
		var supplyRatePerBlock *big.Int
		var borrowRatePerBlock *big.Int
		var totalSupply *big.Int
		var getCash *big.Int
		results := []interface{}{&symbol, &decimals, &supplyRatePerBlock, &borrowRatePerBlock, &totalSupply, &getCash}
		for k, method := range iTokenMethods {
			if err := iTokenABI.UnpackIntoInterface(results[k], method, (*responses)[j+k].ReturnData); err != nil {
				return nil, fmt.Errorf("failed to unpack %v: %v", method, err)
			}
		}
		// Unpack controller calls
		var controllerMarket struct {
			CollateralFactorMantissa *big.Int
			BorrowFactorMantissa     *big.Int
			BorrowCapacity           *big.Int
			SupplyCapacity           *big.Int
			MintPaused               bool
			RedeemPaused             bool
			BorrowPaused             bool
		}
		if err := controllerABI.UnpackIntoInterface(&controllerMarket, "markets", (*responses)[j+factor-2].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack markets: %v", err)
		}
		// Unpack oracle calls
		var underlyingPrice *big.Int
		if err := oracleABI.UnpackIntoInterface(&underlyingPrice, "getUnderlyingPrice", (*responses)[j+factor-1].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack getUnderlyingPrice: %v", err)
		}
		log.Print(symbol, underlyingPrice)

		// Fill out market info
		// LTV
		ltvInt := new(big.Int).Mul(controllerMarket.CollateralFactorMantissa, big.NewInt(100))
		ltv := new(big.Float).SetInt(ltvInt)
		ltv.Quo(ltv, utils.ETHMantissa)
		// APYs
		supplyAPY := utils.ConvertRatePerBlockToAPY(supplyRatePerBlock)
		borrowAPY := utils.ConvertRatePerBlockToAPY(borrowRatePerBlock)
		// Caps
		decimalsFactor := big.NewFloat(math.Pow(10, float64(decimals)))
		supplyCapInt := new(big.Int).Sub(controllerMarket.SupplyCapacity, totalSupply)
		supplyCap := new(big.Float).SetInt(supplyCapInt)
		supplyCap.Quo(supplyCap, decimalsFactor)
		borrowCap := new(big.Float).SetInt(getCash)
		borrowCap.Quo(borrowCap, decimalsFactor)
		// Price, dforce uses diff decimals for diff tokens?...
		price := new(big.Float).SetInt(underlyingPrice)
		commonSymbol := utils.CommonSymbol(symbol)
		if commonSymbol == "USDC" || commonSymbol == "USDT" {
			price.Quo(price, big.NewFloat(math.Pow(10, 30)))
		} else if commonSymbol == "BTC" {
			price.Quo(price, big.NewFloat(math.Pow(10, 28)))
		} else {
			price.Quo(price, utils.ETHMantissa)
		}
		market := &t.MarketInfo{
			Protocol:   DForceName,
			Chain:      d.chain,
			Token:      symbol,
			Decimals:   decimals,
			LTV:        ltv,
			SupplyAPY:  supplyAPY,
			BorrowAPY:  borrowAPY,
			SupplyCap:  supplyCap,
			BorrowCap:  borrowCap,
			PriceInUSD: price,
		}
		supplyMarkets[i] = market
		borrowMarkets[i] = market
	}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolChain{
		Protocol:      DForceName,
		Chain:         d.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
}

// Lends the token to the protocol
func (d *DForce) Supply(from common.Address, token string, amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

// // Withdraws the token from the protocol
// Withdraw(user string, token string, amount *big.Int) error
// Borrows the token from the protocol
func (d *DForce) Borrow(from common.Address, token string, amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

// // Repays the token to the protocol
// Repay(user string, token string, amount *big.Int) error

// // Fetches the user's positions and leverage
// GetAccountData(user string)
