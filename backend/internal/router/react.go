package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//go:embed dist/**
var EmbeddedFiles embed.FS

//go:embed dist/index.html
var indexHTML []byte

// Setups the routes for the react routes
func reactRoutes(mux *mux.Router) {

	distFS, err := fs.Sub(EmbeddedFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(distFS))
	mux.PathPrefix("/assets/").Handler(fileServer)
	mux.PathPrefix("/static/").Handler(fileServer)
	mux.PathPrefix("/favicon.ico").Handler(fileServer)
	mux.PathPrefix("/manifest.json").Handler(fileServer)

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
