package service

import (
	"encoding/json"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

func SearchAll(searchValue string) []byte {
	results := ad.SearchAll(searchValue)

	jsonData, _ := json.Marshal(results)

	return jsonData
}

func SearchAllUsers(searchValue string, attributes ...string) ([]byte, error) {
	users := new([]map[string][]string)
	cError := ad.SearchAllUserNew(users, searchValue, attributes...)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(users)

	return jsonData, nil
}

func SearchAllGroups(searchValue string, attributes ...string) ([]byte, error) {
	groups := new([]map[string][]string)
	cError := ad.SearchAllGroupsNew(groups, searchValue, attributes...)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(groups)

	return jsonData, nil
}

func SearchAllComputers(searchValue string, attributes ...string) ([]byte, error) {
	computers := new([]map[string][]string)
	cError := ad.SearchAllComputersNew(computers, searchValue, attributes...)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(computers)

	return jsonData, nil
}

func SearchAllDrives(searchValue string) ([]byte, error) {
	drives, cError := ad.SearchAllDrives(searchValue)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(drives)

	return jsonData, nil
}

func SearchAllPrinters(searchValue string) ([]byte, error) {
	printers, cError := ad.SearchAllPrinters(searchValue)

	if cError != nil {
		return []byte(""), cError
	}

	jsonData, _ := json.Marshal(printers)

	return jsonData, nil
}
