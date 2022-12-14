package middlewares

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"gitlab.com/tokene/faucet/internal/service/requests"
	"net/http"
)

func Login() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			address, token, err := helpers.ValidateJwt(r)
			if err != nil {
				helpers.Log(r).WithError(err).Error("failed to validate token")
				ape.Render(w, problems.BadRequest(err))
				return
			}

			request, err := requests.NewCreateRequest(r)
			if err != nil {
				helpers.Log(r).WithError(err).Error("failed to parse request")
				ape.Render(w, problems.BadRequest(err))
				return
			}
			if request.Attributes.Recipient.Address != address {
				doorman := helpers.DoormanConnector(r)
				err := doorman.CheckPermissionID("CREATE", "*", token)
				if err == nil {
					next.ServeHTTP(w, r)
				}
				helpers.Log(r).WithError(err).Error("haven't permission for this operation")
				ape.Render(w, problems.BadRequest(err))
				return

			}

			next.ServeHTTP(w, r)
		})
	}
}
