package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20)

	image, header, err := r.FormFile("image")
	var bookmark models.Bookmark
	json.Unmarshal([]byte(r.FormValue("bookmark")), &bookmark)

	if err != nil {
		http.Error(w, "missing image", http.StatusBadRequest)
		return
	}

	defer image.Close()

	bookmark.ImagePath = db.SaveImage(image, header)

	db.AddBookmark(bookmark)

	/*
		err = service.AddBookmark(r)

		if err != nil {
			fmt.Println(err)
			return
		}
	*/

}

func GetGeneralBookmarks(w http.ResponseWriter, r *http.Request) {
	bookmarks := db.GenericBookmarks()

	w.WriteHeader(http.StatusOK)
	w.Write(bookmarks)
}

// Returns the generic bookmarks for
func GetAgentsWithBookmarks(w http.ResponseWriter, r *http.Request) {
	agents := service.GetAgentsWithBookmarks()
	jsonData, _ := json.Marshal(agents)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetBookForAgent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	bookmarks, _ := db.GetBookmark(username)

	jsonData, _ := json.Marshal(bookmarks)

	w.Write(jsonData)
}

func SaveBookmark(w http.ResponseWriter, r *http.Request) {

}
