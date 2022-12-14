package main

import (
	"flag"
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
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", pages.CreatePagesRouter())
	r.Mount("/api", api.CreateApiRouter())

	static.ServeStaticFiles("/static", r)

	fmt.Println("Starting server on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
