package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/vendor116/playgo/internal/generated"
)

func GetChiRouter(server generated.ServerInterface) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/v1/info", server.GetInfo)

	return router
}
