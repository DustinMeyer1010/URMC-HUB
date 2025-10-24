package models

// Used to return on all search
type AllResults struct {
	Users     []UserSimpleInfo     `json:"users"`
	Computers []ComputerSimpleInfo `json:"computers"`
	Groups    []GroupSimpleInfo    `json:"groups"`
	Printers  []PrinterSimpleInfo  `json:"printers"`
	Shares    []DriveSimpleInfo    `json:"drives"`
}
