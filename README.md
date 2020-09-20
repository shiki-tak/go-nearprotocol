# GO-NEARPROTOCOL

### Build
```
❯ make build-rust
```

### Run standalone
```
❯ go run cmd/main.go -api=standalone -wasm-file=./example-contract/target/wasm32-unknown-unknown/release/example_contract.wasm -input='{"name": "shiki"}' -method-name=greet
{"outcome":{"balance":"10000000000000000000000000","storage_usage":145,"return_data":{"Value":"\"Hello, shiki from alice\""},"burnt_gas":564640857849,"used_gas":564640857849,"logs":[]},"err":null,"receipts":[],"state":{"U1RBVEU=":""}}
```
