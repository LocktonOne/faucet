package handlers

import (
	"bytes"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"gitlab.com/tokene/faucet/internal/service/requests"
	"gitlab.com/tokene/faucet/internal/txs"
	"gitlab.com/tokene/faucet/resources"
	"io"
	"net/http"
)

func Faucet(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to parse request")
		ape.Render(w, problems.BadRequest(err))
		return
	}

	signedTx, err := txs.SignTx(r, request, helpers.EthRPCConfig(r).Endpoint, request.Attributes.Recipient.Amount)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}
	rawTx, err := txs.NewCreateRawTx(signedTx)

	response, err := http.Post(helpers.EthRPCConfig(r).Endpoint, "application/json", bytes.NewBuffer(rawTx))
	defer response.Body.Close()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get response from core")
		ape.Render(w, problems.InternalError())
		return
	}

	responseMess, err := io.ReadAll(response.Body)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to parse response")
		ape.Render(w, problems.BadRequest(err))
		return
	}
	parsedTxResponse, err := txs.NewParseResultTx(responseMess)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to parse response ")
		ape.Render(w, problems.BadRequest(err))
		return
	}

	ape.Render(w, resources.TxHashResponse{
		Data: resources.TxHash{
			resources.Key{Type: resources.TX_HASH},
			resources.TxHashAttributes{TxHash: parsedTxResponse.Result},
		},
	})
}
