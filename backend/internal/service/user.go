package service

import (
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
	excel "github.com/xuri/excelize/v2"
)

func BulkUserSearch(files []*multipart.FileHeader) *excel.File {
	for _, header := range files {
		fileExtension := strings.Split(header.Filename, ".")[1]
		file, err := header.Open()

		if err != nil {
			continue
		}

		switch fileExtension {
		case "xlsx":
			f := utils.ParseXLSX(file)
			return f
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

func BulkUserSearchValues(values []string) *excel.File {

	f := utils.ParseValuesArray(values)

	return f
}

/*

 */
