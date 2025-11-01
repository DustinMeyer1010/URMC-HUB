package handlers

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	err := service.AddBookmark(r)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func GenerateGenericBookmarks(w http.ResponseWriter, r *http.Request) {

	err := service.GenerateGenericBookmarks()

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generate generics \n%s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Generated Generics"))
}
