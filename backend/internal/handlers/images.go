package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func StaticPhotos(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	imageName := vars["imageName"]
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)

	image := filepath.Join(exeDir, "URMC_HUB_IMAGES", imageName)

	http.ServeFile(w, r, image)
}
