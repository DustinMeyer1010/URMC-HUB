package service

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func PullPrinterInformation(server, queue string) (models.PrinterSimpleInfo, error) {
	printer := fmt.Sprintf("\\\\%s\\%s", server, queue)
	return ad.PullSinglePrinterInformation(printer)
}
