package db

import "github.com/LostProgrammer1010/URMC-HUB/internal/models"

// Will Create a user account if one does not already exist
func CreateDBForAgent(agent models.Agent) {
	stmt, _ := DBConnection.Prepare("INSERT OR IGNORE INTO agents (username) VALUES (?)")
	stmt.Exec(agent.Username)
}
