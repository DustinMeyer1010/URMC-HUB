package models

type LockOutStatus struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Time  string `json:"time"`
}
