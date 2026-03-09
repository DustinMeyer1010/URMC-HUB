package ad

import (
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/errs"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// REFACTOR: This will be refactored into smaller functions and make it easier to read
func SearchAllDrives(searchValue string) ([]models.DriveSimpleInfo, error) {
	allDrives := make([]models.DriveSimpleInfo, 0)
	mapping, err := getDriveToGroupsMapping()

	if err != nil {
		logger.Error(err)
		return allDrives, &errs.FILE_UNREACHABLE
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
func GetGroupToDrivesMapping() (map[string][]string, error) {
	collection := make(map[string][]string)

	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		logger.Error(err)
		return make(map[string][]string), &errs.FILE_UNREACHABLE
	}

	for scanner.Scan() {
		line := scanner.Text()
		adGroup, sharedrives := seperateGroupFromShareDrives(line)
		collection[adGroup] = append(collection[adGroup], sharedrives...)
	}

	return collection, nil
}

// Return collection of Drive to the Groups that give access to that drive
func getDriveToGroupsMapping() (map[string][]string, error) {
	collection := make(map[string][]string)
	scanner, file, err := openLogonServer()
	defer file.Close()

	if err != nil {
		logger.Error(err)
		return make(map[string][]string), &errs.FILE_UNREACHABLE
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

	for _, group := range foundDrive.Groups {
		groupInfo, _ := PullGroupInfo(group)
		result.Groups = append(result.Groups, groupInfo)
	}

	result.Drive = foundDrive.Drive
	result.LocalPath = foundDrive.LocalPath

	return result

}

// Given a list of groups will return the drives those groups give access to if any and the drives will be the key and the groups will be the values which means those groups give access to that drive
func DriveAccess(groups []string) ([]models.DriveAccess, error) {
	var result []models.DriveAccess
	var driveAccess map[string][]string = make(map[string][]string)
	groupToDrive, cError := GetGroupToDrivesMapping()

	if cError != nil {
		return result, cError
	}

	for _, group := range groups {
		if drives, ok := groupToDrive[group]; ok {
			for _, drive := range drives {
				if _, ok := driveAccess[drive]; ok {
					driveAccess[drive] = append(driveAccess[drive], group)
					continue
				}
				driveAccess[drive] = []string{group}
			}
		}
	}

	for drive, groups := range driveAccess {
		result = append(result, models.DriveAccess{Drive: drive, Groups: groups})
	}

	return result, nil
}
