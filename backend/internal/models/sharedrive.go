package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

type ShareDriveSimpleInfo struct {
	Group     []string `json:"groups"`
	Drive     string   `json:"drive"`
	LocalPath string   `json:"local_path"`
}

func FindShareDrive(searchValue string) (shareDrives []ShareDriveSimpleInfo, err error) {
	shareDrives = make([]ShareDriveSimpleInfo, 0)
	searchValue = strings.ToLower(searchValue)

	file, err := os.Open(global.LOGON)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	collection := make(map[string][]string)

	// Loop through the each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(strings.ToLower(line), searchValue) {
			continue
		}
		adGroup, sharedrives := seperateGroupFromShareDrives(line)
		if _, ok := collection[adGroup]; ok {
			collection[adGroup] = append(collection[adGroup], sharedrives...)
		} else {
			collection[adGroup] = sharedrives
		}
	}

	fmt.Println(collection)

	return
}

func seperateGroupFromShareDrives(line string) (string, []string) {
	splitLine := strings.Split(line, "|")
	adGroup, sharedrives := splitLine[0][8:], seperateShareDrives(splitLine[1])
	return adGroup, sharedrives
}

func seperateShareDrives(line string) []string {
	line = strings.ReplaceAll(line, "~", "")
	sharedrives := strings.Split(line, ",")

	return sharedrives

}
