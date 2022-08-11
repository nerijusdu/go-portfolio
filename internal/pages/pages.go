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
	router.Get("/blog", blogPage)

	return router
}

func homePage(w http.ResponseWriter, r *http.Request) {
	projects, err := data.ReadProjects()
	if err != nil {
		somethingWentWrong(w, err)
		return
	}

	renderPage(
		w,
		"pages/home",
		"partials/projectList",
		HomePageModel{Projects: projects},
	)
}

func projectsPage(w http.ResponseWriter, r *http.Request) {
	projects, err := data.ReadProjects()
	if err != nil {
		somethingWentWrong(w, err)
		return
	}

	renderPage(
		w,
		"pages/projects",
		"partials/projectList",
		HomePageModel{Projects: projects},
	)
}

func blogPage(w http.ResponseWriter, r *http.Request) {
	renderPage(
		w,
		"pages/blog",
	)
}
