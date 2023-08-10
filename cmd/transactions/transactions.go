package transactions

import (
	"context"
	"fmt"
	"log"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

// TODO: Estimate gas price dynamically
// Fetches tx params and builds the tx.
func BuildTransaction(e *ethclient.Client, from, to string, amount *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := e.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Printf("Failed to retrieve nonce for %v: %v", from, err)
		return nil, err
	}
	gasPrice, err := e.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to retrieve gasPrice: %v", err)
		return nil, err
	}
	return types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		amount,
		uint64(300000), // gasLimit
		gasPrice,
		data,
	), nil
}

// Sends the tx and returns the tx hash.
func SendTransaction(e *ethclient.Client, signedTx *types.Transaction) (*common.Hash, error) {
	err := e.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Failed to send tx: %v", err)
		return nil, err
	}
	hash := signedTx.Hash()
	return &hash, nil
}

// Bundles multiple calls into one RPC call and returns the results.
// Uses aggregate3 (https://github.com/mds1/multicall)
func HandleMulticall(e *ethclient.Client, calls *[]Multicall3Call3) (*[]Multicall3Result, error) {
	// Pack aggregated calldata
	multicallABI, _ := Multicall3MetaData.GetAbi()
	aggData, err := multicallABI.Pack("aggregate3", *calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall: %v", err)
	}

	// Call Multicall contract
	response, err := e.CallContract(context.Background(), ethereum.CallMsg{
		To:   &MulticallAddress,
		Data: aggData,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call multicall: %v", err)
	}

	// Unpack into Multicall3Result
	responses := new([]Multicall3Result)
	if err := multicallABI.UnpackIntoInterface(responses, "aggregate3", response); err != nil {
		return nil, fmt.Errorf("failed to unpack into results: %v", err)
	}

	return responses, nil
}
