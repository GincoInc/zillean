package zillean

import (
	"math/big"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRPC(t *testing.T) {
	Convey("returns a new rpc", t, func() {
		So(NewRPC(localNet), ShouldHaveSameTypeAs, &RPC{})
	})
}

func TestRPC_GetBalance(t *testing.T) {
	Convey("returns the balance and nonce of a given address", t, func() {
		result, err := NewRPC(testNet).GetBalance("546c73019def014ff2e363c4bc97de9ef90354fa")
		So(err, ShouldBeNil)
		So(result.Balance, ShouldEqual, "1000000000000")
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
		So(result.Header.GasPrice, ShouldEqual, "1000000000")
		So(result.Header.LeaderPubKey, ShouldEqual, "0x0200783E01447F6F1C4927A6BC700B48D021E32F7B7A8498D1E8785FCDCB9EB992")
		So(result.Header.PoWWinners, ShouldResemble, []string{"0x036731768A27A87FD3C94ED8F5931671D3CC09D92583E2F90FE2273829752CA39C", "0x03682E483077DC721F0DF9BA9EF94666998F72C4F689369324FB67867CC921ABA5", "0x0369DAB66B131E282E871B0777D8B4DE835A8CC9A8076CD08779542D9C1F125431", "0x03728428C88C5783947F8924E91B1632DC2BF318F0A45AC99B7D0D0998C25C1F22", "0x03A142426951FCB3514AFA9E69AA08D4F4DBCD0FA01E5E337AFDAE91CB3DD02B7E", "0x03B0425A969E071DBE04CC0D551227D8D1A05B6B9BA5752BE24C9174AE575B5528", "0x03C758076B7042A658AB78B251D7AA71274C438A2D4B3775609E66D80A37806723", "0x03E71515427EAABC2C11E8A41056C5AE5CF3B2EC5BB69E031FF07A2A45D2BB9603", "0x03F3D05231E69B0992359AE6EDEC7EF1F45CF776C177D3A7FE9462FAB1C234E331", "0x03FB373D0198D1D346B3A3885DE446D71088409ABCB9EF095384AF0E219FD3D166"})
		So(result.Header.Prevhash, ShouldEqual, "ba127538d2c63eec121629011ae8173210589689dca54d1e11904dd82c68e9da")
		So(result.Header.Timestamp, ShouldEqual, "1546434439109004")
		So(result.Signature, ShouldEqual, "7E3932613899D5C7EB80A6A6F8B2F47E677746585CE7E02EE68EF9738BD52EB9AEE8A641DBF22B8F0E4CFC63B44EEF16D29AEA8242F01B46D834498616C59CCB")
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(testNet).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "C4AEDCF53CEEA81DF00DDF80039ADF430E51A95C148E6803B210ECF023456DC2AA1BBCAF5552F8D2AC6113DA85996C324E418A87EBD3237764B3BE36D9975E27")
		So(len(result.Body.MicroBlockInfos), ShouldEqual, 4)
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "2")
		So(result.Header.GasLimit, ShouldEqual, "2000000")
		So(result.Header.GasUsed, ShouldEqual, "0")
		So(result.Header.MbInfoHash, ShouldEqual, "316176bfdac13bf59a7426b6fdc204c4484df5fb68efbb98798ead1b0319c517")
		So(result.Header.MinerPubKey, ShouldEqual, "0x024715D33B7B90F28FD42062CA58F6917038361CA529DCEF9EB9CFABF147EE9FEF")
		So(result.Header.NumMicroBlocks, ShouldEqual, 4)
		So(result.Header.NumTxns, ShouldEqual, 0)
		So(result.Header.PrevBlockHash, ShouldEqual, "037a18de52e1703843120a089d91dded718e81e60fafb4dcc24f627e736b284f")
		So(result.Header.Rewards, ShouldEqual, "0")
		So(result.Header.StateDeltaHash, ShouldEqual, "0000000000000000000000000000000000000000000000000000000000000000")
		So(result.Header.StateRootHash, ShouldEqual, "3296a3424bf9eb173d50466ef432c2a32f8c72959103224e9c6030dcb88c5209")
		So(result.Header.Timestamp, ShouldEqual, "1546443697981005")
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
		So(result.Body.MicroBlockInfos, ShouldNotBeNil)
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
		result, err := NewRPC(testNet).GetTransaction("B25E05C30C53B2E2BEEB67B5AAD483069C18D1901B544FD63301DEC6516873DE")
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, "b25e05c30c53b2e2beeb67b5aad483069c18d1901b544fd63301dec6516873de")
		So(result.Amount, ShouldEqual, "1000000000000")
		So(result.GasLimit, ShouldEqual, "1")
		So(result.GasPrice, ShouldEqual, "1000000000")
		So(result.Nonce, ShouldEqual, "1")
		So(result.Receipt.CumulativeGas, ShouldEqual, "1")
		So(result.Receipt.EpochNum, ShouldEqual, "5128")
		So(result.Receipt.Success, ShouldBeTrue)
		So(result.SenderPubKey, ShouldEqual, "0x02892A6380826988CC46F317310D09F3BAB838B9D8C2407775F20F6AB8BD2A9FFF")
		So(result.Signature, ShouldEqual, "0x44509B5C1408B48268062580E74372106983B75C0E8E070086030E0F2D12D32DE94FC644D76F46D3BD4DA5F55FAF04397879CE32FF40A81B3B1FA43EC25B5C04")
		So(result.ToAddr, ShouldEqual, "546c73019def014ff2e363c4bc97de9ef90354fa")
		So(result.Version, ShouldEqual, "0")
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		zillean := NewZillean(testNet)
		privateKey := "AAFD338492962FAD674EE3BD6EBC57C8373B2C9BADBAC8806D890F1FE8C571DF"
		publicKey, _ := zillean.GetPublicKeyFromPrivateKey(privateKey)
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    1,
			To:       "5568CF7C38334A4E960BC99D8F22C1E90645E5F2",
			Amount:   "1000000000000",
			PubKey:   publicKey,
			GasPrice: big.NewInt(1000000000),
			GasLimit: 1,
		}
		k, _ := GenerateDRN(EncodeTransaction(rawTx))
		signature, _ := zillean.SignTransaction(k, rawTx, privateKey)
		result, err := zillean.RPC.CreateTransaction(rawTx, signature)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "5a09b08a28ca4b6ef1935bb3d00307530ad64e4150fba573cdb9f0dea847d1c7")
	})
}

