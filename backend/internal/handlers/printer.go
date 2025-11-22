package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
	"github.com/gorilla/mux"
)

func PrinterInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	server := vars["server"]

	queue := r.URL.Query().Get("queue")

	if queue == "" {
		http.Error(w, "INVALID_QUEUE", http.StatusBadRequest)
		return
	}

	printer, statusError := service.PullPrinterInformation(server, queue)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, http.StatusBadRequest)
		w.Write([]byte(statusError.Error))
		return
	}

	jsonData, err := json.Marshal(printer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func PingPrinter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	err := service.PingPrinter(ip)

	if err != nil {
		http.Error(w, "PING_FAILED", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "PING_SUCCESSFUL")
}

func RelatedPrinters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	printers, statusError := service.RelatedPrinters(ip)

	if statusError != nil {
		http.Error(w, statusError.ErrorType, statusError.HttpStatus)
		w.Write([]byte(statusError.Error))
		return
	}

	jsonData, err := json.Marshal(printers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
