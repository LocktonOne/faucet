package txs

import (
	"context"
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

func SignTx(r *http.Request, request resources.Send, client *ethclient.Client, amount float32) (*types.Transaction, error) {

	signer := helpers.Signer(r)

	nonce, err := client.PendingNonceAt(context.Background(), signer.Address())
	if err != nil {
		return nil, err
	}

	valueF := big.NewFloat(1)
	valueF.Mul(big.NewFloat(float64(eth)), big.NewFloat(float64(amount))) // in wei (1 eth)

	value := new(big.Int)
	valueF.Int(value)

	gasLimit := uint64(210000) // in units// in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
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
		return nil, err
	}

	signedTx, err := signer.SignTx(tx, chainID)
	if err != nil {
		return nil, err
	}

	return signedTx, nil

}