func TestRPC_GetSmartContracts(t *testing.T) {
	Convey("returns the list of smart contracts created by an address", t, func() {
		result, err := NewRPC(testNet).GetSmartContracts("f3d2005b55102d6588dd9771e9356f1908c9d97f")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Address, ShouldEqual, "83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(result[0].State[0].Type, ShouldEqual, "Uint128")
		So(result[0].State[0].Value, ShouldEqual, "0")
		So(result[0].State[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetSmartContractState(t *testing.T) {
	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractState("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Type, ShouldEqual, "Uint128")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetSmartContractCode(t *testing.T) {
	Convey("returns the Scilla code of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractCode("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "scilla_version 0\n\n    (* HelloWorld contract *)\n    \n    import ListUtils\n    \n    (***************************************************)\n    (*               Associated library                *)\n    (***************************************************)\n    library HelloWorld\n    \n    let one_msg = \n      fun (msg : Message) => \n      let nil_msg = Nil {Message} in\n      Cons {Message} msg nil_msg\n    \n    let not_owner_code = Int32 1\n    let set_hello_code = Int32 2\n    \n    (***************************************************)\n    (*             The contract definition             *)\n    (***************************************************)\n    \n    contract HelloWorld\n    (owner: ByStr20)\n    \n    field welcome_msg : String = \"\"\n    \n    transition setHello (msg : String)\n      is_owner = builtin eq owner _sender;\n      match is_owner with\n      | False =>\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : not_owner_code};\n        msgs = one_msg msg;\n        send msgs\n      | True =>\n        welcome_msg := msg;\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : set_hello_code};\n        msgs = one_msg msg;\n        send msgs\n      end\n    end\n    \n    \n    transition getHello ()\n        r <- welcome_msg;\n        e = {_eventname: \"getHello()\"; msg: r};\n        event e\n    end")
	})
}

func TestRPC_GetSmartContractInit(t *testing.T) {
	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractInit("83536f90ed096b5d14ba2c296a32f37849dd3221")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 3)
		So(result[0].Type, ShouldEqual, "Uint32")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_scilla_version")
		So(result[1].Type, ShouldEqual, "ByStr20")
		So(result[1].Value, ShouldEqual, "0xf3d2005b55102d6588dd9771e9356f1908c9d97f")
		So(result[1].Vname, ShouldEqual, "owner")
		So(result[2].Type, ShouldEqual, "BNum")
		So(result[2].Value, ShouldEqual, "9247")
		So(result[2].Vname, ShouldEqual, "_creation_block")
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
