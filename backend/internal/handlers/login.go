package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		panic(err)
	}

	err = ad.Login(user)

	fmt.Println(err)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Logged in"))
}
