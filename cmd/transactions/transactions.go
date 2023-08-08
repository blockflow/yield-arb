package transactions

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
