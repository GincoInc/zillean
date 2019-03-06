package zillean

import (
	"encoding/hex"
	"fmt"
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

func TestECSchnorr_TrySign(t *testing.T) {
	Convey("returns the signature on a given message", t, func() {
		ecs := NewECSchnorr()
		privKey, _ := hex.DecodeString("B7139607427E6A03436469806FC1167ECEA26130736BDE063A4EED01036DBF03")
		pubKey, _ := hex.DecodeString("02892A6380826988CC46F317310D09F3BAB838B9D8C2407775F20F6AB8BD2A9FFF")
		k, _ := hex.DecodeString("af4ff508ac35fc3f3f66e0745b64dbac9068ce6d023deb4de69173fe50ed2b7d")
		msg, _ := hex.DecodeString("088180b40a10011a14df4b175c78e16eebc05173e5c1f87355622d810422230a2102892a6380826988cc46f317310d09f3bab838b9d8c2407775f20f6ab8bd2a9fff2a120a100000000000000000000000e8d4a5100032120a100000000000000000000000003b9aca003801")
		r, s, err := ecs.trySign(privKey, pubKey, k, msg)
		So(err, ShouldBeNil)
		So(fmt.Sprintf("%x", r), ShouldEqual, "c40bb55b911d0fb3aed13c9c75b560324a6586e422d7993a269dfd2eb96ee41a")
		So(fmt.Sprintf("%x", s), ShouldEqual, "a0eea37e14b690ca8d896a2b5453027beff2f373fbcecd34179b5d098f671281")
	})
}

func TestECSchnorr_Sign_And_Verify(t *testing.T) {
	Convey("sign and validate the signature", t, func() {
		ecs := NewECSchnorr()
		privKey := ecs.GeneratePrivateKey()
		pubKey := ecs.GetPublicKey(privKey, false)
		msg := []byte("message")
		r, s := ecs.Sign(privKey, pubKey, msg)
		So(ecs.Verify(r, s, pubKey, msg), ShouldBeTrue)
	})
}
