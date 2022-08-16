package pages

import (
	"net/http"
	"os"
	"text/template"
)

type NameOrData interface {
	string | any
}

var version = os.Getenv("VERSION_HASH")

func renderPageWithData[T any](w http.ResponseWriter, data PageWithData[T], names ...string) error {
	data.Version = "0"
	if version != "" {
		data.Version = version
	}

	for i := 0; i < len(names); i++ {
		names[i] = "templates/" + names[i] + ".html"
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

	if err = tmpl.Execute(w, data); err != nil {
		somethingWentWrong(w, err)
		return err
	}

	return nil
}

func renderPage(w http.ResponseWriter, names ...string) error {
	return renderPageWithData(w, PageWithData[any]{}, names...)
}