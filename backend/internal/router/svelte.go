package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//go:embed all:build
var frontendFS embed.FS

//go:embed build/200.html
var indexHTML []byte

// Setups the routes for the react routes
func reactSvelte(mux *mux.Router) {
	// 1. Create the sub-filesystem for the build folder
	distFS, err := fs.Sub(frontendFS, "build")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(distFS))

	// 2. Explicitly handle SvelteKit's static assets folder
	// Note: removed the * as PathPrefix handles everything under the prefix
	mux.PathPrefix("/_app/").Handler(fileServer)

	// 3. Handle favicon, robots.txt, and other root-level static files
	// and provide the SPA fallback for everything else.
	mux.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := path.Clean(r.URL.Path)

		// Trim the leading slash for fs.Open
		trimmedPath := requestPath[1:]

		// Try to open the file in the embedded FS
		f, err := distFS.Open(trimmedPath)
		if err == nil {
			// File exists! Close it and let FileServer handle the actual delivery
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		// 4. Fallback: If the file doesn't exist, it's likely a SvelteKit route
		// Serve the index/fallback HTML so SvelteKit's client-side router can take over.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(indexHTML)
	})
}
