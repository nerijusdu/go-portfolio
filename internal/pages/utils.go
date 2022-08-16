package pages

import (
	"fmt"
	"net/http"

	"github.com/elliotchance/pie/v2"
)

func somethingWentWrong(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Something went wrong"))
}

func notFound(w http.ResponseWriter, err string) {
	fmt.Println(err)
	w.Header().Add("Location", "/not-found")
	w.WriteHeader(http.StatusPermanentRedirect)
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
