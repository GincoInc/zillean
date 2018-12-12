package zillean

// Balance ...
type Balance struct {
	Balance int64 `json:"balance,string"`
	Nonce   int64 `json:"nonce,string"`
}

// DsBlock ...
type DsBlock struct {
	Header struct {
		BlockNum     string   `json:"blockNum"`
		Difficulty   int64    `json:"difficulty"`
		DifficultyDS int64    `json:"difficultyDS"`
		GasPrice     string   `json:"gasPrice"`
		LeaderPubKey string   `json:"leaderPubKey"`
		PoWWinners   []string `json:"powWinners"`
		Prevhash     string   `json:"prevhash"`
		Timestamp    string   `json:"timestamp"`
	} `json:"header"`
	Signature string `json:"signature"`
}

// TxBlock ...
type TxBlock struct {
	Body struct {
		HeaderSign string `json:"HeaderSign"`
		//MicroBlockInfos []string `json:"MicroBlockInfos"`
	} `json:"body"`
	Header struct {
		BlockNum       string `json:"BlockNum"`
		DsBlockNum     string `json:"DSBlockNum"`
		GasLimit       string `json:"GasLimit"`
		GasUsed        string `json:"GasUsed"`
		MbInfoHash     string `json:"MbInfoHash"`
		MinerPubKey    string `json:"MinerPubKey"`
		NumMicroBlocks int64  `json:"NumMicroBlocks"`
		NumTxns        int64  `json:"NumTxns"`
		PrevBlockHash  string `json:"prevBlockHash"`
		Rewards        string `json:"Rewards"`
		StateDeltaHash string `json:"StateDeltaHash"`
		StateRootHash  string `json:"StateRootHash"`
		Timestamp      string `json:"Timestamp"`
		TxnHash        string `json:"TxnHash"`
		Type           int64  `json:"type"`
		Version        int64  `json:"version"`
	} `json:"header"`
}

// Transaction ...
type Transaction struct {
	ID           string `json:"ID"`
	Amount       string `json:"amount"`
	Nonce        string `json:"nonce"`
	SenderPubKey string `json:"senderPubKey"`
	Signature    string `json:"signature"`
	ToAddr       string `json:"toAddr"`
	Version      string `json:"version"`
}

// RawTx ...
type RawTx struct {
	Version   int32  `json:"version"`
	Nonce     int32  `json:"nonce"`
	To        string `json:"to"`
	Amount    int64  `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  int32  `json:"gasPrice"`
	GasLimit  int32  `json:"gasLimit"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

// BlockchainInfo ...
type BlockchainInfo struct {
	CurrentDSEpoch    string  `json:"CurrentDSEpoch"`
	CurrentMiniEpoch  string  `json:"CurrentMiniEpoch"`
	DSBlockRate       float64 `json:"DSBlockRate"`
	NumDSBlocks       string  `json:"NumDSBlocks"`
	NumPeers          int64   `json:"NumPeers"`
	NumTransactions   string  `json:"NumTransactions"`
	NumTxBlocks       string  `json:"NumTxBlocks"`
	NumTxnsDSEpoch    string  `json:"NumTxnsDSEpoch"`
	NumTxnsTxEpoch    int64   `json:"NumTxnsTxEpoch"`
	ShardingStructure struct {
		NumPeers []int64 `json:"NumPeers"`
	}
	TransactionRate int64   `json:"TransactionRate"`
	TxBlockRate     float64 `json:"TxBlockRate"`
}

// RecentTransactions ...
type RecentTransactions struct {
	TxnHashes []string `json:"TxnHashes"`
	Number    int64    `json:"number"`
}

// ListedBlocks ...
type ListedBlocks struct {
	Data []struct {
		BlockNum int64  `json:"BlockNum"`
		Hash     string `json:"Hash"`
	} `json:"data"`
	MaxPages int64 `json:"maxPages"`
}

// RawTransaction ...
type RawTransaction struct {
	Version  int32
	Nonce    int32
	To       string
	Amount   string
	PubKey   string
	GasPrice int32
	GasLimit int32
	Code     string
	Data     string
}
