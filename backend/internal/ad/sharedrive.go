package ad

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAllShareDrive(searchValue string) (shareDrives []models.ShareDrive, err error) {
	shareDrives = make([]models.ShareDrive, 0)
	searchValue = strings.ToLower(searchValue)
	networkPath := global.LOGON
	file, err := os.Open(networkPath)

	if err != nil {
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1] // Removes last element which is blank space
	foundDrives := make(map[string][]string)

	for _, line := range lines {
		parse := strings.Split(line, "|")
		group, drives := strings.ToLower(parse[0]), strings.Split(parse[1], ",")
		for _, drive := range drives {
			currentDrive := strings.ToLower(strings.TrimSpace(drive[1:]))
			// Checks to see if the input is container in either the group or share drive
			if strings.Contains(currentDrive, searchValue) || strings.Contains(group, searchValue) {
				foundDrives[currentDrive] = append(foundDrives[currentDrive], group[8:])

			}
		}
	}
	shareDrives, err = findLocalPath(foundDrives)

	return

}

func findLocalPath(foundDrives map[string][]string) (shareDrives []models.ShareDrive, err error) {
	shareDrives = make([]models.ShareDrive, 0)
	var localPath string
	networkPath := global.SHARES
	file, err := os.Open(networkPath)

	if err != nil {
		return
	}

	data, err := io.ReadAll(file)

	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	for drive, groups := range foundDrives {
		for _, line := range lines {
			parse := strings.Split(line, "|")
			if strings.Contains(strings.ToLower(parse[0]), strings.TrimSpace(drive[2:])) {
				localPath = parse[1]
				break
			}
		}
		sort.Strings(groups)
		shareDrives = append(shareDrives, models.ShareDrive{Drive: drive, Group: groups, LocalPath: localPath})
	}
	return
}

func CheckGroupForShareDrive(group string) (*models.ShareDrive, error) {
	shareDrive, err := SearchAllShareDrive(group)

	if err != nil {
		return nil, err
	}

	if len(shareDrive) != 0 {
		return &models.ShareDrive{shareDrive[0].Group, shareDrive[0].Drive, shareDrive[0].LocalPath, shareDrive[0].Type}, nil
	}

	return nil, nil
}

func FindShareDriveInfo(share string) (shareDrive ShareDrivePage, err error) {
	share = strings.ToLower(share)
	networkPath := global.LOGON
	file, err := os.Open(networkPath)

	if err != nil {
		return
	}

	defer file.Close()

	data, _ := io.ReadAll(file)
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	groupsFound := make([]string, 0)
	l, err := ConnectToServer("LDAP://urmc-sh.rochester.edu/")

	if err != nil {
		return
	}

	shareDrive.Sharedrive = share

	var wg sync.WaitGroup

	for _, line := range lines {
		parse := strings.Split(line, "|")
		group, drives := parse[0][8:], strings.Split(parse[1], ",")

		for _, drive := range drives {

			drive = strings.ToLower(strings.TrimSpace(drive))[1:]
			if drive == share {
				fmt.Println(drive)
				wg.Add(1)
				go func() {
					defer wg.Done()
					results, err := GroupInfo(group, l, "urmc-sh")
					if err != nil {
						return
					}
					shareDrive.Groups = append(shareDrive.Groups, results)
					groupsFound = append(groupsFound, group)
				}()
			}
		}

	}
	wg.Wait()

	if len(groupsFound) == 0 {
		return
	}

	return
}
