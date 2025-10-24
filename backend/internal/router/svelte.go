package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//go:embed build/**
var frontendFS embed.FS

//go:embed build/index.html
var indexHTML []byte

// Setups the routes for the react routes
func reactRoutes(mux *mux.Router) {

	distFS, err := fs.Sub(frontendFS, "build")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(distFS))
	mux.PathPrefix("/_app/").Handler(fileServer)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestPath := path.Clean(r.URL.Path)

		// Try to serve the actual file if it exists
		f, err := distFS.Open(requestPath[1:])
		if err == nil {
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		// Otherwise, serve React’s index.html
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(indexHTML)

	})
}
