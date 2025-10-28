package handlers

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

// Verify that the users creds are still valid
func Verify(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Verifing...")

	err := ad.Verify()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("Authorized")
	w.WriteHeader(http.StatusOK)
}
