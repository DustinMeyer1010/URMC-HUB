package models

type GroupSimpleInfo struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	OU          string `json:"OU"`
	Information string `json:"Information"`
	Description string `json:"Description"`
}
