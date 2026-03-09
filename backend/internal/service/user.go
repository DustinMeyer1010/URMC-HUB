package service

import (
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
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
			return utils.ParseXLSX(file)
		case "txt":
			return utils.ParsePlainText(file)
		case "csv":
			return utils.ParseCSV(file)
		default:
			continue
		}
	}

	return nil
}

func BulkUserSearchValues(values []string) *excel.File {

	return utils.ParseValuesArray(values)
}

func GetMemberOf(username string) ([]models.GroupSimpleInfo, error) {
	return ad.PullUserMembersOf(username)
}
