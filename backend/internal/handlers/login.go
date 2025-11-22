package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	if err := service.Login(user); err != nil {
		http.Error(w, "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCESSFUL_LOGIN"))
}
