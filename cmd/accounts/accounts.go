package accounts

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
	sk, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Printf("Failed to load pk: %v", err)
		return nil, err
	}
	return types.SignTx(tx, signer, sk)
}

// Returns the auth object to authenticate abigen transactors.
// TODO: switch to geth account management.
// To get passphrases from terminal: https://pkg.go.dev/golang.org/x/term
func GetAuth(e *ethclient.Client, chainID *big.Int, from common.Address) (*bind.TransactOpts, error) {
	// Use .env pk for now
	pkHex := os.Getenv(PrivateKeyKey)
	if pkHex == "" {
		log.Fatalf("private key not found in .env")
	}
	privateKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		log.Fatalf("failed to convert private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID) // ChainID 1 for mainnet
	if err != nil {
		log.Fatalf("failed to create new keyed transactor: %v", err)
	}

	return auth, nil
}
