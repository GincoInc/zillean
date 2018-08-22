# Zilliqa Golang SDK
A Golang implementation of [Zilliqa](https://github.com/Zilliqa/Zilliqa) API.

The project is still under development. Please note that specifications may change greatly.

## Installation
```sh
go get -u github.com/GincoInc/zillean
```

## Getting started
```go
package main

import (
	"fmt"

	"github.com/GincoInc/zillean"
)

func main() {
  // initialize the Zillean
  zillean := zillean.NewZillean("https://api-scilla.zilliqa.com")
  
  // generate a private key
  privKey, _ := zillean.GeneratePrivateKey()
  fmt.Printf("private key: %s\n", privKey) // private key: b6c00064b10d33c4a9fadb5b473d834b1995f132acdbe4b831ab5343702c174e
  
  // get a public key
  pubKey, _ := zillean.GetPublicKeyFromPrivateKey(privKey)
  fmt.Printf("public key: %s\n", pubKey) // public key: 03dcb21aaaa918f91a708858dc271343b4bee059e53202ce0358b68effa7e64378
  
  // get a address
  address, _ := zillean.GetAddressFromPrivateKey(privKey)
  fmt.Printf("address: %s\n", address) // address: 5f0e26adf701bb6a4535f0485fe3400e6e90c9ae
}
```

## Supports
### Wallet API
- [x] GeneratePrivateKey
- [x] VerifyPrivateKey
- [x] GetPublicKeyFromPrivateKey
- [x] IsPublicjKey
- [x] GetAddressFromPrivateKey
- [x] GetAddressFromPublicKey
- [x] IsAddress
- [ ] CreateTransactionJSON

### JSON-RPC API
- [x] GetBalance
- [x] GetDsBlock
- [x] GetTxBlock
- [x] GetLatestDsBlock
- [x] GetLatestTxBlock
- [x] GetTransaction
- [ ] CreateTransaction
- [ ] GetSmartContracts
- [ ] GetSmartContractState
- [ ] GetSmartContractCode
- [ ] GetSmartContractInit
- [x] GetBlockchainInfo
- [x] GetNetworkId
- [x] GetRecentTransactions
- [x] GetDSBlockListing 
- [x] GetTxBlockListing 
