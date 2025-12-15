package db

import (
	"database/sql"
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	_ "modernc.org/sqlite"
)

func AgentDatabaseInit() error {

	fmt.Printf("%s\n", fmt.Sprintf("%s/%s.db", DBLocation, global.Username))

	db, err := sql.Open("sqlite", fmt.Sprintf("%s/%s.db", DBLocation, global.Username))

	if err != nil {
		logger.Errorf("Failed to Open/Create DB\n Error: %s", err.Error())
		return err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		logger.Errorf("Failed to Ping DB after creation\n Error: %s", err.Error())
		return err
	}

	logger.Infof("%s - Create/Open Successfuly\n", fmt.Sprintf("%s.db", global.Username))

	return nil
}

func addAgent(db *sql.DB) error {
	var count int
	query := "SELECT COUNT(*) FROM agent WHERE username = ?"
	err := db.QueryRow(query, global.Username).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if count > 0 {
		fmt.Println("Agent profile already exists")
		return nil
	}
	fmt.Println("Creating Agent Profile")
	query = "INSERT INTO agent (username) VALUES (?)"

	_, err = db.Exec(query, global.Username)

	if err != nil {
		return err
	}

	return nil

}
