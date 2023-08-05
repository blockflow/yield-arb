# yield-arb

Lending and borrowing are separated by design to allow for deposit only protocols (ex. AMM).

## Add a new protocol

1. Create {protocol_name}.go file in `protocols/`
2. Add to switch statement in `protocols/protocols.go`

## Add a new chain to protocol

1. Create {CHAIN_NAME}.json and place in `utils/configs/`
2. Update protocol to include chain

## Notes

Modules -> Packages -> Files

`nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go`

For Windows:

`nodemon --watch './**/*.go' -e go,json --signal SIGKILL --exec go run .`

`go mod init {MODULE_NAME}` to create project dependency file
`go get .` to automatically track dependencies (based on imports)
`go run .` to run server
`go mode tidy` to remove unused dependencies