package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Add bookmark for specific agent
func AddBookmark(username string, bookmark models.Bookmark) error {
	db, err := OpenAgentDB(username)
	if err != nil {
		return fmt.Errorf("AGENT DATABASE NOT FOUND")
	}
	defer db.Close()

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

// Retrieves the bookmarks for specific agent
func GetBookmark(username string) ([]models.Bookmark, error) {
	bookmarks := []models.Bookmark{}
	db, err := OpenAgentDB(username)
	if err != nil {
		return bookmarks, err
	}
	defer db.Close()

	query := "SELECT * FROM bookmarks"
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return bookmarks, err
	}

	for rows.Next() {
		var bookmark models.Bookmark
		err := rows.Scan(&bookmark.Id, &bookmark.Name, &bookmark.URL, &bookmark.Description, &bookmark.ImagePath)
		fmt.Println(err)
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}
