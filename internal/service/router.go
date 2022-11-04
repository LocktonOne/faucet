package service

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/faucet/internal/config"
	"gitlab.com/tokene/faucet/internal/service/handlers"
	"gitlab.com/tokene/faucet/internal/service/helpers"

	"github.com/go-chi/chi"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxSigner(cfg.Signer()),
			helpers.CtxEthRPCConfig(cfg.EthRPCConfig()),
			helpers.CtxAuthConfig(cfg.AuthConfig()),
		),
	)

	r.Route("/integrations/faucet", func(r chi.Router) {
		r.Post("/", handlers.Faucet)
	})

	return r
}
