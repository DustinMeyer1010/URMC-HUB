package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
)

// Find user based on the dn given and return attributes mapped to their values.
// If no user is found then it will return a NOT_FOUND error
func GetUser(dn string, attributes ...string) ([]byte, error) {
	attrs, cError := ad.LookupUser(dn, attributes...)

	jsonData, _ := json.Marshal(attrs)

	return jsonData, cError
}

// Find user based on DN provided and will return all avaiable attrubtes for
// that specific user. If no user is found than it will return a NOT_FOUND error
func GetUserAvaiableAttributes(dn string) ([]byte, error) {
	attr, _ := ad.LookupUser(dn, "*")

	var allAttributesNames strings.Builder
	for k := range attr {
		allAttributesNames.WriteString(k + "\n")
	}
	return []byte(allAttributesNames.String()), nil
}

// Find user based on DN. Finds all drives they have access to based
// on the groups they have. This will return an empty byte if no groups are
// found. Returns error if any problems are ran into
func GetUserDrives(dn string) ([]byte, error) {

	attr, cError := ad.LookupUser(dn, "memberof")

	if cError != nil {
		return []byte(""), cError
	}

	groupNames := utils.DNsToNamesParser(attr["memberof"])

	driveAccess, cError := ad.DriveAccess(groupNames)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(driveAccess)

	return jsonData, nil
}

// Finds user based on DN. Pulls all the groups for that user and will return
// all the groups they are appart off with attributes provided. If no attributes
// provided it will search with samaccountname, cn, dn, information, description.
func GetUserGroups(dn string, attributes ...string) ([]byte, error) {
	user, cError := ad.LookupUser(dn, "memberof")

	groups := []map[string][]string{}

	if cError != nil {
		return []byte(""), cError
	}

	groupsDN := user["memberof"]

	for _, g := range groupsDN {
		attr, _ := ad.LookupGroup(g, attributes...)
		groups = append(groups, attr)
	}

	jsonData, _ := json.Marshal(groups)

	return jsonData, nil
}

// GetUserLockoutStatus retrieves account lockout and password metadata for a user
// across all configured URMC domain controllers. It returns the results as a
// JSON-encoded byte slice.
func GetUserLockoutStatus(dn string, attributes ...string) []byte {

	lockout := ad.LookupLockoutStatus(dn)

	jsonData, _ := json.Marshal(lockout)

	return jsonData
}

func GroupAddToUser(userDN, groupDN string) ([]byte, error) {

	err := ad.GroupAddToUser(userDN, groupDN)

	if err != nil {
		fmt.Println(err)
		return []byte("group was not added"), err
	}

	return []byte("group was added"), nil

}

func GroupRemoveFromUser(userDN, groupDN string) ([]byte, error) {
	err := ad.GroupRemoveFromUser(userDN, groupDN)

	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return []byte("group was added"), nil
}
