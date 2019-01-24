package bigchaindb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	txn "github.com/bigchaindb/go-bigchaindb-driver/pkg/transaction"
	cl "github.com/bigchaindb/go-bigchaindb-driver/pkg/client"
)

//Get Block
func NewClientCreatesConnectionToTestNetwork() {

	// Make sure to add dbd_key.json file for your own test network
	headerBytes, _ := ioutil.ReadFile("../fixtures/dbd_key.json")

	var h map[string]string
	_ = json.Unmarshal(headerBytes, &h)

	cfg := cl.ClientConfig{
		Url:     URL,
		Headers: h,
	}

	client, err := cl.New(cfg)
	if err != nil {
		errors.New("Could not create client")
	}

	block, err := client.GetBlock("0")
	if err != nil {
		errors.New("Could not retrieve genesis block")
	}

	fmt.Println(block.Height)
}

func NewClientCreatesConnectionLocalhost() (int, error) {

	cfg := cl.ClientConfig{
		Url: URL,
	}

	client, err := cl.New(cfg)
	if err != nil {
		return 0, err
	}
	block, err := client.GetBlock("0")
	if err != nil {
		return 0, err
	}
	return block.Height, nil

}

func PostCreateTransaction() {

	ctxnBytes, _ := ioutil.ReadFile("../fixtures/mock_create_txn.json")

	var createTxn *txn.Transaction
	err := json.Unmarshal(ctxnBytes, &createTxn)

	cfg := cl.ClientConfig{
		Url: URL,
	}

	client, _ := cl.New(cfg)

	err = client.PostTransaction(createTxn)
	if err != nil {
		fmt.Println("创建交易失败:", err)
	}
}
