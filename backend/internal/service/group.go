package service

import (
	"encoding/json"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
	excel "github.com/xuri/excelize/v2"
)

func GetAllMembers(group string) ([]string, *customError.Error) {

	return ad.GetAllMembers(group)
}

func GetExcelOfAllMembers(group string) *excel.File {
	membersDNs, _ := GetAllMembers(group)
	members := ad.PullMembersInformation(membersDNs)
	return utils.CreateAllMembersExcel(members)
}

// Retrieves a group's attributes by DistinguishedName and returns them as JSON.
// If attributes is empty, it returns the DistinguishedName. Returns a NOT_FOUND error, if
// the group is not found
func GetGroup(dn string, attributes ...string) ([]byte, *customError.Error) {

	attr, _ := ad.LookupGroup(dn, attributes...)

	jsonData, _ := json.Marshal(attr)

	return jsonData, nil

}

// Retrieves a group's attributes by DistinguishedName and returns them as []byte.
// Returns a NOT_FOUND error if group is not found
func GetGroupAvaiableAttributes(dn string) ([]byte, *customError.Error) {
	attr, _ := ad.LookupGroup(dn, "*")

	var allAttributesNames strings.Builder
	for k := range attr {
		allAttributesNames.WriteString(k + "\n")
	}
	return []byte(allAttributesNames.String()), nil
}
