package pages

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type NameOrData interface {
	string | any
}

var version = os.Getenv("VERSION_HASH")
var templates = newTemplateCache()

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

	tmpl, err := templates.get(names)

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

var tempalteFuncs = template.FuncMap{
	"repeat": func(n int) []int {
		var res []int
		for i := 0; i < n; i++ {
			res = append(res, i+1)
		}
		return res
	},
	"paragraphs": func(s string) []string {
		return strings.Split(s, "\n")
	},
}
