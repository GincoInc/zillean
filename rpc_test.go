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
		So(result.Header.LeaderPubKey, ShouldEqual, "0x02094D8F49E1201B34901F0B6413BBA3686E93A8A5186C3A4C6D176A91EE784357")
		So(result.Header.MinerPubKey, ShouldEqual, "")
		So(result.Header.Nonce, ShouldEqual, "")
		So(result.Header.Prevhash, ShouldEqual, "9ee70b8ca48d948f985fb81755cb6860ccaf8863837037cd1d2373bf1e844a9e")
		So(result.Header.Timestamp, ShouldEqual, "1540517841494643")
		So(result.Signature, ShouldEqual, "477B9EC60B55A3F16D64F99F11826147FE75E39A374472A09E57ABFA5407789CCE51A53A2074ADFB1C2ED2DA71ED428306AA116CA23546E6F84EB006DC8D42B0")
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(testNet).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "5201003C71A5FD0CEE2DDFD251A646EE3AA65CDED4923DF57B8F63C2DF099C7EC793A7DB601D1F92BD645CEB732657A65D8BC71CE8C59FD94782C70C23BBB08B")
		So(result.Body.MicroBlockEmpty, ShouldResemble, []int64{1, 1, 1})
		So(result.Body.MicroBlockHashes, ShouldResemble, []string{"0000000000000000000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000000000000000000", "0000000000000000000000000000000000000000000000000000000000000000"})
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "6")
		So(result.Header.GasLimit, ShouldEqual, "1500000")
		So(result.Header.GasUsed, ShouldEqual, "0")
		So(result.Header.MinerPubKey, ShouldEqual, "0x027A6E4C15046DD06154D913B9215330371971D7213D05FC3B38DAF329A7619824")
		So(result.Header.NumMicroBlocks, ShouldEqual, 3)
		So(result.Header.NumTxns, ShouldEqual, 0)
		So(result.Header.Rewards, ShouldEqual, "0")
		So(result.Header.StateHash, ShouldEqual, "641735cb76d578cd8c20d68f1d163dbb247f1e4a8307bf7ee93380e6881f34ae")
		So(result.Header.Timestamp, ShouldEqual, "1540520511078714")
		So(result.Header.TxnHash, ShouldEqual, "2ea9ab9198d1638007400cd2c3bef1cc745b864b76011a0e1bc52180ac6452d4")
		So(result.Header.PrevBlockHash, ShouldEqual, "db122f3c36b64520157d6b52fa4763aa7aa83ede735126b82fcf3ba988cb0a1f")
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
		So(result.Body.MicroBlockEmpty, ShouldNotBeEmpty)
		So(result.Body.MicroBlockHashes, ShouldNotBeEmpty)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.DsBlockNum, ShouldNotBeBlank)
		So(result.Header.GasLimit, ShouldNotBeBlank)
		So(result.Header.GasUsed, ShouldNotBeBlank)
		So(result.Header.MinerPubKey, ShouldNotBeBlank)
		So(result.Header.NumMicroBlocks, ShouldBeGreaterThan, 0)
		So(result.Header.NumTxns, ShouldHaveSameTypeAs, int64(0))
		So(result.Header.Rewards, ShouldNotBeBlank)
		So(result.Header.StateHash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Header.TxnHash, ShouldNotBeBlank)
		So(result.Header.PrevBlockHash, ShouldNotBeBlank)
		So(result.Header.Type, ShouldHaveSameTypeAs, int64(1))
		So(result.Header.Version, ShouldHaveSameTypeAs, int64(0))
	})
}

func TestRPC_GetTransaction(t *testing.T) {
	Convey("returns details of a Transaction by its hash", t, func() {
		result, err := NewRPC(testNet).GetTransaction("b66a55133127d9718c12cfb022922c322e755b6e2e4027b96c0a0d9ae234fca5")
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, "b66a55133127d9718c12cfb022922c322e755b6e2e4027b96c0a0d9ae234fca5")
		So(result.Amount, ShouldEqual, "1")
		So(result.Nonce, ShouldEqual, "1")
		So(result.SenderPubKey, ShouldEqual, "0x02E8DE95B63E9598894BF4C3CFBBB0D5333385D306CB06122DB47C292AA57A7281")
		So(result.Signature, ShouldEqual, "0x0922D7C65FB9AA6CF85AC026DFB102288D91D6857487122C078FC04E2B6C991B47051BBBC4ECE99E8CB5C5CE6D0CC154979FEC952D0390E776C382240E44BEE5")
		So(result.ToAddr, ShouldEqual, "c0767be67c4895ff347898d1a9f5266f63936b6a")
		So(result.Version, ShouldEqual, "0")
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		zillean := NewZillean(testNet)
		privateKey := "729a77d87bf12e9445733a73961bbacdb93dbc6e175abaea948066ded6c9490a"
		publicKey, _ := zillean.GetPublicKeyFromPrivateKey(privateKey)
		rawTx := RawTransaction{
			Version:  0,
			Nonce:    2,
			To:       "ddff7b0fc10892deab8862514649bbc4757621f8",
			Amount:   "1",
			PubKey:   publicKey,
			GasPrice: 1,
			GasLimit: 1,
		}
		signature, _ := zillean.SignTransaction(rawTx, privateKey)
		result, err := zillean.RPC.CreateTransaction(rawTx, signature)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "ec69ac01c3b7eb7026d5254c897db96584605f4765d623555582b22bfe4c3bb1")
	})
}

// func TestRPC_GetSmartContracts(t *testing.T) {
// 	Convey("returns the list of smart contracts created by an address", t, func() {
// 		// TODO
// 	})
// }

// func TestRPC_GetSmartContractState(t *testing.T) {
// 	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
// 		// TODO
// 	})
// }

// func TestRPC_GetSmartContractCode(t *testing.T) {
// 	Convey("returns the Scilla code of a smart contract address", t, func() {
// 		// TODO
// 	})
// }

// func TestRPC_GetSmartContractInit(t *testing.T) {
// 	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
// 		// TODO
// 	})
// }

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
