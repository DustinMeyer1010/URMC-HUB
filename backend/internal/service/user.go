package service

import (
	"encoding/json"
	"mime/multipart"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
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

func GetMemberOf(username string) ([]models.GroupSimpleInfo, *customError.Error) {
	return ad.PullUserMembersOf(username)
}

// Find user based on the dn given and return attributes mapped to their values.
// If no user is found then it will return a NOT_FOUND error
func GetUser(dn string, attributes ...string) ([]byte, *customError.Error) {
	attrs, cError := ad.LookupUser(dn, attributes...)

	jsonData, _ := json.Marshal(attrs)

	return jsonData, cError
}

// Find user based on DN provided and will return all avaiable attrubtes for
// that specific user. If no user is found than it will return a NOT_FOUND error
func GetUserAvaiableAttributes(dn string) ([]byte, *customError.Error) {
	attr, _ := ad.LookupUser(dn, "*")

	var allAttributesNames strings.Builder
	for k := range attr {
		allAttributesNames.WriteString(k + "\n")
	}
	return []byte(allAttributesNames.String()), nil
}
