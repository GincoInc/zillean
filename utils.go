package zillean

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	crypto "github.com/GincoInc/go-crypto"
	zillean "github.com/KazutakaNagata/zillean/proto"
	"github.com/golang/protobuf/proto"
)

func publicKeyToAddress(publicKey []byte) string {
	return fmt.Sprintf("%x", crypto.Sha256(publicKey)[12:])
}

func encodeTransaction(rawTx RawTransaction) []byte {
	version := uint32(rawTx.Version)
	nonce := uint64(rawTx.Nonce)
	toAddr, _ := hex.DecodeString(rawTx.To)
	_pubKey, _ := hex.DecodeString(rawTx.PubKey)
	pubKey := zillean.ByteArray{Data: _pubKey}
	_amount := &big.Int{}
	_amount.SetString(rawTx.Amount, 10)
	amount := zillean.ByteArray{Data: bigIntToPaddedBytes(_amount, 32)}
	gasPrice := zillean.ByteArray{Data: bigIntToPaddedBytes(rawTx.GasPrice, 32)}
	gasLimit := uint64(rawTx.GasLimit)
	code := []byte(rawTx.Code)
	data := []byte(rawTx.Data)

	protoTxCoreInfo := zillean.ProtoTransactionCoreInfo{
		Version:      &version,
		Nonce:        &nonce,
		Toaddr:       toAddr,
		Senderpubkey: &pubKey,
		Amount:       &amount,
		Gasprice:     &gasPrice,
		Gaslimit:     &gasLimit,
		Code:         code,
		Data:         data,
	}
	encodedTx, _ := proto.Marshal(&protoTxCoreInfo)

	return encodedTx
}

func bigIntToPaddedBytes(i *big.Int, paddedSize int32) []byte {
	bytes := i.Bytes()
	padded, _ := hex.DecodeString(fmt.Sprintf("%0*x", paddedSize, bytes))
	return padded
}

func hash(Q []byte, pubKey []byte, msg []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(Q)
	buffer.Write(pubKey[:33])
	buffer.Write(msg)

	return crypto.Sha256(buffer.Bytes())
}

func generateDRN(nonce []byte) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.Write(generateRandomBytes(32))
	buffer.WriteString("Schnorr+SHA256  ")
	hmacDRBG := crypto.NewHmacDRBG(generateRandomBytes(32), nonce, buffer.Bytes())

	return hmacDRBG.Generate(int32(32), []byte{})
}

func generateRandomBytes(size int32) []byte {
	randomBytes := make([]byte, size)
	_, _ = rand.Read(randomBytes)
	return randomBytes
}
