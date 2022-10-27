package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/faucet/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/faucet", func(r chi.Router) {

		//r.Post("/", handlers.DoSMTH)
		//r.Get("/ws", handlers.ListenSocket)

	})

	return r
}
