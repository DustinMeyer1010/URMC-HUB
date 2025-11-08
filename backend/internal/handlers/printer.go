package handlers

import (
	"encoding/json"
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
