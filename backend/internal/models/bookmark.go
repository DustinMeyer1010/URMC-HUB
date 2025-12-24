package models

type Bookmark struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
}
