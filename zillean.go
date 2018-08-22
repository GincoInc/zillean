package zillean

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"

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

// VerifyPrivateKey ...
func (z *Zillean) VerifyPrivateKey(privateKey string) bool {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return false
	}

	if _, err := crypto.ToECDSA(privateKeyBytes); err != nil {
		return false
	}

	return true
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

// IsPublicKey ...
func (z *Zillean) IsPublicKey(publicKey string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{66}$`).MatchString(publicKey)
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

// IsAddress ...
func (z *Zillean) IsAddress(address string) bool {
	return regexp.MustCompile(`^[0-9a-fA-F]{40}$`).MatchString(address)
}
