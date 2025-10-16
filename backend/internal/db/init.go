package db

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed schema/*.sql
var schema embed.FS
var DBConnection *sql.DB

func Init() {
	var err error
	dbPath := filepath.Join(getExecutableDir(), "database.db")
	DBConnection, err = sql.Open("sqlite", dbPath)

	if err != nil {
		panic(err)
	}

	err = createTables()
	if err != nil {
		panic(err)
	}

}

func createTables() error {

	tables, _ := schema.ReadFile("schema/init.sql")

	_, err := DBConnection.Exec(string(tables))

	return err
}

func getExecutableDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
