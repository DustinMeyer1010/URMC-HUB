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

	computerResponse, statusError := service.ComputerInfo(computer)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, statusError.HttpStatus)
		w.Write([]byte(statusError.Error))
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
