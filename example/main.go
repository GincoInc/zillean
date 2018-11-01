package main

import (
	"fmt"

	"github.com/GincoInc/zillean"
)

func main() {
	// initialize the Zillean
	zillean := zillean.NewZillean("http://127.0.0.1:4200")

	// generate a private key
	privKey := zillean.GeneratePrivateKey()
	fmt.Printf("private key: %s\n", privKey)

	// get a public key
	pubKey, _ := zillean.GetPublicKeyFromPrivateKey(privKey)
	fmt.Printf("public key: %s\n", pubKey)

	// get a address
	address, _ := zillean.GetAddressFromPrivateKey(privKey)
	fmt.Printf("address: %s\n", address)
}
