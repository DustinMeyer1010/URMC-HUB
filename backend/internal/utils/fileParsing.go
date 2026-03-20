package utils

/*



import (
	"bufio"
	"io"
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	excel "github.com/xuri/excelize/v2"
)

func ParseXLSX(file multipart.File) *excel.File {

	f, err := excel.OpenReader(file)

	if err != nil {
		return nil
	}

	sheetNames := f.GetSheetList()
	rows, err := f.GetRows(sheetNames[0])

	if err != nil {
		return nil
	}

	var values []string = make([]string, 0)

	for _, row := range rows {
		input := strings.Join(row, " ")
		values = append(values, input)
	}

	unique, duplicate, notFound, err := ad.BulkLookup(values)

	if err != nil {
		return nil
	}

	return createExcelFile(unique, duplicate, notFound)
}

func ParseCSV(file multipart.File) *excel.File {

	scanner := bufio.NewScanner(file)

	var values []string
	for scanner.Scan() {
		input := strings.ReplaceAll(scanner.Text(), ",", " ")
		input = strings.Trim(input, " ")
		values = append(values, input)
	}

	unique, duplicate, notFound, err := ad.BulkLookup(values)

	if err != nil {
		return nil
	}

	return createExcelFile(unique, duplicate, notFound)
}

func ParsePlainText(file multipart.File) *excel.File {

	content, err := io.ReadAll(file)

	if err != nil {
		return nil
	}

	values := strings.Split(string(content), "\n")

	unique, duplicate, notFound, err := ad.BulkLookup(values)

	if err != nil {
		return nil
	}

	return createExcelFile(unique, duplicate, notFound)
}

func ParseValuesArray(values []string) *excel.File {

	unique, duplicate, notFound, err := ad.BulkLookup(values)

	if err != nil {
		return nil
	}

	return createExcelFile(unique, duplicate, notFound)
}
*/
