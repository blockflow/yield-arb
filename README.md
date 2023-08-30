# yield-arb

Lending and borrowing are separated by design to allow for deposit only protocols (ex. staking). Uses chi framework for API.

## Ideas

 - Create web-hosted front-end (Near BOS, react)
 - Show historical data (The Graph, airstack)
 - Integrate a wallet dApp (walletconnect, metamask)
 - Manual, user signs tx to execute strat
 - Bundle txs/signs using ERC4337 (biconomy)

## Packages

### protocols

Contains all the blockchain protocol specific logic. Each should conform to the same interface as specified in the main file.

### arbitrage

Contains all the strategy logic. Finds profitable opportunities and sends the transactions necessary to execute them.

### utils

Contains all the helper functions and constants. Should not perform any actions.

### accounts

Contains all the logic relating to wallet management. Provides signers and other wallet interactions required for sending transactions.

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
