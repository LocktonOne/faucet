package middlewares

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/faucet/internal/service/helpers"
	"net/http"
)

func Login() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", helpers.AuthConfig(r).Endpoint, nil)
			req.Header.Set("Authorization", r.Header.Get("Authorization"))
			authResponse, err := client.Do(req)
			defer authResponse.Body.Close()
			if err != nil {
				helpers.Log(r).WithError(err).Error("failed to get response from nonce-auth-svc")
				ape.Render(w, problems.InternalError())
				return
			}
			if authResponse.StatusCode != 200 {
				helpers.Log(r).WithError(err).Error("unauthorized")
				ape.Render(w, problems.Unauthorized())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
