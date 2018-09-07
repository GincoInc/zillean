package zillean

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	crypto "github.com/GincoInc/go-crypto"
)

func publicKeyToAddress(publicKey []byte) string {
	return fmt.Sprintf("%x", crypto.Sha256(publicKey)[12:])
}

func encodeTransaction(rawTx RawTransaction) []byte {
	var buffer bytes.Buffer
	buffer.Write(int32ToPaddedBytes(rawTx.Version, 64))
	buffer.Write(int32ToPaddedBytes(rawTx.Nonce, 64))
	to, _ := hex.DecodeString(rawTx.To)
	buffer.Write(to)
	pubKey, _ := hex.DecodeString(rawTx.PubKey)
	buffer.Write(pubKey)
	amount, _ := hex.DecodeString(fmt.Sprintf("%064s", rawTx.Amount))
	buffer.Write(amount)
	buffer.Write(int32ToPaddedBytes(rawTx.GasPrice, 64))
	buffer.Write(int32ToPaddedBytes(rawTx.GasLimit, 64))
	buffer.Write(int32ToPaddedBytes(int32(len(rawTx.Code)), 8))
	code, _ := hex.DecodeString(rawTx.Code)
	buffer.Write(code)
	buffer.Write(int32ToPaddedBytes(int32(len(rawTx.Data)), 8))
	data, _ := hex.DecodeString(rawTx.Data)
	buffer.Write(data)

	return buffer.Bytes()
}

func int32ToPaddedBytes(i, paddedSize int32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(i))
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

func generateDRN(entropy, nonce []byte) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.Write(make([]byte, 32))
	buffer.WriteString("Schnorr+SHA256  ")
	hmacDRBG := crypto.NewHmacDRBG(entropy, nonce, buffer.Bytes())

	return hmacDRBG.Generate(int32(32), []byte{})
}
