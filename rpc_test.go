package zillean

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	address        = "0x7f8A9ED7bA55A092A74105E5bD0Ec9C98e66051d"
	invalidAddress = "InvalidAddress"
)

func TestNewRPC(t *testing.T) {
	Convey("returns a new rpc", t, func() {
		So(NewRPC(localNet), ShouldHaveSameTypeAs, &RPC{})
	})
}

func TestRPC_GetBalance(t *testing.T) {
	Convey("returns the balance and nonce of a given address", t, func() {
		result, err := NewRPC(testNet).GetBalance("2AF379FF56ABD7432D9C74E4D7B95D1BE2F10C1A")
		So(err, ShouldBeNil)
		So(result.Balance, ShouldEqual, 1000)
		So(result.Nonce, ShouldEqual, 0)
	})
}

func TestRPC_GetDsBlock(t *testing.T) {
	Convey("returns details of a Directory Service block by block number", t, func() {
		result, err := NewRPC(testNet).GetDsBlock("1")
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldEqual, "1")
		So(result.Header.Difficulty, ShouldEqual, 3)
		So(result.Header.DifficultyDS, ShouldEqual, 5)
		So(result.Header.GasPrice, ShouldEqual, "100")
		So(result.Header.LeaderPubKey, ShouldEqual, "0x0200F1C6755F97A1E7AAD3EBDAA8D8B701EA45F9F55F1B58ED69DAFB0BE03DB938")
		So(result.Header.PoWWinners, ShouldResemble, []string{"0x03C3C59382A711549F4F74913A87A28397A5D8ED6A33A45AB6D89F744985582806", "0x03C851127068F19DA1D62EF32A8D4BE9382C517BE59ABFA513FECFC88D3997A48B", "0x03CD5CAE3E0ED85B83150DD8F1F8AC113835A447779E7E5D0AB1772D3FF8AD2501", "0x03CEEC73EBC07B714D1CD6A3CF036B7D87B706E076A6EADD633E9F4A6155661316", "0x03D9D0F93842A04D9AE8FC307C591325019008707A325D394E45FE3D517700DCB6"})
		So(result.Header.Prevhash, ShouldEqual, "54cd7c703b55f4c330ebf5d32708671f9d8aa89fb2b3edf0bac82353cd7c0866")
		So(result.Header.Timestamp, ShouldEqual, "1543852169974761")
		So(result.Signature, ShouldEqual, "F2AE964A3F67FE1E9F5B86FBA52ADBD09716E5E4ACA40CECFF8732E6670FE8E52CE0156450164981A9BFDFCF15797AE9C616E838BEBC6BE31201378859960B09")
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(testNet).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "87A2C5F8720C29CBDDE559350C5FD7EF6953E3E4522B46B3C81C3BB51A9AE0B03C28049D0D36C0FDE8008116F1BFA5B3966F1BE36D175148137D56FBDB037597")
		So(result.Body.MicroBlockInfos, ShouldNotBeEmpty)
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "2")
		So(result.Header.GasLimit, ShouldEqual, "2000000")
		So(result.Header.GasUsed, ShouldEqual, "5500")
		So(result.Header.MbInfoHash, ShouldEqual, "89fc73bfd43dafe7fb51b47fe78d70cc3038bfccd741317f1f28e47f83c69cf9")
		So(result.Header.MinerPubKey, ShouldEqual, "0x0206612F55DEDE5AF2A41CE96ECF37DDEBF8A8A05EA27BA71DAC1D612EE776E9F2")
		So(result.Header.NumMicroBlocks, ShouldEqual, 4)
		So(result.Header.NumTxns, ShouldEqual, 5500)
		So(result.Header.PrevBlockHash, ShouldEqual, "cde180c8c85843a0d6089dfb4d15a4b68071d198f1a8a8aaf072cdaa64025605")
		So(result.Header.Rewards, ShouldEqual, "550000")
		So(result.Header.StateDeltaHash, ShouldEqual, "12568d2307efc3b9e9410e7a439988bd2945e7f23b2049d9b81e7ed8ca0a4084")
		So(result.Header.StateRootHash, ShouldEqual, "d4aef6b5628bb0560332369319e6225dc3270ac462725066e313fa4e6f23fd68")
		So(result.Header.Timestamp, ShouldEqual, "1543859352905765")
		So(result.Header.Type, ShouldEqual, 1)
		So(result.Header.Version, ShouldEqual, 0)
	})
}

