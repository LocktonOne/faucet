package txs

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"gitlab.com/tokene/faucet/resources"
	"math/big"
	"net/http"
)

const eth = 1000000000000000000

type CreateRawTx struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int64    `json:"id"`
}

type ParseResultTx struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int64  `json:"id"`
}

func NewParseResultTx(result []byte) (ParseResultTx, error) {
	parseResultTx := ParseResultTx{}
	err := json.Unmarshal(result, &parseResultTx)
	return parseResultTx, err
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

func SignTx(r *http.Request, request resources.Send, core string, amount float32) (string, error) {

	client, err := ethclient.Dial(core)

	if err != nil {
		return "", err
	}

	signer := helpers.Signer(r)

	nonce, err := client.PendingNonceAt(context.Background(), signer.Address())
	if err != nil {
		return "", err
	}

	valueF := big.NewFloat(1)
	valueF.Mul(big.NewFloat(float64(eth)), big.NewFloat(float64(amount))) // in wei (1 eth)

	value := new(big.Int)
	valueF.Int(value)

	gasLimit := uint64(210000) // in units// in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(request.Attributes.Recipient.Address)
	var data []byte

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := signer.SignTx(tx, chainID)
	if err != nil {
		return "", err
	}

	txs := types.Transactions{signedTx}
	b := new(bytes.Buffer)
	txs.EncodeIndex(0, b)
	rawTxBytes := b.Bytes()

	return hex.EncodeToString(rawTxBytes), nil

}
