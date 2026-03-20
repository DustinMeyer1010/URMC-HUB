package service

import (
	"fmt"
	"os/exec"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func PullPrinterInformation(server, queue string) (models.PrinterSimpleInfo, error) {
	printer := fmt.Sprintf("\\\\%s\\%s", server, queue)
	return ad.PullSinglePrinterInformation(printer)
}

func PingPrinter(ip string) error {
	cmd := exec.Command("ping", "-n", "1", ip)
	_, err := cmd.CombinedOutput()
	logger.Error(err)
	return err
}

func RelatedPrinters(ip string) ([]models.PrinterSimpleInfo, error) {
	return ad.RelatedPrinters(ip)
}
