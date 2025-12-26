package service

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
)

func GetAgentsWithBookmarks() []string {
	return db.GetAgentsWithBookmarks()
}
