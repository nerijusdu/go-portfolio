package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreatePagesRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", homePage)

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "home")
}
