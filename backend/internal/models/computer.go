package models

type ComputerSimpleInfo struct {
	Type            string `json:"Type"`
	Name            string `json:"Name"`
	OU              string `json:"OU"`
	OperatingSystem string `json:"OperatingSystem"`
}
