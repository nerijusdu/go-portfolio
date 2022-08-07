package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nerijusdu/go-portfolio/internal/pages"
	"github.com/nerijusdu/go-portfolio/internal/static"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	r.Mount("/", pages.CreatePagesRouter())

	static.ServeStaticFiles("/static", r)

	fmt.Println("Starting server on http://localhost:3001")
	http.ListenAndServe(":3001", r)
}
