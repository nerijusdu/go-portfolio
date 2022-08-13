package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nerijusdu/go-portfolio/internal/api"
	"github.com/nerijusdu/go-portfolio/internal/pages"
	"github.com/nerijusdu/go-portfolio/internal/static"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/", pages.CreatePagesRouter())
	r.Mount("/api", api.CreateApiRouter())

	static.ServeStaticFiles("/static", r)

	fmt.Println("Starting server on http://localhost:3001")
	http.ListenAndServe(":"+port, r)
}
