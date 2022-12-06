package middlewares

import (
	"net/http"
)

func Login() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//todo parse JWT
			//if address == address from request
			//use doorman.CheckPermission
			next.ServeHTTP(w, r)
		})
	}
}
