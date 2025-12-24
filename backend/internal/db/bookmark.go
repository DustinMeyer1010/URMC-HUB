package db

import (
	"fmt"
	"os"
	"strings"

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

func GetAgentsWithBookmarks() []string {
	files, _ := os.ReadDir(DBLocation)
	agentsWithBookmarks := []string{}
	for _, file := range files {
		name := strings.Split(file.Name(), ".")[0]
		agentsWithBookmarks = append(agentsWithBookmarks, name)
	}

	return agentsWithBookmarks
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
func GetBookmark(id string) ([]models.Bookmark, error) {
	db, err := OpenAgentDB()
	var bookmarks []models.Bookmark
	if err != nil {
		return bookmarks, err
	}

	query := "SELECT * FROM bookmarks"

	rows, _ := db.Query(query)

	for rows.Next() {
		var bookmark models.Bookmark
		err := rows.Scan(&bookmark.Id, &bookmark.Name, &bookmark.URL, &bookmark.Description, &bookmark.ImagePath)
		fmt.Println(err)
		bookmarks = append(bookmarks, bookmark)
	}

	fmt.Println(bookmarks)

	return bookmarks, nil
}