func TestRPC_GetLatestDsBlock(t *testing.T) {
	Convey("returns details of the most recent Directory Service block", t, func() {
		result, err := NewRPC(testNet).GetLatestDsBlock()
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.Difficulty, ShouldBeGreaterThan, 0)
		So(result.Header.LeaderPubKey, ShouldNotBeBlank)
		So(result.Header.Prevhash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Signature, ShouldNotBeBlank)
	})
}

func TestRPC_GetLatestTxBlock(t *testing.T) {
	Convey("returns details of the most recent Transaction block", t, func() {
		result, err := NewRPC(testNet).GetLatestTxBlock()
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldNotBeBlank)
		//So(result.Body.MicroBlockEmpty, ShouldNotBeEmpty)
		//So(result.Body.MicroBlockHashes, ShouldNotBeEmpty)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.DsBlockNum, ShouldNotBeBlank)
		So(result.Header.GasLimit, ShouldNotBeBlank)
		So(result.Header.GasUsed, ShouldNotBeBlank)
		So(result.Header.MbInfoHash, ShouldNotBeBlank)
		So(result.Header.MinerPubKey, ShouldNotBeBlank)
		So(result.Header.NumMicroBlocks, ShouldBeGreaterThan, 0)
		So(result.Header.NumTxns, ShouldHaveSameTypeAs, int64(0))
		So(result.Header.PrevBlockHash, ShouldNotBeBlank)
		So(result.Header.Rewards, ShouldNotBeBlank)
		So(result.Header.StateDeltaHash, ShouldNotBeBlank)
		So(result.Header.StateRootHash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Header.Type, ShouldHaveSameTypeAs, int64(1))
		So(result.Header.Version, ShouldHaveSameTypeAs, int64(0))
	})
}

func TestRPC_GetTransaction(t *testing.T) {
	Convey("returns details of a Transaction by its hash", t, func() {
		result, err := NewRPC(testNet).GetTransaction("a379c65e7373aec44e8c70f80891e135610d1e220329bce790e6d60b0e796854")
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, "a379c65e7373aec44e8c70f80891e135610d1e220329bce790e6d60b0e796854")
		So(result.Amount, ShouldEqual, "1")
		So(result.GasLimit, ShouldEqual, "1000")
		So(result.GasPrice, ShouldEqual, "100")
		So(result.Nonce, ShouldEqual, "2")
		So(result.Receipt.CumulativeGas, ShouldEqual, "1")
		So(result.Receipt.Success, ShouldBeTrue)
		So(result.SenderPubKey, ShouldEqual, "0x02B12D89A6854DD4FAACB368E3580389554040C3CD6E2FCADD4A152B66228E6D0D")
		So(result.Signature, ShouldEqual, "0x74925271BA02F9106F52287CAD89D06B58C9003AA0F03741B5D881FC9E4C79B000F3CD1E529B6B219D4FE6BDCF47F43850526C2ADBD73DB61A30684D1D762B45")
		So(result.ToAddr, ShouldEqual, "88bb4def5d6989706b2f72858d6e5cbcd0331b93")
		So(result.Version, ShouldEqual, "0")
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		zillean := NewZillean(testNet)
		privateKey := "79C4793303CDC5C98A9086AA39BDCA60C4140A4B8BE29897781931F38FB5001C"
		publicKey, _ := zillean.GetPublicKeyFromPrivateKey(privateKey)
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    1,
			To:       "FE90767E34BB8E0D33E9B98529FA34F89280B078",
			Amount:   "1",
			PubKey:   publicKey,
			GasPrice: 100,
			GasLimit: 100,
		}
		k, _ := generateDRN(encodeTransaction(rawTx))
		signature, _ := zillean.SignTransaction(k, rawTx, privateKey)
		result, err := zillean.RPC.CreateTransaction(rawTx, signature)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "adf12a29a86d62c7036253b22f0b6b1d9956fd3171444a578e0532bb04f9b498")
	})
}

func TestRPC_GetSmartContracts(t *testing.T) {
	Convey("returns the list of smart contracts created by an address", t, func() {
		// TODO create smart contract for testing use
		result, err := NewRPC(testNetScilla).GetSmartContracts("1D3FE113A0362BA2D63BF0BF41AFCA5A9921AB52")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 2)
	})
}

