package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

//go:embed schema/*.sql
var schema embed.FS

func CreateDatabase() {

	tables, _ := schema.ReadFile("schema/init.sql")

	dbPath := filepath.Join(getExecutableDir(), "database.db")
	fmt.Println(dbPath)
	db, err := sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(string(tables))

	if err != nil {
		log.Fatal(err)
	}
}

func getExecutableDir() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exePath)
}
