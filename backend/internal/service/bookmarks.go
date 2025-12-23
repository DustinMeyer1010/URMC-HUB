package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func AddBookmark(r *http.Request) error {

	var bookmark models.Bookmark

	err := json.NewDecoder(r.Body).Decode(&bookmark)

	if err != nil {
		return fmt.Errorf("%s", "invalid body")
	}

	return db.AddBookmark(bookmark)
}

func GetAgentsWithBookmarks() []string {
	return db.GetAgentsWithBookmarks()
}
