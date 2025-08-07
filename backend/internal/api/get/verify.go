package get

import (
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Verifing...")

	err := ad.Verify()

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
