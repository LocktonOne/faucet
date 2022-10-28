package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/faucet/resources"
	"net/http"
)

type CreateRequest struct {
	Data resources.Send
}

func NewCreateRequest(r *http.Request) (resources.Send, error) {
	request := CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}
	return request.Data, nil
}

//todo add validate func
