package models

// Used to return on all search
type AllResults struct {
	Users     []UserSimpleInfo     `json:"users"`
	Computers []ComputerSimpleInfo `json:"computers"`
	Groups    []GroupSimpleInfo    `json:"groups"`
	Printers  []PrinterSimpleInfo  `json:"printers"`
	Shares    []DriveSimpleInfo    `json:"drives"`
}

type AllResultsNew struct {
	Users     []map[string][]string `json:"users"`
	Computers []map[string][]string `json:"computers"`
	Groups    []map[string][]string `json:"groups"`
	Printers  []PrinterSimpleInfo   `json:"printers"`
	Drives    []DriveSimpleInfo     `json:"drives"`
}

// Used to show Remove and Add of group results
type GroupModifyResults struct {
	Group      string `json:"group"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}
