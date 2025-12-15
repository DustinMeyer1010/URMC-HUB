package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
)

var DBLocation string

//go:embed schema/*
var schema embed.FS

//go:embed generic_bookmarks.json
var bookmarks []byte

func Init() {

	exe, _ := os.Executable()
	base := filepath.Dir(exe)
	DBLocation = filepath.Join(base, "URMC_DB")

	err := os.MkdirAll(DBLocation, os.ModePerm)

	fmt.Println(DBLocation)

	if err != nil {
		fmt.Println("Error creating Directory: ", err)
		return
	}

}

func createTables(db *sql.DB) error {
	data, _ := fs.ReadFile(schema, "schema/tables.sql")

	_, err := db.Exec(string(data))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func OpenAgentDB() (*sql.DB, error) {
	return sql.Open("sqlite", fmt.Sprintf("%s/%s.db", DBLocation, global.Username))
}
