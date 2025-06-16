package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func CreateDatabase() {
	db, err := sql.Open("sqlite", "./database.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	TestTable := `
	CREATE TABLE IF NOT EXISTS profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	);`

	_, err = db.Exec(TestTable)

	if err != nil {
		log.Fatal(err)
	}
}
