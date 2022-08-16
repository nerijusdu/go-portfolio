package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nerijusdu/go-portfolio/internal/data"
)

func CreatePagesRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/not-found", notFoundPage)
	router.Get("/", homePage)
	router.Get("/projects", projectsPage)
	router.Get("/projects/{slug}", projectDetailsPage)
	router.Get("/blog", blogPage)
	router.Get("/blog/{slug}", blogDetailsPage)

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/home",
		"partials/projectList",
		"partials/blogList",
		HomePageData{
			Projects: getHighlighted(data.Projects),
			Blogs:    getHighlighted(data.Blogs),
		},
	)
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/projects",
		"partials/projectList",
		PageWithData[[]data.Project]{
			Title:       "My projects",
			Description: "A list of all my personal projects",
			Data:        getVisible(data.Projects),
		},
	)
}

func projectDetailsPage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	var p *data.Project
	for _, v := range data.Projects {
		if v.Slug == slug {
			p = &v
			break
		}
	}

	if p == nil {
		notFound(w, "Project not found")
		return
	}

	renderPage(w, "pages/projectDetails", PageWithData[*data.Project]{
		Title:       p.Name,
		Description: p.Description,
		Data:        p,
	})
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/blog",
		"partials/blogList",
		PageWithData[[]data.Blog]{
			Title:       "My blog",
			Description: "I write stuff here",
			Data:        getVisible(data.Blogs),
		},
	)
}

func blogDetailsPage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	var b *data.Blog
	for _, v := range data.Blogs {
		if v.Slug == slug {
			b = &v
			break
		}
	}

	if b == nil {
		notFound(w, "Blog article not found")
		return
	}

	renderPage(w, "pages/blogDetails", PageWithData[*data.Blog]{
		Title:       b.Title,
		Description: b.Description,
		Data:        b,
	})
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderPage(w, "pages/404")
}
