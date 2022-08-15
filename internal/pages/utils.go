package pages

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/elliotchance/pie/v2"
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

type Hideable interface {
	IsHidden() bool
}

func getVisible[T Hideable](items []T) []T {
	return pie.Filter(items, func(x T) bool {
		return !x.IsHidden()
	})
}

type Highlightable interface {
	Hideable
	IsHighlighted() bool
}

func getHighlighted[T Highlightable](items []T) []T {
	return pie.Filter(items, func(x T) bool {
		return x.IsHighlighted() && !x.IsHidden()
	})
}
