package zillean

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"

	crypto "github.com/GincoInc/go-crypto"
)

var (
	curve = crypto.Secp256k1()
)

// Zillean represents the zillean object.
type Zillean struct {
	ECS *ECSchnorr
	RPC *RPC
}

// NewZillean returns a new zilliean.Zillean.
func NewZillean(endpoint string) *Zillean {
	return &Zillean{
		ECS: NewECSchnorr(),
		RPC: NewRPC(endpoint),
	}
}

// GeneratePrivateKey ...
func (z *Zillean) GeneratePrivateKey() string {
	return fmt.Sprintf("%x", z.ECS.GeneratePrivateKey())
}

// VerifyPrivateKey ...
func (z *Zillean) VerifyPrivateKey(privateKey string) (bool, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return false, err
	}

	D := new(big.Int).SetBytes(privKey)

	// The D must < N
	if D.Cmp(z.ECS.Curve.Params().N) >= 0 {
		return false, errors.New("invalid private key, >= N")
	}

	// The D must not be zero or negative.
	if D.Sign() <= 0 {
		return false, errors.New("invalid private key, zero or negative")
	}

	pubx, _ := z.ECS.Curve.ScalarBaseMult(D.Bytes())
	if pubx == nil {
		return false, errors.New("invalid private key")
	}

	return true, nil
}

// GetPublicKeyFromPrivateKey ...
func (z *Zillean) GetPublicKeyFromPrivateKey(privateKey string) (string, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", z.ECS.GetPublicKey(privKey, true)), nil
}

// IsPublicKey ...
func (z *Zillean) IsPublicKey(publicKey string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{66}$`).MatchString(publicKey)
}

// GetAddressFromPrivateKey ...
func (z *Zillean) GetAddressFromPrivateKey(privateKey string) (string, error) {
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	return publicKeyToAddress(z.ECS.GetPublicKey(privKey, true)), nil
}

// GetAddressFromPublicKey ...
func (z *Zillean) GetAddressFromPublicKey(publicKey string) (string, error) {
	publicKeyBytes, err := hex.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	return publicKeyToAddress(publicKeyBytes), nil
}

// IsAddress ...
func (z *Zillean) IsAddress(address string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{40}$`).MatchString(address)
}

// SignTransaction ...
func (z *Zillean) SignTransaction(rawTx RawTransaction, privateKey string) (string, error) {
	privKey, _ := hex.DecodeString(privateKey)
	_privKey := make([]byte, len(privKey))
	copy(_privKey, privKey)
	k, _ := generateDRN(_privKey, encodeTransaction(rawTx))

	pubKey, _ := hex.DecodeString(rawTx.PubKey)
	r, s, err := z.ECS.Sign(privKey, pubKey, k, encodeTransaction(rawTx))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x%x", r, s), nil
}

// VerifySignature ...
func (z *Zillean) VerifySignature(r, s, publicKey, msg []byte) bool {
	return z.ECS.Verify(r, s, publicKey, msg)
}
