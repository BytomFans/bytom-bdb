package bigchaindb

import (
	"fmt"

	cl "github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	"golang.org/x/crypto/ed25519"
)

const (
	URL = "https://test.bigchaindb.com/api/v1"
)

// Create client
func CreateClient() (*cl.Client, error) {
	// Create client
	cfg := cl.ClientConfig{
		Url: URL,
	}
	client, err := cl.New(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// get block
func SendTransaction() {
	client, _ := CreateClient()
	// -- GET BLOCK -- //
	block1, err := client.GetBlock("0")
	fmt.Printf("Block height: %d\n", block1.Height, err)
}

//list transaction
func ListTransaction() {
	client, _ := CreateClient()
	txns, err := client.ListTransactions("1", "CREATE")
	fmt.Printf("txns: %d\n", txns, err)
}

func AssetCreationAndTransferE2E() (*txn.Transaction, error) {

	// -- PREPARE CREATE TRANSACTION -- //
	var keyPairs []*txn.KeyPair

	keyPair, err := txn.NewKeyPair()
	if err != nil {
		return nil, err
	}

	keyPairs = append(keyPairs, keyPair)

	var outputs []txn.Output
	var issuers []ed25519.PublicKey
	for _, keyPair := range keyPairs {
		// Create conditions
		condition := txn.NewEd25519Condition(keyPair.PublicKey)
		// Create output
		output, _ := txn.NewOutput(*condition, "1")
		outputs = append(outputs, output)
		// Create issuers
		issuers = append(issuers, keyPair.PublicKey)
	}

	data := make(map[string]interface{})
	data["assetID"] = "bytom"
	asset := txn.Asset{
		Data: data,
	}
	metadata := make(map[string]interface{})
	metadata["planet"] = "比原链专注资产领域的公有区块链平台"

	// New create transaction
	txn, err := txn.NewCreateTransaction(asset, metadata, outputs, issuers)
	if err != nil {
		return nil, err
	}

	// Sign transaction
	err = txn.Sign(keyPairs)
	fmt.Println(err, "Not able to sign txn")
	fmt.Println(txn.String())

	// Post transaction
	client, err := CreateClient()
	err = client.PostTransaction(txn)

	return txn, nil
}
