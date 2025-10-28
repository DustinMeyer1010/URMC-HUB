package models

type Agent struct {
	Username string
}

type Link struct {
	ID          int    `json"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}
