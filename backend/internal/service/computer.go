package service

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func ComputerInfo(computer string) (models.ComputerPageInfo, error) {
	var isOnline bool = true

	computerInfo, err := ad.PullComputerInformation(computer)

	if err != nil {
		return models.ComputerPageInfo{}, err
	}

	pingOutput, err := Ping(computerInfo.Name)

	if err != nil {
		isOnline = false
	}

	if isOnline {
		parsePingInformation(pingOutput)
	}

	return models.ComputerPageInfo{
		ComputerInfo: computerInfo,
		IsOnline:     isOnline,
		PingResults:  pingOutput,
	}, nil
}

func Ping(host string) (string, error) {
	cmd := exec.Command("ping", "-n", "1", host)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func parsePingInformation(output string) {

	out := strings.Split(output, "\r\n\r\n")

	fmt.Println(out)

	/*
		re := regexp.MustCompile(`\[(.*?)\]`)

		matches := re.FindStringSubmatch(output)

		fmt.Println(matches)
	*/
}
