package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/faucet/resources"
	"net/http"
	"regexp"
)

type CreateRequest struct {
	Data resources.Send
}

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")

func NewCreateRequest(r *http.Request) (resources.Send, error) {
	request := CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to unmarshal")
	}
	return request.Data, request.validate()
}

func (r *CreateRequest) validate() error {
	return validation.Errors{
		"/data/type": validation.Validate(&r.Data.Type, validation.Required, validation.In(resources.FAUCET)),
		"/data/attributes/auth_pair/address": validation.Validate(&r.Data.Attributes.Recipient.Address,
			validation.Required,
			validation.Match(AddressRegexp)),
	}.Filter()
}
