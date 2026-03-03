package parser

import "strings"

// Seperates the attrubutes in the URL query
func QueryArray(values string) []string {

	return strings.Split(values, ",")
}
