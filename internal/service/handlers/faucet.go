package handlers

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"gitlab.com/tokene/faucet/internal/service/requests"
	"gitlab.com/tokene/faucet/internal/txs"
	"gitlab.com/tokene/faucet/resources"
	"net/http"
	"time"
)

func Faucet(w http.ResponseWriter, r *http.Request) {

	client, err := ethclient.Dial(helpers.EthRPCConfig(r).Endpoint)

	request, err := requests.NewCreateRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to parse request")
		ape.Render(w, problems.BadRequest(err))
		return
	}

	signedTx, err := txs.SignTx(r, request, client, request.Attributes.Recipient.Amount)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	for {
		_, isPending, err := client.TransactionByHash(context.TODO(), signedTx.Hash())
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to create tx")
			ape.Render(w, problems.InternalError())
			return
		}
		if isPending {
			time.Sleep(1 * time.Millisecond)
		} else {
			break
		}
	}

	client.SendTransaction(context.TODO(), signedTx)
	ape.Render(w, resources.TxHashResponse{
		Data: resources.TxHash{
			resources.Key{Type: resources.TX_HASH},
			resources.TxHashAttributes{TxHash: signedTx.Hash().String()},
		},
	})
}
