package handlers

import (
	"bytes"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"gitlab.com/tokene/faucet/internal/service/requests"
	"io"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to parse request")
		ape.Render(w, problems.BadRequest(err))
		return
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", helpers.AuthConfig(r).Endpoint, nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	authResponse, _ := client.Do(req)
	defer authResponse.Body.Close()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get response from nonce-auth-svc")
		ape.Render(w, problems.InternalError())
		return
	}
	if authResponse.StatusCode == 200 {
		helpers.Log(r).WithError(err).Error("bad response code")
		ape.Render(w, problems.Unauthorized())
		return
	}

	var signParams string
	if request.Attributes.Recipient.Amount != 0 {
		signParams = requests.SignTx(request, helpers.EthRPCConfig(r).Endpoint, helpers.SenderRPCConfig(r).Address, request.Attributes.Recipient.Amount)

	} else {
		signParams = requests.SignTx(request, helpers.EthRPCConfig(r).Endpoint, helpers.SenderRPCConfig(r).Address, helpers.SenderRPCConfig(r).Amount)

	}

	rawTx, err := requests.NewCreateRawTx(signParams)

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

	w.Write(responseMess)

}
