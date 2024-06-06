package helpers

import (
	"github.com/pkg/errors"
	"net/http"
)

func ValidateJwt(r *http.Request) (string, string, error) {
	doorman := DoormanConnector(r)

	token, err := doorman.GetAuthToken(r)
	if err != nil {
		return "", "", errors.Wrap(err, "invalid token")
	}

	address, err := doorman.ValidateJwt(token)
	if err != nil {
		return "", "", errors.Wrap(err, "user does not have permission")
	}

	return address, token, nil
}
