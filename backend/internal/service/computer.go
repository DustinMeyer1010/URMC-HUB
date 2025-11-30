package service

import (
	"os/exec"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func ComputerInfo(computer string) (models.ComputerPageInfo, *customError.Error) {
	var isOnline bool = true

	computerInfo, cError := ad.PullComputerInformation(computer)

	if cError != nil {
		return models.ComputerPageInfo{}, cError
	}

	pingOutput, err := Ping(computerInfo.Name)

	if err != nil {
		isOnline = false
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

/*
func parsePingInformation(output string) {

	out := strings.SplitSeq(output, "\r\n\r\n")

	for line := range out {
		split := strings.SplitSeq(line, "\r\n")
		for section := range split {
			fmt.Printf("Section: %s\n", section)
		}
	}
		re := regexp.MustCompile(`\[(.*?)\]`)

		matches := re.FindStringSubmatch(output)

		fmt.Println(matches)
}
*/

type PingResults struct {
	DNS        string
	IP         string
	Reply      string
	Statistics string
	Packets    string
}
