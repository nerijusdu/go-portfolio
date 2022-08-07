package pages

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderPage(name string, w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles(
		"templates/pages/"+name+".html",
		"templates/partials/header.html",
		"templates/partials/footer.html",
	)

	if err != nil {
		somethingWentWrong(w, err)
		return err
	}

	if err = tmpl.Execute(w, nil); err != nil {
		somethingWentWrong(w, err)
		return err
	}

	return nil
}

func somethingWentWrong(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
}
