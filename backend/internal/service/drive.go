package service

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func GetDriveAccess(groups []string) []models.DriveAccess {
	var result []models.DriveAccess
	var accessMapping map[string][]string = make(map[string][]string)
	groupToDrive, _ := ad.GetGroupToDrivesMapping()

	for _, group := range groups {
		if drives, ok := groupToDrive[group]; ok {
			for _, drive := range drives {
				if _, ok := accessMapping[drive]; ok {
					accessMapping[drive] = append(accessMapping[drive], group)
					continue
				}
				accessMapping[drive] = []string{group}
			}
		}
	}
	for drive, groups := range accessMapping {
		result = append(result, models.DriveAccess{Drive: drive, Groups: groups})
	}
	return result
}
