package models

type UserSimpleInfo struct {
	Type     string `json:"Type"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	NetID    string `json:"NetID"`
	OU       string `json:"OU"`
}
