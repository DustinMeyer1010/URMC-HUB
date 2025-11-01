package models

import "github.com/google/uuid"

type Bookmark struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	ImagePath   string    `json:"image_path"`
}
