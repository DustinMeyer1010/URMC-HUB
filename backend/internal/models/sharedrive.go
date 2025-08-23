package models

type ShareDrive struct {
	Group     []string `json:"groups"`
	Drive     string   `json:"drive"`
	LocalPath string   `json:"local_path`
}
