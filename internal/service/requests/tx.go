package requests

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokene/faucet/resources"
	"log"
	"math/big"
)

const eth = 10000

type CreateRawTx struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int64    `json:"id"`
}

func NewCreateRawTx(params string) ([]byte, error) {
	tx := CreateRawTx{}
	tx.Jsonrpc = "2.0"
	tx.Method = "eth_sendRawTransaction"
	tx.Params = append(tx.Params, fmt.Sprint("0x", params))
	tx.Id = 1

	rawTX, err := json.Marshal(tx)
	if err != nil {
		return rawTX, err
	}
	return rawTX, nil

}

func SignTx(request resources.Send, core string, sender string, amount float32) string {

	client, err := ethclient.Dial(core)

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("1a6e0674272af1dbffec2eb4188de8b311abc48046a38f5cbf6db55eb4fe9597")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress(sender)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(int64(float64(amount) * eth)) // in wei (1 eth)
	gasLimit := uint64(21000)                         // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(request.Attributes.Recipient.Address)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	b := new(bytes.Buffer)
	ts.EncodeIndex(0, b)
	rawTxBytes := b.Bytes()
	return hex.EncodeToString(rawTxBytes)

}
