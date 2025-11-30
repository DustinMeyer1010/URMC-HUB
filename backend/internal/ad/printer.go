package ad

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Matches printer to the search
func SearchAllPrinters(searchValue string) ([]models.PrinterSimpleInfo, *customError.Error) {
	printers := make([]models.PrinterSimpleInfo, 0)
	printersList, cError := fetchPrinters()

	if cError != nil {
		return printers, cError
	}

	for _, printer := range printersList {
		compare := fmt.Sprintf("\\\\%s\\%s %s %s %s %s %s", printer.Server, printer.Queue, printer.Model, printer.IP, printer.PrintProcessor, printer.Location, printer.Notes)
		if strings.Contains(strings.ToLower(compare), strings.ToLower(searchValue)) {
			printers = append(printers, printer)
		}
	}
	// Only returning the top 100 results
	if len(printers) > 100 {
		printers = printers[0:100]
	}

	return printers, nil
}

// Retrieves the printer queue from the server
func fetchPrinters() ([]models.PrinterSimpleInfo, *customError.Error) {
	printers := make([]models.PrinterSimpleInfo, 0)
	// Make a GET request
	resp, requestError := http.Get("https://apps.mc.rochester.edu/ISD/SIG/PrintQueues/PrintQReport.csv")

	if requestError != nil {
		logger.Error(requestError)
		cError := customError.REQUEST_ERROR.NewError(requestError)
		return printers, &cError
	}
	defer resp.Body.Close()

	file := csv.NewReader(resp.Body)

	records, fileReadError := file.ReadAll()

	if fileReadError != nil {
		logger.Error(fileReadError)
		cError := customError.READ_FILE_ERROR.NewError(fileReadError)
		return printers, &cError
	}

	for _, record := range records {
		printers = append(printers, models.PrinterSimpleInfo{
			Server:         record[0],
			Queue:          record[1],
			Model:          record[2],
			IP:             record[3],
			PrintProcessor: record[4],
			Location:       record[5],
			Notes:          record[6],
		})
	}

	return printers, nil
}

func PullSinglePrinterInformation(printer string) (models.PrinterSimpleInfo, *customError.Error) {
	printersList, cError := fetchPrinters()

	if cError != nil {
		return models.PrinterSimpleInfo{}, cError
	}

	for _, p := range printersList {
		if fmt.Sprintf("\\\\%s\\%s", p.Server, p.Queue) == printer {
			return p, nil
		}
	}

	cError1 := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO PRINTER FOUND FOR: %s", printer))

	return models.PrinterSimpleInfo{}, &cError1

}

func RelatedPrinters(ip string) ([]models.PrinterSimpleInfo, *customError.Error) {
	printerList, cError := fetchPrinters()
	var relatedPrinters []models.PrinterSimpleInfo

	if cError != nil {
		return relatedPrinters, cError
	}

	for _, p := range printerList {
		if p.IP == ip {
			relatedPrinters = append(relatedPrinters, p)
		}
	}

	return relatedPrinters, nil
}
