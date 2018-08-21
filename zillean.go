package zillean

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

// Zillean represents the zillean object.
type Zillean struct {
	RPC *RPC
}

// NewZillean returns a new zilliean.Zillean.
func NewZillean(endpoint string) *Zillean {
	return &Zillean{
		RPC: NewRPC(endpoint),
	}
}

// GeneratePrivateKey ...
func (z *Zillean) GeneratePrivateKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

// GetPublicKeyFromPrivateKey ...
func (z *Zillean) GetPublicKeyFromPrivateKey(privateKey string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", fromECDSAPub(true, &ecdsaPrivateKey.PublicKey)), nil
}

// GetAddressFromPrivateKey ...
func (z *Zillean) GetAddressFromPrivateKey(privateKey string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	ecdsaPrivateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", err
	}

	return publicKeyToAddress(fromECDSAPub(true, &ecdsaPrivateKey.PublicKey)), nil
}

// GetAddressFromPublicKey ...
func (z *Zillean) GetAddressFromPublicKey(publicKey string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	return publicKeyToAddress(publicKeyBytes), nil
}

// utils ///////////////////////////////////////////////////////////////////
//       ///////////////////////////////////////////////////////////////////
//       ///////////////////////////////////////////////////////////////////
//       ///////////////////////////////////////////////////////////////////

func publicKeyToAddress(publicKeyBytes []byte) string {
	return fmt.Sprintf("%x", hashSha256(publicKeyBytes)[12:])
}

func hashSha256(data []byte) []byte {
	sha := sha256.New()
	sha.Write(data)

	return sha.Sum(nil)
}

func fromECDSAPub(compress bool, pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}

	return marshal(compress, pub.X, pub.Y)
}

func marshal(compress bool, x, y *big.Int) []byte {
	byteLen := (crypto.S256().Params().BitSize + 7) >> 3

	if compress {
		ret := make([]byte, 1+byteLen)
		if y.Bit(0) == 0 {
			ret[0] = 2
		} else {
			ret[0] = 3
		}
		xBytes := x.Bytes()
		copy(ret[1+byteLen-len(xBytes):], xBytes)
		return ret
	}

	ret := make([]byte, 1+2*byteLen)
	ret[0] = 4 // uncompressed point
	xBytes := x.Bytes()
	copy(ret[1+byteLen-len(xBytes):], xBytes)
	yBytes := y.Bytes()
	copy(ret[1+2*byteLen-len(yBytes):], yBytes)
	return ret
}
