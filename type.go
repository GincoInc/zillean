package zillean

// Balance ...
type Balance struct {
	Balance int64 `json:"balance"`
	Nonce   int64 `json:"nonce"`
}

// DsBlock ...
type DsBlock struct {
	Header struct {
		BlockNum     string `json:"blockNum"`
		Difficulty   int64  `json:"difficulty"`
		LeaderPubKey string `json:"leaderPubKey"`
		MinerPubKey  string `json:"minerPubKey"`
		Nonce        string `json:"nonce"`
		Prevhash     string `json:"prevhash"`
		Timestamp    string `json:"timestamp"`
	} `json:"header"`
	Signature string `json:"signature"`
}

// TxBlock ...
type TxBlock struct {
	Body struct {
		HeaderSign       string   `json:"HeaderSign"`
		MicroBlockEmpty  []int64  `json:"MicroBlockEmpty"`
		MicroBlockHashes []string `json:"MicroBlockHashes"`
	} `json:"body"`
	Header struct {
		BlockNum       string `json:"BlockNum"`
		DsBlockNum     string `json:"DSBlockNum"`
		GasLimit       string `json:"GasLimit"`
		GasUsed        string `json:"GasUsed"`
		MinerPubKey    string `json:"MinerPubKey"`
		NumMicroBlocks int64  `json:"NumMicroBlocks"`
		NumTxns        int64  `json:"NumTxns"`
		StateHash      string `json:"StateHash"`
		Timestamp      string `json:"Timestamp"`
		TxnHash        string `json:"TxnHash"`
		PrevBlockHash  string `json:"prevBlockHash"`
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
	Version   int64  `json:"version"`
	Nonce     int64  `json:"nonce"`
	To        string `json:"to"`
	Amount    int64  `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  int64  `json:"gasPrice"`
	GasLimit  int64  `json:"gasLimit"`
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
