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
		http.Error(w, "Invaid printer name", http.StatusBadRequest)
		return
	}

	printer, err := service.PullPrinterInformation(server, queue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(printer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func PingPrinter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	err := service.PingPrinter(ip)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Failed to Ping: %s", ip)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully Pinged: %s", ip)
}

func RelatedPrinters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := vars["ip"]

	printers, err := service.RelatedPrinters(ip)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Failed to Ping: %s", ip)
		return
	}

	jsonData, err := json.Marshal(printers)

	if err != nil {
		http.Error(w, "Error Creating Resposne", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
