package models

import (
	"bufio"
	"os"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

type DriveSimpleInfo struct {
	Groups    []string `json:"groups"`
	Drive     string   `json:"drive"`
	LocalPath string   `json:"local_path"`
}

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
