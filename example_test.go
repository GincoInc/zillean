package zillean

import (
	"fmt"
	"math/big"
)

const endpoint = "https://api.zilliqa.com"

func Example() {
	// generate a new account
	toZil := NewZillean(endpoint)
	toPrivKey := toZil.GeneratePrivateKey()
	toPubKey, _ := toZil.GetPublicKeyFromPrivateKey(toPrivKey)
	toAddress, _ := toZil.GetAddressFromPrivateKey(toPrivKey)
	fmt.Printf("private key: %s\n", toPrivKey)
	fmt.Printf("public key: %s\n", toPubKey)
	fmt.Printf("address: %s\n", toAddress)

	// use an existing account
	fromZil := NewZillean(endpoint)
	fromPrivKey := "AAFD338492962FAD674EE3BD6EBC57C8373B2C9BADBAC8806D890F1FE8C571DF"
	fromPubKey, _ := fromZil.GetPublicKeyFromPrivateKey(fromPrivKey)

	// create a transaction
	rawTx := RawTransaction{
		Version:  0,
		Nonce:    1,
		To:       toAddress,
		Amount:   "1000000000000",
		PubKey:   fromPubKey,
		GasPrice: big.NewInt(1000000000),
		GasLimit: 1,
	}
	k, _ := GenerateDRN(EncodeTransaction(rawTx))
	signature, _ := fromZil.SignTransaction(k, rawTx, fromPrivKey)
	txID, _ := fromZil.RPC.CreateTransaction(rawTx, signature)
	fmt.Printf("txID: %s\n", txID)
}
