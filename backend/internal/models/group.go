package models

type GroupSimpleInfo struct {
	Name        string `json:"Name"`
	OU          string `json:"OU"`
	Information string `json:"Information"`
	Description string `json:"Description"`
}
