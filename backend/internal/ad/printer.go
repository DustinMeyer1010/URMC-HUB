package ad

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Matches printer to the search
func SearchAllPrinters(searchValue string) ([]models.PrinterSimpleInfo, error) {
	printers := make([]models.PrinterSimpleInfo, 0)
	printersList, cError := fetchPrinters()

	if cError != nil {
		return printers, cError
	}

	for _, printer := range printersList {
		if checkForMatchingPrinter(printer, searchValue) {
			printers = append(printers, printer)
		}

		// Only want 100 matching printers
		if len(printers) > 100 {
			break
		}

	}

	return printers, nil
}

func checkForMatchingPrinter(printer models.PrinterSimpleInfo, searchValue string) bool {
	compare := fmt.Sprintf("\\\\%s\\%s %s %s %s %s %s", printer.Server, printer.Queue, printer.Model, printer.IP, printer.PrintProcessor, printer.Location, printer.Notes)

	if strings.Contains(strings.ToLower(compare), strings.ToLower(searchValue)) {
		return true
	}

	return false
}

// Retrieves the printer queue from the server
func fetchPrinters() ([]models.PrinterSimpleInfo, error) {

	// Make a GET request
	resp, requestError := http.Get(global.PRINTQUEUEREPORT)

	if requestError != nil {
		logger.Error(requestError)
		cError := errs.REQUEST_ERROR.NewError(requestError)
		return nil, &cError
	}
	defer resp.Body.Close()

	printers, csvError := readPrinterFromExcel(csv.NewReader(resp.Body))

	if csvError != nil {
		return nil, csvError
	}

	return printers, nil
}

func readPrinterFromExcel(file *csv.Reader) ([]models.PrinterSimpleInfo, error) {

	records, fileReadError := file.ReadAll()

	if fileReadError != nil {
		logger.Error(fileReadError)
		cError := errs.READ_FILE_ERROR.NewError(fileReadError)
		return nil, &cError
	}

	var printers []models.PrinterSimpleInfo

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

func PullSinglePrinterInformation(printer string) (models.PrinterSimpleInfo, error) {
	printersList, cError := fetchPrinters()

	if cError != nil {
		return models.PrinterSimpleInfo{}, cError
	}

	for _, p := range printersList {
		if fmt.Sprintf("\\\\%s\\%s", p.Server, p.Queue) == printer {
			return p, nil
		}
	}

	cError1 := errs.NOT_FOUND.NewMessage(fmt.Sprintf("NO PRINTER FOUND FOR: %s", printer))

	return models.PrinterSimpleInfo{}, &cError1

}

func RelatedPrinters(ip string) ([]models.PrinterSimpleInfo, error) {
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
