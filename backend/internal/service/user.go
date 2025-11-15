package service

import (
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

func BulkUserSearch(files []*multipart.FileHeader) error {
	for _, header := range files {
		fileExtension := strings.Split(header.Filename, ".")[1]
		file, err := header.Open()

		if err != nil {
			continue
		}

		switch fileExtension {
		case "xlsx":
			utils.ParseXLSX(file)
		case "txt":
			utils.ParsePlainText(file)
		case "csv":
			utils.ParseCSV(file)
		default:
			continue
		}

	}

	return nil
}

/*
	data, _ := io.ReadAll(file)

	f, err := excelize.OpenReader(bytes.NewReader(data))
	sheetNames := f.GetSheetList()
	rows, err := f.GetRows(sheetNames[0])

	if err != nil {
		log.Fatal(err)
	}

	for i, row := range rows {
		if i == 1 {
			fmt.Printf("Row %d: %v\n", i+1, row)
		}

	}
*/
