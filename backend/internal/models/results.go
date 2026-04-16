package models

type AllResultsNew struct {
	Users     []map[string][]string `json:"users"`
	Computers []map[string][]string `json:"computers"`
	Groups    []map[string][]string `json:"groups"`
	Printers  []PrinterSimpleInfo   `json:"printers"`
	Drives    []DriveSimpleInfo     `json:"drives"`
}
