package utils

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	excel "github.com/xuri/excelize/v2"
)

func createExcelFile(single []models.UserDetails, duplicates [][]models.UserDetails) *excel.File {
	f := excel.NewFile()
	sheetName := "Sheet1"

	f.SetColWidth(sheetName, "A", "C", 50)

	currentRow := 1

	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "Username")
	f.SetCellValue(sheetName, "C1", "Email")

	for _, user := range single {
		currentRow += 1
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), user.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), user.Username)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), user.Email)
	}

	currentRow += 3

	style, _ := f.NewStyle(&excel.Style{
		Alignment: &excel.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	f.MergeCell(sheetName, fmt.Sprintf("A%d", currentRow), fmt.Sprintf("C%d", currentRow))
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", currentRow), fmt.Sprintf("A%d", currentRow), style)
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), "Duplicates Found")

	currentRow += 1

	for _, foundUser := range duplicates {
		for _, user := range foundUser {
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), user.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), user.Username)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), user.Email)
			currentRow += 1
		}
		currentRow += 1
	}

	return f

}
