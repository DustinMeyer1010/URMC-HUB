package parser

import (
	"strings"
)

// Seperates the attrubutes in the URL query
func QueryArray(values string) []string {

	// No Attributes given
	if values == "" {
		return []string{}
	}

	return strings.Split(values, ",")
}
