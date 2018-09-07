package zillean

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"

	crypto "github.com/GincoInc/go-crypto"
)

// See https://docs.zilliqa.com/whitepaper.pdf

// ECSchnorr ...
type ECSchnorr struct {
	Curve elliptic.Curve
}

// NewECSchnorr returns a new ECSchnorr.
func NewECSchnorr() *ECSchnorr {
	return &ECSchnorr{
		Curve: crypto.Secp256k1(),
	}
}

// GeneratePrivateKey ...
func (ecs *ECSchnorr) GeneratePrivateKey() []byte {
	b := make([]byte, 32)
	rand.Read(b)
	return b
}

// GetPublicKey ...
func (ecs *ECSchnorr) GetPublicKey(privKey []byte, compress bool) []byte {
	pubx, puby := ecs.Curve.ScalarBaseMult(privKey)

	return crypto.Marshal(ecs.Curve, pubx, puby, compress)
}

// Sign ...
func (ecs *ECSchnorr) Sign(privKey, pubKey, k, msg []byte) ([]byte, []byte, error) {
	_k := new(big.Int).SetBytes(k)
	if _k.Cmp(big.NewInt(0)) == 0 || _k.Cmp(ecs.Curve.Params().N) >= 0 {
		return nil, nil, errors.New("Invalid k")
	}

	Qx, Qy := ecs.Curve.ScalarBaseMult(k)
	Q := crypto.Compress(ecs.Curve, Qx, Qy)

	r := new(big.Int).SetBytes(hash(Q, pubKey, msg))
	if r.Cmp(big.NewInt(0)) == 0 || r.Cmp(ecs.Curve.Params().N) >= 0 {
		return nil, nil, errors.New("Invalid r")
	}

	sk := new(big.Int).SetBytes(privKey)
	_r := *r
	// k - r * sk mod n
	s := new(big.Int).Mod(_r.Sub(new(big.Int).SetBytes(k), _r.Mul(&_r, sk)), ecs.Curve.Params().N)
	if s.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("Invalid s")
	}

	return r.Bytes(), s.Bytes(), nil
}

// Verify ...
func (ecs *ECSchnorr) Verify(r, s, pubKey, msg []byte) bool {
	pkx, pky := elliptic.Unmarshal(ecs.Curve, pubKey)
	rpkx, rpky := ecs.Curve.ScalarMult(pkx, pky, r)
	sGx, sGy := ecs.Curve.ScalarBaseMult(s)
	Qx, Qy := ecs.Curve.Add(sGx, sGy, rpkx, rpky)
	Q := crypto.Compress(ecs.Curve, Qx, Qy)

	return bytes.Equal(r, hash(Q, pubKey, msg))
}
