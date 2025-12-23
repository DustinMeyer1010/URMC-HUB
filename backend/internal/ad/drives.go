package ad

import (
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Will have to be redone
func SearchAllDrives(searchValue string) ([]models.DriveSimpleInfo, *customError.Error) {
	allDrives := make([]models.DriveSimpleInfo, 0)
	mapping, err := getDriveToGroupsMapping()

	if err != nil {
		logger.Error(err)
		return allDrives, &customError.FILE_UNREACHABLE
	}

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
			return allDrives, nil
		}
	}

	return allDrives, nil
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
func GetGroupToDrivesMapping() (map[string][]string, *customError.Error) {
	collection := make(map[string][]string)

	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		logger.Error(err)
		return make(map[string][]string), &customError.FILE_UNREACHABLE
	}

	for scanner.Scan() {
		line := scanner.Text()
		adGroup, sharedrives := seperateGroupFromShareDrives(line)
		collection[adGroup] = append(collection[adGroup], sharedrives...)
	}

	return collection, nil
}

func getDriveToGroupsMapping() (map[string][]string, *customError.Error) {
	collection := make(map[string][]string)
	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		logger.Error(err)
		return make(map[string][]string), &customError.FILE_UNREACHABLE
	}

	for scanner.Scan() {
		line := scanner.Text()
		adGroup, drives := seperateGroupFromShareDrives(line)
		for _, drive := range drives {
			collection[drive] = append(collection[drive], adGroup)
		}
	}

	return collection, nil

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
