package zillean

import (
	"errors"

	"github.com/GincoInc/jsonrpc"
)

// RPC represents a JSON-RPC API client object.
type RPC struct {
	client *jsonrpc.RPCClient
}

// NewRPC returns a new RPC object.
func NewRPC(endpoint string) *RPC {
	return &RPC{
		client: jsonrpc.NewRPCClient(endpoint),
	}
}

// GetNetworkID returns the network ID of the specified zilliqa node.
func (r *RPC) GetNetworkID() (string, error) {
	resp, err := r.client.Call("GetNetworkId", []interface{}{})
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var result string
	resp.GetObject(&result)
	return result, nil
}

// GetBlockchainInfo returns statistics about the specified zilliqa node.
func (r *RPC) GetBlockchainInfo() (*BlockchainInfo, error) {
	resp, err := r.client.Call("GetBlockchainInfo", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result BlockchainInfo
	resp.GetObject(&result)
	return &result, nil
}

// GetShardingStructure returns the current sharding structure of the network from the specified network's lookup node.
func (r *RPC) GetShardingStructure() (*ShardingStructure, error) {
	resp, err := r.client.Call("GetShardingStructure", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result ShardingStructure
	resp.GetObject(&result)
	return &result, nil
}

// GetDsBlock returns details of a Directory Service block by block number.
func (r *RPC) GetDsBlock(blockNumber string) (*DsBlock, error) {
	resp, err := r.client.Call("GetDsBlock", []interface{}{blockNumber})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result DsBlock
	resp.GetObject(&result)
	return &result, nil
}

// GetLatestDsBlock returns details of the most recent Directory Service block.
func (r *RPC) GetLatestDsBlock() (*DsBlock, error) {
	resp, err := r.client.Call("GetLatestDsBlock", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result DsBlock
	resp.GetObject(&result)
	return &result, nil
}

// GetNumDSBlocks returns the number of Directory Service blocks in the network so far. This is represented as a String.
func (r *RPC) GetNumDSBlocks() (string, error) {
	resp, err := r.client.Call("GetNumDSBlocks", []interface{}{})
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var result string
	resp.GetObject(&result)
	return result, nil

}

// GetDSBlockRate returns the current Directory Service blockrate per second.
func (r *RPC) GetDSBlockRate() (float64, error) {
	resp, err := r.client.Call("GetDSBlockRate", []interface{}{})
	if err != nil {
		return 0, err
	}

	if resp.Error != nil {
		return 0, errors.New(resp.Error.Message)
	}

	var result float64
	resp.GetObject(&result)
	return result, nil

}

// DSBlockListing returns a paginated list of Directory Service blocks.
// Pass in page number as parameter.
// Returns a maxPages variable that specifies the max number of pages.
// 1 - latest blocks, maxPages - oldest blocks.
func (r *RPC) DSBlockListing(pageNumber int64) (*ListedBlocks, error) {
	resp, err := r.client.Call("DSBlockListing", []interface{}{pageNumber})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result ListedBlocks
	resp.GetObject(&result)
	return &result, nil
}

// GetTxBlock returns details of a Transaction block by block number.
func (r *RPC) GetTxBlock(blockNumber string) (*TxBlock, error) {
	resp, err := r.client.Call("GetTxBlock", []interface{}{blockNumber})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result TxBlock
	resp.GetObject(&result)
	return &result, nil
}

// GetLatestTxBlock returns details of the most recent Transaction block.
func (r *RPC) GetLatestTxBlock() (*TxBlock, error) {
	resp, err := r.client.Call("GetLatestTxBlock", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result TxBlock
	resp.GetObject(&result)
	return &result, nil
}

// GetNumTxBlocks returns the number of Transaction blocks in the network so far, this is represented as String.
func (r *RPC) GetNumTxBlocks() {

}

// GetTxBlockRate returns the current Transaction blockrate per second.
func (r *RPC) GetTxBlockRate() {

}

// TxBlockListing returns a paginated list of Transaction blocks.
// Pass in page number as parameter.
// Returns a maxPages variable that specifies the max number of pages.
// 1 - latest blocks, maxPages - oldest blocks.
func (r *RPC) TxBlockListing(pageNumber int64) (*ListedBlocks, error) {
	resp, err := r.client.Call("TxBlockListing", []interface{}{pageNumber})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result ListedBlocks
	resp.GetObject(&result)
	return &result, nil
}

// GetNumTransactions returns the number of Transactions validated in the network so far. This is represented as a String.
func (r *RPC) GetNumTransactions() {

}

// GetTransactionRate returns the current Transaction rate of the network.
func (r *RPC) GetTransactionRate() {

}

// GetCurrentMiniEpoch returns the number of TX epochs in the network so far represented as String.
func (r *RPC) GetCurrentMiniEpoch() {

}

// GetCurrentDSEpoch returns the number of DS epochs in the network so far represented as String.
func (r *RPC) GetCurrentDSEpoch() {

}

// GetPrevDifficulty returns the minimum shard difficulty of the previous block, this is represented as an Number.
func (r *RPC) GetPrevDifficulty() {

}

// GetPrevDSDifficulty returns the minimum DS difficulty of the previous block, this is represented as an Number.
func (r *RPC) GetPrevDSDifficulty() {

}

// CreateTransaction create a new Transaction.
// See https://github.com/Zilliqa/Zilliqa-JavaScript-Library/#createtransactionjson in javascript
// for an example of how to construct the transaction object.
func (r *RPC) CreateTransaction(rawTx RawTransaction, signature string) (string, error) {
	resp, err := r.client.Call("CreateTransaction", []interface{}{RawTransaction{
		Version:   rawTx.Version,
		Nonce:     rawTx.Nonce,
		To:        rawTx.To,
		Amount:    rawTx.Amount,
		PubKey:    rawTx.PubKey,
		GasPrice:  rawTx.GasPrice,
		GasLimit:  rawTx.GasLimit,
		Signature: signature,
	}})

	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var result struct {
		Info   string `json:"Info"`
		TranID string `json:"TranID"`
	}
	resp.GetObject(&result)
	return result.TranID, nil
}

// GetTransaction returns details of a Transaction by its hash.
func (r *RPC) GetTransaction(txHash string) (*Transaction, error) {
	resp, err := r.client.Call("GetTransaction", []interface{}{txHash})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result Transaction
	resp.GetObject(&result)
	return &result, nil
}

// GetRecentTransactions returns  the most recent transactions (upto 100) accepted by the specified zilliqa node.
func (r *RPC) GetRecentTransactions() (*RecentTransactions, error) {
	resp, err := r.client.Call("GetRecentTransactions", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result RecentTransactions
	resp.GetObject(&result)
	return &result, nil
}

// GetTransactionsForTxBlock returns the transactions included within a micro-block created by a specific shard.
func (r *RPC) GetTransactionsForTxBlock() {

}

// GetNumTxnsTxEpoch returns the number of transactions in this Transaction epoch, this is represented as String.
func (r *RPC) GetNumTxnsTxEpoch() {

}

// GetNumTxnsDSEpoch returns the number of transactions in this Directory Service epoch, this is represented as String.
func (r *RPC) GetNumTxnsDSEpoch() {

}

// GetMinimumGasPrice returns the minimum gas price of the last DS epoch represented as String. This is measured in the smallest price unit Qa (10^-12 Zil) in Zilliqa.
func (r *RPC) GetMinimumGasPrice() {

}

// GetSmartContractCode returns the Scilla code of a smart contract address.
func (r *RPC) GetSmartContractCode(contractAddress string) (string, error) {
	resp, err := r.client.Call("GetSmartContractCode", []interface{}{contractAddress})
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var result struct {
		Code string `json:"code"`
	}
	resp.GetObject(&result)
	return result.Code, nil
}

// GetSmartContractInit returns the initialization parameters (immutable) of a given smart contract address.
func (r *RPC) GetSmartContractInit(contractAddress string) ([]SmartContractState, error) {
	resp, err := r.client.Call("GetSmartContractInit", []interface{}{contractAddress})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result []SmartContractState
	resp.GetObject(&result)
	return result, nil
}

// GetSmartContractState returns  the state variables (mutable) of a smart contract address.
func (r *RPC) GetSmartContractState(contractAddress string) ([]SmartContractState, error) {
	resp, err := r.client.Call("GetSmartContractState", []interface{}{contractAddress})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result []SmartContractState
	resp.GetObject(&result)
	return result, nil
}

// GetSmartContracts returns the list of smart contracts created by an address.
func (r *RPC) GetSmartContracts(address string) ([]SmartContract, error) {
	resp, err := r.client.Call("GetSmartContracts", []interface{}{address})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result []SmartContract
	resp.GetObject(&result)
	return result, nil
}

// GetContractAddressFromTransactionID returns a smart contract address of 20 bytes from a transaction ID, represented as a String .
func (r *RPC) GetContractAddressFromTransactionID() {

}

// GetBalance returns the balance and nonce of a given address.
func (r *RPC) GetBalance(address string) (*Balance, error) {
	resp, err := r.client.Call("GetBalance", []interface{}{address})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	var result Balance
	resp.GetObject(&result)
	return &result, nil
}
