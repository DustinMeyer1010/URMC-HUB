package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/gorilla/mux"
)

func Bookmarks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	case "POST":
		AddBookmark(w, r)
		return
	case "UPDATE":
		return
	case "DELETE":
		RemoveBookmark(w, r)
	}
}

func AddBookmark(w http.ResponseWriter, r *http.Request) {

	var link models.Link

	err := json.NewDecoder(r.Body).Decode(&link)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid link"))
		return
	}

	err = db.AddLink(link)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveBookmark(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	varid := vars["id"]

	id, err := strconv.Atoi(varid)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not a valid ID for bookmark"))
		return
	}

	err = db.RemoveLink(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

}
