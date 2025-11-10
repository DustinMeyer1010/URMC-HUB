package service

import (
	"fmt"
	"net"
	"time"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func PullPrinterInformation(server, queue string) (models.PrinterSimpleInfo, error) {
	printer := fmt.Sprintf("\\\\%s\\%s", server, queue)
	return ad.PullSinglePrinterInformation(printer)
}

func PingPrinter(ip string) error {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, "80"), timeout)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func RelatedPrinters(ip string) ([]models.PrinterSimpleInfo, error) {
	return ad.RelatedPrinters(ip)
}
