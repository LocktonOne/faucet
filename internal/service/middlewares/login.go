package middlewares

import (
	"context"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"net/http"
)

func Login() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			address, token, err := helpers.ValidateJwt(r)
			if err != nil {
				helpers.Log(r).WithError(err).Error("failed to validate token")
				ape.Render(w, problems.Unauthorized())
				return
			}
			ctx := context.WithValue(r.Context(), "token", token)
			ctx = context.WithValue(ctx, "address", address)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
