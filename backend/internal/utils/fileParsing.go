package utils

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	excel "github.com/xuri/excelize/v2"
)

// TODO Refactor
func ParseXLSX(file multipart.File) *excel.File {

	f, err := excel.OpenReader(file)

	var duplicate [][]models.UserDetails = make([][]models.UserDetails, 0)
	var single []models.UserDetails = make([]models.UserDetails, 0)

	if err != nil {
		return nil
	}

	sheetNames := f.GetSheetList()
	rows, err := f.GetRows(sheetNames[0])

	if err != nil {
		return nil
	}

	err = ad.CreatePresistantConn()

	if err != nil {
		return nil
	}

	for _, row := range rows {
		input := strings.Join(row, " ")
		results, err := ad.UserDetails(input)

		if err != nil && err.Error() != "NOT_FOUND" {
			fmt.Println("no account found")
			continue
		}

		if len(results) == 0 {
			continue
		}

		if len(results) == 1 {
			single = append(single, results[0])
			continue
		}

		duplicate = append(duplicate, results)

	}

	ad.DisconnectPresistantConn()

	return createExcelFile(single, duplicate)
}

func ParseCSV(file multipart.File) {

}

func ParsePlainText(file multipart.File) {}
