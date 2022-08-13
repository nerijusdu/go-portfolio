package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateApiRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return r
}
