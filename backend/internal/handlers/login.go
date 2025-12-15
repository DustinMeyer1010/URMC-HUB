package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	logger.LogRequestInfo(r.Method, r.URL.Path)
	var user models.UserLogin

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "INVALID_BODY", http.StatusBadRequest)
		return
	}

	if err := service.Login(user); err != nil {
		http.Error(w, "INVALID_CREDENTIALS", http.StatusUnauthorized)
		return
	}

	if !utils.FileExist(fmt.Sprintf("%s/%s.db", db.DBLocation, global.Username)) {
		logger.Infof("Generating DB for %s", global.Username)
		err := db.AgentDatabaseInit()
		if err != nil {
			logger.Errorf("Failed to generate Agent database\n Error: %s", err.Error())
		}

		logger.Infof("Database Generated Successfully Location: %s", db.DBLocation)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    user.Username,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Secure:   false, // must be false on localhost
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCESSFUL_LOGIN"))
}
