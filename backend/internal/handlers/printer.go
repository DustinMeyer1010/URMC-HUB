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

	printer, cError := service.PullPrinterInformation(server, queue)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, jsonError := json.Marshal(printer)

	if jsonError != nil {
		http.Error(w, jsonError.Error(), http.StatusInternalServerError)
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

	printers, cError := service.RelatedPrinters(ip)

	if cError != nil {
		http.Error(w, cError.Type, cError.StatusCode)
		w.Write([]byte(cError.Msg))
		return
	}

	jsonData, _ := json.Marshal(printers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
