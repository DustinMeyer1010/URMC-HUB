package parser

import "strings"

func QueryArray(values string) []string {

	return strings.Split(values, ",")
}
