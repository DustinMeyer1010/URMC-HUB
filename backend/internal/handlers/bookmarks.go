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

	return
}
