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

	computerResponse, cError := service.ComputerInfo(computer)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(computerResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
