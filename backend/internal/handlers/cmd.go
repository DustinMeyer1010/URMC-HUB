package handlers

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	host := query.Get("host")

	data, err := utils.NSLookUp(host)

	fmt.Println(data, err)
}

func NSLookUp(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	host := query.Get("host")

	data, err := utils.NSLookUp(host)

	fmt.Println(data, err)
}
