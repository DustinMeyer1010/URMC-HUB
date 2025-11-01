package db

import (
	"database/sql"
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	_ "modernc.org/sqlite"
)

func AgentDatabaseInit() error {

	fmt.Printf("%s\n", fmt.Sprintf("%s/%s.db", path, global.Username))

	db, err := sql.Open("sqlite", fmt.Sprintf("%s/%s.db", path, global.Username))

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return err
	}

	fmt.Printf("%s - Create/Open Successfuly\n", fmt.Sprintf("%s.db", global.Username))

	if !initalized {
		createTables(db)
		addAgent(db)
		if !genericBookmarksCreated(db) {
			addGenericBookMarks(db)
		}
		initalized = true
	}

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

func genericBookmarksCreated(db *sql.DB) bool {
	var alreadyCreate bool
	query := "SELECT generic_bookmarks FROM agent WHERE username = ?"
	err := db.QueryRow(query, global.Username).Scan(&alreadyCreate)

	if err != nil {
		return true
	}

	if alreadyCreate {
		return true
	}

	return false
}
