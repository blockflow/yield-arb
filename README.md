# yield-arb

## Overview

An API written in Go that allows for the execution of yield arbitrage strategies across multiple blockchains. Aggregates deposit and withdrawal opportunities across multiple protocols and chains to find the most profitable action sets. Calculates effective APYs/cost and adjusts for risk tolerance, base collateral, and liquidity.

```
Example: Bob has $10,000 USDT. The API found two profitable strategies.

// Total APY = 5%, Gas Cost = $0.05
1. Supply $10,000 USDT to AaveV3 on Arbitrum for 5% APY.

// Total APY = 7.4%, Gas Cost = $15.00
1. Supply $10,000 USDT to AaveV3 on Arbitrum for 5% APY.
2. Borrow $8,000 WETH from AaveV3 on Arbitrum for 3% APY.
3. Bridge WETH to Ethereum.
4. Supply $8,000 WETH to CompoundV3 on Ethereum for 6% APY.
```

Deposits and withdrawals are separated by design to allow for deposit only protocols (ex. staking). Uses chi framework for API. Bridge integration in works.

## Usage

```
cd cmd/
go run .
```

## API

### GET /strats

Returns all strategies sorted by APY in descending order.

### POST /transactions

Returns all transactions required to execute the listed strategies.

## Packages

### protocols

Contains all the blockchain protocol specific logic. Each should conform to the same interface as specified in the main file.

### arbitrage

Contains all the strategy logic. Finds profitable opportunities and sends the transactions necessary to execute them.

### utils

Contains all the helper functions and constants. Should not perform any actions.

### accounts

[DEPRECATED in favor of frontend wallet integration.] Contains all the logic relating to wallet management. Provides signers and other wallet interactions required for sending transactions.

## Add a new protocol

1. Create {PROTOCOL} subdirectory in `protocols/`
2. Add in `{PROTOCOL}.go` and other files
3. Import and add to switch statement in `protocols/protocols.go`

## Add a new chain to protocol

1. Create {CHAIN_NAME}.json and place in `utils/configs/`
2. Update protocol(s) to include chain

## Add a new contract

1. Fetch ABI from etherscan or associated block explorer and put in `protocols/{PROTOCOL}/`
2. If needed, build abigen tool (https://geth.ethereum.org/docs/tools/abigen) and place in `protocols/`
3. Generate bindings into `protocols/bindings/` with `./abigen --abi cmd/protocols/PROTOCOL/CONTRACT.json --pkg PROTOCOL --type CONTRACT --out cmd/protocols/PROTOCOL/CONTRACT.go`

## Testing

1. Run a fork with `anvil -f <URL>`
2. Wrap ETH into WETH on arbiscan (point MetaMask to local fork)
3. Change RPC URL for chain temporarily in config
4. Happy testing!

## Notes

Modules -> Packages -> Files

`nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go`

For Windows:

`nodemon --watch './**/*.go' -e go,json --signal SIGKILL --exec go run .`

`go mod init {MODULE_NAME}` to create project dependency file
`go get .` to automatically track dependencies (based on imports)
`go run .` to run server
`go mode tidy` to remove unused dependencies

Uses Multicall to reduce RPC calls. (https://github.com/mds1/multicall)

If need to retrieve non-view function output, just set the stateMutability in the json to view.

Useful GETH repo: https://github.com/everFinance/goether/tree/main

EIP712 TypedData uses lower cases for type names.
