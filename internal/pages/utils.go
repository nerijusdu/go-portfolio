package pages

import (
	"fmt"
	"net/http"
	"text/template"
)

type NameOrData interface {
	string | any
}

func renderPage(w http.ResponseWriter, namesOrData ...NameOrData) error {
	var data any = nil
	templates := make([]string, 0)

	for i := 0; i < len(namesOrData); i++ {
		switch v := namesOrData[i].(type) {
		case string:
			templates = append(templates, "templates/"+v+".html")
		default:
			data = v
		}
	}

	templates = append(templates,
		"templates/partials/header.html",
		"templates/partials/footer.html",
		"templates/partials/navigation.html",
	)

	// TODO: cache parsed templates?
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		somethingWentWrong(w, err)
		return err
	}

	if err = tmpl.Execute(w, data); err != nil {
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
