package accounts

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

const PrivateKeyKey string = "PRIVATE_KEY"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load dotenv: %v", err)
	}
	log.Print("Loaded dotenv")
}

func SignTransaction(from string, chainID big.Int, tx *types.Transaction) (*types.Transaction, error) {
	signer := types.NewCancunSigner(&chainID)
	pk := os.Getenv(PrivateKeyKey)
	log.Print(pk)
	// sk, err := crypto.ToECDSA(common.FromHex(pk))
	sk, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Printf("Failed to load pk: %v", err)
		return nil, err
	}
	return types.SignTx(tx, signer, sk)
}
