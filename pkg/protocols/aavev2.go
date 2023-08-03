package protocols

import (
	"context"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Add notifications for when config changes

var addressProvider common.Address = common.HexToAddress("0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5")

type AaveV2 struct {
	chain       string
	cl          *ethclient.Client
	chainid     big.Int
	lendingpool common.Address
}

func NewAaveV2Protocol() *AaveV2 {
	return &AaveV2{}
}

func (a *AaveV2) GetChains() ([]string, error) {
	return []string{"ethereum"}, nil
}

func (a *AaveV2) Connect(chain string) error {
	// Setup the client
	var cl *ethclient.Client
	switch chain {
	case "ethereum":
		_cl, clErr := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/NiPLhDKdUp9f7e6BPsQeW4lRXAo2rtbZ")
		if clErr != nil {
			log.Printf("Failed to connect to the %v client: %v", chain, clErr)
			return clErr
		}
		cl = _cl
	}

	// Fetch chainid
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
		return err
	}

	// Load addresses provider abi
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "abis", "aavev2", "addresses_provider.json")
	rawABI, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open addresses_provider.json: %v", err)
		return err
	}
	defer rawABI.Close()
	parsedABI, err := abi.JSON(rawABI)
	if err != nil {
		log.Printf("Failed to parse addresses_provider.json: %v", err)
		return err
	}

	// Instantiate addresses provider
	addressesProviderContract := bind.NewBoundContract(addressProvider, parsedABI, cl, cl, cl)

	// Fetch lending pool address
	results := []interface{}{new(common.Address)}
	callOpts := &bind.CallOpts{}
	err = addressesProviderContract.Call(callOpts, &results, "getLendingPool")
	if err != nil {
		log.Printf("Failed to fetch lending pool: %v", err)
		return err
	}

	a.chain = chain
	a.cl = cl
	a.chainid = *chainid
	a.lendingpool = *results[0].(*common.Address)
	log.Printf("Connected to %v (chainid: %v, lendingpool: %v)", a.chain, a.chainid, a.lendingpool)
	return nil
}

func (a *AaveV2) GetLendingTokens() ([]string, error) {
	return []string{"DAI", "USDC", "USDT", "TUSD", "SUSD", "BAT", "ETH", "LINK", "KNC", "REP", "WBTC", "ZRX"}, nil
}
