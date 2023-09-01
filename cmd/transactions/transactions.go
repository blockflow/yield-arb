package transactions

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/utils"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

var approveGasLimits = map[string]uint64{
	"ethereum": 50000,
	"arbitrum": 500000,
	"base":     40000,
}

// TODO: Estimate gas price dynamically
// Fetches tx params and builds the tx.
func BuildTransaction(e *ethclient.Client, from, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := e.PendingNonceAt(context.Background(), from)
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
		to,
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

// Waits until the tx has the specified number of confirmations.
func WaitForConfirmations(cl *ethclient.Client, tx *types.Transaction, confirms uint64) (*types.Receipt, error) {
	// Wait for tx to be mined
	receipt, err := bind.WaitMined(context.Background(), cl, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}
	txBlockNumber := receipt.BlockNumber.Uint64()

	// Wait for confirmations
	latestBlock := uint64(0)
	for latestBlock < txBlockNumber+confirms {
		// Sleep for 1 sec
		time.Sleep(time.Second)
		latestBlock, err = cl.BlockNumber(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get latest block number: %v", err)
		}
	}

	return receipt, nil
}

// Checks to see if the approval is >= amount and approves max if not.
// Will wait for the approval tx if it is sent.
// If approval not needed, will return nil for tx.
func ApproveERC20IfNeeded(cl *ethclient.Client, auth *bind.TransactOpts, token, owner, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tokenContract, err := NewERC20Permit(token, cl)
	if err != nil {
		return nil, fmt.Errorf("failed to get token contract: %v", err)
	}

	// Get allowance
	allowance, err := tokenContract.Allowance(nil, owner, spender)
	if err != nil {
		return nil, fmt.Errorf("failed to get allowance: %v", err)
	}

	// Approve if needed, less than halfmax and less than amount
	halfMax := new(big.Int).Div(utils.MaxUint256, big.NewInt(2))
	if allowance.Cmp(halfMax) == -1 && allowance.Cmp(amount) == -1 {
		tx, err := tokenContract.Approve(auth, spender, utils.MaxUint64)
		if err != nil {
			return nil, fmt.Errorf("failed to approve: %v", err)
		}
		log.Printf("Approving %v %v for %v", utils.MaxUint64, token, spender)
		_, err = WaitForConfirmations(cl, tx, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to wait for confirmations: %v", err)
		}
		log.Printf("Approved %v %v for %v", utils.MaxUint64, token, spender)
		return tx, nil
	}
	return nil, nil
}

func GetApprovalTxIfNeeded(cl *ethclient.Client, chain string, token, owner, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tokenContract, err := NewERC20Permit(token, cl)
	if err != nil {
		return nil, fmt.Errorf("failed to get token contract: %v", err)
	}
	tokenABI, err := ERC20PermitMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get token abi: %v", err)
	}

	// Get allowance
	allowance, err := tokenContract.Allowance(nil, owner, spender)
	if err != nil {
		return nil, fmt.Errorf("failed to get allowance: %v", err)
	}

	// Approve if needed
	if allowance.Cmp(amount) == -1 {
		data, err := tokenABI.Pack("approve", spender, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve: %v", err)
		}
		inner := &types.DynamicFeeTx{
			To:   &token,
			Data: data,
			Gas:  approveGasLimits[chain],
		}
		tx := types.NewTx(inner)
		return tx, nil
	}
	return nil, nil
}
