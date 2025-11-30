package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func ComputerInfo(w http.ResponseWriter, r *http.Request) {

	logger.LogRequestInfo(r.Method, r.URL.Path)

	vars := mux.Vars(r)
	computer := vars["computer"]

	logger.ServerLogger.Infof("Searching For Computer | Input: %s", computer)

	computerResponse, cError := service.ComputerInfo(computer)

	if cError != nil {
		logger.ServerLogger.Errorf("%s %s", cError.Type, cError.Msg)
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	logger.ServerLogger.Infof("Found Computers | Input: %+v", computerResponse)

	jsonData, _ := json.Marshal(computerResponse)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
