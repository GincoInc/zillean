package zillean

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	crypto "github.com/GincoInc/go-crypto"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	localNet = "http://127.0.0.1:4200"
	testNet  = "https://api.zilliqa.com"
	// testNet = "https://testnet-n-api.aws.zilliqa.com"
	// testNet = "https://api-scilla.zilliqa.com"
)

func TestZillean_GeneratePrivateKey(t *testing.T) {
	Convey("returns the new private key from random seed", t, func() {
		result := NewZillean(localNet).GeneratePrivateKey()
		So(result, ShouldHaveLength, 64)
		So(result, ShouldHaveSameTypeAs, "string")
	})
}

func TestZillean_VerifyPrivateKey(t *testing.T) {
	Convey("returns true when the valid private key is given ", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(localNet).VerifyPrivateKey(vector.privateKey)
			So(err, ShouldBeNil)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when the invalid private key is given ", t, func() {
		result, err := NewZillean(localNet).VerifyPrivateKey("invalid private key")
		So(err.Error(), ShouldStartWith, "encoding/hex")
		So(result, ShouldBeFalse)
		result, err = NewZillean(localNet).VerifyPrivateKey("0000000000000000000000000000000000000000000000000000000000000000")
		So(result, ShouldBeFalse)
		result, err = NewZillean(localNet).VerifyPrivateKey("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
		So(result, ShouldBeFalse)
	})
}

func TestZillean_GetPublicKeyFromPrivateKey(t *testing.T) {
	Convey("returns the public key from the private key", t, func() {
		for _, vector := range testVectors[:1] {
			result, err := NewZillean(localNet).GetPublicKeyFromPrivateKey(vector.privateKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.publicKey)
		}
	})
}

func TestZillean_IsPublicKey(t *testing.T) {
	Convey("returns true when the valid public key is given ", t, func() {
		for _, vector := range testVectors {
			result := NewZillean(localNet).IsPublicKey(vector.publicKey)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when the invalid public key is given ", t, func() {
		result := NewZillean(localNet).IsPublicKey("invalid public key")
		So(result, ShouldBeFalse)
	})
}

func TestZillean_GetAddressFromPrivateKey(t *testing.T) {
	Convey("returns the address from the private key", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(localNet).GetAddressFromPrivateKey(vector.privateKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.address)
		}
	})
}

func TestZillean_GetAddressFromPublicKey(t *testing.T) {
	Convey("returns the address from the public key", t, func() {
		for _, vector := range testVectors {
			result, err := NewZillean(localNet).GetAddressFromPublicKey(vector.publicKey)
			So(err, ShouldBeNil)
			So(result, ShouldEqual, vector.address)
		}
	})
}

func TestZillean_IsAddress(t *testing.T) {
	Convey("returns true when the valid address is given ", t, func() {
		for _, vector := range testVectors {
			result := NewZillean(localNet).IsAddress(vector.address)
			So(result, ShouldBeTrue)
		}
	})

	Convey("returns false when the invalid address is given ", t, func() {
		result := NewZillean(localNet).IsAddress("invalid address")
		So(result, ShouldBeFalse)
	})
}

func TestZillean_SignTransaction(t *testing.T) {
	Convey("returns the signature", t, func() {
		privateKey := "79C4793303CDC5C98A9086AA39BDCA60C4140A4B8BE29897781931F38FB5001C"
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    1,
			To:       "FE90767E34BB8E0D33E9B98529FA34F89280B078",
			Amount:   "1",
			PubKey:   "03AD5893983179A55C466D94995DE934140EF3CB610526AEDFAC214DB7EC8E0946",
			GasPrice: big.NewInt(100),
			GasLimit: 100,
		}
		signature, err := NewZillean(localNet).SignTransaction(rawTx, privateKey)
		So(err, ShouldBeNil)
		So(signature, ShouldHaveLength, 128)
	})
}

func TestEncodeTransaction(t *testing.T) {
	Convey("returns the encoded transaction", t, func() {
		rawTx := RawTransaction{
			Version:  10,
			Nonce:    16,
			To:       "FE90767E34BB8E0D33E9B98529FA34F89280B078",
			Amount:   "100",
			PubKey:   "03AD5893983179A55C466D94995DE934140EF3CB610526AEDFAC214DB7EC8E0946",
			GasPrice: big.NewInt(88),
			GasLimit: 888,
			Code:     "aiueo",
			Data:     "abcde",
		}
		encodedTx, _ := hex.DecodeString("080a10101a14fe90767e34bb8e0d33e9b98529fa34f89280b07822230a2103ad5893983179a55c466d94995de934140ef3cb610526aedfac214db7ec8e09462a120a100000000000000000000000000000006432120a100000000000000000000000000000005838f8064205616975656f4a056162636465")
		So(encodeTransaction(rawTx), ShouldResemble, encodedTx)
	})
}

func TestHash(t *testing.T) {
	Convey("returns the hash", t, func() {
		k, _ := hex.DecodeString("eb449eb275abeaf7accce6fd5bb54d0e5b8500d7a9eb25d1e298facda2ed25ac")
		Qx, Qy := crypto.Secp256k1().ScalarBaseMult(k)
		Q := crypto.Compress(crypto.Secp256k1(), Qx, Qy)
		pubKey, _ := hex.DecodeString("04163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299d4cf0efcf07e96e576732dae122b9a8ac142214a6bc133b77aa5b79ba46b3e20")
		encodedTx, _ := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000080e3e927f8be54eb20f4f47baa2f4d2364943359104163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299d4cf0efcf07e96e576732dae122b9a8ac142214a6bc133b77aa5b79ba46b3e200000000000000000000000000000000000000000000000000000000000000378000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000580000000000000000")
		So(fmt.Sprintf("%x", hash(Q, pubKey, encodedTx)), ShouldEqual, "4664d452d23a069d558aece56a00a9a20cbb1ca2d93e886cd706e8f6aee016df")
	})
}
