package service

import (
	"encoding/json"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
)

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
