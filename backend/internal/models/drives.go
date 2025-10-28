package models

import (
	"bufio"
	"os"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

// DriveSimpleInfo contains minimal drive information used for listing or syncing drives.
// - Groups: The names of AD groups or users with access.
// - Drive: The mapped drive letter or identifier.
// - LocalPath: The local system path where the drive is mounted.
type DriveSimpleInfo struct {
	Groups    []string `json:"groups"`
	Drive     string   `json:"drive"`
	LocalPath string   `json:"local_path"`
}

type DriveAccess struct {
	Groups []string `json:"groups"`
	Drive  string   `json:"drive"`
}

// Get the local path for the provided share drive and then return a copy of that share drive with the localPath in it
func (d DriveSimpleInfo) GetLocalPath() DriveSimpleInfo {
	file, err := os.Open(global.SHARES)

	if err != nil {
		return d
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "|")
		drive, localPath := strings.ToLower(splitLine[0]), splitLine[1]
		if strings.Contains(d.Drive, "\\\\") {
			if drive == strings.Split(strings.ToLower(d.Drive), "\\\\")[1] {
				d.LocalPath = localPath
				return d
			}
			continue
		}
		if drive == strings.ToLower(d.Drive) {
			d.LocalPath = localPath
			return d
		}
	}
	return d
}
