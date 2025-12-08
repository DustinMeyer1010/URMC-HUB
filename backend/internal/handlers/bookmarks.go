package handlers

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	err := service.AddBookmark(r)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func GetGeneralBookmarks(w http.ResponseWriter, r *http.Request) {
	bookmarks := db.GenericBookmarks()

	w.WriteHeader(http.StatusOK)
	w.Write(bookmarks)
}
