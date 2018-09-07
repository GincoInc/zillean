package zillean

import (
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestECSchnorr_GeneratePrivateKey(t *testing.T) {
	Convey("returns the new private key from random seed", t, func() {
		result := NewECSchnorr().GeneratePrivateKey()
		So(result, ShouldHaveLength, 32)
		So(result, ShouldHaveSameTypeAs, []uint8{})
	})
}

func TestECSchnorr_GetPublicKey(t *testing.T) {
	Convey("returns the public key from the private key", t, func() {
		for _, vector := range testVectors {
			privKey, _ := hex.DecodeString(vector.privateKey)
			pubKey, _ := hex.DecodeString(vector.publicKey)
			So(NewECSchnorr().GetPublicKey(privKey, true), ShouldResemble, pubKey)
		}
	})
}

func TestECSchnorr_Sign_And_Verify(t *testing.T) {
	Convey("sign and validate the signature", t, func() {
		ecs := NewECSchnorr()
		privKey := ecs.GeneratePrivateKey()
		pubKey := ecs.GetPublicKey(privKey, false)
		msg := []byte("message")
		r, s, err := ecs.Sign(privKey, pubKey, []byte{0x01}, msg)
		So(err, ShouldBeNil)
		So(ecs.Verify(r, s, pubKey, msg), ShouldBeTrue)
	})
}
