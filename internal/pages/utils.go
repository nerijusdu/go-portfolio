package pages

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderPage(w http.ResponseWriter, names ...string) error {
	for i := 0; i < len(names); i++ {
		names[i] = "templates/pages/" + names[i] + ".html"
	}

	names = append(names,
		"templates/partials/header.html",
		"templates/partials/footer.html",
		"templates/partials/navigation.html",
	)

	// TODO: cache parsed templates?
	tmpl, err := template.ParseFiles(names...)

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
