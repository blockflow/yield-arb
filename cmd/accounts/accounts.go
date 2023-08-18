package accounts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/joho/godotenv"
)

type SignedPermit struct {
	Deadline  *big.Int
	V         uint8
	R         []byte
	S         []byte
	Signature []byte
}

const PrivateKeyKey string = "PRIVATE_KEY"

var EIP2612Types = apitypes.Types{
	"EIP712Domain": []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "chainId", Type: "uint256"},
		{Name: "verifyingContract", Type: "address"},
	},
	"Permit": []apitypes.Type{
		{Name: "owner", Type: "address"},
		{Name: "spender", Type: "address"},
		{Name: "value", Type: "uint256"},
		{Name: "nonce", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
	},
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load dotenv: %v", err)
	}
	log.Print("Loaded dotenv")
}

// Signs a permit for the given token.
// Used https://github.com/everFinance/goether/tree/main as reference.
func SignEIP2612Permit(cl *ethclient.Client, chainID *big.Int, chain, token string, owner, spender common.Address, value *big.Int) (*SignedPermit, error) {
	tokenAddress, err := utils.ConvertSymbolToAddress(chain, token)
	if err != nil {
		return nil, err
	}

	// Get nonce
	erc20PermitContract, err := NewERC20Permit(common.HexToAddress(tokenAddress), cl)
	if err != nil {
		return nil, fmt.Errorf("failed to get erc20 permit contract: %v", err)
	}
	nonce, err := erc20PermitContract.Nonces(nil, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	// Set deadline to 1 day from now
	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %v", err)
	}
	timestamp := big.NewInt(0).SetUint64(block.Time())
	deadline := big.NewInt(0).Add(timestamp, big.NewInt(86400))

	domain := apitypes.TypedDataDomain{
		Name:              token,
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
		VerifyingContract: tokenAddress,
	}
	message := apitypes.TypedDataMessage{
		"owner":    owner.Hex(),
		"spender":  spender.Hex(),
		"value":    value.String(),
		"nonce":    nonce.String(),
		"deadline": deadline.String(),
	}

	typedData := &apitypes.TypedData{
		Types:       EIP2612Types,
		PrimaryType: "Permit",
		Domain:      domain,
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", domain.Map())
	if err != nil {
		return nil, fmt.Errorf("failed to hash domain: %v", err)
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, message)
	if err != nil {
		return nil, fmt.Errorf("failed to hash message: %v", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256(rawData)

	privKey, err := GetPrivKey(owner.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to get priv key: %v", err)
	}
	signature, err := crypto.Sign(hash, privKey)
	if err != nil {
		log.Fatalf("Failed to sign message: %v", err)
	}

	return &SignedPermit{
		Deadline:  deadline,
		V:         signature[64] + 27,
		R:         signature[:32],
		S:         signature[32:64],
		Signature: signature,
	}, nil
}

func SignTransaction(from string, chainID big.Int, tx *types.Transaction) (*types.Transaction, error) {
	signer := types.NewCancunSigner(&chainID)
	pk := os.Getenv(PrivateKeyKey)
	sk, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("failed to load pk: %v", err)
	}
	return types.SignTx(tx, signer, sk)
}

func GetPrivKey(from string) (*ecdsa.PrivateKey, error) {
	pk := os.Getenv(PrivateKeyKey)
	sk, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("failed to load pk: %v", err)
	}
	return sk, nil
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
