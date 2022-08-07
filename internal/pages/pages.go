package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreatePagesRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", homePage)

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println("Error reading tempalte", err)
		somethingWentWrong(w)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		somethingWentWrong(w)
		return
	}
}

func somethingWentWrong(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
}
