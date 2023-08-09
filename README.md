# yield-arb

Lending and borrowing are separated by design to allow for deposit only protocols (ex. AMM).

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

## Notes

Modules -> Packages -> Files

`nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go`

For Windows:

`nodemon --watch './**/*.go' -e go,json --signal SIGKILL --exec go run .`

`go mod init {MODULE_NAME}` to create project dependency file
`go get .` to automatically track dependencies (based on imports)
`go run .` to run server
`go mode tidy` to remove unused dependencies