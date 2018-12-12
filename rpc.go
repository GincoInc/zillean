package zillean

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/GincoInc/jsonrpc"
)

// RPC represents a JSON-RPC API client object.
type RPC struct {
	client *jsonrpc.RPCClient
}

// NewRPC returns a new zilliean.RPC.
func NewRPC(endpoint string) *RPC {
	return &RPC{
		client: jsonrpc.NewRPCClient(endpoint),
	}
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

// GetLatestTxBlock returns details of the most recent Transaction block.
func (r *RPC) GetLatestTxBlock() (*TxBlock, error) {
	resp, err := r.client.Call("GetLatestTxBlock", []interface{}{})
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}
	fmt.Println(resp)

	var result TxBlock
	resp.GetObject(&result)
	return &result, nil
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

// CreateTransaction create a new Transaction.
// See https://github.com/Zilliqa/Zilliqa-JavaScript-Library/#createtransactionjson in javascript
// for an example of how to construct the transaction object.
func (r *RPC) CreateTransaction(rawTx RawTransaction, signature string) (string, error) {
	amount, _ := strconv.ParseInt(rawTx.Amount, 16, 64)
	resp, err := r.client.Call("CreateTransaction", []interface{}{RawTx{
		Version:   rawTx.Version,
		Nonce:     rawTx.Nonce,
		To:        rawTx.To,
		Amount:    amount,
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

// GetSmartContracts returns the list of smart contracts created by an address.
// TODO
func (r *RPC) GetSmartContracts(contractAddress string) (*Transaction, error) {
	resp, err := r.client.Call("GetSmartContracts", []interface{}{contractAddress})
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

// GetSmartContractState returns  the state variables (mutable) of a smart contract address.
// TODO
func (r *RPC) GetSmartContractState(contractAddress string) (*Transaction, error) {
	resp, err := r.client.Call("GetSmartContractState", []interface{}{contractAddress})
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

// GetSmartContractCode returns the Scilla code of a smart contract address.
// TODO
func (r *RPC) GetSmartContractCode(contractAddress string) (*Transaction, error) {
	resp, err := r.client.Call("GetSmartContractCode", []interface{}{contractAddress})
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

// GetSmartContractInit returns the initialization parameters (immutable) of a given smart contract address.
// TODO
func (r *RPC) GetSmartContractInit(contractAddress string) (*Transaction, error) {
	resp, err := r.client.Call("GetSmartContractInit", []interface{}{contractAddress})
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
