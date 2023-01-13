package service

import (
	"github.com/Swapica/relayer-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxChains(s.cfg.Chains()),
			handlers.CtxTransactor(s.cfg.Transactor()),
		),
	)
	r.Route("/integrations/relayer-svc", func(r chi.Router) {
		r.Post("/transaction", handlers.CallContract)
	})

	return r
}
