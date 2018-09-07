package zillean

import (
	"encoding/hex"
	"fmt"
	"testing"

	crypto "github.com/GincoInc/go-crypto"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	localNet = "http://127.0.0.1:4200"
	testNet  = "https://scilla-test-api.aws.z7a.xyz"
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

func TestZillean_SignTransaction_And_Verify(t *testing.T) {
	Convey("returns the signature", t, func() {
		privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
		rawTx := RawTransaction{
			Version:  8,
			Nonce:    8,
			To:       "0e3e927f8be54eb20f4f47baa2f4d23649433591",
			Amount:   "378",
			PubKey:   "04163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299d4cf0efcf07e96e576732dae122b9a8ac142214a6bc133b77aa5b79ba46b3e20",
			GasPrice: 8,
			GasLimit: 88,
		}

		signature, err := NewZillean(localNet).SignTransaction(rawTx, privateKey)
		So(err, ShouldBeNil)
		So(signature[:64], ShouldEqual, "4664d452d23a069d558aece56a00a9a20cbb1ca2d93e886cd706e8f6aee016df")
		So(signature[64:], ShouldEqual, "e80f47b43e6711c2495fbe618cc86bf045bc93d0a1c107ed89a7d45da3451f59")
	})
}

func TestEncodeTransaction(t *testing.T) {
	Convey("returns the encoded transaction", t, func() {
		rawTx := RawTransaction{
			Version:  8,
			Nonce:    8,
			To:       "cdaf201ca7057f1135bef25c41dda40d68032a4f",
			Amount:   "378",
			PubKey:   "02076f5b8511a3ad45a4856681ab66c0b8a979f44640036e752231298ed75ad48e",
			GasPrice: 8,
			GasLimit: 88,
		}
		encodedTx, _ := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000008cdaf201ca7057f1135bef25c41dda40d68032a4f02076f5b8511a3ad45a4856681ab66c0b8a979f44640036e752231298ed75ad48e0000000000000000000000000000000000000000000000000000000000000378000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000580000000000000000")
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
