package utils

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	excel "github.com/xuri/excelize/v2"
)

func createExcelFile(single []ad.BulkSearchResult, duplicates [][]ad.BulkSearchResult, notFound []ad.BulkSearchResult) *excel.File {
	f := excel.NewFile()
	sheetName := "Sheet1"

	f.SetColWidth(sheetName, "A", "D", 50)

	currentRow := 1

	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "Username")
	f.SetCellValue(sheetName, "C1", "Email")
	f.SetCellValue(sheetName, "D1", "Department")

	for _, user := range single {
		currentRow += 1
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), user.UserDetails.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), user.UserDetails.Username)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), user.UserDetails.Email)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", currentRow), user.UserDetails.Department)
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
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), "Search Value: ")
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), foundUser[0].Value)
		currentRow += 1
		for _, user := range foundUser {
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), user.UserDetails.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), user.UserDetails.Username)
			f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), user.UserDetails.Email)
			f.SetCellValue(sheetName, fmt.Sprintf("D%d", currentRow), user.UserDetails.Department)
			currentRow += 1
		}
		currentRow += 1
	}

	currentRow += 1

	f.MergeCell(sheetName, fmt.Sprintf("A%d", currentRow), fmt.Sprintf("C%d", currentRow))
	f.SetCellStyle(sheetName, fmt.Sprintf("A%d", currentRow), fmt.Sprintf("A%d", currentRow), style)
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), "Not Found")
	currentRow += 1

	for _, n := range notFound {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), "Search Value: ")
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), n.Value)
		currentRow += 1
	}

	return f

}

/*
func CreateAllMembersExcel(members []models.UserSimpleInfo) *excel.File {

	f := excel.NewFile()
	sheetName := "Sheet1"

	f.SetColWidth(sheetName, "A", "D", 50)

	currentRow := 1

	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "Username")
	f.SetCellValue(sheetName, "C1", "NET_ID")
	f.SetCellValue(sheetName, "D1", "EMAIL")
	f.SetCellValue(sheetName, "D1", "OU")

	for _, member := range members {
		currentRow += 1
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), member.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), member.Username)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), member.NetID)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", currentRow), member.Email)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", currentRow), member.OU)
	}

	return f

}
*/
