package dforce

import (
	"context"
	"fmt"
	"log"
	"math/big"
	t "yield-arb/cmd/protocols/types"
	txs "yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DForce struct {
	chain              string
	cl                 *ethclient.Client
	chainID            *big.Int
	controllerAddress  common.Address
	controllerContract *Controller
}

const DForceName = "dforce"

var controllerAddresses = map[string]string{
	"arbitrum": "0x8E7e9eA9023B81457Ae7E6D2a51b003D421E5408",
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
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Fetch chainid
	// TODO: Fetch this in multicall with pool address.
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
		return err
	}

	// Instantiate controller
	controllerAddress := common.HexToAddress(controllerAddresses[chain])
	controllerContract, err := NewController(controllerAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate AddressesProvider: %v", err)
		return err
	}

	d.chain = chain
	d.cl = cl
	d.chainID = chainid
	d.controllerAddress = controllerAddress
	d.controllerContract = controllerContract
	log.Printf("%v connected to %v (chainid: %v, controller: %v)", DForceName, d.chain, d.chainID, controllerAddress)
	return nil
}

// Returns the markets for the protocol
func (d *DForce) GetMarkets() (*t.ProtocolChain, error) {
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
	var calls []txs.Multicall3Call3
	iTokenMethods := []string{"symbol", "decimals", "supplyRatePerBlock", "borrowRatePerBlock", "totalSupply", "totalBorrows"}
	for _, itoken := range allITokens {
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
		data, err := controllerABI.Pack("markets", itoken)
		if err != nil {
			return nil, fmt.Errorf("failed to pack markets calldata: %v", err)
		}
		calls = append(calls, txs.Multicall3Call3{
			Target:   d.controllerAddress,
			CallData: data,
		})
	}

	// Perform multicall
	responses, err := txs.HandleMulticall(d.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall asset info: %v", err)
	}

	// Unpack
	supplyMarkets := make([]*t.MarketInfo, len(allITokens))
	borrowMarkets := make([]*t.MarketInfo, len(allITokens))
	type ReturnData struct {
		TotalSupplyAsset *big.Int
		Reserved         *big.Int
	}
	factor := len(iTokenMethods) + 1
	// iTokenMethods := []string{"symbol", "decimals", "supplyRatePerBlock", "borrowRatePerBlock", "totalSupply", "totalBorrows"}
	for i, _ := range allITokens {
		j := i * factor
		// Unpack itoken calls
		var symbol string
		var decimals uint8
		var supplyRatePerBlock *big.Int
		var borrowRatePerBlock *big.Int
		var totalSupply *big.Int
		var totalBorrows *big.Int
		results := []interface{}{&symbol, &decimals, supplyRatePerBlock, borrowRatePerBlock, totalSupply, totalBorrows}
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
		if err := controllerABI.UnpackIntoInterface(&controllerMarket, "markets", (*responses)[j+factor-1].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack markets: %v", err)
		}

		// Fill out market info
		// LTV
		ltvInt := new(big.Int).Mul(controllerMarket.CollateralFactorMantissa, big.NewInt(100))
		ltv := new(big.Float).SetInt(ltvInt)
		ltv.Quo(ltv, utils.ETHMantissa)
		// APYs
		supplyAPYInt := new(big.Int).Exp(supplyRatePerBlock, utils.BlocksPerDay[d.chain], nil)
		supplyAPY := new(big.Float).SetInt(supplyAPYInt)
		borrowAPYInt := new(big.Int).Exp(borrowRatePerBlock, utils.BlocksPerDay[d.chain], nil)
		borrowAPY := new(big.Float).SetInt(borrowAPYInt)
		// Caps
		supplyCap := new(big.Int).Sub(controllerMarket.SupplyCapacity, totalSupply)
		borrowCap := new(big.Int).Sub(totalSupply, totalBorrows)
		market := &t.MarketInfo{
			Protocol:  DForceName,
			Chain:     d.chain,
			Token:     symbol,
			Decimals:  decimals,
			LTV:       ltv,
			SupplyAPY: supplyAPY,
			BorrowAPY: borrowAPY,
			SupplyCap: new(big.Float).SetInt(supplyCap),
			BorrowCap: new(big.Float).SetInt(borrowCap),
			// PriceInUSD: ,
		}
		supplyMarkets[i] = market
		borrowMarkets[i] = market
	}

	return &t.ProtocolChain{
		Protocol:      DForceName,
		Chain:         d.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
}

// // Lends the token to the protocol
// Supply(from common.Address, token string, amount *big.Int) (*types.Transaction, error)
// // Withdraws the token from the protocol
// Withdraw(user string, token string, amount *big.Int) error
// // Borrows the token from the protocol
// Borrow(from common.Address, token string, amount *big.Int) (*types.Transaction, error)
// // Repays the token to the protocol
// Repay(user string, token string, amount *big.Int) error

// // Fetches the user's positions and leverage
// GetAccountData(user string)
