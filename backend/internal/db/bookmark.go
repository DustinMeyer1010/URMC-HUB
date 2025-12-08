package db

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func AddBookmark(bookmark models.Bookmark) error {
	db, err := OpenAgentDB()

	if err != nil {
		return nil
	}
	query := "INSERT INTO bookmarks (name, description, url, image_path) VALUES (?, ?, ?, ?)"

	_, err = db.Exec(query, bookmark.Name, bookmark.Description, bookmark.URL, bookmark.ImagePath)

	if err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}

func GenericBookmarks() []byte {
	return bookmarks
}

// TODO
func GetAllGenericBookmarks() ([]models.Bookmark, error) {
	return []models.Bookmark{}, nil
}

// TODO
func GetAllBookmarks(username string) ([]models.Bookmark, error) {
	return []models.Bookmark{}, nil
}

// TODO
func GetBookmark(id string) error {
	return nil
}
