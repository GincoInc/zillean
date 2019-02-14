package zillean

import (
	"fmt"
	"math/big"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRPC(t *testing.T) {
	Convey("returns a new rpc", t, func() {
		So(NewRPC(localNet), ShouldHaveSameTypeAs, &RPC{})
	})
}

func TestRPC_GetNetworkID(t *testing.T) {
	Convey("returns the network ID of the specified zilliqa node", t, func() {
		result, err := NewRPC(testNet).GetNetworkID()
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "333")
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

func TestRPC_GetShardingStructure(t *testing.T) {
	Convey("returns the current sharding structure of the network from the specified network's lookup node", t, func() {
		result, err := NewRPC(testNet).GetShardingStructure()
		So(err, ShouldBeNil)
		fmt.Println(result)
		So(len(result.NumPeers), ShouldBeGreaterThan, 0)
		So(result.NumPeers, ShouldHaveSameTypeAs, []int64{})
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
		So(result.Header.LeaderPubKey, ShouldEqual, "0x02081DCD3D93A4406E6D90241931A4D8A28553EC7BA28AB5B51D35D992CA2C7383")
		So(result.Header.PoWWinners, ShouldResemble, []string{"0x027409E2C105498DE346980A7BD917E93574D86CB3A13B3CE3C989B2E2A96D5A69"})
		So(result.Header.Prevhash, ShouldEqual, "0f00e9d3175300fc287812d201edcfbfcb8165809606545595bf53700c524648")
		So(result.Header.Timestamp, ShouldEqual, "1549265830654931")
		So(result.Signature, ShouldEqual, "CF8A45F50153BC860582DAFEEE074CC3D027DB94D839102BD777EF8AB1F5753163F2223E405B88F3A27D878465B4B9087BDB0D551C1EF954010FC99E8EB265A1")
	})
}

func TestRPC_GetLatestDsBlock(t *testing.T) {
	Convey("returns details of the most recent Directory Service block", t, func() {
		result, err := NewRPC(testNet).GetLatestDsBlock()
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.Difficulty, ShouldBeGreaterThan, 0)
		So(result.Header.DifficultyDS, ShouldBeGreaterThan, 0)
		So(result.Header.GasPrice, ShouldNotBeBlank)
		So(result.Header.LeaderPubKey, ShouldNotBeBlank)
		So(result.Header.Prevhash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Signature, ShouldNotBeBlank)
	})
}

func TestRPC_GetNumDSBlocks(t *testing.T) {
	Convey("returns the number of Directory Service blocks in the network so far. This is represented as a String", t, func() {
		result, err := NewRPC(testNet).GetNumDSBlocks()
		So(err, ShouldBeNil)
		So(result, ShouldNotBeBlank)
	})
}

func TestRPC_GetDSBlockRate(t *testing.T) {
	Convey("returns the current Directory Service blockrate per second", t, func() {
		result, err := NewRPC(testNet).GetDSBlockRate()
		So(err, ShouldBeNil)
		So(result, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_DSBlockListing(t *testing.T) {
	Convey("returns a paginated list of Directory Service blocks", t, func() {
		result, err := NewRPC(testNet).DSBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(testNet).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "07968762C6819E0D17B8761B31F68A26D4AA189547213B4D74E00A09A3B7EECCC4AF8D87C74DE9188A69F25A15D524781BDCD3CE0BE9C594E0D4DBFE00ABAC2A")
		So(len(result.Body.MicroBlockInfos), ShouldEqual, 4)
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "2")
		So(result.Header.GasLimit, ShouldEqual, "200000")
		So(result.Header.GasUsed, ShouldEqual, "0")
		So(result.Header.MbInfoHash, ShouldEqual, "db311f58e5c43b043f9143c0b8efd62ceaf98eaeedf58b8b8c5300f0df780da1")
		So(result.Header.MinerPubKey, ShouldEqual, "0x0238EA7FD93C9E0F30EB8F95BC2B22D7C998D76CFB1620172638B998A4BE01C5F0")
		So(result.Header.NumMicroBlocks, ShouldEqual, 4)
		So(result.Header.NumTxns, ShouldEqual, 0)
		So(result.Header.PrevBlockHash, ShouldEqual, "bb6ba0e008f272037c2fad24965a0b67380885ef1a853fe12d07714357a8f541")
		So(result.Header.Rewards, ShouldEqual, "0")
		So(result.Header.StateDeltaHash, ShouldEqual, "0000000000000000000000000000000000000000000000000000000000000000")
		So(result.Header.StateRootHash, ShouldEqual, "c9065f6fd1520e6ed6174a2ae4c587acdfbd7346fcd4419d483cb5bb7b343ef5")
		So(result.Header.Timestamp, ShouldEqual, "1549267096666600")
		So(result.Header.Version, ShouldEqual, 1)
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
		So(result.Header.Version, ShouldHaveSameTypeAs, int64(0))
	})
}

func TestRPC_GetNumTxBlocks(t *testing.T) {
	Convey("returns the number of Transaction blocks in the network so far, this is represented as String", t, func() {
		result, err := NewRPC(testNet).GetNumTxBlocks()
		So(err, ShouldBeNil)
		So(result, ShouldNotBeBlank)
	})
}

func TestRPC_GetTxBlockRate(t *testing.T) {
	Convey("returns the current Transaction blockrate per second", t, func() {
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

func TestRPC_GetNumTransactions(t *testing.T) {
	Convey("returns the number of Transactions validated in the network so far. This is represented as a String", t, func() {
	})
}

func TestRPC_GetTransactionRate(t *testing.T) {
	Convey("returns the current Transaction rate of the network", t, func() {
	})
}

func TestRPC_GetCurrentMiniEpoch(t *testing.T) {
	Convey("returns the number of TX epochs in the network so far represented as String", t, func() {
	})
}

func TestRPC_GetCurrentDSEpoch(t *testing.T) {
	Convey("returns the number of DS epochs in the network so far represented as String", t, func() {
	})
}

func TestRPC_GetPrevDifficulty(t *testing.T) {
	Convey("returns the minimum shard difficulty of the previous block, this is represented as an Number", t, func() {
	})
}

func TestRPC_GetPrevDSDifficulty(t *testing.T) {
	Convey("returns the minimum DS difficulty of the previous block, this is represented as an Number", t, func() {
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		zillean := NewZillean(testNet)
		privateKey := "B7139607427E6A03436469806FC1167ECEA26130736BDE063A4EED01036DBF03"
		publicKey, _ := zillean.GetPublicKeyFromPrivateKey(privateKey)
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    1,
			To:       "546c73019def014ff2e363c4bc97de9ef90354fa",
			Amount:   "1000000000000",
			PubKey:   publicKey,
			GasPrice: big.NewInt(1000000000),
			GasLimit: 1,
		}
		signature, _ := zillean.SignTransaction(rawTx, privateKey)
		result, err := zillean.RPC.CreateTransaction(rawTx, signature)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "b25e05c30c53b2e2beeb67b5aad483069c18d1901b544fd63301dec6516873de")
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

func TestRPC_GetRecentTransactions(t *testing.T) {
	Convey("returns  the most recent transactions (upto 100) accepted by the specified zilliqa node.", t, func() {
		result, err := NewRPC(testNet).GetRecentTransactions()
		So(err, ShouldBeNil)
		So(len(result.TxnHashes), ShouldEqual, result.Number)
	})
}

func TestRPC_GetTransactionsForTxBlock(t *testing.T) {
	Convey("returns the transactions included within a micro-block created by a specific shard", t, func() {
	})
}

func TestRPC_GetNumTxnsTxEpoch(t *testing.T) {
	Convey("returns the number of transactions in this Transaction epoch, this is represented as String", t, func() {
	})
}

func TestRPC_GetNumTxnsDSEpoch(t *testing.T) {
	Convey("returns the number of transactions in this Directory Service epoch, this is represented as String", t, func() {
	})
}

func TestRPC_GetMinimumGasPrice(t *testing.T) {
	Convey("returns the minimum gas price of the last DS epoch represented as String", t, func() {
	})
}

func TestRPC_GetSmartContractCode(t *testing.T) {
	Convey("returns the Scilla code of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractCode("2112d6eda5d539826e2a2f175d96a181791a34ab")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "scilla_version 0\n\n    (* HelloWorld contract *)\n    \n    import ListUtils\n    \n    (***************************************************)\n    (*               Associated library                *)\n    (***************************************************)\n    library HelloWorld\n    \n    let one_msg = \n      fun (msg : Message) => \n      let nil_msg = Nil {Message} in\n      Cons {Message} msg nil_msg\n    \n    let not_owner_code = Int32 1\n    let set_hello_code = Int32 2\n    \n    (***************************************************)\n    (*             The contract definition             *)\n    (***************************************************)\n    \n    contract HelloWorld\n    (owner: ByStr20)\n    \n    field welcome_msg : String = \"\"\n    \n    transition setHello (msg : String)\n      is_owner = builtin eq owner _sender;\n      match is_owner with\n      | False =>\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : not_owner_code};\n        msgs = one_msg msg;\n        send msgs\n      | True =>\n        welcome_msg := msg;\n        msg = {_tag : \"Main\"; _recipient : _sender; _amount : Uint128 0; code : set_hello_code};\n        msgs = one_msg msg;\n        send msgs\n      end\n    end\n    \n    \n    transition getHello ()\n        r <- welcome_msg;\n        e = {_eventname: \"getHello()\"; msg: r};\n        event e\n    end")
	})
}

func TestRPC_GetSmartContractInit(t *testing.T) {
	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractInit("2112d6eda5d539826e2a2f175d96a181791a34ab")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 3)
		So(result[0].Type, ShouldEqual, "Uint32")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_scilla_version")
		So(result[1].Type, ShouldEqual, "ByStr20")
		So(result[1].Value, ShouldEqual, "0xf49f1306bc8fb0cd8167a58a3550c1443072e96b")
		So(result[1].Vname, ShouldEqual, "owner")
		So(result[2].Type, ShouldEqual, "BNum")
		So(result[2].Value, ShouldEqual, "5143")
		So(result[2].Vname, ShouldEqual, "_creation_block")
	})
}

func TestRPC_GetSmartContractState(t *testing.T) {
	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
		result, err := NewRPC(testNet).GetSmartContractState("2112d6eda5d539826e2a2f175d96a181791a34ab")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Type, ShouldEqual, "Uint128")
		So(result[0].Value, ShouldEqual, "0")
		So(result[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetSmartContracts(t *testing.T) {
	Convey("returns the list of smart contracts created by an address", t, func() {
		result, err := NewRPC(testNet).GetSmartContracts("f49f1306bc8fb0cd8167a58a3550c1443072e96b")
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(result[0].Address, ShouldEqual, "2112d6eda5d539826e2a2f175d96a181791a34ab")
		So(result[0].State[0].Type, ShouldEqual, "Uint128")
		So(result[0].State[0].Value, ShouldEqual, "0")
		So(result[0].State[0].Vname, ShouldEqual, "_balance")
	})
}

func TestRPC_GetContractAddressFromTransactionID(t *testing.T) {
	Convey("returns a smart contract address of 20 bytes from a transaction ID, represented as a String", t, func() {
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
