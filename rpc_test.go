package zillean

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	baseURL = "https://api-scilla.zilliqa.com"
)

var (
	address        = "0x7f8A9ED7bA55A092A74105E5bD0Ec9C98e66051d"
	invalidAddress = "InvalidAddress"
)

func TestNewRPC(t *testing.T) {
	Convey("returns a new rpc", t, func() {
		So(NewRPC(baseURL), ShouldHaveSameTypeAs, &RPC{})
	})
}

func TestRPC_GetBalance(t *testing.T) {
	Convey("returns the balance and nonce of a given address", t, func() {
		result, err := NewRPC(baseURL).GetBalance("c5a829596fb06a59e2b1ddb6589811c759025d52")
		So(err, ShouldBeNil)
		So(result.Balance, ShouldEqual, 0)
		So(result.Nonce, ShouldEqual, 0)
	})
}

func TestRPC_GetDsBlock(t *testing.T) {
	Convey("returns details of a Directory Service block by block number", t, func() {
		result, err := NewRPC(baseURL).GetDsBlock("1")
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldEqual, "1")
		So(result.Header.Difficulty, ShouldEqual, 20)
		So(result.Header.LeaderPubKey, ShouldEqual, "0x02038BED2B924F908C8BD29980FC88E58D55C315254B15F7E91E18A1CBC7FF8E12")
		So(result.Header.MinerPubKey, ShouldEqual, "0x0336D5398D16FCB89331ED85F5F70649FD9570D069FF137A54F183F95C1056F72D")
		So(result.Header.Nonce, ShouldEqual, "1533615055")
		So(result.Header.Prevhash, ShouldEqual, "02ee29d8420e6f5d333e79927a32ac357d75889908db597fbadd6ec02ea1d7bb")
		So(result.Header.Timestamp, ShouldEqual, "1533614826240737")
		So(result.Signature, ShouldEqual, "DF68B28E8D51A43E6D2BECD3D96B4FF73F724FAC48652B9E88B61BF9C33C838D4848ABBADD4C8E327EDD31EC887EF10D49C9A1EB5EB2F4093A3ADB103BFB9021")
	})
}

func TestRPC_GetTxBlock(t *testing.T) {
	Convey("returns details of a Transaction block by block number.", t, func() {
		result, err := NewRPC(baseURL).GetTxBlock("100")
		So(err, ShouldBeNil)
		So(result.Body.HeaderSign, ShouldEqual, "E11C3969F4086AB1BF118C20B9F87CD81CFDD182D8B85A555C6D224011F9072A1F31231F17E8C3BE7FB73C687214EDC909A32FE96BC114BBB5A100F6DF9431ED")
		So(result.Body.MicroBlockEmpty, ShouldResemble, []int64{1})
		So(result.Body.MicroBlockHashes, ShouldResemble, []string{"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"})
		So(result.Header.BlockNum, ShouldEqual, "100")
		So(result.Header.DsBlockNum, ShouldEqual, "3")
		So(result.Header.GasLimit, ShouldEqual, "100")
		So(result.Header.GasUsed, ShouldEqual, "1")
		So(result.Header.MinerPubKey, ShouldEqual, "0x0384ED39FBB23CB6C0547D6F7C8B8440E1C3A9647FA4ECCF6362C076C87188A926")
		So(result.Header.NumMicroBlocks, ShouldEqual, 1)
		So(result.Header.NumTxns, ShouldEqual, 0)
		So(result.Header.StateHash, ShouldEqual, "0000000000000000000000000000000000000000000000000000000000000000")
		So(result.Header.Timestamp, ShouldEqual, "1533616392915940")
		So(result.Header.TxnHash, ShouldEqual, "5df6e0e2761359d30a8275058e299fcc0381534545f55cf43e41983f5d4c9456")
		So(result.Header.PrevBlockHash, ShouldEqual, "6cf5f7625c8018e077c27a2e6b1930c9cba3c2458fc5144133fa8b2a12e7ac06")
		So(result.Header.Type, ShouldEqual, 1)
		So(result.Header.Version, ShouldEqual, 0)
	})
}

func TestRPC_GetLatestDsBlock(t *testing.T) {
	Convey("returns details of the most recent Directory Service block", t, func() {
		result, err := NewRPC(baseURL).GetLatestDsBlock()
		So(err, ShouldBeNil)
		So(result.Header.BlockNum, ShouldNotBeBlank)
		So(result.Header.Difficulty, ShouldBeGreaterThan, 0)
		So(result.Header.LeaderPubKey, ShouldNotBeBlank)
		So(result.Header.MinerPubKey, ShouldNotBeBlank)
		So(result.Header.Nonce, ShouldNotBeBlank)
		So(result.Header.Prevhash, ShouldNotBeBlank)
		So(result.Header.Timestamp, ShouldNotBeBlank)
		So(result.Signature, ShouldNotBeBlank)
	})
}

func TestRPC_GetLatestTxBlock(t *testing.T) {
	Convey("returns details of the most recent Transaction block", t, func() {
		result, err := NewRPC(baseURL).GetLatestTxBlock()
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
		// TODO
	})
}

func TestRPC_CreateTransaction(t *testing.T) {
	Convey("returns a hash of created Transaction", t, func() {
		// TODO
	})
}

func TestRPC_GetSmartContracts(t *testing.T) {
	Convey("returns the list of smart contracts created by an address", t, func() {
		// TODO
	})
}

func TestRPC_GetSmartContractState(t *testing.T) {
	Convey("returns the state variables (mutable) of a smart contract address", t, func() {
		// TODO
	})
}

func TestRPC_GetSmartContractCode(t *testing.T) {
	Convey("returns the Scilla code of a smart contract address", t, func() {
		// TODO
	})
}

func TestRPC_GetSmartContractInit(t *testing.T) {
	Convey("returns the initialization parameters (immutable) of a given smart contract address", t, func() {
		// TODO
	})
}

func TestRPC_GetBlockchainInfo(t *testing.T) {
	Convey("returns statistics about the specified zilliqa node", t, func() {
		result, err := NewRPC(baseURL).GetBlockchainInfo()
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
		result, err := NewRPC(baseURL).GetNetworkID()
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "TestNet")
	})
}

func TestRPC_GetRecentTransactions(t *testing.T) {
	Convey("returns  the most recent transactions (upto 100) accepted by the specified zilliqa node.", t, func() {
		result, err := NewRPC(baseURL).GetRecentTransactions()
		So(err, ShouldBeNil)
		So(len(result.TxnHashes), ShouldEqual, result.Number)
	})
}

func TestRPC_GetDSBlockListing(t *testing.T) {
	Convey("returns a paginated list of Directory Service blocks", t, func() {
		result, err := NewRPC(baseURL).DSBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}

func TestRPC_GetTxBlockListing(t *testing.T) {
	Convey("returns a paginated list of Transaction blocks", t, func() {
		result, err := NewRPC(baseURL).TxBlockListing(1)
		So(err, ShouldBeNil)
		So(len(result.Data), ShouldEqual, 10)
		So(result.MaxPages, ShouldBeGreaterThan, 0)
	})
}
