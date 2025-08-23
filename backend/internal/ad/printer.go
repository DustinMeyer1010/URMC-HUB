package ad

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Matches printer to the search
func SearchAllPrinters(searchValue string) (printers []models.Printer, err error) {
	printers = make([]models.Printer, 0)
	printersList, err := fetchPrinters()

	if err != nil {
		return
	}

	for _, printer := range printersList {
		compare := fmt.Sprintf("\\\\%s\\%s %s %s %s %s %s", printer.Server, printer.Queue, printer.Model, printer.IP, printer.PrintProcessor, printer.Location, printer.Notes)
		if strings.Contains(strings.ToLower(compare), strings.ToLower(searchValue)) {
			printers = append(printers, printer)
		}
	}
	return
}

// Retrieves the printer queue from the server
func fetchPrinters() (printers []models.Printer, err error) {
	printers = make([]models.Printer, 0)
	// Make a GET request
	resp, err := http.Get("https://apps.mc.rochester.edu/ISD/SIG/PrintQueues/PrintQReport.csv")

	if err != nil {
		return
	}
	defer resp.Body.Close()

	file := csv.NewReader(resp.Body)

	records, err := file.ReadAll()

	if err != nil {
		return
	}

	for _, record := range records {
		printers = append(printers, models.Printer{
			Server:         record[0],
			Queue:          record[1],
			Model:          record[2],
			IP:             record[3],
			PrintProcessor: record[4],
			Location:       record[5],
			Notes:          record[6],
		})
	}

	return
}
