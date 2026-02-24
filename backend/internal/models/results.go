package models

// Used to return on all search
type AllResults struct {
	Users     []UserSimpleInfo     `json:"users"`
	Computers []ComputerSimpleInfo `json:"computers"`
	Groups    []GroupSimpleInfo    `json:"groups"`
	Printers  []PrinterSimpleInfo  `json:"printers"`
	Shares    []DriveSimpleInfo    `json:"drives"`
}

// Used to show Remove and Add of group results
type GroupModifyResults struct {
	Group      string `json:"group"`
	Message    string `json:"message"`
	Successful bool   `json:"successful"`
}
