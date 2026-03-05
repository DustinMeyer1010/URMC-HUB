package utils

import "strings"

// Pull the commonName of the DN
func DNToNameParser(dn string) string {
	categories := strings.Split(dn, ",")

	if len(categories) == 0 {
		return ""
	}
	name := categories[0]

	return strings.ReplaceAll(name, "CN=", "")
}

// Takes multiple DN and will parse them to get the commonName out of them
func DNsToNamesParser(dns []string) []string {
	names := []string{}
	for _, dn := range dns {
		names = append(names, DNToNameParser(dn))
	}

	return names
}
