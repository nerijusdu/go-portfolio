package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nerijusdu/go-portfolio/internal/data"
)

func CreatePagesRouter() chi.Router {
	router := chi.NewRouter()

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
		HomePageModel{
			Projects: data.Projects,
			Blogs:    data.Blogs,
		},
	)
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/projects",
		"partials/projectList",
		data.Projects,
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
		w.WriteHeader(404)
		w.Write([]byte("project not found"))
		return
	}

	renderPage(w, "pages/projectDetails", p)
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/blog",
		"partials/blogList",
		data.Blogs,
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
		w.WriteHeader(404)
		w.Write([]byte("blog not found"))
		return
	}

	renderPage(w, "pages/blogDetails", b)
}
