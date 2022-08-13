package static

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func ServeStaticFiles(route string, r chi.Router) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "www"))
	serveDataAssets(route, r)
	fileServer(r, route, filesDir)
	serveFavicon(route, r)
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func serveFavicon(staticRoute string, r chi.Router) {
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", staticRoute+"/favicon.ico")
		w.WriteHeader(http.StatusPermanentRedirect)
	})
}

func serveDataAssets(staticRoute string, r chi.Router) {
	workdir, _ := os.Getwd()
	publicDataDir := http.Dir(filepath.Join(workdir, "/data/public"))
	fileServer(r, "/data", publicDataDir)
}
