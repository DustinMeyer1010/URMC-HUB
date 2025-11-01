package db

import (
	"encoding/json"
	"fmt"
	"log"

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

func GenerateGenericBookmarks() error {
	db, err := OpenAgentDB()
	if err != nil {
		return nil
	}

	err = createGenericAgent(db)

	if err != nil {
		return err
	}

	var genericBookmarks []models.Bookmark
	json.Unmarshal(bookmarks, &genericBookmarks)
	stmt, err := db.Prepare("INSERT INTO bookmarks (name, url, image_path, description) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, b := range genericBookmarks {
		_, err = stmt.Exec(b.Name, b.URL, b.ImagePath, b.Description)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil

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
