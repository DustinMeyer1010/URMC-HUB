package get

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

func PullUserInformation(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)

	user, err := ad.PullUserInformation("dmeyer20")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed"))
		return
	}

	jsonData, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
