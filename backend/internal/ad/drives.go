package ad

import (
	"fmt"
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

// Returns collection of Group to the Drive that the specific group access to
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

// Return collection of Drive to the Groups that give access to that drive
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

func GetShareDriveGroups(drive string) models.DrivePageInfo {
	drives, _ := SearchAllDrives(drive)
	foundDrive := models.DriveSimpleInfo{}
	result := models.DrivePageInfo{}

	for _, cur := range drives {
		if cur.Drive == drive {
			foundDrive = cur
		}
	}

	fmt.Println(foundDrive)

	for _, group := range foundDrive.Groups {
		groupInfo, _ := PullGroupInfo(group)
		result.Groups = append(result.Groups, groupInfo)
	}

	result.Drive = foundDrive.Drive
	result.LocalPath = foundDrive.LocalPath

	return result

}
