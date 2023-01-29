package pages

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/nerijusdu/go-portfolio/internal/data"
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
		names[i] = "templates/" + names[i] + ".go.html"
	}

	names = append(names,
		"templates/partials/header.go.html",
		"templates/partials/footer.go.html",
		"templates/partials/navigation.go.html",
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
	"sliceSkills": func(arr []data.Skill, length, part int) []data.Skill {
		var res []data.Skill
		l := len(arr) / length
		s := (part - 1) * l
		for i := s; i < s+l; i++ {
			res = append(res, arr[i])
		}
		return res
	},
}
