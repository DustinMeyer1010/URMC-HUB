package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func ComputerInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	computer := vars["computer"]

	computerResponse, err := service.ComputerInfo(computer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(computerResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
