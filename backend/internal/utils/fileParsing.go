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

	if err := f.SaveAs("C:\\Users\\dmeyer20\\Downloads\\example.xlsx"); err != nil {
		fmt.Println("Error saving file:", err)
	}

	return f

}

func ParseCSV(file multipart.File) {

}

func ParsePlainText(file multipart.File) {}
