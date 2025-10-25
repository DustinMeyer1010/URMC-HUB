package ad

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Will have to be redone
func SearchAllDrives(searchValue string) ([]models.DriveSimpleInfo, error) {
	allDrives := make([]models.DriveSimpleInfo, 0)
	mapping, err := getDriveToGroupsMapping()

	fmt.Println(err)

	if err != nil {
		return allDrives, err
	}

	fmt.Println(searchValue)

	for drive, groups := range mapping {
		if strings.Contains(strings.ToLower(drive), strings.ToLower(searchValue)) {
			allDrives = append(
				allDrives,
				models.DriveSimpleInfo{
					Groups: groups,
					Drive:  drive,
				}.GetLocalPath(),
			)
		} else if checkForGroupMatch(searchValue, groups) {
			allDrives = append(
				allDrives,
				models.DriveSimpleInfo{
					Groups: groups,
					Drive:  drive,
				}.GetLocalPath(),
			)
		}
		if len(allDrives) >= 100 {
			return allDrives, err
		}
	}

	return allDrives, err
}

func checkForGroupMatch(searchValue string, groups []string) bool {
	for _, group := range groups {
		if strings.Contains(strings.ToLower(group), strings.ToLower(searchValue)) {
			return true
		}
	}
	return false
}

// Finds all share drives that match the searchValues
func getGroupToDrivesMapping() (map[string][]string, error) {

	collection := make(map[string][]string)

	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		return make(map[string][]string), err
	}

	for scanner.Scan() {
		line := scanner.Text()
		adGroup, sharedrives := seperateGroupFromShareDrives(line)
		collection[adGroup] = append(collection[adGroup], sharedrives...)
	}

	return collection, err
}

func getDriveToGroupsMapping() (map[string][]string, error) {
	collection := make(map[string][]string)
	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		return make(map[string][]string), err
	}

	for scanner.Scan() {
		line := scanner.Text()
		adGroup, drives := seperateGroupFromShareDrives(line)
		for _, drive := range drives {
			collection[drive] = append(collection[drive], adGroup)
		}
	}

	return collection, err

}

// Takes a line and sperates the group from the share drives on that line
func seperateGroupFromShareDrives(line string) (string, []string) {
	splitLine := strings.Split(line, "|")
	adGroup, sharedrives := splitLine[0][8:], seperateShareDrives(splitLine[1])
	return adGroup, sharedrives
}

// Seperates the share drives from eachother on the current line
func seperateShareDrives(line string) []string {
	line = strings.ReplaceAll(line, "~", "")
	sharedrives := strings.Split(line, ",")

	return sharedrives

}

func openLogonServer() (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(global.LOGON)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, file, err
}
