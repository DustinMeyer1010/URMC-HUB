package models

// Brief summary of the printer information for the search screen
type PrinterSimpleInfo struct {
	Server         string `json:"server"`
	Queue          string `json:"queue"`
	Model          string `json:"model"`
	IP             string `json:"IP"`
	PrintProcessor string `json:"print_processor"`
	Location       string `json:"location"`
	Notes          string `json:"notes"`
}
