package service

import (
	"encoding/json"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

// GetComputer retrieves a computer's attributes by DistinguishedName and returns them as JSON.
// If attributes is empty, it returns the DistinguishedName. Returns a NOT_FOUND error, if
// the computer is not found
func GetComputer(dn string, attributes ...string) ([]byte, error) {

	attr, _ := ad.LookupComputer(dn, attributes...)

	jsonData, _ := json.Marshal(attr)

	return jsonData, nil

}

// GetComputer retrieves a computer's attributes by DistinguishedName and returns them as []byte.
// Returns a NOT_FOUND error if computer is not found
func GetComputerAvaiableAttributes(dn string) ([]byte, error) {
	attr, _ := ad.LookupComputer(dn, "*")

	var allAttributesNames strings.Builder
	for k := range attr {
		allAttributesNames.WriteString(k + "\n")
	}
	return []byte(allAttributesNames.String()), nil
}
