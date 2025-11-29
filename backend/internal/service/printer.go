package service

import (
	"fmt"
	"os/exec"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func PullPrinterInformation(server, queue string) (models.PrinterSimpleInfo, *customError.Error) {
	printer := fmt.Sprintf("\\\\%s\\%s", server, queue)
	return ad.PullSinglePrinterInformation(printer)
}

func PingPrinter(ip string) error {
	cmd := exec.Command("ping", "-n", "1", ip)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	fmt.Println(err)
	return err
}

func RelatedPrinters(ip string) ([]models.PrinterSimpleInfo, *customError.Error) {
	return ad.RelatedPrinters(ip)
}
