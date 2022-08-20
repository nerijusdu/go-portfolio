package pages

import (
	"flag"
	"fmt"
	"html/template"
	"path"
	"strings"
)

var skipCache = flag.Bool("skip-cache", false, "skip template cache")

type templateCache struct {
	cache map[string]*template.Template
}

func newTemplateCache() *templateCache {
	return &templateCache{
		cache: make(map[string]*template.Template),
	}
}

func (c *templateCache) get(names []string) (*template.Template, error) {
	key := strings.Join(names, ";")
	tmpl, ok := c.cache[key]
	if ok {
		return tmpl, nil
	}

	tmpl, err := template.
		New(path.Base(names[0])).
		Funcs(template.FuncMap{
			"repeat": func(n int) []int {
				var res []int
				for i := 0; i < n; i++ {
					res = append(res, i+1)
				}
				return res
			},
		}).
		ParseFiles(names...)
	if err != nil {
		return nil, err
	}

	if !*skipCache {
		fmt.Println("caching", key)
		c.cache[key] = tmpl
	}
	return tmpl, nil
}