func TestRPC_GetSmartContractState(t *testing.T) {
	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
		// TODO create smart contract for testing use
		result, err := NewRPC(testNetScilla).GetSmartContractState("dbe59ad379c07b3f50187fb91e8472a34fa4a33f")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 2)
	})
}

func TestRPC_GetSmartContractCode(t *testing.T) {
	Convey("returns the Scilla code of a smart contract address", t, func() {
		// TODO create smart contract for testing use
		result, err := NewRPC(testNetScilla).GetSmartContractCode("dbe59ad379c07b3f50187fb91e8472a34fa4a33f")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "scilla_version 0\n\n    (* HelloWorld contract *)\n\n    import ListUtils\n\n    (***************************************************)\n    (*               Associated library                *)\n    (***************************************************)\n    library HelloWorld\n\n    let one_msg = \n      fun (msg : Message) => \n      let nil_msg = Nil {Message} in\n      Cons {Message} msg nil_msg\n\n    let not_owner_code = Int32 1\n    let set_hello_code = Int32 2\n\n    (***************************************************)\n    (*             The contract definition             *)\n    (***************************************************)\n\n    contract HelloWorld\n    (owner: ByStr20)\n\n    field welcome_msg : String = \"\"\n\n    transition setHello (msg : String)\n      is_owner = builtin eq owner _sender;\n      match is_owner with\n      | False =>\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : not_owner_code};\n        msgs = one_msg msg;\n        send msgs\n      | True =>\n        welcome_msg := msg;\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : set_hello_code};\n        msgs = one_msg msg;\n        send msgs\n      end\n    end\n\n\n    transition getHello ()\n        r <- welcome_msg;\n        e = {_eventname: \"getHello()\"; msg: r};\n        event e\n    end")
	})
}

func TestRPC_GetSmartContractInit(t *testing.T) {
	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
		// TODO create smart contract for testing use
		result, err := NewRPC(testNetScilla).GetSmartContractInit("dbe59ad379c07b3f50187fb91e8472a34fa4a33f")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 3)
	})
}

func TestRPC_GetBlockchainInfo(t *testing.T) {
	Convey("returns statistics about the specified zilliqa node", t, func() {
		result, err := NewRPC(testNet).GetBlockchainInfo()
		So(err, ShouldBeNil)
		So(result.CurrentDSEpoch, ShouldNotBeBlank)
		So(result.CurrentMiniEpoch, ShouldNotBeBlank)
		So(result.DSBlockRate, ShouldBeGreaterThan, 0)
		So(result.NumDSBlocks, ShouldNotBeBlank)
		So(result.NumPeers, ShouldBeGreaterThan, 0)
		So(result.NumTransactions, ShouldNotBeBlank)
		So(result.NumTxBlocks, ShouldNotBeBlank)
		So(result.NumTxnsDSEpoch, ShouldNotBeBlank)
		So(result.NumTxnsTxEpoch, ShouldHaveSameTypeAs, int64(0))
		So(result.CurrentDSEpoch, ShouldNotBeBlank)
		So(result.ShardingStructure.NumPeers, ShouldHaveSameTypeAs, []int64{})
		So(len(result.ShardingStructure.NumPeers), ShouldBeGreaterThan, 0)
		So(result.TransactionRate, ShouldHaveSameTypeAs, int64(0))
		So(result.TxBlockRate, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetNetworkID(t *testing.T) {
	Convey("returns the network ID of the specified zilliqa node", t, func() {
		result, err := NewRPC(testNet).GetNetworkID()
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "TestNet")
	})
}

func TestRPC_GetRecentTransactions(t *testing.T) {
	Convey("returns  the most recent transactions (upto 100) accepted by the specified zilliqa node.", t, func() {
		result, err := NewRPC(testNet).GetRecentTransactions()
		So(err, ShouldBeNil)
		So(len(result.TxnHashes), ShouldEqual, result.Number)
	})
}

func TestRPC_GetDSBlockListing(t *testing.T) {
	Convey("returns a paginated list of Directory Service blocks", t, func() {
		result, err := NewRPC(testNet).DSBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetTxBlockListing(t *testing.T) {
	Convey("returns a paginated list of Transaction blocks", t, func() {
		result, err := NewRPC(testNet).TxBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}
